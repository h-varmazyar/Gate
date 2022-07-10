package main

import (
	"github.com/h-varmazyar/Gate/pkg/httpext"
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/brokerage/configs"
	"github.com/h-varmazyar/Gate/services/brokerage/internal/app/brokerages"
	"github.com/h-varmazyar/Gate/services/brokerage/internal/app/functions"
	"github.com/h-varmazyar/Gate/services/brokerage/internal/pkg/repository"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func main() {
	configs.Load()
	repository.InitializingDB()

	service.Serve(configs.Variables.GrpcPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		brokerages.NewService().RegisterServer(server)
		functions.NewService().RegisterServer(server)
		return server.Serve(lst)
	})

	service.Serve(configs.Variables.HttpPort, func(lst net.Listener) error {
		router := muxext.NewRouter(true)
		return http.Serve(lst, httpext.DefaultCors.Handler(router))
	})

	service.Start(configs.Variables.ServiceName, configs.Variables.Version)
}

//func createStrategy() error {
//	list := make([]*repository.Strategy, 0)
//	var err error
//	if list, err = repository.Strategies.List(); err != nil {
//		return err
//	}
//	if len(list) == 0 {
//		strategy := &repository.Strategy{
//			Name:        "خودکار",
//			Description: "انجام معاملات به صورت خودکار",
//		}
//		err = repository.Strategies.Save(strategy)
//		if err != nil {
//			return err
//		}
//		var bytes []byte
//		{
//			bytes, err = json.Marshal(struct {
//				Length int `json:"length"`
//			}{
//				Length: 7,
//			})
//			if err != nil {
//				return err
//			}
//			if err = repository.Indicators.Save(&repository.Indicator{
//				StrategyRefer: strategy.ID,
//				Name:          brokerageApi.IndicatorNames_RSI,
//				Description:   "fast rsi",
//				Configs:       bytes,
//			}); err != nil {
//				return err
//			}
//		}
//
//		{
//			bytes, err = json.Marshal(struct {
//				Length  int `json:"length"`
//				SmoothK int `json:"smooth_k"`
//				SmoothD int `json:"smooth_d"`
//			}{
//				Length:  14,
//				SmoothK: 3,
//				SmoothD: 10,
//			})
//			if err != nil {
//				return err
//			}
//			if err = repository.Indicators.Save(&repository.Indicator{
//				StrategyRefer: strategy.ID,
//				Name:          brokerageApi.IndicatorNames_Stochastic,
//				Description:   "slow stochastic",
//				Configs:       bytes,
//			}); err != nil {
//				return err
//			}
//		}
//
//		{
//			bytes, err = json.Marshal(struct {
//				Length    int    `json:"length"`
//				Deviation int    `json:"deviation"`
//				Source    string `json:"source"`
//			}{
//				Length:    20,
//				Deviation: 2,
//				Source:    "close",
//			})
//			if err != nil {
//				return err
//			}
//			if err = repository.Indicators.Save(&repository.Indicator{
//				StrategyRefer: strategy.ID,
//				Name:          brokerageApi.IndicatorNames_BollingerBands,
//				Description:   "regular bollinger bands",
//				Configs:       bytes,
//			}); err != nil {
//				return err
//			}
//		}
//	}
//	return nil
//}
