package indicators

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/buffer"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/repository"
	"math"
)

type stochastic struct {
	basicConfig
	Length  int
	SmoothK int
	SmoothD int
}

func NewStochastic(length, smoothK, smoothD int, marketName string) *stochastic {
	return &stochastic{
		basicConfig: basicConfig{
			MarketName: marketName,
			id:         uuid.New(),
		},
		Length:  length,
		SmoothK: smoothK,
		SmoothD: smoothD,
	}
}

func (conf *stochastic) GetID() string {
	return conf.id.String()
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

func (conf *stochastic) Calculate(candles []*repository.Candle, response interface{}) error {
	if err := conf.validateStochastic(len(candles)); err != nil {
		return err
	}
	values := make([]*StochasticResponse, len(candles))
	for i := conf.Length - 1; i < len(candles); i++ {
		lowest := math.MaxFloat64
		highest := float64(0)
		from := i - (conf.Length - 1)
		for j := from; j < from+conf.Length; j++ {
			if candles[j].Low < lowest {
				lowest = candles[j].Low
			}
			if candles[j].High > highest {
				highest = candles[j].High
			}
		}
		values[i].FastK = 100 * ((candles[i].Close - lowest) / (highest - lowest))

		sum := float64(0)
		for j := i - conf.SmoothK + 1; j <= i; j++ {
			sum += values[j].FastK
		}
		values[i].IndexK = sum / float64(conf.SmoothK)
		values[i].IndexD = calculateIndexD(values[i-conf.SmoothD+1 : i+1])
	}
	response = interface{}(values)
	return nil
}

func (conf *stochastic) Update() interface{} {
	candles := buffer.Markets.GetLastNCandles(conf.MarketName, conf.Length)
	values := make([]*StochasticResponse, conf.Length)
	for i, resp := range buffer.Markets.GetLastNIndicatorValue(conf.MarketName, conf.GetID(), conf.Length) {
		values[i] = resp.(*StochasticResponse)
	}
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
		sum += values[j].FastK
	}

	return &StochasticResponse{
		IndexK: sum / float64(conf.SmoothK),
		IndexD: calculateIndexD(values[lastIndex-conf.SmoothD+1:]),
		FastK:  fastK,
	}
}

func calculateIndexD(values []*StochasticResponse) float64 {
	sum := float64(0)
	for _, value := range values {
		sum += value.IndexK
	}
	return sum / float64(len(values))
}
