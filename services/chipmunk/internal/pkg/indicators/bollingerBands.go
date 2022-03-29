package indicators

import (
	"errors"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
	"math"
)

type bollingerBands struct {
	basicConfig
	Deviation int
	Source    Source
}

func NewBollingerBands(length, deviation int, source Source, marketName string) *bollingerBands {
	return &bollingerBands{
		basicConfig: basicConfig{
			MarketName: marketName,
			id:         uuid.New(),
			length:     length,
		},
		Deviation: deviation,
		Source:    source,
	}
}

func (conf *bollingerBands) GetID() uuid.UUID {
	return conf.id
}

func (conf *bollingerBands) GetType() IndicatorType {
	return BollingerBands
}

func (conf *bollingerBands) GetLength() int {
	return conf.length
}

func (conf *bollingerBands) Calculate(candles []*repository.Candle) error {
	if err := conf.validateBollingerBand(len(candles)); err != nil {
		return err
	}
	cloned := cloneCandles(candles)
	smaConf := movingAverage{
		basicConfig: basicConfig{
			MarketName: conf.MarketName,
			id:         uuid.New(),
			length:     conf.length,
		},
		Source: conf.Source,
	}
	sma, err := smaConf.sma(cloned)
	if err != nil {
		return err
	}
	for i := conf.length - 1; i < len(candles); i++ {
		variance := float64(0)
		ma := sma[i]
		for j := 1 + i - conf.length; j <= i; j++ {
			sum := float64(0)
			switch conf.Source {
			case SourceOpen:
				sum = candles[j].Open
			case SourceHigh:
				sum = candles[j].High
			case SourceLow:
				sum = candles[j].Low
			case SourceClose:
				sum = candles[j].Close
			case SourceOHLC4:
				sum = (candles[j].Open + candles[j].High + candles[j].Low + candles[j].Close) / 4
			case SourceHLC3:
				sum = (candles[j].Low + candles[j].High + candles[j].Close) / 3
			case SourceHL2:
				sum = (candles[j].Low + candles[j].High) / 2
			}
			variance += math.Pow(ma-sum, 2)
		}
		variance /= float64(conf.length)

		candles[i].BollingerBands[conf.id] = repository.BollingerBandsValue{
			UpperBand: ma + float64(conf.Deviation)*math.Sqrt(variance),
			LowerBand: ma - float64(conf.Deviation)*math.Sqrt(variance),
			MA:        ma,
		}
	}
	return nil
}

func (conf *bollingerBands) Update(candles []*repository.Candle) *repository.IndicatorValue {
	smaConf := movingAverage{
		basicConfig: basicConfig{
			MarketName: conf.MarketName,
			id:         uuid.New(),
			length:     conf.length,
		},
		Source: conf.Source,
	}
	sma, err := smaConf.sma(candles)
	if err != nil {
		return nil
	}
	variance := float64(0)
	ma := sma[len(candles)-1]
	for j := 0; j < len(candles); j++ {
		sum := float64(0)
		switch conf.Source {
		case SourceOpen:
			sum = candles[j].Open
		case SourceHigh:
			sum = candles[j].High
		case SourceLow:
			sum = candles[j].Low
		case SourceClose:
			sum = candles[j].Close
		case SourceOHLC4:
			sum = (candles[j].Open + candles[j].High + candles[j].Low + candles[j].Close) / 4
		case SourceHLC3:
			sum = (candles[j].Low + candles[j].High + candles[j].Close) / 3
		case SourceHL2:
			sum = (candles[j].Low + candles[j].High) / 2
		}
		variance += math.Pow(ma-sum, 2)
	}
	variance /= float64(conf.length)
	return &repository.IndicatorValue{
		BB: &repository.BollingerBandsValue{
			UpperBand: ma + float64(conf.Deviation)*math.Sqrt(variance),
			LowerBand: ma - float64(conf.Deviation)*math.Sqrt(variance),
			MA:        ma,
		},
	}
}

func (conf *bollingerBands) validateBollingerBand(length int) error {
	if conf.length != conf.length {
		return errors.New("bollinger band length must be equal to moving average length")
	}
	if length < conf.length {
		return errors.New("length must be bigger than or equal to candle length")
	}
	if conf.Deviation < 1 {
		return errors.New("deviation value must be positive")
	}
	return nil
}
