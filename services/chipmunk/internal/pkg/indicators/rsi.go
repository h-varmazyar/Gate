package indicators

import (
	"errors"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/repository"
	"math"
)

type RSI struct {
	Length int
	Values []*RSIResponse
}

func (conf *RSI) Calculate(candles []*repository.Candle) error {
	if err := conf.validateRSI(len(candles)); err != nil {
		return err
	}
	conf.Values = make([]*RSIResponse, len(candles))
	conf.Values[conf.Length] = conf.firstRsi(candles)
	for i := conf.Length + 1; i < len(candles); i++ {
		conf.Values[i] = conf.Update(candles[:i+1])
	}
	return nil
}

func (conf *RSI) Update(candles []*repository.Candle) *RSIResponse {
	i := len(candles) - 1

	gain := float64(0)
	loss := float64(0)
	change := candles[i].Close - candles[i-1].Close

	if change > 0 {
		gain = change
	} else {
		loss = change * -1
	}

	gain = (conf.Values[i-1].Gain*(float64(conf.Length-1)) + gain) / float64(conf.Length)
	loss = (conf.Values[i-1].Loss*(float64(conf.Length-1)) + loss) / float64(conf.Length)

	//conf.lock.Lock()
	//candles[i].RSI.Gain = gain
	//candles[i].RSI.Loss = loss
	//candles[i].RSI.RSI = 100 - (100 / (1 + gain/loss))
	//conf.lock.Unlock()

	return &RSIResponse{
		Gain: gain,
		Loss: loss,
		RSI:  100 - (100 / (1 + gain/loss)),
	}
}

func (conf *RSI) firstRsi(candles []*repository.Candle) *RSIResponse {
	gain := float64(0)
	loss := float64(0)

	for i := 1; i <= conf.Length; i++ {
		change := candles[i].Close - candles[i-1].Close
		if change > 0 {
			gain += change
		} else {
			loss += change
		}
	}

	loss = math.Abs(loss)
	rs := gain / loss

	//candles[conf.Length].RSI.Gain = gain / float64(conf.Length)
	//candles[conf.Length].RSI.Loss = loss / float64(conf.Length)
	//candles[conf.Length].RSI.RSI = 100 - (100 / (1 + rs))
	return &RSIResponse{
		Gain: gain / float64(conf.Length),
		Loss: loss / float64(conf.Length),
		RSI:  100 - (100 / (1 + rs)),
	}
}

func (conf *RSI) validateRSI(length int) error {
	if length-1 < conf.Length {
		return errors.New("candles length must bigger than indicator period length")
	}
	return nil
}
