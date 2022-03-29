package main

import (
	"fmt"
	"github.com/h-varmazyar/Gate/pkg/envext"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	"github.com/h-varmazyar/Gate/pkg/muxext"
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/ganjeh/internal/app"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

type Configs struct {
	GrpcPort       uint16 `env:"GRPC_PORT,required"`
	HttpPort       uint16 `env:"HTTP_PORT,required"`
	MaxLogsPerPage int64  `env:"MAX_LOGS_PER_PAGE,required"`
}

var (
	Name    = "ganjeh"
	Version = "v0.1.1"
)

func loadConfig() (*Configs, error) {
	configs := new(Configs)
	if err := envext.Load(configs); err != nil {
		return nil, err
	}
	return configs, nil
}

func main() {
	configs, err := loadConfig()
	if err != nil {
		log.WithError(err).Fatal("can not load service configs")
	}

	service.Serve(configs.GrpcPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		app.NewService().RegisterServer(server)
		return server.Serve(lst)
	})

	service.Serve(configs.HttpPort, func(lst net.Listener) error {
		connection := grpcext.NewConnection(fmt.Sprintf(":%v", configs.GrpcPort))
		router := muxext.NewRouter(true)
		app.NewController(connection).RegisterRouter(router)
		return http.Serve(lst, httpext.DefaultCors.Handler(router))
	})

	service.Start(Name, Version)
}
