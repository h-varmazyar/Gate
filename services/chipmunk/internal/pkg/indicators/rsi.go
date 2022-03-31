package indicators

import (
	"errors"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
)

type rsi struct {
	basicConfig
}

func NewRSI(length int, marketName string) *rsi {
	return &rsi{
		basicConfig: basicConfig{
			MarketName: marketName,
			id:         uuid.New(),
			length:     length,
		},
	}
}

func (conf *rsi) GetID() uuid.UUID {
	return conf.id
}

func (conf *rsi) GetType() IndicatorType {
	return RSI
}

func (conf *rsi) GetLength() int {
	return conf.length
}

func (conf *rsi) Calculate(candles []*repository.Candle) error {
	if err := conf.validateRSI(len(candles)); err != nil {
		return err
	}

	{ //first RSI
		gain := float64(0)
		loss := float64(0)
		for i := 1; i <= conf.length; i++ {
			if change := candles[i].Close - candles[i-1].Close; change > 0 {
				gain += change
			} else {
				loss += change
			}
		}
		loss *= -1
		gain = gain / float64(conf.length)
		loss = loss / float64(conf.length)
		rs := gain / loss
		candles[conf.length].RSIs[conf.id] = &repository.RSIValue{
			Gain: gain,
			Loss: loss,
			RSI:  100 - (100 / (1 + rs)),
		}
	}

	for i := conf.length + 1; i < len(candles); i++ {
		gain := float64(0)
		loss := float64(0)
		if change := candles[i].Close - candles[i-1].Close; change > 0 {
			gain = change
		} else {
			loss = change * -1
		}
		avgGain := (candles[i-1].RSIs[conf.id].Gain*float64(conf.length-1) + gain) / float64(conf.length)
		avgLoss := (candles[i-1].RSIs[conf.id].Loss*float64(conf.length-1) + loss) / float64(conf.length)
		rs := avgGain / avgLoss
		rsiValue := 100 - (100 / (1 + rs))

		candles[i].RSIs[conf.id] = &repository.RSIValue{
			Gain: avgGain,
			Loss: avgLoss,
			RSI:  rsiValue,
		}
	}
	return nil
}

func (conf *rsi) Update(candles []*repository.Candle) *repository.IndicatorValue {
	gain, loss := float64(0), float64(0)
	last := len(candles) - 1
	if last < 1 {
		return nil
	}
	if change := candles[last].Close - candles[last-1].Close; change > 0 {
		gain = change
	} else {
		loss = change
	}

	avgGain := (candles[last-1].RSIs[conf.id].Gain*float64(conf.length-1) + gain) / float64(conf.length)
	avgLoss := (candles[last-1].RSIs[conf.id].Loss*float64(conf.length-1) - loss) / float64(conf.length)
	rs := avgGain / avgLoss
	rsiValue := 100 - (100 / (1 + rs))

	return &repository.IndicatorValue{
		RSI: &repository.RSIValue{
			Gain: avgGain,
			Loss: avgLoss,
			RSI:  rsiValue,
		}}
}

func (conf *rsi) validateRSI(length int) error {
	if length-1 < conf.length {
		return errors.New("candles length must bigger than indicator period length")
	}
	return nil
}
