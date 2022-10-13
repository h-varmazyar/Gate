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

	//rsi := repository.Indicator{
	//	UniversalModel: gormext.UniversalModel{},
	//	Type:           chipmunkApi.Indicator_RSI,
	//	Configs: &repository.IndicatorConfigs{
	//		RSI: &repository.RsiConfigs{Length: 7},
	//	},
	//}
	//
	//log.Infof("rsi creation: %v", repository.Indicators.Create(&rsi))

	////id, _ := uuid.Parse("d4cc478a-7783-47f6-9c73-987b6c675d5d")
	//id, _ := uuid.Parse("1917ae7b-a3cf-4e59-8e78-d1a44c073c81")
	//rsi, err := repository.Indicators.Return(id)
	//log.Errorf("rsi fetch err: %v", err)
	//log.Infof("rsi fetch: %v", rsi)
	//log.Warnf("rsi fetch: %v", rsi.Configs)

	service.Start(configs.Variables.ServiceName, configs.Variables.Version)
}
