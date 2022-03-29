package indicators

import (
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/models"
	"sync"
)

type Source string

type Configuration struct {
	//moving average
	MovingAverageSource Source `yaml:"MovingAverageSource"`
	MovingAverageLength int    `yaml:"MovingAverageLength"`
	//bollinger band
	BollingerDeviation int `yaml:"BollingerDeviation"`
	BollingerLength    int `yaml:"BollingerLength"`
	//macd
	MacdSignalLength int    `yaml:"MacdSignalLength"`
	MacdFastLength   int    `yaml:"MacdFastLength"`
	MacdSlowLength   int    `yaml:"MacdSlowLength"`
	MacdSource       Source `yaml:"MacdSource"`
	From             int
	//RSI
	RsiLength int `yaml:"RsiLength"`
	//stochastic
	StochasticLength  int `yaml:"StochasticLength"`
	StochasticSmoothK int `yaml:"StochasticSmoothK"`
	StochasticSmoothD int `yaml:"StochasticSmoothD"`
	//ADX/ATR
	AdxAtrLength int `yaml:"AdxAtrLength"`
	//parabolic sar
	maxAcceleration float64 `yaml:"maxAcceleration"`
	Acceleration    float64 `yaml:"Acceleration"`
	af              float64
	xp              float64
	lock            *sync.Mutex
}

const (
	SourceCustom Source = "custom"
	SourceOHLC4  Source = "ohlc4"
	SourceClose  Source = "close"
	SourceOpen   Source = "open"
	SourceHigh   Source = "high"
	SourceHLC3   Source = "hlc3"
	SourceLow    Source = "low"
	SourceHL2    Source = "hl2"
)

func cloneCandles(candles []*models.Candle) []*models.Candle {
	var cloned []*models.Candle
	for _, candle := range candles {
		c := *candle
		cloned = append(cloned, &c)
	}
	return cloned
}
