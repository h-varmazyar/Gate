package indicators

import (
	"errors"
	"github.com/google/uuid"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
	"math"
)

type stochastic struct {
	id uuid.UUID
	repository.StochasticConfigs
}

func NewStochastic(id uuid.UUID, configs *repository.StochasticConfigs) (*stochastic, error) {
	if err := validateStochasticConfigs(configs); err != nil {
		return nil, err
	}
	return &stochastic{
		id:                id,
		StochasticConfigs: *configs,
	}, nil
}

func (conf *stochastic) GetType() chipmunkApi.IndicatorType {
	return chipmunkApi.IndicatorType_Stochastic
}

func (conf *stochastic) GetLength() int {
	return conf.Length
}

func (conf *stochastic) validateStochastic(length int) error {
	if length <= conf.Length {
		return errors.New("candles length must bigger or equal than indicator period length")
	}
	if conf.SmoothK >= conf.Length {
		return errors.New("smoothK parameter must be smaller than indicator period length")
	}
	if conf.SmoothD >= conf.Length {
		return errors.New("smoothD parameter must be smaller than indicator period length")
	}
	return nil
}

func (conf *stochastic) Calculate(candles []*repository.Candle) error {
	if err := conf.validateStochastic(len(candles)); err != nil {
		return err
	}
	for _, candle := range candles {
		if candle.Stochastics[conf.id] == nil {
			candle.Stochastics[conf.id] = new(repository.StochasticValue)
		}
	}
	for i := conf.Length - 1; i < len(candles); i++ {
		lowest := math.MaxFloat64
		highest := float64(0)
		for j := i - (conf.Length - 1); j <= i; j++ {
			if candles[j].Low < lowest {
				lowest = candles[j].Low
			}
			if candles[j].High > highest {
				highest = candles[j].High
			}
		}
		stochasticValue := new(repository.StochasticValue)
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
	for i := len(candles) - conf.Length; i < len(candles); i++ {
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

func validateStochasticConfigs(indicator *repository.StochasticConfigs) error {
	return nil
}
