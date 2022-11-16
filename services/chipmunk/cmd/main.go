package main

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/assets"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/indicators"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/resolutions"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/wallets"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net"
)

func main() {
	ctx := context.Background()
	conf := loadConfigs()
	logger := log.New()
	dbInstance, err := loadDB(ctx, conf.DB)
	if err != nil {
		logger.Panicf("failed to initiate databases with error %v", err)
	}

	if err = amqpext.InitializeAMQP(conf.AMQPConfigs); err != nil {
		logger.Panicf("failed to initialize amqp: %v", err)
	}

	initializeAndRegisterApps(ctx, logger, dbInstance, conf)
}

func loadConfigs() *Configs {
	configs := new(Configs)
	confBytes, err := ioutil.ReadFile("../configs/local.yaml")
	if err != nil {
		log.WithError(err).Fatal("can not load yaml file")
	}
	if err = yaml.Unmarshal(confBytes, configs); err != nil {
		log.WithError(err).Fatal("can not unmarshal yaml file")
	}
	return configs
}

func loadDB(ctx context.Context, configs *db.Configs) (*db.DB, error) {
	return db.NewDatabase(ctx, configs)
}

func initializeAndRegisterApps(ctx context.Context, logger *log.Logger, dbInstance *db.DB, configs *Configs) {
	var err error
	var assetsApp *assets.App
	assetsApp, err = assets.NewApp(ctx, logger, dbInstance, configs.AssetsApp)
	if err != nil {
		logger.WithError(err).Panicf("failed to initiate assets app")
	}

	var marketsApp *markets.App
	marketsApp, err = markets.NewApp(ctx, logger, dbInstance, configs.MarketsApp)
	if err != nil {
		logger.WithError(err).Panicf("failed to initiate markets app")
	}

	var candlesApp *candles.App
	candlesApp, err = candles.NewApp(ctx, logger, dbInstance, configs.CandlesApp)
	if err != nil {
		logger.WithError(err).Panicf("failed to initiate markets app")
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

	var walletsApp *wallets.App
	walletsApp, err = wallets.NewApp(ctx, logger, configs.ResolutionsApp)
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
