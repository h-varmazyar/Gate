package indicators

import (
	"github.com/mrNobody95/Gate/models"
	"sync"
)

type Source string

const (
	SourceOpen  = "open"
	SourceClose = "close"
	SourceHigh  = "high"
	SourceLow   = "low"
	SourceHL2   = "hl2"
	SourceHLC3  = "hlc3"
	SourceOHLC4 = "ohlc4"
)

type Configuration struct {
	Candles []models.Candle
	//moving average
	MovingAverageLength int
	MovingAverageSource Source
	//bollinger band
	BollingerLength    int
	BollingerDeviation int
	//macd
	MacdFastLength   int
	MacdSlowLength   int
	MacdSignalLength int

	//Length  int
	//stochastic
	StochasticSmoothD int
	StochasticSmoothK int
	//parabolic sar
	acceleration       float64
	maxAcceleration    float64
	accelerationFactor float64
	startAcceleration  float64
	extremePoint       float64
	trend              models.Trend
	//moving average
	//source Source
}

var indicatorLock sync.Mutex

func cloneCandles(candles []models.Candle) []models.Candle {
	return append([]models.Candle{}, candles...)
}
