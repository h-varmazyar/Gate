package indicators

import (
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/models"
	"math"
)

func (conf *Configuration) validateATR(length int) error {
	if length <= conf.AdxAtrLength {
		return errors.New(fmt.Sprintf("candles length must be bigger or equal than %d", conf.AdxAtrLength))
	}
	return nil
}

func (conf *Configuration) CalculateATR(candles []models.Candle) error {
	if err := conf.validateATR(len(candles)); err != nil {
		return err
	}
	for i := 1; i < len(candles); i++ {
		method1 := candles[i].High - candles[i].Low
		method2 := math.Abs(candles[i].High - candles[i-1].Close)
		method3 := math.Abs(candles[i-1].Close - candles[i].Low)
		candles[i].ATR.TR = math.Max(method1, math.Max(method2, method3))
	}

	sumTR := float64(0)
	for i := 1; i <= conf.AdxAtrLength; i++ {
		sumTR += candles[i].ATR.TR
	}
	candles[conf.AdxAtrLength].ATR.ATR = sumTR / float64(conf.AdxAtrLength)
	for i := conf.AdxAtrLength + 1; i < len(candles); i++ {
		candles[i].ATR.ATR = (candles[i-1].ATR.ATR*float64(conf.AdxAtrLength-1) + candles[i].ATR.TR) / float64(conf.AdxAtrLength)
	}
	return nil
}

func (conf *Configuration) UpdateATR(candles []models.Candle) {
	i := len(candles) - 1
	method1 := candles[i].High - candles[i].Low
	method2 := math.Abs(candles[i].High - candles[i-1].Close)
	method3 := math.Abs(candles[i-1].Close - candles[i].Low)
	candles[i].ATR.TR = math.Max(method1, math.Max(method2, method3))
	candles[i].ATR.ATR = (candles[i-1].ATR.ATR*float64(conf.AdxAtrLength-1) + candles[i].ATR.TR) / float64(conf.AdxAtrLength)
}
