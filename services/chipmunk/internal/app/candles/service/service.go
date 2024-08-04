package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Service struct {
	db               repository.CandleRepository
	logger           *log.Logger
	configs          *Configs
	functionsService coreApi.FunctionsServiceClient
	strategyService  eagleApi.StrategyServiceClient
}

var (
	GrpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, configs *Configs, db repository.CandleRepository) *Service {
	if GrpcService == nil {
		coreConn := grpcext.NewConnection(configs.CoreAddress)
		eagleConn := grpcext.NewConnection(configs.EagleAddress)
		GrpcService = &Service{
			db:               db,
			logger:           logger,
			configs:          configs,
			functionsService: coreApi.NewFunctionsServiceClient(coreConn),
			strategyService:  eagleApi.NewStrategyServiceClient(eagleConn),
		}
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	chipmunkApi.RegisterCandleServiceServer(server, s)
}

func (s *Service) List(_ context.Context, req *chipmunkApi.CandleListReq) (*chipmunkApi.Candles, error) {
	_, err := uuid.Parse(req.MarketID)
	if err != nil {
		return nil, err
	}
	_, err = uuid.Parse(req.ResolutionID)
	if err != nil {
		return nil, err
	}
	candles := buffer.CandleBuffer.ReturnCandles(req.MarketID, req.ResolutionID, int(req.Count))
	response := new(chipmunkApi.Candles)

	for _, candle := range candles {
		element := new(chipmunkApi.Candle)
		mapper.Struct(candle, element)
		//element.IndicatorValues = make(map[string]*chipmunkApi.IndicatorValue)
		//for key, value := range candle.RSIs {
		//	element.IndicatorValues[key.String()] = &chipmunkApi.IndicatorValue{
		//		Type: chipmunkApi.Indicator_RSI,
		//		Value: &chipmunkApi.IndicatorValue_RSI{
		//			RSI: &chipmunkApi.RSI{
		//				RSI: value.RSI,
		//			},
		//		},
		//	}
		//}
		//for key, value := range candle.Stochastics {
		//	element.IndicatorValues[key.String()] = &chipmunkApi.IndicatorValue{
		//		Type: chipmunkApi.Indicator_Stochastic,
		//		Value: &chipmunkApi.IndicatorValue_Stochastic{
		//			Stochastic: &chipmunkApi.Stochastic{
		//				IndexK: value.IndexK,
		//				IndexD: value.IndexD,
		//			},
		//		},
		//	}
		//}
		//for key, value := range candle.BollingerBands {
		//	element.IndicatorValues[key.String()] = &chipmunkApi.IndicatorValue{
		//		Type: chipmunkApi.Indicator_BollingerBands,
		//		Value: &chipmunkApi.IndicatorValue_BollingerBands{
		//			BollingerBands: &chipmunkApi.BollingerBands{
		//				UpperBand: value.UpperBand,
		//				LowerBand: value.LowerBand,
		//				MA:        value.MA,
		//			},
		//		},
		//	}
		//}
		//for key, value := range candle.MovingAverages {
		//	element.IndicatorValues[key.String()] = &chipmunkApi.IndicatorValue{
		//		Type: chipmunkApi.Indicator_MovingAverage,
		//		Value: &chipmunkApi.IndicatorValue_MovingAverage{
		//			MovingAverage: &chipmunkApi.MovingAverage{
		//				Simple:      value.Simple,
		//				Exponential: value.Exponential,
		//			},
		//		},
		//	}
		//}
		response.Elements = append(response.Elements, element)
	}
	return response, nil
}

func (s *Service) Update(ctx context.Context, _ *chipmunkApi.CandleUpdateReq) (*chipmunkApi.Candle, error) {
	return nil, errors.New(ctx, codes.Unimplemented)
}

//
//func (s *Service) BulkUpdate(ctx context.Context, req *chipmunkApi.CandleBulkUpdateReq) (*api.Void, error) {
//	resolutionList, err := s.resolutionService.List(ctx, &chipmunkApi.ResolutionListReq{
//		Platform: req.Platform,
//	})
//
//	s.logger.Infof("update bulk")
//	if err != nil {
//		s.logger.WithError(err).Errorf("failed to get resolutionList")
//		return nil, err
//	}
//	for key, ticker := range req.Tickers {
//		marketID, err := uuid.Parse(key)
//		if err != nil {
//			s.logger.WithError(err).Error("invalid market id in bulk update")
//			return nil, err
//		}
//		for _, resolution := range resolutionList.Elements {
//			if ticker.Close == 0 {
//				continue
//			}
//			resolutionID, err := uuid.Parse(resolution.ID)
//			if err != nil {
//				continue
//			}
//
//			last := s.buffer.ReturnCandles(key, resolution.ID, 1)
//			if last == nil || len(last) == 0 {
//				continue
//			}
//
//			lastTime := last[len(last)-1].Time
//			lastAddedResolution := time.Unix(last[len(last)-1].Time.Unix(), 0).Add(time.Duration(resolution.Duration))
//			lastAdded2xResolution := time.Unix(last[len(last)-1].Time.Unix(), 0).Add(time.Duration(resolution.Duration * 2))
//			reqTime := time.Unix(req.Date, 0)
//
//			lastVol := ticker.Volume
//			for i := 0; i < len(last)-1; i++ {
//				lastVol -= last[i].Volume
//			}
//
//			c := &entities.Candle{
//				Time:         last[len(last)-1].Time,
//				Open:         last[len(last)-1].Open,
//				High:         last[len(last)-1].High,
//				Low:          last[len(last)-1].Low,
//				Close:        last[len(last)-1].Close,
//				Volume:       lastVol,
//				Amount:       0,
//				MarketID:     marketID,
//				ResolutionID: resolutionID,
//			}
//
//			if reqTime.After(lastTime) && reqTime.Before(lastAddedResolution) {
//				c.Close = ticker.Close
//				if ticker.Close < c.Low {
//					c.Low = ticker.Close
//				}
//				if ticker.Close > c.High {
//					c.High = ticker.Close
//				}
//			} else if reqTime.After(lastAddedResolution) && reqTime.Before(lastAdded2xResolution) {
//				c.Time = lastAddedResolution
//				c.Open = ticker.Close
//				c.High = ticker.Close
//				c.Low = ticker.Close
//				c.Close = ticker.Close
//			} else {
//				continue
//			}
//			if err = s.db.Save(c); err != nil {
//				s.logger.WithError(err).Errorf("failed to save candle")
//				continue
//			}
//			s.buffer.Push(c)
//		}
//	}
//	return new(api.Void), nil
//}

//func (s *Service) StartWorkers(ctx context.Context, req *chipmunkApi.CandleWorkerStartReq) (*api.Void, error) {
//	if err := s.validateDownloadPrimaryCandlesRequest(ctx, req); err != nil {
//		return nil, err
//	}
//
//	indicators, err := s.indicatorService.List(ctx, &chipmunkApi.IndicatorListReq{Type: chipmunkApi.Indicator_All})
//	if err != nil {
//		s.logger.WithError(err).Error("failed to fetch all indicatorService")
//		return nil, err
//	}
//
//	loadedIndicators, err := s.loadIndicators(indicators.Elements)
//	if err != nil {
//		return nil, err
//	}
//
//	if !s.candleReaderWorker.IsStarted() {
//		s.candleReaderWorker.Start(loadedIndicators)
//	}
//
//	go func() {
//		for _, market := range req.Markets.Elements {
//			s.preparePrimaryDataRequests(req.Platform, market, req.Resolutions, loadedIndicators)
//		}
//
//		if !s.lastCandleWorker.Started {
//			s.lastCandleWorker.Start(req.Markets.Elements, req.Resolutions.Elements, loadedIndicators)
//		}
//	}()
//
//	s.missedCandlesWorker.Start(req.Markets.Elements, req.Resolutions.Elements)
//	s.redundantRemoverWorker.Start(req.Markets.Elements, req.Resolutions.Elements)
//	return new(api.Void), nil
//}
//
//func (s *Service) StopWorkers(_ context.Context, _ *chipmunkApi.CandleWorkerStopReq) (*api.Void, error) {
//	//todo: must be implement based on platform
//	if s.candleReaderWorker.IsStarted() {
//		s.candleReaderWorker.Stop()
//	}
//
//	if s.lastCandleWorker.Started {
//		s.lastCandleWorker.Stop()
//	}
//
//	if s.missedCandlesWorker.Started {
//		s.missedCandlesWorker.Stop()
//	}
//
//	if s.redundantRemoverWorker.Started {
//		s.redundantRemoverWorker.Stop()
//	}
//	return new(api.Void), nil
//}
