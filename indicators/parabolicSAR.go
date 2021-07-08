package indicators

import (
	"errors"
	"math"
)

type Trend int

const (
	Long  Trend = 1
	Short Trend = -1
)

func (conf *IndicatorConfig) CalculatePSAR() error {
	rangeCounter := 1
	if conf.Candles[1].High >= conf.Candles[0].High || conf.Candles[0].Low <= conf.Candles[1].Low {
		conf.trend = Long
		conf.Candles[1].ParabolicSAR.SAR = conf.Candles[0].Low
		conf.extremePoint = conf.Candles[0].High
	} else {
		conf.trend = Short
		conf.extremePoint = conf.Candles[0].Low
	}
	conf.Candles[1].ParabolicSAR.Trend = conf.trend
	if err := conf.validatePSAR(); err != nil {
		return err
	}

	for i, candle := range conf.Candles[rangeCounter : len(conf.Candles)-1] {
		nextSar := candle.ParabolicSAR.SAR
		if conf.trend == Long {
			if candle.High > conf.extremePoint {
				conf.extremePoint = candle.High
				conf.accelerationFactor = math.Min(conf.maxAcceleration, conf.acceleration+conf.accelerationFactor)
			}
			tmpSar := nextSar + conf.accelerationFactor*(conf.extremePoint-nextSar)
			nextSar = math.Min(math.Min(conf.Candles[rangeCounter+i-1].Low, candle.Low), tmpSar)
			if nextSar > conf.Candles[rangeCounter+i+1].Low {
				conf.trend = Short
				nextSar = conf.extremePoint
				conf.extremePoint = conf.Candles[rangeCounter+i+1].Low
				conf.accelerationFactor = conf.startAcceleration
			}
		} else if conf.trend == Short {
			if candle.Low < conf.extremePoint {
				conf.extremePoint = candle.Low
				conf.accelerationFactor = math.Min(conf.maxAcceleration, conf.acceleration+conf.accelerationFactor)
			}
			tmpSar := nextSar + conf.accelerationFactor*(conf.extremePoint-nextSar)
			nextSar = math.Max(math.Max(conf.Candles[rangeCounter+i-1].High, candle.High), tmpSar)
			if nextSar < conf.Candles[rangeCounter+i+1].High {
				conf.trend = Long
				nextSar = conf.extremePoint
				conf.extremePoint = conf.Candles[rangeCounter+i+1].High
				conf.accelerationFactor = conf.startAcceleration
			}
		}
		conf.Candles[rangeCounter+i+1].ParabolicSAR.SAR = nextSar
		conf.Candles[rangeCounter+i+1].ParabolicSAR.Trend = conf.trend
		conf.Candles[rangeCounter+i+1].ParabolicSAR.TrendFlipped = candle.ParabolicSAR.Trend != conf.trend
	}
	return nil
}

func (conf *IndicatorConfig) UpdatePSAR() error {
	lastIndex := len(conf.Candles) - 1
	if conf.Candles[lastIndex].High >= conf.Candles[lastIndex-1].High || conf.Candles[lastIndex-1].Low <= conf.Candles[lastIndex].Low {
		conf.trend = Long
	} else {
		conf.trend = Short
	}
	nextSar := conf.Candles[lastIndex-1].ParabolicSAR.SAR
	if conf.trend == Long {
		if conf.Candles[lastIndex-1].High > conf.extremePoint {
			conf.extremePoint = conf.Candles[lastIndex-1].High
			conf.accelerationFactor = math.Min(conf.maxAcceleration, conf.acceleration+conf.accelerationFactor)
		}
		tmpSar := nextSar + conf.accelerationFactor*(conf.extremePoint-nextSar)
		nextSar = math.Min(math.Min(conf.Candles[lastIndex-2].Low, conf.Candles[lastIndex-1].Low), tmpSar)
		if nextSar > conf.Candles[lastIndex].Low {
			conf.trend = Short
			nextSar = conf.extremePoint
			conf.extremePoint = conf.Candles[lastIndex].Low
			conf.accelerationFactor = conf.startAcceleration
		}
	} else if conf.trend == Short {
		if conf.Candles[lastIndex-1].Low < conf.extremePoint {
			conf.extremePoint = conf.Candles[lastIndex-1].Low
			conf.accelerationFactor = math.Min(conf.maxAcceleration, conf.acceleration+conf.accelerationFactor)
		}
		tmpSar := nextSar + conf.accelerationFactor*(conf.extremePoint-nextSar)
		nextSar = math.Max(math.Max(conf.Candles[lastIndex-2].High, conf.Candles[lastIndex-1].High), tmpSar)
		if nextSar < conf.Candles[lastIndex].High {
			conf.trend = Long
			nextSar = conf.extremePoint
			conf.extremePoint = conf.Candles[lastIndex].High
			conf.accelerationFactor = conf.startAcceleration
		}
	}
	indicatorLock.Lock()
	conf.Candles[lastIndex].ParabolicSAR.SAR = nextSar
	conf.Candles[lastIndex].ParabolicSAR.Trend = conf.trend
	conf.Candles[lastIndex].ParabolicSAR.TrendFlipped = conf.Candles[lastIndex-1].ParabolicSAR.Trend != conf.trend
	indicatorLock.Unlock()
	return nil
}

func (conf *IndicatorConfig) validatePSAR() error {
	if len(conf.Candles) < 2 {
		return errors.New("candle length must be more than 2")
	}
	return nil
}
