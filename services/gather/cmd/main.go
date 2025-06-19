package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron/v2"
	"github.com/h-varmazyar/Gate/pkg/db"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	"github.com/h-varmazyar/Gate/pkg/service"
	v1 "github.com/h-varmazyar/Gate/services/gather/api/rest/v1"
	candlesHandler "github.com/h-varmazyar/Gate/services/gather/api/rest/v1/candles"
	"github.com/h-varmazyar/Gate/services/gather/configs"
	"github.com/h-varmazyar/Gate/services/gather/internal/adapters/coinex"
	"github.com/h-varmazyar/Gate/services/gather/internal/adapters/core"
	"github.com/h-varmazyar/Gate/services/gather/internal/adapters/sahamyab"
	"github.com/h-varmazyar/Gate/services/gather/internal/pkg/buffer"
	"github.com/h-varmazyar/Gate/services/gather/internal/repositories"
	"github.com/h-varmazyar/Gate/services/gather/internal/services/candles"
	"github.com/h-varmazyar/Gate/services/gather/internal/workers/candleTicker"
	"github.com/h-varmazyar/Gate/services/gather/internal/workers/lastCandle"
	"github.com/h-varmazyar/Gate/services/gather/internal/workers/marketUpdate"
	"github.com/h-varmazyar/Gate/services/gather/internal/workers/sahamyabArchive"
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"net"
	"net/http"
	"time"
)

const (
	serviceName = "gather"
	version     = "v0.0.1"
)

func main() {
	ctx := context.Background()
	logger := log.New()
	cfg, err := configs.Read()
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(cfg.GinMode)

	logger.Infof("starting %v(%v)", serviceName, version)

	dbInstance, err := db.NewDatabase(ctx, cfg.Database)
	if err != nil {
		logger.Panicf("failed to initiate databases with error %v", err)
	}

	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	buffer.InitializeCandleBuffer(cfg.CandleBuffer)

	natsConnection, err := nats.Connect(cfg.Nats.URL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer natsConnection.Close()

	//natsProducer := candlesProducer.NewProducer(logger, natsConnection)

	assetsRepo := repositories.NewAssetRepository(logger, dbInstance.DB)
	candlesRepo := repositories.NewCandleRepository(logger, dbInstance.DB)
	resolutionsRepo := repositories.NewResolutionRepository(logger, dbInstance.DB)
	marketsRepo := repositories.NewMarketRepository(logger, dbInstance.DB)
	postRepo := repositories.NewPostRepository(logger, dbInstance.DB)

	sahamyabAdapter := sahamyab.NewSahamyab(cfg.SahamyabAdapter)
	coreAdapter := core.NewCore(cfg.CoreAdapter)
	coinexAdapter, err := coinex.NewCoinex(cfg.CoinexAdapter)
	if err != nil {
		panic(err)
	}

	candlesService := candles.NewService(logger, candlesRepo)

	lastCandleWorker := lastCandle.NewWorker(logger, cfg.LastCandleWorker, coreAdapter, coinexAdapter, candlesRepo, marketsRepo, resolutionsRepo)
	candleTickerWorker := candleTicker.NewWorker(logger, cfg.TickerWorker, coinexAdapter, marketsRepo)
	marketUpdateWorker := marketUpdate.NewWorker(logger, cfg.MarketUpdateWorker, assetsRepo, candlesRepo, marketsRepo, resolutionsRepo, coreAdapter, coinexAdapter, candleTickerWorker, lastCandleWorker)
	sahamyabArchiveWorker, err := sahamyabArchive.NewWorker(logger, cfg.SahamyabArchive, postRepo, sahamyabAdapter)
	if err != nil {
		log.WithError(err).Panicf("failed to initiate sahamyab archive worker")
	}
	//sahamyabStreamWorker := sahamyabStream.NewWorker(logger, cfg.SahamyabStream, natsProducer, postRepo)

	if cfg.WarmupWorker.NeedWarmup {
		marketUpdateWorker.ImmediateUpdate()
	}
	//
	//fmt.Println(int64(time.Minute * 15))
	//
	//if err = lastCandleWorker.Run(); err != nil {
	//	panic(err)
	//}
	////if err = candleTickerWorker.Run(); err != nil {
	////	panic(err)
	////}
	//if err = marketUpdateWorker.Run(scheduler); err != nil {
	//	panic(err)
	//}

	if err = sahamyabArchiveWorker.Start(); err != nil {
		panic(err)
	}

	//if err = sahamyabStreamWorker.Start(); err != nil {
	//	panic(err)
	//}

	scheduler.Start()
	defer scheduler.Shutdown()

	candlesHandler := candlesHandler.New(logger, candlesService)

	service.Serve(uint16(cfg.HTTP.APIPort), func(lst net.Listener) error {
		router := gin.New()

		router.GET("/ping", func(c *gin.Context) {
			pong := struct {
				Time string
			}{
				Time: time.Now().String(),
			}
			httpext.SendGinModel(c, http.StatusOK, pong)
		})

		apiRouter := router.Group("/api")

		v1.NewRouter(logger, candlesHandler).RegisterRoutes(apiRouter)

		return http.Serve(lst, router)
	})

	service.Start(serviceName, version)
}
