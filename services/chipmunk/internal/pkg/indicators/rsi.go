package indicators

import (
	"errors"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
	"math"
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

	{
		gain := float64(0)
		loss := float64(0)
		for i := 1; i <= conf.length; i++ {
			if change := candles[i].Close - candles[i-1].Close; change > 0 {
				gain += change
			} else {
				loss += change
			}
		}
		loss = math.Abs(loss)
		rs := gain / loss
		candles[conf.length].RSIs[conf.id] = repository.RSIValue{
			Gain: gain / float64(conf.length),
			Loss: loss / float64(conf.length),
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
		rs := gain / loss
		candles[i].RSIs[conf.id] = repository.RSIValue{
			Gain: (candles[i-1].RSIs[conf.id].Gain*(float64(conf.length-1)) + gain) / float64(conf.length),
			Loss: (candles[i-1].RSIs[conf.id].Loss*(float64(conf.length-1)) + loss) / float64(conf.length),
			RSI:  100 - (100 / (1 + rs)),
		}
	}
	return nil
}

func (conf *rsi) Update(candles []*repository.Candle) *repository.IndicatorValue {
	gain := float64(0)
	loss := float64(0)
	change := candles[1].Close - candles[0].Close

	if change > 0 {
		gain = change
	} else {
		loss = change * -1
	}

	return &repository.IndicatorValue{RSI: &repository.RSIValue{
		Gain: (candles[0].RSIs[conf.id].Gain*(float64(conf.length-1)) + gain) / float64(conf.length),
		Loss: (candles[0].RSIs[conf.id].Loss*(float64(conf.length-1)) + loss) / float64(conf.length),
		RSI:  100 - (100 / (1 + gain/loss)),
	}}
}

func (conf *rsi) validateRSI(length int) error {
	if length-1 < conf.length {
		return errors.New("candles length must bigger than indicator period length")
	}
	return nil
}
