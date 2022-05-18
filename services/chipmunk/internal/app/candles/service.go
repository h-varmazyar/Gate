package candles

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api"
	"google.golang.org/grpc"
)

type Service struct {
	brokerageService brokerageApi.BrokerageServiceClient
	marketService    chipmunkApi.MarketServiceClient
	strategyService  eagleApi.StrategyServiceClient
}

var (
	GrpcService *Service
)

func NewService() *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		brokerageConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)
		chipmunkConn := grpcext.NewConnection(fmt.Sprintf(":%v", configs.Variables.GrpcPort))
		eagleConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Eagle)
		GrpcService.brokerageService = brokerageApi.NewBrokerageServiceClient(brokerageConn)
		GrpcService.marketService = chipmunkApi.NewMarketServiceClient(chipmunkConn)
		GrpcService.strategyService = eagleApi.NewStrategyServiceClient(eagleConn)
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	chipmunkApi.RegisterCandleServiceServer(server, s)
}

//func (s *Service) AddMarket(ctx context.Context, req *chipmunkApi.AddMarketRequest) (*api.Void, error) {
//	settings := new(WorkerSettings)
//	var (
//		err       error
//		brokerage *brokerageApi.Brokerage
//		strategy  *eagleApi.Strategy
//	)
//
//	settings.Market, err = s.marketService.Return(ctx, &chipmunkApi.ReturnMarketRequest{
//		ID: req.MarketID,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	brokerage, err = s.brokerageService.Return(ctx, &brokerageApi.ReturnBrokerageReq{ID: req.BrokerageID})
//	if err != nil {
//		return nil, err
//	}
//
//	strategy, err = s.strategyService.Return(ctx, &eagleApi.ReturnStrategyReq{
//		ID: brokerage.StrategyID,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	if settings.Indicators, err = loadIndicators(ctx, strategy); err != nil {
//		log.WithError(err).Error("failed to parse indicators")
//		return nil, err
//	}
//	settings.Resolution = brokerage.Resolution
//	Worker.AddMarket(settings)
//	return &api.Void{}, nil
//}

func (s *Service) ReturnLastNCandles(_ context.Context, req *chipmunkApi.BufferedCandlesRequest) (*chipmunkApi.Candles, error) {
	marketID, err := uuid.Parse(req.MarketID)
	if err != nil {
		return nil, err
	}
	candles := buffer.Markets.GetLastNCandles(marketID, int(req.Count))
	response := new(chipmunkApi.Candles)
	mapper.Slice(candles, &response.Elements)
	return response, nil
}

//func (s *Service) DeleteMarket(_ context.Context, req *chipmunkApi.DeleteMarketRequest) (*api.Void, error) {
//	marketID, err := uuid.Parse(req.MarketID)
//	if err != nil {
//		return nil, err
//	}
//	return new(api.Void), Worker.DeleteMarket(marketID)
//}
