package indicators

import (
	"errors"
	"math"
)

func (conf *Configuration) CalculateBollingerBand() error {
	if err := conf.validateBollingerBand(); err != nil {
		return err
	}
	err := conf.calculateSMA()
	if err != nil {
		return err
	}
	for i := conf.BollingerLength - 1; i < len(conf.Candles); i++ {
		variance := float64(0)
		ma := conf.Candles[i].MovingAverage.Simple
		for j := 1 + i - conf.BollingerLength; j <= i; j++ {
			sum := (conf.Candles[j].Close + conf.Candles[j].High + conf.Candles[j].Low) / 3
			variance += math.Pow(ma-sum, 2)
		}
		variance /= float64(conf.BollingerLength)
		indicatorLock.Lock()
		conf.Candles[i].BollingerBand.MA = ma
		conf.Candles[i].BollingerBand.UpperBond = ma + float64(conf.BollingerDeviation)*math.Sqrt(variance)
		conf.Candles[i].BollingerBand.LowerBond = ma - float64(conf.BollingerDeviation)*math.Sqrt(variance)
		indicatorLock.Unlock()
	}
	return nil
}

func (conf *Configuration) UpdateBollingerBand() error {
	err := conf.updateSMA()
	if err != nil {
		return err
	}
	variance := float64(0)
	ma := conf.Candles[len(conf.Candles)-1].MovingAverage.Simple
	for j := len(conf.Candles) - conf.BollingerLength; j < len(conf.Candles); j++ {
		sum := (conf.Candles[j].Close + conf.Candles[j].High + conf.Candles[j].Low) / 3
		variance += math.Pow(ma-sum, 2)
	}
	variance /= float64(conf.BollingerLength)
	indicatorLock.Lock()
	conf.Candles[len(conf.Candles)-1].BollingerBand.MA = ma
	conf.Candles[len(conf.Candles)-1].BollingerBand.UpperBond = ma + float64(conf.BollingerDeviation)*math.Sqrt(variance)
	conf.Candles[len(conf.Candles)-1].BollingerBand.LowerBond = ma - float64(conf.BollingerDeviation)*math.Sqrt(variance)
	indicatorLock.Unlock()
	return nil
}

func (conf *Configuration) validateBollingerBand() error {
	if conf.BollingerLength != conf.MovingAverageLength {
		return errors.New("bollinger band length must be equal to moving average length")
	}
	if len(conf.Candles) < conf.BollingerLength {
		return errors.New("candles BollingerLength must bigger or equal than indicator period BollingerLength")
	}
	if conf.BollingerDeviation < 1 {
		return errors.New("deviation value must be positive")
	}
	return nil
}
