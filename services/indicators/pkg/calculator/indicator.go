package calculator

import (
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	"golang.org/x/net/context"
)

type Indicator interface {
	Calculate(ctx context.Context, candles []*chipmunkAPI.Candle) (*indicatorsAPI.IndicatorValues, error)
	UpdateLast(ctx context.Context, candles *chipmunkAPI.Candle) *indicatorsAPI.IndicatorValue
	GetMarket() *chipmunkAPI.Market
	GetResolution() *chipmunkAPI.Resolution
	GetId() uint
}
