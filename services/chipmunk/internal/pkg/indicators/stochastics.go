package indicators

import (
	"errors"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
	"math"
)

type stochastic struct {
	basicConfig
	SmoothK int
	SmoothD int
}

func NewStochastic(length, smoothK, smoothD int, marketName string) *stochastic {
	return &stochastic{
		basicConfig: basicConfig{
			MarketName: marketName,
			id:         uuid.New(),
			length:     length,
		},
		SmoothK: smoothK,
		SmoothD: smoothD,
	}
}

func (conf *stochastic) GetID() uuid.UUID {
	return conf.id
}

func (conf *stochastic) GetType() IndicatorType {
	return Stochastic
}

func (conf *stochastic) GetLength() int {
	return conf.length
}

func (conf *stochastic) validateStochastic(length int) error {
	if length <= conf.length {
		return errors.New("candles length must bigger or equal than indicator period length")
	}
	if conf.SmoothK >= conf.length {
		return errors.New("smoothK parameter must be smaller than indicator period length")
	}
	if conf.SmoothD >= conf.length {
		return errors.New("smoothD parameter must be smaller than indicator period length")
	}
	return nil
}

func (conf *stochastic) Calculate(candles []*repository.Candle) error {
	if err := conf.validateStochastic(len(candles)); err != nil {
		return err
	}
	for i := conf.length - 1; i < len(candles); i++ {
		lowest := math.MaxFloat64
		highest := float64(0)
		from := i - (conf.length - 1)
		for j := from; j < from+conf.length; j++ {
			if candles[j].Low < lowest {
				lowest = candles[j].Low
			}
			if candles[j].High > highest {
				highest = candles[j].High
			}
		}
		var stochasticValue repository.StochasticValue
		{
		}
		stochasticValue.FastK = 100 * ((candles[i].Close - lowest) / (highest - lowest))

		sum := stochasticValue.FastK
		for j := i - conf.SmoothK + 1; j < i; j++ {
			sum += candles[j].Stochastics[conf.id].FastK
		}
		stochasticValue.IndexK = sum / float64(conf.SmoothK)
		stochasticValue.IndexD = calculateIndexD(conf.id, candles[i-conf.SmoothD+1:i+1])

		candles[i].Stochastics[conf.id] = stochasticValue
	}
	return nil
}

func (conf *stochastic) Update(candles []*repository.Candle) *repository.IndicatorValue {
	lowest := math.MaxFloat64
	highest := float64(0)
	lastIndex := len(candles) - 1
	for i := len(candles) - conf.length; i < len(candles); i++ {
		if candles[i].Low < lowest {
			lowest = candles[i].Low
		}
		if candles[i].High > highest {
			highest = candles[i].High
		}
	}
	fastK := 100 * ((candles[lastIndex].Close - lowest) / (highest - lowest))

	sum := fastK
	for j := len(candles) - conf.SmoothK; j < len(candles)-1; j++ {
		sum += candles[j].Stochastics[conf.id].FastK
	}

	return &repository.IndicatorValue{
		Stochastic: &repository.StochasticValue{
			IndexK: sum / float64(conf.SmoothK),
			IndexD: calculateIndexD(conf.id, candles[lastIndex-conf.SmoothD+1:]),
			FastK:  fastK,
		},
	}
}

func calculateIndexD(id uuid.UUID, candles []*repository.Candle) float64 {
	sum := float64(0)
	for _, candle := range candles {
		sum += candle.Stochastics[id].IndexK
	}
	return sum / float64(len(candles))
}
