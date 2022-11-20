package service

import (
	"context"
	"github.com/google/uuid"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/workers"
	indicators "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/indicators/service"
	resolutions "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/resolutions/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

type Service struct {
	db                repository.CandleRepository
	buffer            *buffer.CandleBuffer
	logger            *log.Logger
	configs           *Configs
	functionsService  coreApi.FunctionsServiceClient
	strategyService   eagleApi.StrategyServiceClient
	resolutionService *resolutions.Service
	indicatorService  *indicators.Service
	worker            *workers.PrimaryData
}

type Dependencies struct {
	Buffer            *buffer.CandleBuffer
	ResolutionService *resolutions.Service
	IndicatorService  *indicators.Service
	Worker            *workers.PrimaryData
}

var (
	GrpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, configs *Configs, db repository.CandleRepository, dependencies *Dependencies) *Service {
	if GrpcService == nil {
		coreConn := grpcext.NewConnection(configs.CoreAddress)
		eagleConn := grpcext.NewConnection(configs.EagleAddress)
		GrpcService = new(Service)
		GrpcService.db = db
		GrpcService.logger = logger
		GrpcService.configs = configs
		GrpcService.functionsService = coreApi.NewFunctionsServiceClient(coreConn)
		GrpcService.strategyService = eagleApi.NewStrategyServiceClient(eagleConn)
		GrpcService.indicatorService = dependencies.IndicatorService
		GrpcService.resolutionService = dependencies.ResolutionService
		GrpcService.buffer = dependencies.Buffer
		GrpcService.worker = dependencies.Worker
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	chipmunkApi.RegisterCandleServiceServer(server, s)
}

func (s *Service) List(_ context.Context, req *chipmunkApi.CandleListReq) (*chipmunkApi.Candles, error) {
	marketID, err := uuid.Parse(req.MarketID)
	if err != nil {
		return nil, err
	}
	resolutionID, err := uuid.Parse(req.ResolutionID)
	if err != nil {
		return nil, err
	}
	candles := s.buffer.ReturnCandles(marketID, resolutionID, int(req.Count))
	response := new(chipmunkApi.Candles)

	for _, candle := range candles {
		element := new(chipmunkApi.Candle)
		mapper.Struct(candle, element)
		element.IndicatorValues = make(map[string]*chipmunkApi.IndicatorValue)
		for key, value := range candle.RSIs {
			element.IndicatorValues[key.String()] = &chipmunkApi.IndicatorValue{
				Type: chipmunkApi.Indicator_RSI,
				Value: &chipmunkApi.IndicatorValue_RSI{
					RSI: &chipmunkApi.RSI{
						RSI: value.RSI,
					},
				},
			}
		}
		for key, value := range candle.Stochastics {
			element.IndicatorValues[key.String()] = &chipmunkApi.IndicatorValue{
				Type: chipmunkApi.Indicator_Stochastic,
				Value: &chipmunkApi.IndicatorValue_Stochastic{
					Stochastic: &chipmunkApi.Stochastic{
						IndexK: value.IndexK,
						IndexD: value.IndexD,
					},
				},
			}
		}
		for key, value := range candle.BollingerBands {
			element.IndicatorValues[key.String()] = &chipmunkApi.IndicatorValue{
				Type: chipmunkApi.Indicator_BollingerBands,
				Value: &chipmunkApi.IndicatorValue_BollingerBands{
					BollingerBands: &chipmunkApi.BollingerBands{
						UpperBand: value.UpperBand,
						LowerBand: value.LowerBand,
						MA:        value.MA,
					},
				},
			}
		}
		for key, value := range candle.MovingAverages {
			element.IndicatorValues[key.String()] = &chipmunkApi.IndicatorValue{
				Type: chipmunkApi.Indicator_MovingAverage,
				Value: &chipmunkApi.IndicatorValue_MovingAverage{
					MovingAverage: &chipmunkApi.MovingAverage{
						Simple:      value.Simple,
						Exponential: value.Exponential,
					},
				},
			}
		}
		response.Elements = append(response.Elements, element)
	}
	return response, nil
}

func (s *Service) Update(ctx context.Context, _ *chipmunkApi.CandleUpdateReq) (*chipmunkApi.Candle, error) {
	return nil, errors.New(ctx, codes.Unimplemented)
}

func (s *Service) BulkUpdate(ctx context.Context, req *chipmunkApi.CandleBulkUpdateReq) (*api.Void, error) {
	resolutionList, err := s.resolutionService.List(ctx, &chipmunkApi.ResolutionListReq{
		Platform: req.Platform,
	})
	if err != nil {
		s.logger.WithError(err).Errorf("failed to get resolutionList")
		return nil, err
	}
	for key, ticker := range req.Tickers {
		marketID, err := uuid.Parse(key)
		if err != nil {
			s.logger.WithError(err).Error("invalid market id in bulk update")
			return nil, err
		}
		for _, resolution := range resolutionList.Elements {
			resolutionID, err := uuid.Parse(resolution.ID)
			if err != nil {
				continue
			}
			c := &entity.Candle{
				Open:         ticker.Open,
				High:         ticker.High,
				Low:          ticker.Low,
				Close:        ticker.Close,
				Volume:       ticker.Volume,
				MarketID:     marketID,
				ResolutionID: resolutionID,
			}
			last := s.buffer.ReturnCandles(marketID, resolutionID, 1)
			if last == nil {
				continue
			}
			if last[0].Time.Add(time.Duration(resolution.Duration)).After(time.Unix(req.Date, 0)) {
				c.Time = last[0].Time
			} else {
				if last[0].Time.Add(time.Duration(resolution.Duration)).Before(time.Unix(req.Date, 0).Add(time.Duration(resolution.Duration) * -1)) {
					continue
				}
				c.Time = last[0].Time.Add(time.Duration(resolution.Duration))
			}
			if err = s.db.Save(c); err != nil {
				s.logger.WithError(err).Errorf("failed to save candle")
				continue
			}
			s.buffer.Push(c)
		}
	}
	return new(api.Void), nil
}

func (s *Service) DownloadPrimaryCandles(ctx context.Context, req *chipmunkApi.DownloadPrimaryCandlesReq) (*api.Void, error) {
	if err := s.validateDownloadPrimaryCandlesRequest(ctx, req); err != nil {
		return nil, err
	}

	strategyID, err := s.prepareDownloadPrimaryCandles(req)
	if err != nil {
		return nil, err
	}

	for _, market := range req.Markets.Elements {
		go s.preparePrimaryDataRequests(req.Platform, market, req.Resolutions, strategyID)
	}
	return new(api.Void), nil
}
