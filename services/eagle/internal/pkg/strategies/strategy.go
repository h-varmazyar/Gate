package strategies

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

type Strategy interface {
	RsiSignal(candles ...*models.Candle) float64
	StochasticSignal(candles ...*models.Candle) float64
	BollingerBandSignal(makerFeeRate, takerFeeRate float64, candles ...*models.Candle) float64
}
