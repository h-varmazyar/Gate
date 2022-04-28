package indicators

import (
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
)

type Indicator interface {
	Calculate([]*repository.Candle) error
	Update([]*repository.Candle) *repository.IndicatorValue
	GetType() chipmunkApi.IndicatorType
	GetLength() int
}

//type basicConfig struct {
//	MarketName string
//	id         uuid.UUID
//	length     int
//}
//
//type IndicatorType string
//type Source string
//
//const (
//	RSI            IndicatorType = "rsi"
//	Stochastic     IndicatorType = "stochastic"
//	BollingerBands IndicatorType = "bollingerBand"
//	MovingAverage  IndicatorType = "movingAverage"
//)
//
//const (
//	SourceCustom Source = "custom"
//	SourceOHLC4  Source = "ohlc4"
//	SourceClose  Source = "close"
//	SourceOpen   Source = "open"
//	SourceHigh   Source = "high"
//	SourceHLC3   Source = "hlc3"
//	SourceLow    Source = "low"
//	SourceHL2    Source = "hl2"
//)
