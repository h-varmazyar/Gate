package entity

import (
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"time"
)

type Candle struct {
	gormext.UniversalModel
	Time            time.Time
	Open            float64
	High            float64
	Low             float64
	Close           float64
	Volume          float64
	Amount          float64
	MarketID        uuid.UUID
	ResolutionID    uuid.UUID
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
