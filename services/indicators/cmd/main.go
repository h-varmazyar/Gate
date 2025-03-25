package main

import (
	"bytes"
	"context"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/indicators/configs"
	"github.com/h-varmazyar/Gate/services/indicators/internal/adapters/chipmunk"
	candlesConsumer "github.com/h-varmazyar/Gate/services/indicators/internal/brokers/consumer/candles"
	tickersConsumer "github.com/h-varmazyar/Gate/services/indicators/internal/brokers/consumer/tickers"
	"github.com/h-varmazyar/Gate/services/indicators/internal/repository"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/db"
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

const (
	serviceName = "indicators"
	version     = "0.0.1"
)

func main() {
	ctx := context.Background()
	logger := log.New()

	cfg, err := configs.Read()
	if err != nil {
		log.Panic("failed to read configs")
	}

	logger.Infof("starting %v(%v)", serviceName, version)

	dbInstance, err := loadDB(ctx, cfg.Database)
	if err != nil {
		logger.Panicf("failed to initiate databases with error %v", err)
	}

	natsConnection, err := nats.Connect(cfg.NatsURL)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}

	indicatorsRepo := repository.NewIndicatorRepository(dbInstance.DB)

	chipmunkAdapter := chipmunk.NewChipmunk(cfg.ChipmunkAdapter)

	defer natsConnection.Close()
	candlesConsumer, err := candlesConsumer.NewConsumer(logger, natsConnection, indicatorsRepo)
	if err != nil {
		log.Fatalf("Error creating candles consumer: %v", err)
	}
	tickersConsumer, err := tickersConsumer.NewConsumer(logger, natsConnection, indicatorsRepo, chipmunkAdapter)
	if err != nil {
		log.Fatalf("Error creating tickers consumer: %v", err)
	}
	candlesConsumer.StartListening()
	tickersConsumer.StartListening()

	initializeAndRegisterApps(ctx, logger, dbInstance, cfg)
}

func loadConfigs(defaultConfig bool) (*Configs, error) {
	log.Infof("reding configs...")

	if defaultConfig {
		viper.SetConfigType("yaml")
		log.Infof("reading deafult configs")
		err := viper.ReadConfig(bytes.NewBuffer(configs.DefaultConfig))
		if err != nil {
			log.WithError(err).Error("read from default configs failed")
			return nil, err
		}
	} else {
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Warnf("failed to read from env: %v", err)
			viper.AddConfigPath("./configs")  //path for docker compose configs
			viper.AddConfigPath("../configs") //path for local configs
			viper.SetConfigName("config")
			viper.SetConfigType("yaml")
			if err = viper.ReadInConfig(); err != nil {
				log.Warnf("failed to read from yaml: %v", err)
				localErr := viper.ReadConfig(bytes.NewBuffer(configs.DefaultConfig))
				if localErr != nil {
					log.WithError(localErr).Error("read from default configs failed")
					return nil, localErr
				}
			}
		}
	}

	conf := new(Configs)
	if err := viper.Unmarshal(conf); err != nil {
		log.Errorf("failed unmarshal")
		return nil, err
	}

	return conf, nil
}

func loadDB(ctx context.Context, configs gormext.Configs) (*db.DB, error) {
	return db.NewDatabase(ctx, configs)
}

func initializeAndRegisterApps(ctx context.Context, logger *log.Logger, dbInstance *db.DB, configs *configs.Config) {
	//walletDependencies := &internal.AppDependencies{}
	//
	//app, err := internal.NewApp(ctx, logger, configs.AppConfigs, walletDependencies, dbInstance)
	//if err != nil {
	//	logger.WithError(err).Panicf("failed to initiate wallets app")
	//}

	service.Serve(uint16(configs.GRPC.Port), func(lst net.Listener) error {
		server := grpc.NewServer()
		//app.Service.RegisterServer(server)
		return server.Serve(lst)
	})

	service.Start(serviceName, version)
}
