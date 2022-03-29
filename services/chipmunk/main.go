package main

import (
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/ohlc"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/wallets"
	"google.golang.org/grpc"
	"net"
)

func main() {
	service.Serve(configs.Variables.GrpcPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		ohlc.NewService().RegisterServer(server)
		wallets.NewService().RegisterServer(server)
		return server.Serve(lst)
	})

	service.Start(configs.Variables.ServiceName, configs.Variables.Version)
}
