package ohlc

import (
	"context"
	"encoding/json"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/indicators"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
	log "github.com/sirupsen/logrus"
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
	settings := new(WorkerSettings)
	if req.Market == nil {
		return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "market is nil")
	}
	brokerage, err := s.brokerageService.Get(ctx, &brokerageApi.BrokerageIDReq{ID: req.BrokerageID})
	if err != nil {
		return nil, err
	}
	if settings.Indicators, err = parseIndicators(ctx, req.Market.Name, brokerage.Strategy); err != nil {
		log.WithError(err).Error("failed to parse indicators")
		return nil, err
	}
	settings.Market = req.Market
	settings.Resolution = brokerage.Resolution
	Worker.AddMarket(settings)
	return &api.Void{}, nil
}

func (s *Service) ReturnCandles(_ context.Context, req *chipmunkApi.BufferedCandlesRequest) (*api.Candles, error) {
	response := make([]*api.Candle, 0)
	limit := 1000
	for i := 0; ; i += limit {
		list, err := repository.Candles.ReturnList(req.MarketID, req.ResolutionID, limit, i)
		if err != nil {
			return nil, err
		}
		tmp := make([]*api.Candle, 0)
		mapper.Slice(list, &tmp)
		response = append(response, tmp...)
		if len(list) < limit {
			break
		}
	}
	return &api.Candles{Candles: response}, nil
}

func (s *Service) ReturnLastCandle(_ context.Context, req *chipmunkApi.BufferedCandlesRequest) (*api.Candle, error) {
	candle, err := repository.Candles.ReturnLast(req.MarketID, req.ResolutionID)
	if err != nil {
		return nil, err
	}
	response := new(api.Candle)
	mapper.Struct(candle, &response)
	return response, nil
}

func (s *Service) CancelWorker(_ context.Context, req *chipmunkApi.CancelWorkerRequest) (*api.Void, error) {
	return new(api.Void), Worker.CancelWorker(req.MarketID, req.ResolutionID)
}

func parseIndicators(ctx context.Context, market string, strategy *brokerageApi.Strategy) ([]indicators.Indicator, error) {
	response := make([]indicators.Indicator, 0)
	for _, indicator := range strategy.Indicators {
		var err error
		if indicator == nil {
			continue
		}
		switch indicator.Name {
		case brokerageApi.IndicatorNames_RSI:
			rsiSettings := &struct {
				Length int `json:"length"`
			}{}
			if err = json.Unmarshal(indicator.Configs, rsiSettings); err != nil {
				return nil, err
			}
			response = append(response, indicators.NewRSI(rsiSettings.Length, market))
		case brokerageApi.IndicatorNames_Stochastic:
			stochasticSettings := &struct {
				Length  int `json:"length"`
				SmoothK int `json:"smooth_k"`
				SmoothD int `json:"smooth_d"`
			}{}
			if err = json.Unmarshal(indicator.Configs, stochasticSettings); err != nil {
				return nil, err
			}
			response = append(response, indicators.NewStochastic(stochasticSettings.Length, stochasticSettings.SmoothK, stochasticSettings.SmoothD, market))
		case brokerageApi.IndicatorNames_MovingAverage:
			movingAverageSettings := &struct {
				Length int    `json:"length"`
				Source string `json:"source"`
			}{}
			if err = json.Unmarshal(indicator.Configs, movingAverageSettings); err != nil {
				return nil, err
			}
			response = append(response, indicators.NewMovingAverage(movingAverageSettings.Length, indicators.Source(movingAverageSettings.Source), market))
		case brokerageApi.IndicatorNames_BollingerBands:
			bollingerBandsSettings := &struct {
				Length    int    `json:"length"`
				Deviation int    `json:"deviation"`
				Source    string `json:"source"`
			}{}
			if err = json.Unmarshal(indicator.Configs, bollingerBandsSettings); err != nil {
				return nil, err
			}
			response = append(response, indicators.NewBollingerBands(bollingerBandsSettings.Length, bollingerBandsSettings.Deviation,
				indicators.Source(bollingerBandsSettings.Source), market))
		default:
			return nil, errors.NewWithSlug(ctx, codes.NotFound, "no indicators found")
		}
	}
	return response, nil
}
