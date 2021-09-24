package indicators

import (
	"errors"
	"fmt"
	"math"
)

func (conf *Configuration) CalculateADX() error {
	if err := conf.validateADX(); err != nil {
		return err
	}
	for i := 1; i < len(conf.Candles); i++ {
		upMove := conf.Candles[i].High - conf.Candles[i-1].High
		downMove := conf.Candles[i-1].Low - conf.Candles[i].Low
		indicatorLock.Lock()
		if upMove > downMove && upMove > 0 {
			conf.Candles[i].DmPositive = upMove
			conf.Candles[i].DmNegative = 0
		} else if downMove > 0 {
			conf.Candles[i].DmNegative = downMove
			conf.Candles[i].DmPositive = 0
		}
		method1 := conf.Candles[i].High - conf.Candles[i].Low
		method2 := math.Abs(conf.Candles[i].High - conf.Candles[i-1].Close)
		method3 := math.Abs(conf.Candles[i-1].Close - conf.Candles[i].Low)
		indicatorLock.Lock()
		conf.Candles[i].ADX.TR = math.Max(method1, math.Max(method2, method3))
		indicatorLock.Unlock()
	}
	sumTR := float64(0)
	sumDP := float64(0)
	sumDN := float64(0)
	for i := 0; i < conf.AdxAtrLength; i++ {
		sumTR += conf.Candles[i].ADX.TR
		sumDP += conf.Candles[i].DmPositive
		sumDN += conf.Candles[i].DmNegative
	}
	indicatorLock.Lock()
	conf.Candles[conf.AdxAtrLength-1].ADX.TR = sumTR
	conf.Candles[conf.AdxAtrLength-1].ADX.DmPositive = sumDP
	conf.Candles[conf.AdxAtrLength-1].ADX.DmNegative = sumDN
	conf.Candles[conf.AdxAtrLength-1].ADX.DIPositive = 100 * sumDP / sumTR
	conf.Candles[conf.AdxAtrLength-1].ADX.DINegative = 100 * sumDN / sumTR
	conf.Candles[conf.AdxAtrLength-1].ADX.DX =
		100 * math.Abs(conf.Candles[conf.AdxAtrLength-1].ADX.DIPositive-conf.Candles[conf.AdxAtrLength-1].ADX.DINegative) /
			(conf.Candles[conf.AdxAtrLength-1].ADX.DIPositive + conf.Candles[conf.AdxAtrLength-1].ADX.DINegative)
	for i := conf.AdxAtrLength; i < len(conf.Candles); i++ {
		conf.Candles[i].ADX.TR = (float64(conf.AdxAtrLength-1)/float64(conf.AdxAtrLength))*conf.Candles[i-1].ADX.TR + conf.Candles[i].ADX.TR
		conf.Candles[i].DmPositive = (float64(conf.AdxAtrLength-1)/float64(conf.AdxAtrLength))*conf.Candles[i-1].DmPositive + conf.Candles[i].DmPositive
		conf.Candles[i].DmNegative = (float64(conf.AdxAtrLength-1)/float64(conf.AdxAtrLength))*conf.Candles[i-1].DmNegative + conf.Candles[i].DmNegative
		conf.Candles[i].ADX.DIPositive = 100 * conf.Candles[i].DmPositive / conf.Candles[i].ADX.TR
		conf.Candles[i].ADX.DINegative = 100 * conf.Candles[i].DmNegative / conf.Candles[i].ADX.TR
		conf.Candles[i].ADX.DX = 100 * math.Abs(conf.Candles[i].ADX.DIPositive-conf.Candles[i].ADX.DINegative) /
			(conf.Candles[i].ADX.DIPositive + conf.Candles[i].ADX.DINegative)
		indicatorLock.Unlock()
	}
	sumDX := float64(0)
	for i := conf.AdxAtrLength; i < 2*conf.AdxAtrLength; i++ {
		sumDX += conf.Candles[i].ADX.DX
	}
	indicatorLock.Lock()
	conf.Candles[2*conf.AdxAtrLength-1].ADX.ADX = sumDX / float64(conf.AdxAtrLength)
	for i := 2 * conf.AdxAtrLength; i < len(conf.Candles); i++ {
		conf.Candles[i].ADX.ADX = (conf.Candles[i-1].ADX.ADX*float64(conf.AdxAtrLength-1) + conf.Candles[i].DX) / float64(conf.AdxAtrLength)
	}
	indicatorLock.Unlock()
	return nil
}

func (conf *Configuration) UpdateADX() {
	i := len(conf.Candles) - 1
	indicatorLock.Lock()
	upMove := conf.Candles[i].High - conf.Candles[i-1].High
	downMove := conf.Candles[i-1].Low - conf.Candles[i].Low
	if upMove > downMove && upMove > 0 {
		conf.Candles[i].DmPositive = upMove
	} else {
		conf.Candles[i].DmPositive = 0
	}
	if downMove > upMove && downMove > 0 {
		conf.Candles[i].DmNegative = downMove
	} else {
		conf.Candles[i].DmNegative = 0
	}
	indicatorLock.Unlock()
	conf.UpdateATR()
	sumTR := float64((conf.AdxAtrLength-1)/conf.AdxAtrLength)*conf.Candles[i-1].ADX.TR + conf.Candles[i].ADX.TR
	sumDP := float64((conf.AdxAtrLength-1)/conf.AdxAtrLength)*conf.Candles[i-1].DmNegative + conf.Candles[i].DmPositive
	sumDN := float64((conf.AdxAtrLength-1)/conf.AdxAtrLength)*conf.Candles[i-1].DmNegative + conf.Candles[i].DmNegative
	indicatorLock.Lock()
	conf.Candles[i].ADX.DIPositive = 100 * sumDP / sumTR
	conf.Candles[i].ADX.DINegative = 100 * sumDN / sumTR
	conf.Candles[i].ADX.DX = 100 * math.Abs(conf.Candles[i].ADX.DIPositive-conf.Candles[i].ADX.DINegative) /
		(conf.Candles[i].ADX.DIPositive + conf.Candles[i].ADX.DINegative)
	conf.Candles[i].ADX.ADX = float64((conf.AdxAtrLength-1)/conf.AdxAtrLength)*conf.Candles[i-1].ADX.ADX + conf.Candles[i].DX
	indicatorLock.Unlock()
}

func (conf *Configuration) validateADX() error {
	if len(conf.Candles) < conf.AdxAtrLength*2 {
		return errors.New(fmt.Sprintf("candles length must be bigger or equal than %d", conf.AdxAtrLength*2))
	}
	return nil
}
