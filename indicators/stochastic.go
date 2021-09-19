package indicators

import (
	"errors"
	"github.com/mrNobody95/Gate/models"
	"math"
)

func (conf *Configuration) validateStochastic() error {
	if len(conf.Candles) < conf.StochasticLength {
		return errors.New("candles length must bigger or equal than indicator period length")
	}
	if conf.StochasticSmoothK >= conf.StochasticLength {
		return errors.New("smoothK parameter must be smaller than indicator period length")
	}
	if conf.StochasticSmoothD >= conf.StochasticLength {
		return errors.New("smoothD parameter must be smaller than indicator period length")
	}
	return nil
}

/*
 path then candles to the function for calculating stochastic value. if appendCandles false it calculate
 stochastic value for candles[conf.Length:] (candles from 0 to conf.length-1 stochastic values not defined
 and have zero value) otherwise is calculate the stochastic for all new input candles.
 finally stochastic values save to the conf.candles[index].stochastic parameter or input candles.
*/
func (conf *Configuration) CalculateStochastic() error {
	if err := conf.validateStochastic(); err != nil {
		return err
	}

	for i := conf.StochasticLength - 1; i < len(conf.Candles); i++ {
		lowest := math.MaxFloat64
		highest := float64(0)
		for j := i - (conf.StochasticLength - 1); j < conf.StochasticLength; j++ {
			if conf.Candles[j].Low < lowest {
				lowest = conf.Candles[j].Low
			}
			if conf.Candles[j].High > highest {
				highest = conf.Candles[j].High
			}
		}
		indicatorLock.Lock()
		conf.Candles[i].Stochastic.FastK = 100 * ((conf.Candles[i].Close - lowest) / (highest - lowest))
		indicatorLock.Unlock()

		sum := float64(0)
		for j := i - conf.StochasticSmoothK + 1; j <= i; j++ {
			sum += conf.Candles[j].FastK
		}
		indicatorLock.Lock()
		conf.Candles[i].Stochastic.IndexK = sum / float64(conf.StochasticSmoothK)
		conf.Candles[i].Stochastic.IndexD = calculateIndexD(conf.Candles[i-conf.StochasticSmoothD+1 : i+1])
		indicatorLock.Unlock()
	}
	return nil
}

func (conf *Configuration) UpdateStochastic() {
	lowest := math.MaxFloat64
	highest := float64(0)
	lastIndex := len(conf.Candles) - 1
	for i := len(conf.Candles) - conf.StochasticLength; i < len(conf.Candles); i++ {
		if conf.Candles[i].Low < lowest {
			lowest = conf.Candles[i].Low
		}
		if conf.Candles[i].High > highest {
			highest = conf.Candles[i].High
		}
	}
	indicatorLock.Lock()
	conf.Candles[lastIndex].FastK = 100 * ((conf.Candles[lastIndex].Close - lowest) / (highest - lowest))
	indicatorLock.Unlock()

	sum := float64(0)
	for j := len(conf.Candles) - conf.StochasticSmoothK; j < len(conf.Candles); j++ {
		sum += conf.Candles[j].FastK
	}
	indicatorLock.Lock()
	conf.Candles[lastIndex].IndexK = sum / float64(conf.StochasticSmoothK)
	conf.Candles[lastIndex].IndexD = calculateIndexD(conf.Candles[lastIndex-conf.StochasticSmoothD+1:])
	indicatorLock.Unlock()
}

func calculateIndexD(candles []models.Candle) float64 {
	sum := float64(0)
	for _, candle := range candles {
		sum += candle.Stochastic.IndexK
	}
	return sum / float64(len(candles))
}
