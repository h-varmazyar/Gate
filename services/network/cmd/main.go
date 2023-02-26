package main

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/network/internal/app/IPs"
	"github.com/h-varmazyar/Gate/services/network/internal/app/rateLimiters"
	"github.com/h-varmazyar/Gate/services/network/internal/app/requests"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/db"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net"
)

func main() {
	ctx := context.Background()
	logger := log.New()
	configs := loadConfigs(logger)

	dbInstance, err := loadDB(ctx, configs.DB)
	if err != nil {
		logger.WithError(err).Panic("failed to initiate databases")
	}

	if err := amqpext.InitializeAMQP(configs.AMQPConfigs); err != nil {
		logger.WithError(err).Panic("can not initialize AMQP")
	}

	initializeAndRegisterApps(ctx, logger, dbInstance, configs)
}

func loadConfigs(logger *log.Logger) *Configs {
	configs := new(Configs)
	confBytes, err := ioutil.ReadFile("../configs/config.yaml")
	if err != nil {
		logger.WithError(err).Fatal("can not load yaml file")
	}
	if err = yaml.Unmarshal(confBytes, configs); err != nil {
		logger.WithError(err).Fatal("can not unmarshal yaml file")
	}
	return configs
}

func loadDB(ctx context.Context, configs *db.Configs) (*db.DB, error) {
	return db.NewDatabase(ctx, configs)
}

func initializeAndRegisterApps(ctx context.Context, logger *log.Logger, db *db.DB, configs *Configs) {
	var err error
	var ipsApp *IPs.App
	ipsApp, err = IPs.NewApp(ctx, logger, db, configs.IPsApp)
	if err != nil {
		logger.WithError(err).Panicf("failed to initiate IPs service")
	}

	var rateLimitersApp *rateLimiters.App
	rateLimitersApp, err = rateLimiters.NewApp(ctx, logger, db, configs.RateLimitersApp)
	if err != nil {
		logger.WithError(err).Panic("failed to initiate rate limiters service")
	}

	var requestsApp *requests.App
	requestsApp, err = requests.NewApp(ctx, logger, configs.RequestsApp, rateLimitersApp.Service, ipsApp.Service)
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
