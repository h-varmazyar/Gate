package automatedStrategy

import "github.com/mrNobody95/Gate/services/eagle/internal/pkg/models"

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 09.12.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type Automated struct {
	MinGainPercentage float64
}

func (s *Automated) RsiSignal(candles ...*models.Candle) float64 {
	if len(candles) != 2 {
		return 0
	}
	if candles[0].RSI.RSI < 30 && candles[1].RSI.RSI >= 30 {
		return 1
	}
	return 0
}

func (s *Automated) StochasticSignal(candles ...*models.Candle) float64 {
	if len(candles) != 2 {
		return 0
	}
	if candles[1].IndexD > 20 || candles[0].IndexK > 20 {
		return 0
	}
	if candles[0].IndexK < candles[1].IndexK {
		return 1
	}
	return 0
}

func (s *Automated) BollingerBandSignal(makerFeeRate, takerFeeRate float64, candles ...*models.Candle) float64 {
	if len(candles) != 2 {
		return 0
	}
	if candles[0].Low > candles[0].LowerBand {
		return 0
	}
	price := candles[1].Close * (1 + makerFeeRate/100) * (1 + s.MinGainPercentage/100) * (1 + takerFeeRate/100)
	if price < candles[1].UpperBand {
		return 1
	}
	return 0
}
