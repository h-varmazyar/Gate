package ohlc

import (
	"context"
	"encoding/json"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/errors"
	"github.com/mrNobody95/Gate/pkg/grpcext"
	"github.com/mrNobody95/Gate/pkg/mapper"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	chipmunkApi "github.com/mrNobody95/Gate/services/chipmunk/api"
	"github.com/mrNobody95/Gate/services/chipmunk/configs"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/buffer"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/indicators"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/repository"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/workers"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/workers/OHLC"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Service struct {
	brokerageService brokerageApi.BrokerageServiceClient
}

var (
	GrpcService *Service
)

func NewService() *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		brokerageConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)
		GrpcService.brokerageService = brokerageApi.NewBrokerageServiceClient(brokerageConn)
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	chipmunkApi.RegisterOhlcServiceServer(server, s)
}

func (s *Service) AddMarket(ctx context.Context, req *chipmunkApi.AddMarketRequest) (*api.Void, error) {
	settings := new(OHLC.WorkerSettings)
	if req.Market == nil {
		return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "market is nil")
	}
	brokerage, err := s.brokerageService.Get(ctx, &brokerageApi.BrokerageIDReq{ID: req.BrokerageID})
	if err != nil {
		return nil, err
	}
	strategyIndicators, err := parseIndicators(ctx, req.Market.Name, brokerage.Strategy)
	if err != nil {
		return nil, err
	}
	settings.Market = req.Market
	settings.Resolution = brokerage.Resolution
	settings.Indicators = strategyIndicators
	workers.OHLCWorker.AddMarket(settings)
	return &api.Void{}, nil
}

func (s *Service) ReturnCandles(_ context.Context, req *chipmunkApi.BufferedCandlesRequest) (*api.Candles, error) {
	response := make([]*api.Candle, 0)
	for i := 0; ; i += 1000 {
		list, err := repository.Candles.ReturnList(req.MarketID, req.ResolutionID, i)
		if err != nil {
			return nil, err
		}
		tmp := make([]*api.Candle, 0)
		mapper.Slice(list, &tmp)
		response = append(response, tmp...)
		if len(list) < 1000 {
			break
		}
	}
	return &api.Candles{Candles: response}, nil
}

func (s *Service) ReturnLastCandle(_ context.Context, req *chipmunkApi.BufferedCandlesRequest) (*api.Candle, error) {
	candle := buffer.Candles.Last(req.MarketID, req.ResolutionID)
	response := new(api.Candle)
	mapper.Struct(candle, &response)
	return response, nil
}

func (s *Service) CancelWorker(_ context.Context, req *chipmunkApi.CancelWorkerRequest) (*api.Void, error) {
	return new(api.Void), workers.OHLCWorker.CancelWorker(req.MarketID, req.ResolutionID)
}

func parseIndicators(ctx context.Context, market string, strategy *brokerageApi.Strategy) ([]indicators.Indicator, error) {
	response := make([]indicators.Indicator, 0)
	for _, indicator := range strategy.Indicators {
		var err error
		switch indicator.Name {
		case brokerageApi.IndicatorNames_RSI:
			tmp := &struct {
				Length int `json:"length"`
			}{}
			if err = json.Unmarshal(indicator.Configs, tmp); err != nil {
				return nil, err
			}
			response = append(response, indicators.NewRSI(tmp.Length, market))
		case brokerageApi.IndicatorNames_Stochastic:
			tmp := &struct {
				Length  int `json:"length"`
				SmoothK int `json:"smooth_k"`
				SmoothD int `json:"smooth_d"`
			}{}
			if err = json.Unmarshal(indicator.Configs, tmp); err != nil {
				return nil, err
			}
			response = append(response, indicators.NewStochastic(tmp.Length, tmp.SmoothK, tmp.SmoothD, market))
		case brokerageApi.IndicatorNames_MovingAverage:
		case brokerageApi.IndicatorNames_BollingerBands:
		default:
			return nil, errors.NewWithSlug(ctx, codes.NotFound, "no indicators found")
		}
	}
}
