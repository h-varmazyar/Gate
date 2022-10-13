package main

import (
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/eagle/configs"
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/signals"
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/strategies"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/repository"
	"google.golang.org/grpc"
	"net"
)

func main() {
	//initializing
	configs.LoadVariables()
	repository.InitializingDB()

	registerGrpcServer()

	service.Start(configs.Variables.ServiceName, configs.Variables.Version)
}

func registerGrpcServer() {
	service.Serve(configs.Variables.GrpcPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		strategies.NewService().RegisterServer(server)
		signals.NewService().RegisterServer(server)
		return server.Serve(lst)
	})
}
