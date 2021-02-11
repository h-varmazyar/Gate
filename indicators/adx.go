package indicators

import (
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/models"
	"math"
)

type adxConfig struct {
	*basicConfig
}

func NewAdxConfig(length int) *adxConfig {
	return &adxConfig{&basicConfig{Length: length}}
}

func (conf *adxConfig) CalculateADX(candles []models.Candle, appendCandles bool) error {
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
	if err := conf.validate(); err != nil {
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

func (conf *adxConfig) validate() error {
	if len(conf.Candles) < conf.Length*2 {
		return errors.New(fmt.Sprintf("candles length must be bigger or equal than %d", conf.Length*2))
	}
	return nil
}
