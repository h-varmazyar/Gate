package calculator

import (
	"errors"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/entities"
	"golang.org/x/net/context"
	"math"
)

type Stochastic struct {
	id uint
	*entities.StochasticConfigs
	Market      *chipmunkAPI.Market
	Resolution  *chipmunkAPI.Resolution
	fasts       []float64
	lastCandles []*chipmunkAPI.Candle
	lastValues  []*indicatorsAPI.IndicatorValue
}

func NewStochastic(id uint, configs *entities.StochasticConfigs, market *chipmunkAPI.Market, resolution *chipmunkAPI.Resolution) (*Stochastic, error) {
	return &Stochastic{
		id:                id,
		StochasticConfigs: configs,
		Market:            market,
		Resolution:        resolution,
		fasts:             make([]float64, configs.SmoothK),
		lastCandles:       make([]*chipmunkAPI.Candle, configs.Period),
		lastValues:        make([]*indicatorsAPI.IndicatorValue, configs.SmoothD),
	}, nil
}

func (conf *Stochastic) Calculate(_ context.Context, candles []*chipmunkAPI.Candle) (*indicatorsAPI.IndicatorValues, error) {
	if err := conf.validateStochastic(len(candles)); err != nil {
		return nil, err
	}

	values := &indicatorsAPI.IndicatorValues{
		Values: make([]*indicatorsAPI.IndicatorValue, len(candles)),
	}

	for i := 0; i < conf.Period-1; i++ {
		values.Values[i] = &indicatorsAPI.IndicatorValue{
			Time: candles[i].Time,
		}
	}

	for i := conf.Period - 1; i < len(candles); i++ {
		lowest := math.MaxFloat64
		highest := float64(0)
		for j := i - (conf.Period - 1); j <= i; j++ {
			if candles[j].Low < lowest {
				lowest = candles[j].Low
			}
			if candles[j].High > highest {
				highest = candles[j].High
			}
		}

		fastK := 100 * ((candles[i].Close - lowest) / (highest - lowest))
		values.Values[i] = &indicatorsAPI.IndicatorValue{
			Time: candles[i].Time,
			Value: &indicatorsAPI.IndicatorValue_Stochastic{
				Stochastic: new(indicatorsAPI.StochasticValue),
			},
		}

		lastKIndex := float64(0)
		if values.Values[i-1].GetStochastic() != nil {
			lastKIndex = values.Values[i-1].GetStochastic().IndexK
		}

		values.Values[i].GetStochastic().IndexK = (fastK-conf.fasts[0])/float64(conf.SmoothK) + lastKIndex
		values.Values[i].GetStochastic().IndexD = calculateIndexD(values.Values[i-conf.SmoothD+1 : i+1])

		conf.fasts = conf.fasts[1:]
		conf.fasts = append(conf.fasts, fastK)
	}

	conf.lastCandles = candles[len(candles)-conf.Period:]
	conf.lastValues = values.Values[len(values.Values)-conf.SmoothD:]

	return values, nil
}

func (conf *Stochastic) UpdateLast(_ context.Context, candle *chipmunkAPI.Candle) *indicatorsAPI.IndicatorValue {
	if candle.Time > conf.lastCandles[len(conf.lastCandles)-1].Time {
		conf.fasts = conf.fasts[1:]
		conf.fasts = append(conf.fasts, 0)

		conf.lastCandles = conf.lastCandles[1:]
		conf.lastCandles = append(conf.lastCandles, candle)

		conf.lastValues = conf.lastValues[1:]
		newValue := &indicatorsAPI.IndicatorValue{
			Time: candle.Time,
			Value: &indicatorsAPI.IndicatorValue_Stochastic{
				Stochastic: new(indicatorsAPI.StochasticValue),
			},
		}
		conf.lastValues = append(conf.lastValues, newValue)
	}

	lowest := math.MaxFloat64
	highest := float64(0)
	for _, lastCandle := range conf.lastCandles {
		if lastCandle.Low < lowest {
			lowest = lastCandle.Low
		}
		if lastCandle.High > highest {
			highest = lastCandle.High
		}
	}

	fastK := 100 * ((candle.Close - lowest) / (highest - lowest))
	conf.lastValues[len(conf.lastValues)-1].GetStochastic().IndexK = (fastK-conf.fasts[0])/float64(conf.SmoothK) + conf.lastValues[len(conf.lastValues)-1].GetStochastic().IndexK
	conf.lastValues[len(conf.lastValues)-1].GetStochastic().IndexD = calculateIndexD(conf.lastValues)

	conf.fasts[len(conf.fasts)-1] = fastK
	conf.lastCandles[len(conf.lastCandles)-1] = candle

	return conf.lastValues[len(conf.lastValues)-1]
}

func (conf *Stochastic) GetMarket() *chipmunkAPI.Market {
	return conf.Market
}

func (conf *Stochastic) GetResolution() *chipmunkAPI.Resolution {
	return conf.Resolution
}

func (conf *Stochastic) GetId() uint {
	return conf.id
}

func (conf *Stochastic) validateStochastic(period int) error {
	if period <= conf.Period {
		return errors.New("rateLimiters length must bigger or equal than indicators period length")
	}
	if conf.SmoothK >= conf.Period {
		return errors.New("smoothK parameter must be smaller than indicators period length")
	}
	if conf.SmoothD >= conf.Period {
		return errors.New("smoothD parameter must be smaller than indicators period length")
	}
	return nil
}

func calculateIndexD(values []*indicatorsAPI.IndicatorValue) float64 {
	sum := float64(0)
	for _, value := range values {
		if value.GetStochastic() == nil {
			return 0
		}
		sum += value.GetStochastic().IndexK
	}
	return sum / float64(len(values))
}
