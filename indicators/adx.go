package indicators

import (
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/models"
	"math"
)

func (conf *Configuration) CalculateADX(candles []models.Candle) error {
	if err := conf.validateADX(len(candles)); err != nil {
		return err
	}
	for i := 1; i < len(candles); i++ {
		upMove := candles[i].High - candles[i-1].High
		downMove := candles[i-1].Low - candles[i].Low
		if upMove > downMove && upMove > 0 {
			candles[i].DmPositive = upMove
			candles[i].DmNegative = 0
		} else if downMove > 0 {
			candles[i].DmNegative = downMove
			candles[i].DmPositive = 0
		}
		method1 := candles[i].High - candles[i].Low
		method2 := math.Abs(candles[i].High - candles[i-1].Close)
		method3 := math.Abs(candles[i-1].Close - candles[i].Low)
		candles[i].ADX.TR = math.Max(method1, math.Max(method2, method3))
	}
	sumTR := float64(0)
	sumDP := float64(0)
	sumDN := float64(0)
	for i := 0; i < conf.AdxAtrLength; i++ {
		sumTR += candles[i].ADX.TR
		sumDP += candles[i].DmPositive
		sumDN += candles[i].DmNegative
	}
	candles[conf.AdxAtrLength-1].ADX.TR = sumTR
	candles[conf.AdxAtrLength-1].ADX.DmPositive = sumDP
	candles[conf.AdxAtrLength-1].ADX.DmNegative = sumDN
	candles[conf.AdxAtrLength-1].ADX.DIPositive = 100 * sumDP / sumTR
	candles[conf.AdxAtrLength-1].ADX.DINegative = 100 * sumDN / sumTR
	candles[conf.AdxAtrLength-1].ADX.DX =
		100 * math.Abs(candles[conf.AdxAtrLength-1].ADX.DIPositive-candles[conf.AdxAtrLength-1].ADX.DINegative) /
			(candles[conf.AdxAtrLength-1].ADX.DIPositive + candles[conf.AdxAtrLength-1].ADX.DINegative)
	for i := conf.AdxAtrLength; i < len(candles); i++ {
		candles[i].ADX.TR = (float64(conf.AdxAtrLength-1)/float64(conf.AdxAtrLength))*candles[i-1].ADX.TR + candles[i].ADX.TR
		candles[i].DmPositive = (float64(conf.AdxAtrLength-1)/float64(conf.AdxAtrLength))*candles[i-1].DmPositive + candles[i].DmPositive
		candles[i].DmNegative = (float64(conf.AdxAtrLength-1)/float64(conf.AdxAtrLength))*candles[i-1].DmNegative + candles[i].DmNegative
		candles[i].ADX.DIPositive = 100 * candles[i].DmPositive / candles[i].ADX.TR
		candles[i].ADX.DINegative = 100 * candles[i].DmNegative / candles[i].ADX.TR
		candles[i].ADX.DX = 100 * math.Abs(candles[i].ADX.DIPositive-candles[i].ADX.DINegative) /
			(candles[i].ADX.DIPositive + candles[i].ADX.DINegative)
	}
	sumDX := float64(0)
	for i := conf.AdxAtrLength; i < 2*conf.AdxAtrLength; i++ {
		sumDX += candles[i].ADX.DX
	}
	candles[2*conf.AdxAtrLength-1].ADX.ADX = sumDX / float64(conf.AdxAtrLength)
	for i := 2 * conf.AdxAtrLength; i < len(candles); i++ {
		candles[i].ADX.ADX = (candles[i-1].ADX.ADX*float64(conf.AdxAtrLength-1) + candles[i].DX) / float64(conf.AdxAtrLength)
	}
	return nil
}

func (conf *Configuration) UpdateADX(candles []models.Candle) {
	i := len(candles) - 1
	upMove := candles[i].High - candles[i-1].High
	downMove := candles[i-1].Low - candles[i].Low
	if upMove > downMove && upMove > 0 {
		candles[i].DmPositive = upMove
	} else {
		candles[i].DmPositive = 0
	}
	if downMove > upMove && downMove > 0 {
		candles[i].DmNegative = downMove
	} else {
		candles[i].DmNegative = 0
	}
	conf.UpdateATR(candles)
	sumTR := float64((conf.AdxAtrLength-1)/conf.AdxAtrLength)*candles[i-1].ADX.TR + candles[i].ADX.TR
	sumDP := float64((conf.AdxAtrLength-1)/conf.AdxAtrLength)*candles[i-1].DmNegative + candles[i].DmPositive
	sumDN := float64((conf.AdxAtrLength-1)/conf.AdxAtrLength)*candles[i-1].DmNegative + candles[i].DmNegative
	candles[i].ADX.DIPositive = 100 * sumDP / sumTR
	candles[i].ADX.DINegative = 100 * sumDN / sumTR
	candles[i].ADX.DX = 100 * math.Abs(candles[i].ADX.DIPositive-candles[i].ADX.DINegative) /
		(candles[i].ADX.DIPositive + candles[i].ADX.DINegative)
	candles[i].ADX.ADX = float64((conf.AdxAtrLength-1)/conf.AdxAtrLength)*candles[i-1].ADX.ADX + candles[i].DX
}

func (conf *Configuration) validateADX(length int) error {
	if length < conf.AdxAtrLength*2 {
		return errors.New(fmt.Sprintf("candles length must be bigger or equal than %d", conf.AdxAtrLength*2))
	}
	return nil
}
