package indicators

import (
	"errors"
	"github.com/mrNobody95/Gate/models"
	"math"
)

const (
	Long  models.Trend = 1
	Short models.Trend = -1
)

func (conf *Configuration) CalculatePSAR() error {
	if err := conf.validatePSAR(); err != nil {
		return err
	}
	var trend models.Trend
	var nextSar float64
	var xp float64
	af := float64(0)
	if conf.Candles[1].High >= conf.Candles[0].High || conf.Candles[0].Low <= conf.Candles[1].Low {
		trend = Long
		nextSar = conf.Candles[1].Low
		xp = math.Max(conf.Candles[0].High, conf.Candles[1].High)
	} else {
		trend = Short
		nextSar = conf.Candles[1].High
		xp = math.Min(conf.Candles[0].Low, conf.Candles[1].Low)
	}
	indicatorLock.Lock()
	conf.Candles[1].ParabolicSAR.Trend = trend
	conf.Candles[1].ParabolicSAR.SAR = nextSar
	indicatorLock.Unlock()

	for i := 1; i < len(conf.Candles)-1; i++ {
		nextSar = conf.Candles[i].ParabolicSAR.SAR
		if conf.Candles[i].Trend == Long {
			if conf.Candles[i].High > xp {
				xp = conf.Candles[i].High
				af = math.Min(conf.maxAcceleration, conf.Acceleration+af)
			}
			nextSar = math.Min(math.Min(conf.Candles[i].Low, conf.Candles[i-1].Low), nextSar+af*(xp-nextSar))
			if nextSar > conf.Candles[i+1].Low {
				trend = Short
				nextSar = xp
				xp = conf.Candles[i+1].Low
				af = conf.Acceleration
			}
		} else if conf.Candles[i].Trend == Short {
			if conf.Candles[i].Low < xp {
				xp = conf.Candles[i].Low
				af = math.Min(conf.maxAcceleration, conf.Acceleration+af)
			}
			nextSar = math.Max(math.Max(conf.Candles[i-1].High, conf.Candles[i].High), nextSar+af*(xp-nextSar))
			if nextSar < conf.Candles[i+1].High {
				trend = Long
				nextSar = xp
				xp = conf.Candles[i+1].High
				af = conf.Acceleration
			}
		}
		indicatorLock.Lock()
		conf.Candles[i+1].ParabolicSAR.SAR = nextSar
		conf.Candles[i+1].ParabolicSAR.Trend = trend
		conf.Candles[i+1].ParabolicSAR.TrendFlipped = conf.Candles[i].ParabolicSAR.Trend != trend
		indicatorLock.Unlock()
	}
	indicatorLock.Lock()
	conf.af = af
	conf.xp = xp
	indicatorLock.Unlock()
	return nil
}

func (conf *Configuration) UpdatePSAR() error {
	var trend models.Trend
	lastIndex := len(conf.Candles) - 1
	xp := conf.xp
	af := conf.af
	nextSar := conf.Candles[lastIndex-1].ParabolicSAR.SAR
	if conf.Candles[lastIndex-1].Trend == Long {
		if conf.Candles[lastIndex-1].High > xp {
			xp = conf.Candles[lastIndex-1].High
			af = math.Min(conf.maxAcceleration, conf.Acceleration+af)
		}
		nextSar = math.Min(math.Min(conf.Candles[lastIndex-2].Low, conf.Candles[lastIndex-1].Low), nextSar+af*(xp-nextSar))
		if nextSar > conf.Candles[lastIndex].Low {
			trend = Short
			nextSar = xp
			xp = conf.Candles[lastIndex].Low
			af = conf.Acceleration
		}
	} else if conf.Candles[lastIndex-1].Trend == Short {
		if conf.Candles[lastIndex-1].Low < xp {
			xp = conf.Candles[lastIndex-1].Low
			af = math.Min(conf.maxAcceleration, conf.Acceleration+af)
		}
		nextSar = math.Max(math.Max(conf.Candles[lastIndex-2].High, conf.Candles[lastIndex-1].High), nextSar+af*(xp-nextSar))
		if nextSar < conf.Candles[lastIndex].High {
			trend = Long
			nextSar = xp
			xp = conf.Candles[lastIndex].High
			af = conf.Acceleration
		}
	}
	indicatorLock.Lock()
	conf.xp = xp
	conf.af = af
	conf.Candles[lastIndex].ParabolicSAR.SAR = nextSar
	conf.Candles[lastIndex].ParabolicSAR.Trend = trend
	conf.Candles[lastIndex].ParabolicSAR.TrendFlipped = conf.Candles[lastIndex-1].ParabolicSAR.Trend != trend
	indicatorLock.Unlock()
	return nil
}

func (conf *Configuration) validatePSAR() error {
	if len(conf.Candles) < 2 {
		return errors.New("candle length must be more than 2")
	}
	return nil
}
