package calculator

import (
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	"golang.org/x/net/context"
	"time"
)

type Indicator interface {
	Calculate(ctx context.Context, candles []Candle) (*indicatorsAPI.IndicatorValues, error)
	UpdateLast(ctx context.Context, candles Candle) *indicatorsAPI.IndicatorValue
	GetMarket() *chipmunkAPI.Market
	GetResolution() *chipmunkAPI.Resolution
	GetId() uint
}

type Candle struct {
	Time   time.Time
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}
