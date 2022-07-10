package main

import (
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/telegramBot/configs"
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/app"
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/pkg/repository"
	"google.golang.org/grpc"
	"net"
)

func main() {
	//initializing
	configs.LoadVariables()
	repository.InitializingDB()

	service.Serve(configs.Variables.GrpcPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		app.NewService().RegisterServer(server)
		return server.Serve(lst)
	})

	service.Start(configs.Variables.ServiceName, configs.Variables.Version)
}
