package main

import (
	"bytes"
	"context"
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

	conf, err := loadConfigs()
	if err != nil {
		log.Panic("failed to read configs")
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

func loadConfigs() (*Configs, error) {
	viper.SetConfigName("config")
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
