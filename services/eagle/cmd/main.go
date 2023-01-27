package main

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/service"
	strategies "github.com/h-varmazyar/Gate/services/eagle/internal/app/strategies"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/db"
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
	strategiesApp, err := strategies.NewApp(ctx, logger, dbInstance, configs.StrategiesApp)
	if err != nil {
		logger.Panicf("failed to initiate strategies service with error: %v", err)
	}

	//dependencies := &signals.AppDependencies{
	//	ServiceDependencies: &signalsService.Dependencies{
	//		StrategyService: strategiesApp.Service,
	//	},
	//}
	//signalsApp, err := signals.NewApp(ctx, logger, configs.SignalsApp, dependencies)
	//if err != nil {
	//	logger.Panicf("failed to initiate brokerages service with error %v", err)
	//}

	service.Serve(configs.GRPCPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		strategiesApp.Service.RegisterServer(server)
		//signalsApp.Service.RegisterServer(server)
		return server.Serve(lst)
	})

	service.Start(configs.ServiceName, configs.Version)
}
