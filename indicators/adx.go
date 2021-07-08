package indicators

import (
	"errors"
	"fmt"
	"math"
)

func (conf *IndicatorConfig) CalculateADX() error {
	rangeCounter := conf.Length
	dmCounter := 1
	adxCounter := (conf.Length * 2) - 1
	if err := conf.validateADX(); err != nil {
		return err
	}
	for i, candle := range conf.Candles[dmCounter:] {
		candle.ADX.DmPositive = candle.High - conf.Candles[dmCounter-i-1].High
		candle.ADX.DmNegative = conf.Candles[dmCounter-i-1].Low - candle.Low
		if candle.ADX.DmPositive > candle.ADX.DmNegative {
			candle.ADX.DmNegative = 0
		} else {
			candle.ADX.DmPositive = 0
		}
		candle.ADX.TR = math.Max(candle.High-candle.Low, math.Max(candle.High-conf.Candles[dmCounter-i-1].Close, conf.Candles[dmCounter-i-1].Close-candle.Low))
	}
	for i, candle := range conf.Candles[rangeCounter:] {
		smoothedDmPositive := float64(0)
		smoothedDmNegative := float64(0)
		smoothedTR := float64(0)
		for _, innerCandle := range conf.Candles[rangeCounter+i+1-conf.Length : rangeCounter+i+1] {
			smoothedDmPositive += innerCandle.ADX.DmPositive
			smoothedDmNegative += innerCandle.ADX.DmNegative
			smoothedTR += innerCandle.ADX.TR
		}
		smoothedDmPositive = smoothedDmPositive - (smoothedDmPositive / float64(conf.Length)) + candle.ADX.DmPositive
		smoothedDmNegative = smoothedDmNegative - (smoothedDmNegative / float64(conf.Length)) + candle.ADX.DmNegative
		smoothedTR = smoothedTR - (smoothedTR / float64(conf.Length)) + candle.ADX.TR
		candle.ADX.DIPositive = 100 * smoothedDmPositive / smoothedTR
		candle.ADX.DINegative = 100 * smoothedDmNegative / smoothedTR
		candle.ADX.DX = 100 * math.Abs(candle.ADX.DIPositive-candle.ADX.DINegative) / (candle.ADX.DIPositive + candle.ADX.DINegative)
	}
	sum := float64(0)
	for _, candle := range conf.Candles[rangeCounter:adxCounter] {
		sum += candle.ADX.DX
	}
	conf.Candles[adxCounter-1].ADX.ADX = sum / float64(adxCounter-rangeCounter)

	for i, candle := range conf.Candles[adxCounter:] {
		candle.ADX.ADX = (float64(conf.Length-1)*conf.Candles[adxCounter-i-1].ADX.ADX + candle.ADX.DX) / float64(conf.Length)
	}
	return nil
}

func (conf *IndicatorConfig) UpdateADX() {
	lastIndex := len(conf.Candles) - 1
	length := float64(conf.Length)
	lastCandle := conf.Candles[lastIndex]
	preLastCandle := conf.Candles[lastIndex-1]

	lastCandle.ADX.DmPositive = lastCandle.High - preLastCandle.High
	lastCandle.ADX.DmNegative = preLastCandle.Low - lastCandle.Low
	if lastCandle.ADX.DmPositive > lastCandle.ADX.DmNegative {
		lastCandle.ADX.DmNegative = 0
	} else {
		lastCandle.ADX.DmPositive = 0
	}
	lastCandle.ADX.TR = math.Max(lastCandle.High-lastCandle.Low,
		math.Max(lastCandle.High-preLastCandle.Close, preLastCandle.Close-lastCandle.Low))

	smoothedDmPositive := ((preLastCandle.ADX.DmPositive * (length - 1)) + lastCandle.ADX.DmPositive) / length
	smoothedDmNegative := ((preLastCandle.ADX.DmNegative * (length - 1)) + lastCandle.ADX.DmNegative) / length
	smoothedTR := ((preLastCandle.ADX.TR * (length - 1)) + lastCandle.ADX.TR) / length
	smoothedDmPositive = smoothedDmPositive - (smoothedDmPositive / length) + lastCandle.ADX.DmPositive
	smoothedDmNegative = smoothedDmNegative - (smoothedDmNegative / length) + lastCandle.ADX.DmNegative
	smoothedTR = smoothedTR - (smoothedTR / length) + lastCandle.ADX.TR

	lastCandle.ADX.DIPositive = 100 * smoothedDmPositive / smoothedTR
	lastCandle.ADX.DINegative = 100 * smoothedDmNegative / smoothedTR
	lastCandle.ADX.DX = 100 * math.Abs(lastCandle.ADX.DIPositive-lastCandle.ADX.DINegative) / (lastCandle.ADX.DIPositive + lastCandle.ADX.DINegative)

	lastCandle.ADX.ADX = ((length-1)*preLastCandle.ADX.ADX + lastCandle.ADX.DX) / length

	indicatorLock.Lock()
	conf.Candles[lastIndex] = lastCandle
	indicatorLock.Unlock()
}

func (conf *IndicatorConfig) validateADX() error {
	if len(conf.Candles) < conf.Length*2 {
		return errors.New(fmt.Sprintf("candles length must be bigger or equal than %d", conf.Length*2))
	}
	return nil
}
