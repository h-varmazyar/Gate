package indicators

import (
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/models/todo"
	"sync"
)

type Configuration struct {
	//moving average
	MovingAverageLength int
	MovingAverageSource Source

	Candles []models.Candle
	Length  int
	//stochastic
	StochasticSmoothD int
	StochasticSmoothK int
	//parabolic sar
	acceleration       float64
	maxAcceleration    float64
	accelerationFactor float64
	startAcceleration  float64
	extremePoint       float64
	trend              todo.Trend
	//moving average
	source Source
	//macd
	MacdFastLength   int
	MacdSlowLength   int
	MacdSignalLength int
	//bollinger band
	BollingerDeviation int
}

var indicatorLock sync.Mutex

func cloneCandles(candles []models.Candle) []models.Candle {
	return append([]models.Candle{}, candles...)
}
