package indicators

import (
	"errors"
	"github.com/mrNobody95/Gate/models"
)

type stochasticConfig struct {
	*basicConfig
	Length  int
	SmoothD int
	SmoothK int
}

type Stochastic struct {
	IndexK float64
	IndexD float64
	Candle *models.Candle
}

func NewStochasticConfig(candles []models.Candle, length, smoothD, smoothK int) (*stochasticConfig, error) {
	conf := stochasticConfig{
		basicConfig: &basicConfig{Candles: candles},
		Length:      length,
		SmoothD:     smoothD,
		SmoothK:     smoothK,
	}

	return conf.validate()
}

func (conf *stochasticConfig) validate() (*stochasticConfig, error) {
	if len(conf.Candles) < conf.Length {
		return nil, errors.New("candles length must bigger or equal than indicator period length")
	}
	if conf.SmoothD >= conf.Length {
		return nil, errors.New("smoothD parameter must be smaller than indicator period length")
	}
	return conf, nil
}

func (conf *stochasticConfig) Calculate() []Stochastic {
	lowest := float64(0)
	highest := float64(0)
	stochastics := make([]Stochastic, len(conf.Candles))
	//todo: check indices
	for i, candle := range conf.Candles[conf.Length-1:] {
		for _, innerCandle := range conf.Candles[i : conf.Length+i] {
			if innerCandle.Low < lowest {
				lowest = innerCandle.Low
			}
			if innerCandle.High > highest {
				highest = innerCandle.High
			}
		}
		counter := conf.Length + i - 1
		stochastics[counter].Candle = &candle
		stochastics[counter].IndexK = 100 * ((candle.Close - lowest) / (highest - lowest))
		stochastics[counter].IndexD = calculateIndexD(stochastics[counter-conf.SmoothD : counter+1])
	}
	return nil
}

func calculateIndexD(st []Stochastic) float64 {
	sum := float64(0)
	for _, s := range st {
		sum += s.IndexK
	}
	return sum / float64(len(st))
}
