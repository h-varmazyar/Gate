package indicators

import (
	"errors"
)

func (conf *IndicatorConfig) CalculateRSI(firstOfSeries bool) error {
	rangeCounter := 1
	if err := conf.validateRSI(firstOfSeries); err != nil {
		return err
	}
	if firstOfSeries {
		conf.firstRsi()
		rangeCounter = conf.Length + 1
	}
	for i, candle := range conf.Candles[rangeCounter:] {
		candle.RSI.RSI = 100 - (100 / (1 + conf.smoothedRs(rangeCounter+i)))
	}
	return nil
}

func (conf *IndicatorConfig) UpdateRSI() error {
	lastIndex := len(conf.Candles) - 1
	indicatorLock.Lock()
	conf.Candles[lastIndex].RSI.RSI = 100 - (100 / (1 + conf.smoothedRs(lastIndex)))
	indicatorLock.Unlock()
	return nil
}

func (conf *IndicatorConfig) firstRsi() {
	gain := float64(0)
	loss := float64(0)

	for i, candle := range conf.Candles[1 : conf.Length+1] {
		change := candle.Close - conf.Candles[i-1].Close
		if change > 0 {
			gain += change
		} else {
			loss += change
		}
	}

	rs := gain / loss

	conf.Candles[conf.Length].RSI.Gain = gain / float64(conf.Length)
	conf.Candles[conf.Length].RSI.Loss = loss / float64(conf.Length)
	conf.Candles[conf.Length].RSI.RSI = 100 - (100 / (1 + rs))
}

func (conf *IndicatorConfig) smoothedRs(currentIndex int) float64 {
	change := conf.Candles[currentIndex].Close - conf.Candles[currentIndex-1].Close
	indicatorLock.Lock()
	if change > 0 {
		conf.Candles[currentIndex].RSI.Gain = change
		conf.Candles[currentIndex].RSI.Loss = float64(0)
	} else {
		conf.Candles[currentIndex].RSI.Gain = float64(0)
		conf.Candles[currentIndex].RSI.Loss = change
	}
	indicatorLock.Unlock()
	gain := (conf.Candles[currentIndex-1].RSI.Gain*(float64(conf.Length-1)) + conf.Candles[currentIndex].RSI.Gain) / float64(conf.Length)
	loss := (conf.Candles[currentIndex-1].RSI.Loss*(float64(conf.Length-1)) + conf.Candles[currentIndex].RSI.Loss) / float64(conf.Length)
	return gain / loss
}

func (conf *IndicatorConfig) validateRSI(firstOfSeries bool) error {
	if firstOfSeries && len(conf.Candles)-1 < conf.Length {
		return errors.New("candles length must bigger than indicator period length")
	}
	return nil
}
