package indicators

import (
	"errors"
	"github.com/mrNobody95/Gate/models"
)

type RSIConfig struct {
	*basicConfig
}

func NewRSIConfig(length int) *RSIConfig {
	return &RSIConfig{
		basicConfig: &basicConfig{
			Length: length,
		},
	}
}

func (conf *RSIConfig) CalculateRsi(candles []models.Candle, firstOfSeries, appendCandles bool) error {
	var rangeCounter int
	if appendCandles {
		rangeCounter = len(conf.Candles)
		conf.Candles = append(conf.Candles, candles...)
	} else {
		conf.Candles = candles
		rangeCounter = 1
	}
	if err := conf.validate(firstOfSeries); err != nil {
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

func (conf *RSIConfig) firstRsi() {
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

func (conf *RSIConfig) smoothedRs(currentIndex int) float64 {
	change := conf.Candles[currentIndex].Close - conf.Candles[currentIndex-1].Close
	if change > 0 {
		conf.Candles[currentIndex].RSI.Gain = change
		conf.Candles[currentIndex].RSI.Loss = float64(0)
	} else {
		conf.Candles[currentIndex].RSI.Gain = float64(0)
		conf.Candles[currentIndex].RSI.Loss = change
	}
	gain := (conf.Candles[currentIndex-1].RSI.Gain*(float64(conf.Length-1)) + conf.Candles[currentIndex].RSI.Gain) / float64(conf.Length)
	loss := (conf.Candles[currentIndex-1].RSI.Loss*(float64(conf.Length-1)) + conf.Candles[currentIndex].RSI.Loss) / float64(conf.Length)
	return gain / loss
}

func (conf *RSIConfig) validate(firstOfSeries bool) error {
	if firstOfSeries && len(conf.Candles)-1 < conf.Length {
		return errors.New("candles length must bigger than indicator period length")
	}
	return nil
}
