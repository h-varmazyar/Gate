package main

import (
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"github.com/h-varmazyar/Gate/pkg/service"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/eagle/configs"
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/signals"
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/strategies"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/repository"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"time"
)

func main() {
	//initializing
	configs.LoadVariables()
	repository.InitializingDB()

	////testStrategy()
	//restReturn()
	//return

	registerGrpcServer()

	service.Start(configs.Variables.ServiceName, configs.Variables.Version)
}

func restReturn() {
	if st, err := repository.Strategies.Return(uuid.MustParse("f59e72c5-a842-4843-a769-216c6e8b6caf")); err != nil {
		log.WithError(err)
		return
	} else {
		log.Infof("st: %v", st)
		for _, indicator := range st.Indicators {
			log.Infof("ind: %v", indicator)
		}
	}
}

func testStrategy() {
	id := uuid.New()
	strategy := &repository.Strategy{
		UniversalModel: gormext.UniversalModel{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:                  "test",
		Description:           "test indicators",
		MinDailyProfitRate:    0,
		MinProfitPerTradeRate: 0,
		MaxFundPerTrade:       0,
		MaxFundPerTradeRate:   0,
		WorkingResolutionID:   uuid.MustParse("ab28acd0-3517-483f-b3a1-7bd879fa85d0"),
		Indicators: []*repository.StrategyIndicator{
			{
				StrategyID:  id,
				IndicatorID: uuid.MustParse("f93e0959-e55b-4afe-b233-8b30a244cdc8"),
				Type:        chipmunkApi.Indicator_RSI,
			},
		},
	}

	repository.Strategies.Save(strategy)
}

func registerGrpcServer() {
	service.Serve(configs.Variables.GrpcPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		strategies.NewService().RegisterServer(server)
		signals.NewService().RegisterServer(server)
		return server.Serve(lst)
	})
}
