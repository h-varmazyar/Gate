package indicators

import (
	"errors"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/models"
	"math"
)

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
* Date: 03.12.21
* Github: https://github.com/h-varmazyar
* Email: hossein.varmazyar@yahoo.com
**/

func (conf *Configuration) CalculateRSI(candles []*models.Candle) error {
	if err := conf.validateRSI(len(candles)); err != nil {
		return err
	}
	conf.firstRsi(candles)
	for i := conf.RsiLength + 1; i < len(candles); i++ {
		conf.UpdateRSI(candles[:i+1])
	}
	return nil
}

func (conf *Configuration) UpdateRSI(candles []*models.Candle) {
	i := len(candles) - 1

	gain := float64(0)
	loss := float64(0)
	change := candles[i].Close - candles[i-1].Close

	if change > 0 {
		gain = change
	} else {
		loss = change * -1
	}

	gain = (candles[i-1].RSI.Gain*(float64(conf.RsiLength-1)) + gain) / float64(conf.RsiLength)
	loss = (candles[i-1].RSI.Loss*(float64(conf.RsiLength-1)) + loss) / float64(conf.RsiLength)

	conf.lock.Lock()
	candles[i].RSI.Gain = gain
	candles[i].RSI.Loss = loss
	candles[i].RSI.RSI = 100 - (100 / (1 + gain/loss))
	conf.lock.Unlock()
}

func (conf *Configuration) firstRsi(candles []*models.Candle) {
	gain := float64(0)
	loss := float64(0)

	for i := 1; i <= conf.RsiLength; i++ {
		change := candles[i].Close - candles[i-1].Close
		if change > 0 {
			gain += change
		} else {
			loss += change
		}
	}

	loss = math.Abs(loss)
	rs := gain / loss

	candles[conf.RsiLength].RSI.Gain = gain / float64(conf.RsiLength)
	candles[conf.RsiLength].RSI.Loss = loss / float64(conf.RsiLength)
	candles[conf.RsiLength].RSI.RSI = 100 - (100 / (1 + rs))
}

func (conf *Configuration) validateRSI(length int) error {
	if length-1 < conf.RsiLength {
		return errors.New("candles length must bigger than indicator period length")
	}
	return nil
}
