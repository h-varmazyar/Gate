package indicators

import (
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/repository"
)

type MovingAverage struct {
	Source Source
	Length int
	Values []*MovingAverageResponse
}

func (conf *MovingAverage) sma(candles []*repository.Candle) error {
	if err := conf.validateMA(len(candles)); err != nil {
		return err
	}
	for i := conf.Length - 1; i < len(candles); i++ {
		sum := float64(0)
		for _, innerCandle := range candles[i-conf.Length+1 : i+1] {
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
		if conf.Values[i] == nil {
			conf.Values[i] = new(MovingAverageResponse)
		}
		conf.Values[i].Simple = sum / float64(conf.Length)
	}
	return nil
}

func (conf *MovingAverage) updateSMA(candles []*repository.Candle) float64 {
	smaConf := MovingAverage{
		Source: conf.Source,
		Length: conf.Length,
		Values: make([]*MovingAverageResponse, len(candles)),
	}
	if err := smaConf.sma(candles); err != nil {
		return float64(0)
	}
	return smaConf.Values[len(candles)-1].Simple
}

func (conf *MovingAverage) Calculate(candles []*repository.Candle) error {
	conf.Values = make([]*MovingAverageResponse, len(candles))
	if err := conf.sma(candles); err != nil {
		return err
	}
	conf.Values[conf.Length-1].Exponential = conf.Values[conf.Length-1].Simple

	factor := 2 / float64(conf.Length+1)
	for i := conf.Length; i < len(candles); i++ {
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
		conf.Values[i].Exponential = price*factor + conf.Values[i-1].Exponential*(1-factor)
	}
	return nil
}

func (conf *MovingAverage) Update(candles []*repository.Candle) *MovingAverageResponse {
	i := len(candles) - 1
	price := float64(0)
	factor := 2 / float64(conf.Length+1)

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
	return &MovingAverageResponse{
		Exponential: price*factor + conf.Values[i-1].Exponential*(1-factor),
		Simple:      conf.updateSMA(candles),
	}
}

func (conf *MovingAverage) validateMA(length int) error {
	if length <= conf.Length {
		return errors.New(fmt.Sprintf("candles length must be grater or equal than %d", conf.Length))
	}
	switch conf.Source {
	case SourceClose, SourceOpen, SourceHigh, SourceLow, SourceHL2, SourceHLC3, SourceOHLC4, SourceCustom:
		return nil
	default:
		return errors.New(fmt.Sprintf("moving average source not valid: %s", conf.Source))
	}
}
