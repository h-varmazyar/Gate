package indicators

import (
	"errors"
	"github.com/mrNobody95/Gate/models"
)

func (conf *IndicatorConfig) validateStochastic() error {
	if len(conf.Candles) < conf.Length {
		return errors.New("candles length must bigger or equal than indicator period length")
	}
	if conf.SmoothD >= conf.Length {
		return errors.New("smoothD parameter must be smaller than indicator period length")
	}
	return nil
}

// path then candles to the function for calculating stochastic value. if appendCandles false it calculate
// stochastic value for candles[conf.Length:] (candles from 0 to conf.length-1 stochastic values not defined
// and have zero value) otherwise is calculate the stochastic for all new input candles.
// finally stochastic values save to the conf.candles[index].stochastic parameter or input candles.
func (conf *IndicatorConfig) CalculateStochastic(candles []models.Candle, appendCandles bool) error {
	rangeCounter := conf.Length - 1
	if appendCandles {
		rangeCounter = len(conf.Candles)
		conf.Candles = append(conf.Candles, candles...)
	} else {
		conf.Candles = candles
	}
	if err := conf.validateStochastic(); err != nil {
		return err
	}

	lowest := float64(0)
	highest := float64(0)
	//todo: check indices
	for i, candle := range conf.Candles[rangeCounter:] {
		for _, innerCandle := range conf.Candles[rangeCounter-conf.Length+i+1 : rangeCounter+i+1] {
			if innerCandle.Low < lowest {
				lowest = innerCandle.Low
			}
			if innerCandle.High > highest {
				highest = innerCandle.High
			}
		}
		counter := conf.Length + i - 1
		candle.Stochastic.IndexK = 100 * ((candle.Close - lowest) / (highest - lowest))
		candle.Stochastic.IndexD = calculateIndexD(conf.Candles[counter-conf.SmoothD : counter+1])
	}
	return nil
}

func (conf *IndicatorConfig) UpdateStochastic() error {
	lowest := float64(0)
	highest := float64(0)
	lastIndex := len(conf.Candles) - 1
	for _, innerCandle := range conf.Candles[lastIndex-conf.Length+1:] {
		if innerCandle.Low < lowest {
			lowest = innerCandle.Low
		}
		if innerCandle.High > highest {
			highest = innerCandle.High
		}
	}
	conf.Candles[lastIndex].Stochastic.IndexK = 100 * ((conf.Candles[lastIndex].Close - lowest) / (highest - lowest))
	conf.Candles[lastIndex].Stochastic.IndexD = calculateIndexD(conf.Candles[lastIndex-conf.SmoothD+1:])
	return nil
}

func calculateIndexD(candles []models.Candle) float64 {
	sum := float64(0)
	for _, candle := range candles {
		sum += candle.Stochastic.IndexK
	}
	return sum / float64(len(candles))
}
