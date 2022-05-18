package main

import (
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/assets"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/indicators"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/resolutions"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/wallets"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
	"google.golang.org/grpc"
	"net"
)

func main() {
	//initializing
	configs.Load()
	repository.InitializingDB()
	buffer.NewMarketInstance()

	service.Serve(configs.Variables.GrpcPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		assets.NewService().RegisterServer(server)
		candles.NewService().RegisterServer(server)
		wallets.NewService().RegisterServer(server)
		indicators.NewService().RegisterServer(server)
		markets.NewService().RegisterServer(server)
		resolutions.NewService().RegisterServer(server)
		return server.Serve(lst)
	})

	markets.InitializeWorker()
	wallets.InitializeWorker()

	service.Start(configs.Variables.ServiceName, configs.Variables.Version)
}
