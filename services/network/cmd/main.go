package main

import (
	"bytes"
	"context"
	"flag"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/network/configs"
	"github.com/h-varmazyar/Gate/services/network/internal/app/IPs"
	"github.com/h-varmazyar/Gate/services/network/internal/app/rateLimiters"
	"github.com/h-varmazyar/Gate/services/network/internal/app/requests"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/db"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

func main() {
	ctx := context.Background()
	logger := log.New()

	defaultConf := flag.Bool("default-configs", false, "run program with default config")
	flag.Parse()

	conf, err := loadConfigs(*defaultConf)
	if err != nil {
		log.Panicf("failed to read configs: %v", err)
	}

	logger.Infof("starting %v(%v)", conf.ServiceName, conf.Version)

	dbInstance, err := loadDB(ctx, conf.DB)
	if err != nil {
		logger.WithError(err).Panic("failed to initiate databases")
	}

	if err := amqpext.InitializeAMQP(conf.AMQPConfigs); err != nil {
		logger.WithError(err).Panic("can not initialize AMQP")
	}

	initializeAndRegisterApps(ctx, logger, dbInstance, conf)
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
		log.Errorf("faeiled unmarshal")
		return nil, err
	}

	return conf, nil
}

func loadDB(ctx context.Context, configs gormext.Configs) (*db.DB, error) {
	return db.NewDatabase(ctx, configs)
}

func initializeAndRegisterApps(ctx context.Context, logger *log.Logger, db *db.DB, configs *Configs) {
	var err error
	var ipsApp *IPs.App
	ipsApp, err = IPs.NewApp(ctx, logger, db)
	if err != nil {
		logger.WithError(err).Panicf("failed to initiate IPs service")
	}

	var rateLimitersApp *rateLimiters.App
	rateLimitersApp, err = rateLimiters.NewApp(ctx, logger, db)
	if err != nil {
		logger.WithError(err).Panic("failed to initiate rate limiters service")
	}

	var requestsApp *requests.App
	requestsApp, err = requests.NewApp(ctx, logger, rateLimitersApp.Service, ipsApp.Service)
	if err != nil {
		logger.WithError(err).Panic("failed to initiate requests service")
	}

	service.Serve(configs.GRPCPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		requestsApp.Service.RegisterServer(server)
		ipsApp.Service.RegisterServer(server)
		rateLimitersApp.Service.RegisterServer(server)
		return server.Serve(lst)
	})

	service.Start(configs.ServiceName, configs.Version)
}
