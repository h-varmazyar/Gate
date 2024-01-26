package calculator

import (
	"github.com/h-varmazyar/Gate/pkg/errors"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/entity"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
)

func NewIndicator(ctx context.Context, entityIndicator *entity.Indicator, market *chipmunkAPI.Market, resolution *chipmunkAPI.Resolution) (Indicator, error) {
	var (
		indicator Indicator
		err       error
	)

	switch entityIndicator.Type {
	case indicatorsAPI.Type_RSI:
		indicator, err = NewRSI(entityIndicator.ID, entityIndicator.Configs.RSI, market, resolution)
	case indicatorsAPI.Type_STOCHASTIC:
		//indicator, err = NewStochastic(entityIndicator.Configs.Stochastic, market, resolution)
	case indicatorsAPI.Type_SMA:
		indicator, err = NewSMA(entityIndicator.ID, entityIndicator.Configs.SMA, market, resolution)
	case indicatorsAPI.Type_EMA:
		//indicator, err = NewEMA(entityIndicator.Configs.SMA, market, resolution)
	case indicatorsAPI.Type_BOLLINGER_BANDS:
		indicator, err = NewBollingerBands(entityIndicator.ID, entityIndicator.Configs.BollingerBands, market, resolution)
	default:
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "invalid_indicator_type")
	}
	if err != nil {
		return nil, err
	}
	return indicator, nil
}

func cloneCandle(from *chipmunkAPI.Candle) *chipmunkAPI.Candle {
	if from == nil {
		return nil
	}

	return &chipmunkAPI.Candle{
		UpdatedAt:    from.UpdatedAt,
		CreatedAt:    from.CreatedAt,
		Volume:       from.Volume,
		Amount:       from.Amount,
		Close:        from.Close,
		Open:         from.Open,
		Time:         from.Time,
		High:         from.High,
		Low:          from.Low,
		ID:           from.ID,
		MarketID:     from.MarketID,
		ResolutionID: from.ResolutionID,
	}
}

func cloneCandles(from []*chipmunkAPI.Candle) []*chipmunkAPI.Candle {
	if from == nil {
		return nil
	}

	to := make([]*chipmunkAPI.Candle, len(from))
	for i, candle := range from {
		to[i] = cloneCandle(candle)
	}
	return to
}
