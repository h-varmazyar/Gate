package indicators

import (
	"github.com/mrNobody95/Gate/models"
	"sync"
)

type Source string

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
	//RSI
	RsiLength int
	//stochastic
	StochasticLengthK int
	StochasticLengthD int
	//ADX/ATR
	AdxAtrLength int
	//parabolic sar
	maxAcceleration float64
	Acceleration    float64
	af              float64
	xp              float64
}

const (
	SourceOpen  = "open"
	SourceClose = "close"
	SourceHigh  = "high"
	SourceLow   = "low"
	SourceHL2   = "hl2"
	SourceHLC3  = "hlc3"
	SourceOHLC4 = "ohlc4"
)

var indicatorLock sync.Mutex

func DefaultConfig() *Configuration {
	return &Configuration{}
}

func cloneCandles(candles []models.Candle) []models.Candle {
	return append([]models.Candle{}, candles...)
}
