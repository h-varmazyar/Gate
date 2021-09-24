package indicators

import (
	"errors"
	"fmt"
	"math"
)

func (conf *Configuration) validateATR() error {
	if len(conf.Candles) <= conf.AdxAtrLength {
		return errors.New(fmt.Sprintf("candles length must be bigger or equal than %d", conf.AdxAtrLength))
	}
	return nil
}

func (conf *Configuration) CalculateATR() error {
	if err := conf.validateATR(); err != nil {
		return err
	}
	indicatorLock.Lock()
	indicatorLock.Unlock()
	for i := 1; i < len(conf.Candles); i++ {
		method1 := conf.Candles[i].High - conf.Candles[i].Low
		method2 := math.Abs(conf.Candles[i].High - conf.Candles[i-1].Close)
		method3 := math.Abs(conf.Candles[i-1].Close - conf.Candles[i].Low)
		indicatorLock.Lock()
		conf.Candles[i].ATR.TR = math.Max(method1, math.Max(method2, method3))
		indicatorLock.Unlock()
	}

	sumTR := float64(0)
	for i := 1; i <= conf.AdxAtrLength; i++ {
		sumTR += conf.Candles[i].ATR.TR
	}
	indicatorLock.Lock()
	conf.Candles[conf.AdxAtrLength].ATR.ATR = sumTR / float64(conf.AdxAtrLength)
	for i := conf.AdxAtrLength + 1; i < len(conf.Candles); i++ {
		conf.Candles[i].ATR.ATR = (conf.Candles[i-1].ATR.ATR*float64(conf.AdxAtrLength-1) + conf.Candles[i].ATR.TR) / float64(conf.AdxAtrLength)
	}
	indicatorLock.Unlock()
	return nil
}

func (conf *Configuration) UpdateATR() {
	i := len(conf.Candles) - 1
	method1 := conf.Candles[i].High - conf.Candles[i].Low
	method2 := math.Abs(conf.Candles[i].High - conf.Candles[i-1].Close)
	method3 := math.Abs(conf.Candles[i-1].Close - conf.Candles[i].Low)
	indicatorLock.Lock()
	conf.Candles[i].ATR.TR = math.Max(method1, math.Max(method2, method3))
	indicatorLock.Unlock()
	conf.Candles[i].ATR.ATR = (conf.Candles[i-1].ATR.ATR*float64(conf.AdxAtrLength-1) + conf.Candles[i].ATR.TR) / float64(conf.AdxAtrLength)
}
