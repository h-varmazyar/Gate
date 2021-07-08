package indicators

import (
	"github.com/mrNobody95/Gate/models"
	"sync"
)

type IndicatorConfig struct {
	Candles []models.Candle
	Length  int
	//stochastic
	SmoothD int
	SmoothK int
	//parabolic sar
	acceleration       float64
	maxAcceleration    float64
	accelerationFactor float64
	startAcceleration  float64
	extremePoint       float64
	trend              Trend
	//moving average
	source Source
	//macd
	fastLength   int
	slowLength   int
	signalLength int
	//bollinger band
	Deviation int
}

var indicatorLock sync.Mutex

func cloneCandles(candles []models.Candle) []models.Candle {
	return append([]models.Candle{}, candles...)
}
