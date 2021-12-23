package main

import (
	"fmt"
	"github.com/mrNobody95/Gate/pkg/envext"
	"github.com/mrNobody95/Gate/pkg/grpcext"
	"github.com/mrNobody95/Gate/pkg/httpext"
	"github.com/mrNobody95/Gate/pkg/muxext"
	"github.com/mrNobody95/Gate/pkg/service"
	"github.com/mrNobody95/Gate/services/brokerage/configs"
	"github.com/mrNobody95/Gate/services/brokerage/internal/app/assets"
	"github.com/mrNobody95/Gate/services/brokerage/internal/app/brokerages"
	"github.com/mrNobody95/Gate/services/brokerage/internal/app/candles"
	"github.com/mrNobody95/Gate/services/brokerage/internal/app/markets"
	"github.com/mrNobody95/Gate/services/brokerage/internal/app/resolutions"
	"github.com/mrNobody95/Gate/services/brokerage/internal/app/wallets"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/repository"
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

var (
	Name    = "brokerage"
	Version = "v0.5.0"
)

func loadConfig() (*configs.Configs, error) {
	configs := new(configs.Configs)
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
	repository.LoadRepositories(configs.DatabaseConnection)
	service.Serve(configs.GrpcPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		assets.NewService().RegisterServer(server)
		brokerages.NewService().RegisterServer(server)
		markets.NewService().RegisterServer(server)
		wallets.NewService(configs).RegisterServer(server)
		resolutions.NewService().RegisterServer(server)
		candles.NewService(configs).RegisterServer(server)
		return server.Serve(lst)
	})

	service.Serve(configs.HttpPort, func(lst net.Listener) error {
		connection := grpcext.NewConnection(fmt.Sprintf(":%v", configs.GrpcPort))
		router := muxext.NewRouter(true)
		wallets.NewController(connection).RegisterRouter(router)
		return http.Serve(lst, httpext.DefaultCors.Handler(router))
	})

	service.Start(Name, Version)
}
