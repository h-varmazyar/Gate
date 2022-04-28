package main

import (
	"encoding/json"
	"github.com/h-varmazyar/Gate/pkg/envext"
	"github.com/h-varmazyar/Gate/pkg/httpext"
	"github.com/h-varmazyar/Gate/pkg/muxext"
	"github.com/h-varmazyar/Gate/pkg/service"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	"github.com/h-varmazyar/Gate/services/brokerage/configs"
	"github.com/h-varmazyar/Gate/services/brokerage/internal/app/assets"
	"github.com/h-varmazyar/Gate/services/brokerage/internal/app/brokerages"
	"github.com/h-varmazyar/Gate/services/brokerage/internal/app/candles"
	"github.com/h-varmazyar/Gate/services/brokerage/internal/app/markets"
	"github.com/h-varmazyar/Gate/services/brokerage/internal/app/resolutions"
	"github.com/h-varmazyar/Gate/services/brokerage/internal/app/wallets"
	"github.com/h-varmazyar/Gate/services/brokerage/internal/pkg/repository"
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/strategy"
	repository2 "github.com/h-varmazyar/Gate/services/eagle/internal/pkg/repository"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

var (
	Name    = "brokerage"
	Version = "v0.7.0"
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

	if err = createStrategy(); err != nil {
		log.WithError(err).Fatal("failed to create strategy")
	}

	service.Serve(configs.GrpcPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		assets.NewService().RegisterServer(server)
		brokerages.NewService(configs).RegisterServer(server)
		markets.NewService(configs).RegisterServer(server)
		wallets.NewService(configs).RegisterServer(server)
		resolutions.NewService().RegisterServer(server)
		strategy.NewService().RegisterServer(server)
		candles.NewService(configs).RegisterServer(server)
		return server.Serve(lst)
	})

	service.Serve(configs.HttpPort, func(lst net.Listener) error {
		router := muxext.NewRouter(true)
		return http.Serve(lst, httpext.DefaultCors.Handler(router))
	})

	service.Start(Name, Version)
}

func createStrategy() error {
	list := make([]*repository2.Strategy, 0)
	var err error
	if list, err = repository.Strategies.List(); err != nil {
		return err
	}
	if len(list) == 0 {
		strategy := &repository2.Strategy{
			Name:        "خودکار",
			Description: "انجام معاملات به صورت خودکار",
		}
		err = repository.Strategies.Save(strategy)
		if err != nil {
			return err
		}
		var bytes []byte
		{
			bytes, err = json.Marshal(struct {
				Length int `json:"length"`
			}{
				Length: 7,
			})
			if err != nil {
				return err
			}
			if err = repository.Indicators.Save(&repository.Indicator{
				StrategyRefer: strategy.ID,
				Name:          brokerageApi.IndicatorNames_RSI,
				Description:   "fast rsi",
				Configs:       bytes,
			}); err != nil {
				return err
			}
		}

		{
			bytes, err = json.Marshal(struct {
				Length  int `json:"length"`
				SmoothK int `json:"smooth_k"`
				SmoothD int `json:"smooth_d"`
			}{
				Length:  14,
				SmoothK: 3,
				SmoothD: 10,
			})
			if err != nil {
				return err
			}
			if err = repository.Indicators.Save(&repository.Indicator{
				StrategyRefer: strategy.ID,
				Name:          brokerageApi.IndicatorNames_Stochastic,
				Description:   "slow stochastic",
				Configs:       bytes,
			}); err != nil {
				return err
			}
		}

		{
			bytes, err = json.Marshal(struct {
				Length    int    `json:"length"`
				Deviation int    `json:"deviation"`
				Source    string `json:"source"`
			}{
				Length:    20,
				Deviation: 2,
				Source:    "close",
			})
			if err != nil {
				return err
			}
			if err = repository.Indicators.Save(&repository.Indicator{
				StrategyRefer: strategy.ID,
				Name:          brokerageApi.IndicatorNames_BollingerBands,
				Description:   "regular bollinger bands",
				Configs:       bytes,
			}); err != nil {
				return err
			}
		}
	}
	return nil
}
