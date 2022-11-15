package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Service struct {
	logger *log.Logger
}

var (
	GrpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, _ *Configs, _ repository.CandleRepository) *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		GrpcService.logger = logger
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	chipmunkApi.RegisterCandleServiceServer(server, s)
}

func (s *Service) ReturnLastNCandles(_ context.Context, req *chipmunkApi.BufferedCandlesRequest) (*chipmunkApi.Candles, error) {
	marketID, err := uuid.Parse(req.MarketID)
	if err != nil {
		return nil, err
	}
	candles := buffer.Markets.GetLastNCandles(marketID, int(req.Count))
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
