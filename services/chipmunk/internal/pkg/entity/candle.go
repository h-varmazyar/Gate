package entity

import (
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"time"
)

//DROP INDEX idx_candles_market_id;
//DROP INDEX idx_candles_resolution_id;
//
//CREATE INDEX idx_candles_market_resolution_time ON public.candles (market_id, resolution_id, time DESC);

type Candle struct {
	gormext.UniversalModel
	Time            time.Time
	Open            float64
	High            float64
	Low             float64
	Close           float64
	Volume          float64
	Amount          float64
	MarketID        uuid.UUID `gorm:"index"`
	ResolutionID    uuid.UUID `gorm:"index"`
	IndicatorValues `gorm:"-"`
}

type IndicatorValues struct {
	BollingerBands map[uuid.UUID]*BollingerBandsValue
	MovingAverages map[uuid.UUID]*MovingAverageValue
	Stochastics    map[uuid.UUID]*StochasticValue
	RSIs           map[uuid.UUID]*RSIValue
}

func NewIndicatorValues() IndicatorValues {
	return IndicatorValues{
		BollingerBands: make(map[uuid.UUID]*BollingerBandsValue),
		MovingAverages: make(map[uuid.UUID]*MovingAverageValue),
		Stochastics:    make(map[uuid.UUID]*StochasticValue),
		RSIs:           make(map[uuid.UUID]*RSIValue),
	}
}
