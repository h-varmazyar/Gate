package main

import (
	"bytes"
	"context"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/core/configs"
	"github.com/h-varmazyar/Gate/services/core/internal/app/brokerages"
	"github.com/h-varmazyar/Gate/services/core/internal/app/functions"
	"github.com/h-varmazyar/Gate/services/core/internal/app/platforms"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages/coinex"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/db"
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

	logger.Infof("running %v(%v)", conf.ServiceName, conf.Version)

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

func loadConfigs() (*Configs, error) {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")  //path for docker compose configs
	viper.AddConfigPath("../configs") //path for local configs
	if err := viper.ReadInConfig(); err != nil {
		localErr := viper.ReadConfig(bytes.NewBuffer(configs.DefaultConfig))
		if localErr != nil {
			return nil, localErr
		}
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

	var platformsApp *platforms.App
	platformsApp, err = platforms.NewApp(ctx, logger, configs.PlatformsApp)
	if err != nil {
		logger.WithError(err).Panicf("failed to initiate assets app")
	}

	service.Serve(configs.GRPCPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		brokeragesApp.Service.RegisterServer(server)
		functionsApp.Service.RegisterServer(server)
		platformsApp.Service.RegisterServer(server)
		return server.Serve(lst)
	})

	service.Start(configs.ServiceName, configs.Version)
}
