package candle

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/indicators"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Service struct {
	brokerageService brokerageApi.BrokerageServiceClient
	strategyService  brokerageApi.StrategyServiceClient
}

var (
	GrpcService *Service
)

func NewService() *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		brokerageConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)
		eagleConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Eagle)
		GrpcService.brokerageService = brokerageApi.NewBrokerageServiceClient(brokerageConn)
		GrpcService.strategyService = brokerageApi.NewStrategyServiceClient(eagleConn)
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	chipmunkApi.RegisterCandleServiceServer(server, s)
}

func (s *Service) AddMarket(ctx context.Context, req *chipmunkApi.AddMarketRequest) (*api.Void, error) {
	settings := new(WorkerSettings)
	if req.Market == nil {
		return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "market is nil")
	}
	brokerage, err := s.brokerageService.Get(ctx, &brokerageApi.BrokerageIDReq{ID: req.BrokerageID})
	if err != nil {
		return nil, err
	}

	var strategy *brokerageApi.Strategy
	strategy, err = s.strategyService.Return(ctx, &brokerageApi.ReturnStrategyReq{
		ID: brokerage.StrategyID,
	})
	if err != nil {
		return nil, err
	}
	if settings.Indicators, err = loadIndicators(ctx, strategy); err != nil {
		log.WithError(err).Error("failed to parse indicators")
		return nil, err
	}
	settings.Market = req.Market
	settings.Resolution = brokerage.Resolution
	Worker.AddMarket(settings)
	return &api.Void{}, nil
}

func (s *Service) ReturnLastNCandles(_ context.Context, req *chipmunkApi.BufferedCandlesRequest) (*api.Candles, error) {
	marketID, err := uuid.Parse(req.MarketID)
	if err != nil {
		return nil, err
	}
	candles := buffer.Markets.GetLastNCandles(marketID, int(req.Count))
	response := new(api.Candles)
	mapper.Slice(candles, &response.Candles)
	return response, nil
}

func (s *Service) CancelWorker(_ context.Context, req *chipmunkApi.CancelWorkerRequest) (*api.Void, error) {
	marketID, err := uuid.Parse(req.MarketID)
	if err != nil {
		return nil, err
	}
	resolutionID, err := uuid.Parse(req.ResolutionID)
	if err != nil {
		return nil, err
	}
	return new(api.Void), Worker.CancelWorker(marketID, resolutionID)
}

func loadIndicators(_ context.Context, strategy *brokerageApi.Strategy) (map[uuid.UUID]indicators.Indicator, error) {
	response := make(map[uuid.UUID]indicators.Indicator)
	for _, strategyIndicator := range strategy.Indicators {
		id, err := uuid.Parse(strategyIndicator.IndicatorID)
		if err != nil {
			continue
		}
		indicator, err := repository.Indicators.Return(id)
		if err != nil {
			return nil, err
		}
		var indicatorCalculator indicators.Indicator
		switch indicator.Type {
		case chipmunkApi.IndicatorType_RSI:
			indicatorCalculator, err = indicators.NewRSI(indicator.ID, indicator.Configs.RSI)
		case chipmunkApi.IndicatorType_Stochastic:
			indicatorCalculator, err = indicators.NewStochastic(indicator.ID, indicator.Configs.Stochastic)
		case chipmunkApi.IndicatorType_MovingAverage:
			indicatorCalculator, err = indicators.NewMovingAverage(indicator.ID, indicator.Configs.MovingAverage)
		case chipmunkApi.IndicatorType_BollingerBands:
			indicatorCalculator, err = indicators.NewBollingerBands(indicator.ID, indicator.Configs.BollingerBands)
		}
		if err != nil {
			return nil, err
		}
		response[indicator.ID] = indicatorCalculator
	}
	return response, nil
}
