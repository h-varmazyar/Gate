package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron/v2"
	"github.com/h-varmazyar/Gate/services/gather/configs"
	"github.com/h-varmazyar/Gate/services/gather/internal/adapters/core"
	"github.com/h-varmazyar/Gate/services/gather/internal/pkg/buffer"
	"github.com/h-varmazyar/Gate/services/gather/internal/repositories"
	"github.com/h-varmazyar/Gate/services/gather/internal/workers/lastCandle"
	"github.com/h-varmazyar/Gate/services/gather/internal/workers/marketUpdate"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/db"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"syscall"
)

const (
	serviceName = "Gathering"
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

	assetsRepo := repositories.NewAssetRepository(logger, dbInstance.DB)
	candlesRepo := repositories.NewCandleRepository(logger, dbInstance.DB)
	resolutionsRepo := repositories.NewResolutionRepository(logger, dbInstance.DB)
	marketsRepo := repositories.NewMarketRepository(logger, dbInstance.DB)

	coreAdapter := core.NewCore(cfg.CoreAdapter)

	if err = marketUpdate.NewWorker(logger, cfg.MarketUpdateWorker, assetsRepo, candlesRepo, marketsRepo, resolutionsRepo, coreAdapter).Run(scheduler); err != nil {
		panic(err)
	}
	if err = lastCandle.NewWorker(logger, cfg.LastCandleWorker, coreAdapter, candlesRepo, marketsRepo, resolutionsRepo).Run(); err != nil {
		panic(err)
	}
	scheduler.Start()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	<-done

	scheduler.Shutdown()
}
