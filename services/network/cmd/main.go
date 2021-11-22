package main

import (
	"github.com/mrNobody95/Gate/pkg/envext"
	"github.com/mrNobody95/Gate/pkg/service"
	"github.com/mrNobody95/Gate/services/network/internal/app"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
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
	GrpcPort uint16 `env:"GRPC_PORT,required"`
}

var (
	Name    = "network"
	Version = "v0.2.0"
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

	service.Start(Name, Version)
}
