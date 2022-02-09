package indicators

import (
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/repository"
	"sync"
)

type Configuration struct {
	//moving average
	//bollinger band

	//macd
	MacdSignalLength int    `yaml:"MacdSignalLength"`
	MacdFastLength   int    `yaml:"MacdFastLength"`
	MacdSlowLength   int    `yaml:"MacdSlowLength"`
	MacdSource       Source `yaml:"MacdSource"`
	From             int
	//RSI

	//stochastic

	//ADX/ATR
	AdxAtrLength int `yaml:"AdxAtrLength"`
	//parabolic sar
	maxAcceleration float64 `yaml:"maxAcceleration"`
	Acceleration    float64 `yaml:"Acceleration"`
	af              float64
	xp              float64
	lock            *sync.Mutex
}

func cloneCandles(input []*repository.Candle) []*repository.Candle {
	var cloned []*repository.Candle
	for _, candle := range input {
		c := *candle
		cloned = append(cloned, &c)
	}
	return cloned
}
