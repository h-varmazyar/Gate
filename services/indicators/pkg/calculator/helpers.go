package calculator

import (
	"github.com/h-varmazyar/Gate/pkg/errors"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/entities"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
)

func NewIndicator(ctx context.Context, entityIndicator *entities.Indicator, market *chipmunkAPI.Market, resolution *chipmunkAPI.Resolution) (Indicator, error) {
	var (
		indicator Indicator
		err       error
	)

	switch entityIndicator.Type {
	case entities.IndicatorTypeRSI:
		indicator, err = NewRSI(entityIndicator.ID, entityIndicator.Configs.RSI, market, resolution)
	case entities.IndicatorTypeStochastic:
		//indicator, err = NewStochastic(entityIndicator.Configs.Stochastic, market, resolution)
	case entities.IndicatorTypeSMA:
		indicator, err = NewSMA(entityIndicator.ID, entityIndicator.Configs.SMA, market, resolution)
	case entities.IndicatorTypeEMA:
		//indicator, err = NewEMA(entityIndicator.Configs.SMA, market, resolution)
	case entities.IndicatorTypeBollingerBands:
		indicator, err = NewBollingerBands(entityIndicator.ID, entityIndicator.Configs.BollingerBands, market, resolution)
	default:
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "invalid_indicator_type")
	}
	if err != nil {
		return nil, err
	}
	return indicator, nil
}

//func cloneCandle(from *chipmunkAPI.Candle) *Candle {
//	if from == nil {
//		return nil
//	}
//
//	return Candle{
//		Volume: from.Volume,
//		Close:  from.Close,
//		Open:   from.Open,
//		Time:   time.Unix(from.Time, 0),
//		High:   from.High,
//		Low:    from.Low,
//	}
//}
//
//func cloneCandles(from []*chipmunkAPI.Candle) []*chipmunkAPI.Candle {
//	if from == nil {
//		return nil
//	}
//
//	to := make([]*chipmunkAPI.Candle, len(from))
//	for i, candle := range from {
//		to[i] = cloneCandle(candle)
//	}
//	return to
//}
