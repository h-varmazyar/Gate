package indicators

import (
	"errors"
	"github.com/mrNobody95/Gate/models"
	"math"
)

func (conf *Configuration) CalculateBollingerBand(candles []models.Candle) error {
	if err := conf.validateBollingerBand(len(candles)); err != nil {
		return err
	}
	err := conf.calculateSMA(candles)
	if err != nil {
		return err
	}
	for i := conf.BollingerLength - 1; i < len(candles); i++ {
		variance := float64(0)
		ma := candles[i].MovingAverage.Simple
		for j := 1 + i - conf.BollingerLength; j <= i; j++ {
			sum := float64(0)
			switch conf.MovingAverageSource {
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
		variance /= float64(conf.BollingerLength)
		candles[i].BollingerBand.MA = ma
		candles[i].BollingerBand.UpperBond = ma + float64(conf.BollingerDeviation)*math.Sqrt(variance)
		candles[i].BollingerBand.LowerBond = ma - float64(conf.BollingerDeviation)*math.Sqrt(variance)
	}
	return nil
}

func (conf *Configuration) UpdateBollingerBand(candles []models.Candle) error {
	err := conf.updateSMA(candles) //conf.Candles[len(conf.Candles)-conf.MovingAverageLength:]
	if err != nil {
		return err
	}
	variance := float64(0)
	ma := candles[len(candles)-1].MovingAverage.Simple
	for j := 0; j < len(candles); j++ {
		sum := float64(0)
		switch conf.MovingAverageSource {
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
	variance /= float64(conf.BollingerLength)
	candles[len(candles)-1].BollingerBand.MA = ma
	candles[len(candles)-1].BollingerBand.UpperBond = ma + float64(conf.BollingerDeviation)*math.Sqrt(variance)
	candles[len(candles)-1].BollingerBand.LowerBond = ma - float64(conf.BollingerDeviation)*math.Sqrt(variance)
	return nil
}

func (conf *Configuration) validateBollingerBand(length int) error {
	if conf.BollingerLength != conf.MovingAverageLength {
		return errors.New("bollinger band length must be equal to moving average length")
	}
	if length < conf.BollingerLength {
		return errors.New("BollingerLength must be bigger than or equal to candle length")
	}
	if conf.BollingerDeviation < 1 {
		return errors.New("deviation value must be positive")
	}
	return nil
}
