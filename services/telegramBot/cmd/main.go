package main

import (
	"bytes"
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/telegramBot/configs"
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/app"
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/pkg/repository"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

func main() {
	logger := log.New()

	conf, err := loadConfigs()
	if err != nil {
		log.Panic("failed to read configs")
	}

	logger.Infof("starting %v(%v)", conf.ServiceName, conf.Version)

	repository.InitializingDB(conf.DB)

	service.Serve(conf.GrpcPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		app.NewService(conf.ServiceConfigs).RegisterServer(server)
		return server.Serve(lst)
	})

	service.Start(conf.ServiceName, conf.Version)
}

func loadConfigs() (*Configs, error) {
	log.Infof("reding configs...")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Warnf("failed to read from env: %v", err)
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./configs")  //path for docker compose configs
		viper.AddConfigPath("../configs") //path for local configs
		if err = viper.ReadInConfig(); err != nil {
			log.Warnf("failed to read from yaml: %v", err)
			localErr := viper.ReadConfig(bytes.NewBuffer(configs.DefaultConfig))
			if localErr != nil {
				log.Infof("read from default env")
				return nil, localErr
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
