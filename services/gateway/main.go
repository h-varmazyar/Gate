package main

import (
	"github.com/h-varmazyar/Gate/pkg/muxext"
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/gateway/configs"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/brokerage"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/chipmunk"
	"github.com/h-varmazyar/Gate/services/gateway/internal/app/eagle"
	"google.golang.org/grpc"
	"net"
)

func main() {
	configs.Load()

	service.Serve(configs.Variables.GrpcPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		router := muxext.NewRouter(true)
		//router.Use(httpext.Authorization)

		brokerage.RegisterRoutes(router)
		chipmunk.RegisterRoutes(router)
		eagle.RegisterRoutes(router)
		return server.Serve(lst)
	})

	service.Start(configs.Variables.ServiceName, configs.Variables.Version)
}
