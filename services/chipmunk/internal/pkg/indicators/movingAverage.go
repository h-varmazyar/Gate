package indicators

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
)

type movingAverage struct {
	basicConfig
	Source Source
}

func NewMovingAverage(length int, source Source, marketName string) *movingAverage {
	return &movingAverage{
		basicConfig: basicConfig{
			MarketName: marketName,
			id:         uuid.New(),
			length:     length,
		},
		Source: source,
	}
}

func (conf *movingAverage) GetID() uuid.UUID {
	return conf.id
}

func (conf *movingAverage) GetType() IndicatorType {
	return MovingAverage
}

func (conf *movingAverage) GetLength() int {
	return conf.length
}

func (conf *movingAverage) sma(candles []*repository.Candle) ([]float64, error) {
	response := make([]float64, len(candles))
	if err := conf.validateMA(len(candles)); err != nil {
		return nil, err
	}
	for i := conf.length - 1; i < len(candles); i++ {
		sum := float64(0)
		for _, innerCandle := range candles[i-conf.length+1 : i+1] {
			switch conf.Source {
			case SourceClose:
				sum += innerCandle.Close
			case SourceOpen:
				sum += innerCandle.Open
			case SourceHigh:
				sum += innerCandle.High
			case SourceLow:
				sum += innerCandle.Low
			case SourceHL2:
				sum += (innerCandle.High + innerCandle.Low) / 2
			case SourceHLC3:
				sum += (innerCandle.High + innerCandle.Low + innerCandle.Close) / 3
			case SourceOHLC4:
				sum += (innerCandle.Open + innerCandle.Close + innerCandle.High + innerCandle.Low) / 4
			}
		}
		response[i] = sum / float64(conf.length)
	}
	return response, nil
}

func (conf *movingAverage) updateSMA(candles []*repository.Candle) float64 {
	smaConf := movingAverage{
		basicConfig: basicConfig{
			MarketName: conf.MarketName,
			id:         uuid.New(),
			length:     conf.length,
		},
		Source: conf.Source,
	}
	if sma, err := smaConf.sma(candles); err != nil {
		return float64(0)
	} else {
		return sma[len(candles)-1]
	}
}

func (conf *movingAverage) Calculate(candles []*repository.Candle) error {
	values := make([]*repository.MovingAverageValue, len(candles))
	var sma []float64
	var err error
	if sma, err = conf.sma(candles); err != nil {
		return err
	}
	for i := range values {
		values[i].Simple = sma[i]
	}
	values[conf.length-1].Exponential = values[conf.length-1].Simple

	factor := 2 / float64(conf.length+1)
	for i := conf.length; i < len(candles); i++ {
		price := float64(0)
		switch conf.Source {
		case SourceClose:
			price = candles[i].Close
		case SourceOpen:
			price = candles[i].Open
		case SourceHigh:
			price = candles[i].High
		case SourceLow:
			price = candles[i].Low
		case SourceHL2:
			price = (candles[i].High + candles[i].Low) / 2
		case SourceHLC3:
			price = (candles[i].High + candles[i].Low + candles[i].Close) / 3
		case SourceOHLC4:
			price = (candles[i].Open + candles[i].Close + candles[i].High + candles[i].Low) / 4
		}
		values[i].Exponential = price*factor + values[i-1].Exponential*(1-factor)
	}

	for i := 0; i < len(candles); i++ {
		candles[i].MovingAverages[conf.id] = &repository.MovingAverageValue{
			Simple:      values[i].Simple,
			Exponential: values[i].Exponential,
		}
	}
	return nil
}

func (conf *movingAverage) Update(candles []*repository.Candle) *repository.IndicatorValue {
	//candles := buffer.Markets.GetLastNCandles(conf.MarketName, 2)
	//values := buffer.Markets.GetLastNIndicatorValue(conf.MarketName, conf.GetID(), 2)

	i := len(candles) - 1
	price := float64(0)
	factor := 2 / float64(conf.length+1)

	switch conf.Source {
	case SourceClose:
		price = candles[i].Close
	case SourceOpen:
		price = candles[i].Open
	case SourceHigh:
		price = candles[i].High
	case SourceLow:
		price = candles[i].Low
	case SourceHL2:
		price = (candles[i].High + candles[i].Low) / 2
	case SourceHLC3:
		price = (candles[i].High + candles[i].Low + candles[i].Close) / 3
	case SourceOHLC4:
		price = (candles[i].Open + candles[i].Close + candles[i].High + candles[i].Low) / 4
	}
	return &repository.IndicatorValue{
		MA: &repository.MovingAverageValue{
			Exponential: price*factor + candles[i-1].MovingAverages[conf.id].Exponential*(1-factor),
			Simple:      conf.updateSMA(candles),
		},
	}
}

func (conf *movingAverage) validateMA(length int) error {
	if length < conf.length {
		return errors.New(fmt.Sprintf("candles length must be grater or equal than %d", conf.length))
	}
	switch conf.Source {
	case SourceClose, SourceOpen, SourceHigh, SourceLow, SourceHL2, SourceHLC3, SourceOHLC4, SourceCustom:
		return nil
	default:
		return errors.New(fmt.Sprintf("moving average source not valid: %s", conf.Source))
	}
}
