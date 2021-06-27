package indicators

import (
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/models"
	"math"
)

func (conf *IndicatorConfig) CalculateADX(candles []models.Candle, appendCandles bool) error {
	rangeCounter := conf.Length
	dmCounter := 1
	adxCounter := (conf.Length * 2) - 1
	if appendCandles {
		rangeCounter = len(conf.Candles)
		dmCounter = len(conf.Candles)
		adxCounter = len(conf.Candles)
		conf.Candles = append(conf.Candles, candles...)
	} else {
		conf.Candles = candles
	}
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
	if !appendCandles {
		sum := float64(0)
		for _, candle := range conf.Candles[rangeCounter:adxCounter] {
			sum += candle.ADX.DX
		}
		conf.Candles[adxCounter-1].ADX.ADX = sum / float64(adxCounter-rangeCounter)
	}
	for i, candle := range conf.Candles[adxCounter:] {
		candle.ADX.ADX = (float64(conf.Length-1)*conf.Candles[adxCounter-i-1].ADX.ADX + candle.ADX.DX) / float64(conf.Length)
	}
	return nil
}

func (conf *IndicatorConfig) UpdateADX() error {
	lastIndex := len(conf.Candles) - 1
	conf.Candles[lastIndex].ADX.DmPositive = conf.Candles[lastIndex].High - conf.Candles[lastIndex-1].High
	conf.Candles[lastIndex].ADX.DmNegative = conf.Candles[lastIndex-1].Low - conf.Candles[lastIndex].Low
	if conf.Candles[lastIndex].ADX.DmPositive > conf.Candles[lastIndex].ADX.DmNegative {
		conf.Candles[lastIndex].ADX.DmNegative = 0
	} else {
		conf.Candles[lastIndex].ADX.DmPositive = 0
	}
	conf.Candles[lastIndex].ADX.TR = math.Max(conf.Candles[lastIndex].High-conf.Candles[lastIndex].Low,
		math.Max(conf.Candles[lastIndex].High-conf.Candles[lastIndex-1].Close,
			conf.Candles[lastIndex-1].Close-conf.Candles[lastIndex].Low))

	smoothedDmPositive := ((conf.Candles[lastIndex-1].ADX.DmPositive * (float64(conf.Length - 1))) + conf.Candles[lastIndex].ADX.DmPositive) / float64(conf.Length)
	smoothedDmNegative := ((conf.Candles[lastIndex-1].ADX.DmNegative * (float64(conf.Length - 1))) + conf.Candles[lastIndex].ADX.DmNegative) / float64(conf.Length)
	smoothedTR := ((conf.Candles[lastIndex-1].ADX.TR * (float64(conf.Length - 1))) + conf.Candles[lastIndex].ADX.TR) / float64(conf.Length)
	smoothedDmPositive = smoothedDmPositive - (smoothedDmPositive / float64(conf.Length)) + conf.Candles[lastIndex].ADX.DmPositive
	smoothedDmNegative = smoothedDmNegative - (smoothedDmNegative / float64(conf.Length)) + conf.Candles[lastIndex].ADX.DmNegative
	smoothedTR = smoothedTR - (smoothedTR / float64(conf.Length)) + conf.Candles[lastIndex].ADX.TR

	conf.Candles[lastIndex].ADX.DIPositive = 100 * smoothedDmPositive / smoothedTR
	conf.Candles[lastIndex].ADX.DINegative = 100 * smoothedDmNegative / smoothedTR
	conf.Candles[lastIndex].ADX.DX = 100 * math.Abs(conf.Candles[lastIndex].ADX.DIPositive-conf.Candles[lastIndex].ADX.DINegative) / (conf.Candles[lastIndex].ADX.DIPositive + conf.Candles[lastIndex].ADX.DINegative)

	conf.Candles[lastIndex].ADX.ADX = (float64(conf.Length-1)*conf.Candles[lastIndex-1].ADX.ADX + conf.Candles[lastIndex].ADX.DX) / float64(conf.Length)
	return nil
}

func (conf *IndicatorConfig) validateADX() error {
	if len(conf.Candles) < conf.Length*2 {
		return errors.New(fmt.Sprintf("candles length must be bigger or equal than %d", conf.Length*2))
	}
	return nil
}
