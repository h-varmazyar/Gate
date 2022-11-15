package main

import (
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/app"
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/pkg/repository"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net"
)

func main() {
	//initializing
	configs := loadConfigs()

	repository.InitializingDB(configs.DBConfigs)

	service.Serve(configs.GrpcPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		app.NewService(configs.ServiceConfigs).RegisterServer(server)
		return server.Serve(lst)
	})

	service.Start(configs.ServiceName, configs.Version)
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
