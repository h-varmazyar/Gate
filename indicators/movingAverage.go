package indicators

import (
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/models"
)

//type movingAverageConfig struct {
//	*basicConfig
//	source Source
//}
//
//func NewMovingAverageConfig(length int, source Source) *movingAverageConfig {
//	return &movingAverageConfig{
//		basicConfig: &basicConfig{
//			Length: length,
//		},
//		source: source,
//	}
//}

func sma(candles []models.Candle, length int, source Source) {
	for i := length - 1; i < len(candles); i++ {
		sum := float64(0)
		for _, innerCandle := range candles[i-length+1 : i+1] {
			switch source {
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
		candles[i].MovingAverage.Simple = sum / float64(length)
	}
}

func (conf *Configuration) CalculateSMA() error {
	if err := conf.validateMA(); err != nil {
		return err
	}

	sma(conf.Candles, conf.MovingAverageLength, conf.MovingAverageSource)
	return nil
}

func (conf *Configuration) UpdateSMA() error {
	if err := conf.validateMA(); err != nil {
		return err
	}

	sma(conf.Candles[len(conf.Candles)-conf.MovingAverageLength:], conf.MovingAverageLength, conf.MovingAverageSource)
	return nil
}

func (conf *Configuration) CalculateEMA() error {
	if err := conf.validateMA(); err != nil {
		return err
	}
	sma(conf.Candles[:conf.MovingAverageLength], conf.MovingAverageLength, conf.MovingAverageSource)
	conf.Candles[conf.MovingAverageLength-1].MovingAverage.Exponential = conf.Candles[conf.MovingAverageLength-1].MovingAverage.Simple

	factor := 2 / float64(conf.MovingAverageLength+1)
	for i := conf.MovingAverageLength; i < len(conf.Candles); i++ {
		price := float64(0)
		switch conf.source {
		case SourceClose:
			price = conf.Candles[i].Close
		case SourceOpen:
			price = conf.Candles[i].Open
		case SourceHigh:
			price = conf.Candles[i].High
		case SourceLow:
			price = conf.Candles[i].Low
		case SourceHL2:
			price = (conf.Candles[i].High + conf.Candles[i].Low) / 2
		case SourceHLC3:
			price = (conf.Candles[i].High + conf.Candles[i].Low + conf.Candles[i].Close) / 3
		case SourceOHLC4:
			price = (conf.Candles[i].Open + conf.Candles[i].Close + conf.Candles[i].High + conf.Candles[i].Low) / 4
		}
		conf.Candles[i].MovingAverage.Exponential = price*factor + conf.Candles[i-1].MovingAverage.Exponential*(1-factor)
	}
	return nil
}

func (conf *Configuration) UpdateEMA() error {
	if err := conf.validateMA(); err != nil {
		return err
	}
	i := len(conf.Candles) - 1
	price := float64(0)
	factor := 2 / float64(conf.MovingAverageLength+1)

	switch conf.source {
	case SourceClose:
		price = conf.Candles[i].Close
	case SourceOpen:
		price = conf.Candles[i].Open
	case SourceHigh:
		price = conf.Candles[i].High
	case SourceLow:
		price = conf.Candles[i].Low
	case SourceHL2:
		price = (conf.Candles[i].High + conf.Candles[i].Low) / 2
	case SourceHLC3:
		price = (conf.Candles[i].High + conf.Candles[i].Low + conf.Candles[i].Close) / 3
	case SourceOHLC4:
		price = (conf.Candles[i].Open + conf.Candles[i].Close + conf.Candles[i].High + conf.Candles[i].Low) / 4
	}
	conf.Candles[i].MovingAverage.Exponential = price*factor + conf.Candles[i-1].MovingAverage.Exponential*(1-factor)
	return nil
}

func (conf *Configuration) validateMA() error {
	if len(conf.Candles) <= conf.MovingAverageLength {
		return errors.New(fmt.Sprintf("candles length must be grater or equal than %d", conf.MovingAverageLength))
	}
	switch conf.MovingAverageSource {
	case SourceClose, SourceOpen, SourceHigh, SourceLow, SourceHL2, SourceHLC3, SourceOHLC4:
		return nil
	default:
		return errors.New(fmt.Sprintf("moving average source not valid: %s", conf.MovingAverageSource))
	}
}
