package indicators

import (
	"errors"
	"math"
)

func (conf *Configuration) CalculateRSI() error {
	if err := conf.validateRSI(); err != nil {
		return err
	}
	conf.firstRsi()
	for i := conf.RsiLength + 1; i < len(conf.Candles); i++ {
		conf.smoothedRs(i)
	}
	return nil
}

func (conf *Configuration) UpdateRSI() {
	conf.smoothedRs(len(conf.Candles) - 1)
}

func (conf *Configuration) firstRsi() {
	gain := float64(0)
	loss := float64(0)

	for i := 1; i <= conf.RsiLength; i++ {
		change := conf.Candles[i].Close - conf.Candles[i-1].Close
		if change > 0 {
			gain += change
		} else {
			loss += change
		}
	}

	loss = math.Abs(loss)
	rs := gain / loss

	indicatorLock.Lock()
	conf.Candles[conf.RsiLength].RSI.Gain = gain / float64(conf.RsiLength)
	conf.Candles[conf.RsiLength].RSI.Loss = loss / float64(conf.RsiLength)
	conf.Candles[conf.RsiLength].RSI.RSI = 100 - (100 / (1 + rs))
	indicatorLock.Unlock()
}

func (conf *Configuration) smoothedRs(index int) {
	gain := float64(0)
	loss := float64(0)
	change := conf.Candles[index].Close - conf.Candles[index-1].Close

	if change > 0 {
		gain = change
	} else {
		loss = change * -1
	}

	gain = (conf.Candles[index-1].RSI.Gain*(float64(conf.RsiLength-1)) + gain) / float64(conf.RsiLength)
	loss = (conf.Candles[index-1].RSI.Loss*(float64(conf.RsiLength-1)) + loss) / float64(conf.RsiLength)

	indicatorLock.Lock()
	conf.Candles[index].RSI.Gain = gain
	conf.Candles[index].RSI.Loss = loss
	conf.Candles[index].RSI.RSI = 100 - (100 / (1 + gain/loss))
	indicatorLock.Unlock()
}

func (conf *Configuration) validateRSI() error {
	if len(conf.Candles)-1 < conf.RsiLength {
		return errors.New("candles length must bigger than indicator period length")
	}
	return nil
}
