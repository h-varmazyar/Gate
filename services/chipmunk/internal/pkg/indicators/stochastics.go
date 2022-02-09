package indicators

import (
	"errors"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/repository"
	"math"
)

type Stochastic struct {
	Length  int
	SmoothK int
	SmoothD int
	Values  []*StochasticResponse
}

func (conf *Stochastic) validateStochastic(length int) error {
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

func (conf *Stochastic) Calculate(candles []*repository.Candle) error {
	if err := conf.validateStochastic(len(candles)); err != nil {
		return err
	}

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
		conf.Values[i].FastK = 100 * ((candles[i].Close - lowest) / (highest - lowest))

		sum := float64(0)
		for j := i - conf.SmoothK + 1; j <= i; j++ {
			sum += conf.Values[j].FastK
		}
		conf.Values[i].IndexK = sum / float64(conf.SmoothK)
		conf.Values[i].IndexD = calculateIndexD(conf.Values[i-conf.SmoothD+1 : i+1])
	}
	return nil
}

func (conf *Stochastic) Update(candles []*repository.Candle) *StochasticResponse {
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
		sum += conf.Values[j].FastK
	}

	//conf.lock.Lock()
	//candles[lastIndex].FastK = fastK
	//candles[lastIndex].IndexK = sum / float64(conf.SmoothK)
	//candles[lastIndex].IndexD = calculateIndexD(candles[lastIndex-conf.SmoothD+1:])
	//conf.lock.Unlock()

	return &StochasticResponse{
		IndexK: sum / float64(conf.SmoothK),
		IndexD: calculateIndexD(conf.Values[lastIndex-conf.SmoothD+1:]),
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
