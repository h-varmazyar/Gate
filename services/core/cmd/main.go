package main

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/service"
	"github.com/h-varmazyar/Gate/services/core/internal/app/brokerages"
	"github.com/h-varmazyar/Gate/services/core/internal/app/functions"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages/coinex"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/db"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net"
)

func main() {
	ctx := context.Background()
	conf := loadConfigs()
	logger := log.New()
	dbInstance, err := loadDB(ctx, conf.DB)
	if err != nil {
		logger.Panicf("failed to initiate databases with error %v", err)
	}

	if err = amqpext.InitializeAMQP(conf.AMQPConfigs); err != nil {
		logger.Panicf("failed to initialize amqp: %v", err)
	}

	if err = initializeAsyncHandlers(conf); err != nil {
		logger.WithError(err).Panicf("failed to initialize async handlers")
	}

	initializeAndRegisterApps(ctx, logger, dbInstance, conf)
}

func loadConfigs() *Configs {
	configs := new(Configs)
	confBytes, err := ioutil.ReadFile("../configs/local.yaml")
	if err != nil {
		log.WithError(err).Fatal("can not load yaml file")
	}
	if err = yaml.Unmarshal(confBytes, configs); err != nil {
		log.WithError(err).Fatal("can not unmarshal yaml file")
	}
	return configs
}

func loadDB(ctx context.Context, configs *db.Configs) (*db.DB, error) {
	return db.NewDatabase(ctx, configs)
}

func initializeAsyncHandlers(configs *Configs) error {
	if err := coinex.ListenCallbacks(configs.CoinexConfigs); err != nil {
		log.WithError(err).Error("failed to initialize coinex callback listener")
		return err
	}
	return nil
}

func initializeAndRegisterApps(ctx context.Context, logger *log.Logger, dbInstance *db.DB, configs *Configs) {
	brokeragesApp, err := brokerages.NewApp(ctx, logger, dbInstance, configs.BrokeragesApp)
	if err != nil {
		logger.Panicf("failed to initiate brokerages service with error %v", err)
	}
	functionsApp, err := functions.NewApp(ctx, logger, configs.FunctionsApp, brokeragesApp.Service)
	if err != nil {
		logger.Panicf("failed to initiate functions service with error %v", err)
	}

	service.Serve(configs.GRPCPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		brokeragesApp.Service.RegisterServer(server)
		functionsApp.Service.RegisterServer(server)
		return server.Serve(lst)
	})

	service.Start(configs.ServiceName, configs.Version)
}

//func main() {
//	configs.Load()
//	repository.InitializingDB()
//
//	service.Serve(configs.Variables.GrpcPort, func(lst net.Listener) error {
//		server := grpc.NewServer()
//		service2.NewService().RegisterServer(server)
//		functions.NewService().RegisterServer(server)
//		return server.Serve(lst)
//	})
//
//	service.Serve(configs.Variables.HttpPort, func(lst net.Listener) error {
//		router := mux.NewRouter()
//		return http.Serve(lst, httpext.DefaultCors.Handler(router))
//	})
//
//	service.Start(configs.Variables.ServiceName, configs.Variables.Version)
//}

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
