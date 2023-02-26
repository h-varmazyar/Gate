package main

import (
	"bytes"
	"context"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/eagle/configs"
	strategies "github.com/h-varmazyar/Gate/services/eagle/internal/app/strategies"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/db"
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

func initializeAndRegisterApps(ctx context.Context, logger *log.Logger, dbInstance *db.DB, configs *Configs) {
	strategiesApp, err := strategies.NewApp(ctx, logger, dbInstance, configs.StrategiesApp)
	if err != nil {
		logger.Panicf("failed to initiate strategies service with error: %v", err)
	}

	service.Serve(configs.GRPCPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		strategiesApp.Service.RegisterServer(server)
		//signalsApp.Service.RegisterServer(server)
		return server.Serve(lst)
	})

	service.Start(configs.ServiceName, configs.Version)
}
