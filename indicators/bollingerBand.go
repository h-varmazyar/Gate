package indicators

import (
	"errors"
	"github.com/mrNobody95/Gate/models"
	"math"
)

type bollingerBandConfig struct {
	*basicConfig
	Deviation int
}

func NewBollingerBandConfig(deviation, length int) *bollingerBandConfig {
	return &bollingerBandConfig{
		basicConfig: &basicConfig{
			Length: length,
		},
		Deviation: deviation,
	}
}

func (conf *bollingerBandConfig) CalculateBollingerBand(candles []models.Candle, appendCandles bool) error {
	rangeCounter := conf.Length
	if appendCandles {
		rangeCounter = len(conf.Candles)
		conf.Candles = append(conf.Candles, candles...)
	} else {
		conf.Candles = candles
	}
	if err := conf.validate(); err != nil {
		return err
	}

	mac := NewMovingAverageConfig(conf.Length, SourceHLC3)
	err := mac.CalculateSimple(conf.Candles[rangeCounter-conf.Length:], false)
	if err != nil {
		return err
	}
	for i, candle := range conf.Candles[rangeCounter-1:] {
		variance := float64(0)
		//ma := TypicalPriceMovingAverage(conf.Candles[rangeCounter-conf.Length+i : rangeCounter+i])
		ma := mac.Candles[rangeCounter+i].MovingAverage.Simple
		for _, innerCandle := range conf.Candles[rangeCounter-conf.Length+i : rangeCounter+i] {
			sum := (innerCandle.Close + innerCandle.High + innerCandle.Low) / 3
			variance += math.Pow(ma-sum, 2)
		}
		variance /= float64(len(candles))
		candle.BollingerBand = &models.BollingerBand{
			UpperBond: ma + float64(conf.Deviation)*math.Sqrt(variance),
			LowerBond: ma - float64(conf.Deviation)*math.Sqrt(variance),
			MA:        ma,
		}
	}
	return nil
}

func (conf *bollingerBandConfig) validate() error {
	if len(conf.Candles) < conf.Length {
		return errors.New("candles length must bigger or equal than indicator period length")
	}
	if conf.Deviation < 1 {
		return errors.New("deviation value must be positive")
	}
	return nil
}
