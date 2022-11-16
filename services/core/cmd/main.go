package main

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/core/internal/app/brokerages"
	"github.com/h-varmazyar/Gate/services/core/internal/app/functions"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages/coinex"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/db"
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

	if err = initializeAsyncHandlers(conf); err != nil {
		logger.WithError(err).Panicf("failed to initialize async handlers")
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

func initializeAsyncHandlers(configs *Configs) error {
	if err := coinex.ListenCallbacks(configs.CoinexConfigs); err != nil {
		log.WithError(err).Error("failed to initialize coinex callback listener")
		return err
	}
	return nil
}

func initializeAndRegisterApps(ctx context.Context, logger *log.Logger, dbInstance *db.DB, configs *Configs) {
	brokeragesApp, err := brokerages.NewApp(ctx, logger, dbInstance, configs.BrokeragesApp)
	if err != nil {
		logger.Panicf("failed to initiate brokerages service with error %v", err)
	}
	functionsApp, err := functions.NewApp(ctx, logger, configs.FunctionsApp, brokeragesApp.Service)
	if err != nil {
		logger.Panicf("failed to initiate functions service with error %v", err)
	}

	service.Serve(configs.GRPCPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		brokeragesApp.Service.RegisterServer(server)
		functionsApp.Service.RegisterServer(server)
		return server.Serve(lst)
	})

	service.Start(configs.ServiceName, configs.Version)
}