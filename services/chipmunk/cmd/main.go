package main

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/assets"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles"
	candlesService "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/indicators"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets"
	marketsService "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/resolutions"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/wallets"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

func main() {
	ctx := context.Background()
	logger := log.New()
	conf, err := loadConfigs()
	if err != nil {
		log.Panic("failed to read configs")
	}

	dbInstance, err := loadDB(ctx, conf.DB)
	if err != nil {
		logger.Panicf("failed to initiate databases with error %v", err)
	}

	if err = amqpext.InitializeAMQP(conf.AMQPConfigs); err != nil {
		logger.Panicf("failed to initialize amqp: %v", err)
	}

	initializeAndRegisterApps(ctx, logger, dbInstance, conf)
}

func loadConfigs() (*Configs, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../configs")
	viper.AddConfigPath("/app/configs")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
		//localErr := viper.ReadConfig(bytes.NewBuffer(configs.DefaultConfig))
		//if localErr != nil {
		//	return nil, localErr
		//}
	}

	conf := new(Configs)
	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}

func loadDB(ctx context.Context, configs gormext.Configs) (*db.DB, error) {
	return db.NewDatabase(ctx, configs)
}

func initializeAndRegisterApps(ctx context.Context, logger *log.Logger, dbInstance *db.DB, configs *Configs) {
	var err error
	var assetsApp *assets.App
	assetsApp, err = assets.NewApp(ctx, logger, dbInstance)
	if err != nil {
		logger.WithError(err).Panicf("failed to initiate assets app")
	}

	var indicatorsApp *indicators.App
	indicatorsApp, err = indicators.NewApp(ctx, logger, dbInstance, configs.IndicatorsApp)
	if err != nil {
		logger.WithError(err).Panicf("failed to initiate markets app")
	}

	var resolutionsApp *resolutions.App
	resolutionsApp, err = resolutions.NewApp(ctx, logger, dbInstance, configs.ResolutionsApp)
	if err != nil {
		logger.WithError(err).Panicf("failed to initiate markets app")
	}

	marketDependencies := &markets.AppDependencies{
		ServiceDependencies: &marketsService.Dependencies{
			AssetsService:      assetsApp.Service,
			IndicatorsService:  indicatorsApp.Service,
			ResolutionsService: resolutionsApp.Service,
		},
	}

	var marketsApp *markets.App
	marketsApp, err = markets.NewApp(ctx, logger, dbInstance, configs.MarketsApp, marketDependencies)
	if err != nil {
		logger.WithError(err).Panicf("failed to initiate markets app")
	}

	candlesDependencies := &candles.AppDependencies{
		ServiceDependencies: &candlesService.Dependencies{
			ResolutionService: resolutionsApp.Service,
			IndicatorService:  indicatorsApp.Service,
		},
		IndicatorService:  indicatorsApp.Service,
		ResolutionService: resolutionsApp.Service,
		MarketService:     marketsApp.Service,
	}
	var candlesApp *candles.App
	candlesApp, err = candles.NewApp(ctx, logger, dbInstance, configs.CandlesApp, candlesDependencies)
	if err != nil {
		logger.WithError(err).Panicf("failed to initiate markets app")
	}

	walletDependencies := &wallets.AppDependencies{
		MarketService: marketsApp.Service,
	}

	var walletsApp *wallets.App
	walletsApp, err = wallets.NewApp(ctx, logger, configs.WalletsApp, walletDependencies)
	if err != nil {
		logger.WithError(err).Panicf("failed to initiate markets app")
	}

	service.Serve(configs.GRPCPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		assetsApp.Service.RegisterServer(server)
		marketsApp.Service.RegisterServer(server)
		candlesApp.Service.RegisterServer(server)
		indicatorsApp.Service.RegisterServer(server)
		resolutionsApp.Service.RegisterServer(server)
		walletsApp.Service.RegisterServer(server)
		return server.Serve(lst)
	})

	service.Start(configs.ServiceName, configs.Version)
}
