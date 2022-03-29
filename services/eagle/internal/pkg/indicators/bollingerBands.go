package indicators

import (
	"errors"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/models"
	"math"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 03.12.21
* Github: https://github.com/h-varmazyar
* Email: hossein.varmazyar@yahoo.com
**/

func (conf *Configuration) CalculateBollingerBand(candles []*models.Candle) error {
	if err := conf.validateBollingerBand(len(candles)); err != nil {
		return err
	}
	cloned := cloneCandles(candles)
	err := conf.calculateSMA(cloned)
	if err != nil {
		return err
	}
	for i := conf.BollingerLength - 1; i < len(candles); i++ {
		variance := float64(0)
		ma := cloned[i].MovingAverage.Simple
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
		candles[i].BollingerBands.MA = ma
		candles[i].BollingerBands.UpperBand = ma + float64(conf.BollingerDeviation)*math.Sqrt(variance)
		candles[i].BollingerBands.LowerBand = ma - float64(conf.BollingerDeviation)*math.Sqrt(variance)
	}
	return nil
}

func (conf *Configuration) UpdateBollingerBand(candles []*models.Candle) error {
	cloned := cloneCandles(candles[len(candles)-conf.BollingerLength:])
	err := conf.updateSMA(cloned)
	if err != nil {
		return err
	}
	variance := float64(0)
	ma := cloned[len(cloned)-1].MovingAverage.Simple
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
	conf.lock.Lock()
	candles[len(candles)-1].BollingerBands.MA = ma
	candles[len(candles)-1].BollingerBands.UpperBand = ma + float64(conf.BollingerDeviation)*math.Sqrt(variance)
	candles[len(candles)-1].BollingerBands.LowerBand = ma - float64(conf.BollingerDeviation)*math.Sqrt(variance)
	conf.lock.Unlock()
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
