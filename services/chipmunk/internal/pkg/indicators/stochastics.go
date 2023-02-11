package indicators

import (
	"errors"
	"github.com/google/uuid"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	"math"
)

type stochastic struct {
	id uuid.UUID
	entity.StochasticConfigs
}

func NewStochastic(id uuid.UUID, configs *entity.StochasticConfigs) (*stochastic, error) {
	if err := validateStochasticConfigs(configs); err != nil {
		return nil, err
	}
	return &stochastic{
		id:                id,
		StochasticConfigs: *configs,
	}, nil
}

func (conf *stochastic) GetType() chipmunkApi.IndicatorType {
	return chipmunkApi.Indicator_Stochastic
}

func (conf *stochastic) GetLength() int {
	return conf.Length
}

func (conf *stochastic) validateStochastic(length int) error {
	if length <= conf.Length {
		return errors.New("candles length must bigger or equal than indicators period length")
	}
	if conf.SmoothK >= conf.Length {
		return errors.New("smoothK parameter must be smaller than indicators period length")
	}
	if conf.SmoothD >= conf.Length {
		return errors.New("smoothD parameter must be smaller than indicators period length")
	}
	return nil
}

func (conf *stochastic) Calculate(candles []*entity.Candle) error {
	if err := conf.validateStochastic(len(candles)); err != nil {
		return err
	}
	for _, candle := range candles {
		if candle.Stochastics[conf.id] == nil {
			candle.Stochastics[conf.id] = new(entity.StochasticValue)
		}
	}

	conf.calculateStochastic(candles)

	return nil
}

func (conf *stochastic) Update(candles []*entity.Candle) {
	first := candles[0]
	start := buffer.CandleBuffer.Before(first.MarketID.String(), first.ResolutionID.String(), first.Time, conf.Length)

	internalCandles := append(start, candles...)

	conf.calculateStochastic(internalCandles)
}

func (conf *stochastic) calculateStochastic(candles []*entity.Candle) {
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
		stochasticValue := new(entity.StochasticValue)
		stochasticValue.FastK = 100 * ((candles[i].Close - lowest) / (highest - lowest))

		sum := stochasticValue.FastK
		for j := i - conf.SmoothK + 1; j < i; j++ {
			sum += candles[j].Stochastics[conf.id].FastK
		}
		stochasticValue.IndexK = sum / float64(conf.SmoothK)
		stochasticValue.IndexD = calculateIndexD(conf.id, candles[i-conf.SmoothD+1:i+1])

		candles[i].Stochastics[conf.id] = stochasticValue
	}
}

//func (conf *stochastic) Update(candles []*entity.Candle) {
//	lowest := math.MaxFloat64
//	highest := float64(0)
//	lastIndex := len(candles) - 1
//	for i := len(candles) - conf.Length; i < len(candles); i++ {
//		if candles[i].Low < lowest {
//			lowest = candles[i].Low
//		}
//		if candles[i].High > highest {
//			highest = candles[i].High
//		}
//	}
//	fastK := 100 * ((candles[lastIndex].Close - lowest) / (highest - lowest))
//
//	sum := fastK
//	for j := len(candles) - conf.SmoothK; j < len(candles)-1; j++ {
//		sum += candles[j].Stochastics[conf.id].FastK
//	}
//
//	return &entity.IndicatorValue{
//		Stochastic: &entity.StochasticValue{
//			IndexK: sum / float64(conf.SmoothK),
//			IndexD: calculateIndexD(conf.id, candles[lastIndex-conf.SmoothD+1:]),
//			FastK:  fastK,
//		},
//	}
//}

func calculateIndexD(id uuid.UUID, candles []*entity.Candle) float64 {
	sum := float64(0)
	for _, candle := range candles {
		sum += candle.Stochastics[id].IndexK
	}
	return sum / float64(len(candles))
}

func validateStochasticConfigs(indicator *entity.StochasticConfigs) error {
	return nil
}
