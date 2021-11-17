package main

import (
	"fmt"
	"github.com/mrNobody95/Gate/pkg/envext"
	"github.com/mrNobody95/Gate/pkg/grpcext"
	"github.com/mrNobody95/Gate/pkg/httpext"
	"github.com/mrNobody95/Gate/pkg/muxext"
	"github.com/mrNobody95/Gate/pkg/service"
	"github.com/mrNobody95/Gate/services/ganjeh/internal/app"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 12.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type Configs struct {
	GrpcPort       uint16 `env:"GRPC_PORT,required"`
	HttpPort       uint16 `env:"HTTP_PORT,required"`
	MaxLogsPerPage int64  `env:"MAX_LOGS_PER_PAGE,required"`
}

var (
	Name    = "ganjeh"
	Version = "v0.1.0"
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
