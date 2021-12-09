package indicators

import (
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/services/eagle/internal/pkg/models"
)

func sma(candles []*models.Candle, length int, source Source) {
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

func (conf *Configuration) calculateSMA(candles []*models.Candle) error {
	if err := conf.validateMA(len(candles)); err != nil {
		return err
	}

	sma(candles, conf.MovingAverageLength, conf.MovingAverageSource)
	return nil
}

func (conf *Configuration) updateSMA(candles []*models.Candle) error {
	if err := conf.validateMA(len(candles)); err != nil {
		return err
	}

	sma(candles, conf.MovingAverageLength, conf.MovingAverageSource)
	return nil
}

func (conf *Configuration) CalculateEMA(candles []*models.Candle) error {
	if err := conf.validateMA(len(candles)); err != nil {
		return err
	}
	sma(candles[conf.From:conf.MovingAverageLength+conf.From], conf.MovingAverageLength, conf.MovingAverageSource)
	candles[conf.MovingAverageLength+conf.From-1].MovingAverage.Exponential = candles[conf.MovingAverageLength+conf.From-1].MovingAverage.Simple

	factor := 2 / float64(conf.MovingAverageLength+1)
	for i := conf.MovingAverageLength + conf.From; i < len(candles); i++ {
		price := float64(0)
		switch conf.MovingAverageSource {
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
		candles[i].Exponential = price*factor + candles[i-1].Exponential*(1-factor)
	}
	return nil
}

func (conf *Configuration) UpdateEMA(candles []*models.Candle) error {
	i := len(candles) - 1
	price := float64(0)
	factor := 2 / float64(conf.MovingAverageLength+1)

	switch conf.MovingAverageSource {
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
	candles[i].Exponential = price*factor + candles[i-1].Exponential*(1-factor)
	return nil
}

func (conf *Configuration) validateMA(length int) error {
	if length <= conf.MovingAverageLength {
		return errors.New(fmt.Sprintf("candles length must be grater or equal than %d", conf.MovingAverageLength))
	}
	switch conf.MovingAverageSource {
	case SourceClose, SourceOpen, SourceHigh, SourceLow, SourceHL2, SourceHLC3, SourceOHLC4, SourceCustom:
		return nil
	default:
		return errors.New(fmt.Sprintf("moving average source not valid: %s", conf.MovingAverageSource))
	}
}
