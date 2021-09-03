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
		}
		if downMove > upMove && downMove > 0 {
			conf.Candles[i].DmNegative = downMove
		}
		indicatorLock.Unlock()
	}
	//first TR[length]
	sumTR := make([]float64, len(conf.Candles))
	sumDP := make([]float64, len(conf.Candles))
	sumDN := make([]float64, len(conf.Candles))
	for i := 1; i <= conf.AdxAtrLength; i++ {
		sumTR[conf.AdxAtrLength] += conf.Candles[i].TR
		sumDP[conf.AdxAtrLength] += conf.Candles[i].DmPositive
		sumDN[conf.AdxAtrLength] += conf.Candles[i].DmNegative
	}
	indicatorLock.Lock()
	conf.Candles[conf.AdxAtrLength].ADX.DIPositive = 100 * sumDP[conf.AdxAtrLength] / sumTR[conf.AdxAtrLength]
	conf.Candles[conf.AdxAtrLength].ADX.DINegative = 100 * sumDN[conf.AdxAtrLength] / sumTR[conf.AdxAtrLength]
	conf.Candles[conf.AdxAtrLength].ADX.DX = 100 * math.Abs(
		conf.Candles[conf.AdxAtrLength].ADX.DIPositive-conf.Candles[conf.AdxAtrLength].ADX.DINegative) /
		(conf.Candles[conf.AdxAtrLength].ADX.DIPositive + conf.Candles[conf.AdxAtrLength].ADX.DINegative)
	indicatorLock.Unlock()
	for i := conf.AdxAtrLength + 1; i < len(conf.Candles); i++ {
		sumTR[i] = float64((conf.AdxAtrLength-1)/conf.AdxAtrLength)*conf.Candles[i-1].TR + conf.Candles[i].TR
		sumDP[i] = float64((conf.AdxAtrLength-1)/conf.AdxAtrLength)*conf.Candles[i-1].DmNegative + conf.Candles[i].DmPositive
		sumDN[i] = float64((conf.AdxAtrLength-1)/conf.AdxAtrLength)*conf.Candles[i-1].DmNegative + conf.Candles[i].DmNegative

		indicatorLock.Lock()
		conf.Candles[i].ADX.DIPositive = 100 * sumDP[i] / sumTR[i]
		conf.Candles[i].ADX.DINegative = 100 * sumDN[i] / sumTR[i]
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
	for i := conf.AdxAtrLength + 1; i < len(conf.Candles); i++ {
		conf.Candles[i].ADX.ADX = float64((conf.AdxAtrLength-1)/conf.AdxAtrLength)*conf.Candles[i-1].ADX.ADX + conf.Candles[i].DX
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
	sumTR := float64((conf.AdxAtrLength-1)/conf.AdxAtrLength)*conf.Candles[i-1].TR + conf.Candles[i].TR
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
