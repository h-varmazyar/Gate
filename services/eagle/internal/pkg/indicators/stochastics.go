package indicators

import (
	"errors"
	"github.com/mrNobody95/Gate/services/eagle/internal/pkg/models"
	"math"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 03.12.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

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

func (conf *Configuration) CalculateStochastic(candles []*models.Candle) error {
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

func (conf *Configuration) UpdateStochastic(candles []*models.Candle) {
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
	fastK := 100 * ((candles[lastIndex].Close - lowest) / (highest - lowest))

	sum := fastK
	for j := len(candles) - conf.StochasticSmoothK; j < len(candles)-1; j++ {
		sum += candles[j].FastK
	}

	conf.lock.Lock()
	candles[lastIndex].FastK = fastK
	candles[lastIndex].IndexK = sum / float64(conf.StochasticSmoothK)
	candles[lastIndex].IndexD = calculateIndexD(candles[lastIndex-conf.StochasticSmoothD+1:])
	conf.lock.Unlock()
}

func calculateIndexD(candles []*models.Candle) float64 {
	sum := float64(0)
	for _, candle := range candles {
		sum += candle.Stochastic.IndexK
	}
	return sum / float64(len(candles))
}
