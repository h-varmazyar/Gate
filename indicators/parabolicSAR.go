package indicators

import (
	"errors"
	"github.com/mrNobody95/Gate/models"
	"math"
	"sync"
)

const (
	Long  models.Trend = 1
	Short models.Trend = -1
)

func (conf *Configuration) CalculatePSAR(candles []models.Candle) error {
	if err := conf.validatePSAR(len(candles)); err != nil {
		return err
	}
	var trend models.Trend
	var nextSar float64
	var xp float64
	af := float64(0)
	if candles[1].High >= candles[0].High || candles[0].Low <= candles[1].Low {
		trend = Long
		nextSar = candles[1].Low
		xp = math.Max(candles[0].High, candles[1].High)
	} else {
		trend = Short
		nextSar = candles[1].High
		xp = math.Min(candles[0].Low, candles[1].Low)
	}
	candles[1].ParabolicSAR.Trend = trend
	candles[1].ParabolicSAR.SAR = nextSar

	for i := 1; i < len(candles)-1; i++ {
		nextSar = candles[i].ParabolicSAR.SAR
		if candles[i].Trend == Long {
			if candles[i].High > xp {
				xp = candles[i].High
				af = math.Min(conf.maxAcceleration, conf.Acceleration+af)
			}
			nextSar = math.Min(math.Min(candles[i].Low, candles[i-1].Low), nextSar+af*(xp-nextSar))
			if nextSar > candles[i+1].Low {
				trend = Short
				nextSar = xp
				xp = candles[i+1].Low
				af = conf.Acceleration
			}
		} else if candles[i].Trend == Short {
			if candles[i].Low < xp {
				xp = candles[i].Low
				af = math.Min(conf.maxAcceleration, conf.Acceleration+af)
			}
			nextSar = math.Max(math.Max(candles[i-1].High, candles[i].High), nextSar+af*(xp-nextSar))
			if nextSar < candles[i+1].High {
				trend = Long
				nextSar = xp
				xp = candles[i+1].High
				af = conf.Acceleration
			}
		}
		candles[i+1].ParabolicSAR.SAR = nextSar
		candles[i+1].ParabolicSAR.Trend = trend
		candles[i+1].ParabolicSAR.TrendFlipped = candles[i].ParabolicSAR.Trend != trend
	}
	conf.af = af
	conf.xp = xp
	return nil
}

func (conf *Configuration) UpdatePSAR(candles []models.Candle) {
	var trend models.Trend
	lastIndex := len(candles) - 1
	xp := conf.xp
	af := conf.af
	nextSar := candles[lastIndex-1].ParabolicSAR.SAR
	if candles[lastIndex-1].Trend == Long {
		if candles[lastIndex-1].High > xp {
			xp = candles[lastIndex-1].High
			af = math.Min(conf.maxAcceleration, conf.Acceleration+af)
		}
		nextSar = math.Min(math.Min(candles[lastIndex-2].Low, candles[lastIndex-1].Low), nextSar+af*(xp-nextSar))
		if nextSar > candles[lastIndex].Low {
			trend = Short
			nextSar = xp
			xp = candles[lastIndex].Low
			af = conf.Acceleration
		}
	} else if candles[lastIndex-1].Trend == Short {
		if candles[lastIndex-1].Low < xp {
			xp = candles[lastIndex-1].Low
			af = math.Min(conf.maxAcceleration, conf.Acceleration+af)
		}
		nextSar = math.Max(math.Max(candles[lastIndex-2].High, candles[lastIndex-1].High), nextSar+af*(xp-nextSar))
		if nextSar < candles[lastIndex].High {
			trend = Long
			nextSar = xp
			xp = candles[lastIndex].High
			af = conf.Acceleration
		}
	}
	indicatorLock := sync.Mutex{}
	indicatorLock.Lock()
	conf.xp = xp
	conf.af = af
	indicatorLock.Unlock()
	candles[lastIndex].ParabolicSAR.SAR = nextSar
	candles[lastIndex].ParabolicSAR.Trend = trend
	candles[lastIndex].ParabolicSAR.TrendFlipped = candles[lastIndex-1].ParabolicSAR.Trend != trend
}

func (conf *Configuration) validatePSAR(length int) error {
	if length < 2 {
		return errors.New("candle length must be more than 2")
	}
	return nil
}
