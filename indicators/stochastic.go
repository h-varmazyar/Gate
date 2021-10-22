package indicators

import (
	"errors"
	"github.com/mrNobody95/Gate/models"
	"math"
)

func (conf *Configuration) validateStochastic(length int) error {
	if length <= conf.StochasticLength {
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
func (conf *Configuration) CalculateStochastic(candles []models.Candle) error {
	if err := conf.validateStochastic(len(candles)); err != nil {
		return err
	}

	for i := conf.StochasticLength - 1; i < len(candles); i++ {
		lowest := math.MaxFloat64
		highest := float64(0)
		from := i - (conf.StochasticLength - 1)
		for j := from; j < from+conf.StochasticLength; j++ {
			if candles[j].Low < lowest {
				lowest = candles[j].Low
			}
			if candles[j].High > highest {
				highest = candles[j].High
			}
		}
		candles[i].Stochastic.FastK = 100 * ((candles[i].Close - lowest) / (highest - lowest))

		sum := float64(0)
		for j := i - conf.StochasticSmoothK + 1; j <= i; j++ {
			sum += candles[j].FastK
		}
		candles[i].Stochastic.IndexK = sum / float64(conf.StochasticSmoothK)
		candles[i].Stochastic.IndexD = calculateIndexD(candles[i-conf.StochasticSmoothD+1 : i+1])
	}
	return nil
}

func (conf *Configuration) UpdateStochastic(candles []models.Candle) {
	lowest := math.MaxFloat64
	highest := float64(0)
	lastIndex := len(candles) - 1
	for i := len(candles) - conf.StochasticLength; i < len(candles); i++ {
		if candles[i].Low < lowest {
			lowest = candles[i].Low
		}
		if candles[i].High > highest {
			highest = candles[i].High
		}
	}
	candles[lastIndex].FastK = 100 * ((candles[lastIndex].Close - lowest) / (highest - lowest))

	sum := float64(0)
	for j := len(candles) - conf.StochasticSmoothK; j < len(candles); j++ {
		sum += candles[j].FastK
	}
	candles[lastIndex].IndexK = sum / float64(conf.StochasticSmoothK)
	candles[lastIndex].IndexD = calculateIndexD(candles[lastIndex-conf.StochasticSmoothD+1:])
}

func calculateIndexD(candles []models.Candle) float64 {
	sum := float64(0)
	for _, candle := range candles {
		sum += candle.Stochastic.IndexK
	}
	return sum / float64(len(candles))
}
