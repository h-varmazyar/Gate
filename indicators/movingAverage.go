package indicators

import (
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/models"
)

type movingAverageConfig struct {
	*basicConfig
	source Source
}

func NewMovingAverageConfig(length int, source Source) *movingAverageConfig {
	return &movingAverageConfig{
		basicConfig: &basicConfig{
			Length: length,
		},
		source: source,
	}
}

func (conf *movingAverageConfig) CalculateSimple(candles []models.Candle, appendCandles bool) error {
	var rangeCounter int
	if appendCandles {
		rangeCounter = len(conf.Candles) + 1
		conf.Candles = append(conf.Candles, candles...)
	} else {
		conf.Candles = candles
		rangeCounter = conf.Length
	}
	if err := conf.validate(); err != nil {
		return err
	}

	sma(conf.Candles[rangeCounter-conf.Length:], conf.Length, conf.source)
	return nil
}

func sma(candles []models.Candle, length int, source Source) {
	for i, candle := range candles[length-1:] {
		sum := float64(0)
		for _, innerCandle := range candles[i : i+length] {
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
		candle.MovingAverage.Simple = sum / float64(length)
	}
}

func (conf *movingAverageConfig) CalculateExponential(candles []models.Candle, appendCandles bool) error {
	var rangeCounter int
	if appendCandles {
		rangeCounter = len(conf.Candles)
		conf.Candles = append(conf.Candles, candles...)
	} else {
		conf.Candles = candles
		rangeCounter = conf.Length

		sma(conf.Candles[0:conf.Length], conf.Length, conf.source)
		conf.Candles[conf.Length-1].MovingAverage.Exponential = conf.Candles[conf.Length-1].MovingAverage.Simple
	}
	if err := conf.validate(); err != nil {
		return err
	}
	factor := 2 / float64(conf.Length+1)
	for i, candle := range conf.Candles[rangeCounter:] {
		price := float64(0)
		switch conf.source {
		case SourceClose:
			price = candle.Close
		case SourceOpen:
			price = candle.Open
		case SourceHigh:
			price = candle.High
		case SourceLow:
			price = candle.Low
		case SourceHL2:
			price = (candle.High + candle.Low) / 2
		case SourceHLC3:
			price = (candle.High + candle.Low + candle.Close) / 3
		case SourceOHLC4:
			price = (candle.Open + candle.Close + candle.High + candle.Low) / 4
		}
		candle.MovingAverage.Exponential = price*factor + conf.Candles[rangeCounter+i-1].MovingAverage.Exponential*(1-factor)
	}
	return nil
}

func (conf *movingAverageConfig) validate() error {
	if len(conf.Candles) < conf.Length {
		return errors.New(fmt.Sprintf("candles length must be grater or equal than %d", conf.Length))
	}
	return nil
}
