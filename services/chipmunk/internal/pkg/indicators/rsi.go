package indicators

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/buffer"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/repository"
	"math"
)

type rsi struct {
	basicConfig
	Length int
}

func NewRSI(length int, marketName string) *rsi {
	return &rsi{
		basicConfig: basicConfig{
			MarketName: marketName,
			id:         uuid.New(),
		},
		Length: length,
	}
}

func (conf *rsi) GetID() string {
	return conf.id.String()
}

func (conf *rsi) Calculate(candles []*repository.Candle, response interface{}) error {
	if err := conf.validateRSI(len(candles)); err != nil {
		return err
	}
	values := make([]*RSIResponse, len(candles))

	{
		gain := float64(0)
		loss := float64(0)
		for i := 1; i <= conf.Length; i++ {
			if change := candles[i].Close - candles[i-1].Close; change > 0 {
				gain += change
			} else {
				loss += change
			}
		}
		loss = math.Abs(loss)
		rs := gain / loss
		values[conf.Length] = &RSIResponse{
			Gain: gain / float64(conf.Length),
			Loss: loss / float64(conf.Length),
			RSI:  100 - (100 / (1 + rs)),
		}
	}

	for i := conf.Length + 1; i < len(candles); i++ {
		gain := float64(0)
		loss := float64(0)
		if change := candles[i].Close - candles[i-1].Close; change > 0 {
			gain = change
		} else {
			loss = change * -1
		}
		rs := gain / loss
		values[i] = &RSIResponse{
			Gain: (values[i-1].Gain*(float64(conf.Length-1)) + gain) / float64(conf.Length),
			Loss: (values[i-1].Loss*(float64(conf.Length-1)) + loss) / float64(conf.Length),
			RSI:  100 - (100 / (1 + rs)),
		}
	}
	response = interface{}(values)
	return nil
}

func (conf *rsi) Update() interface{} {
	candles := buffer.Markets.GetLastNCandles(conf.MarketName, 2)
	values := buffer.Markets.GetLastNIndicatorValue(conf.MarketName, conf.GetID(), 2)

	gain := float64(0)
	loss := float64(0)
	change := candles[1].Close - candles[0].Close

	if change > 0 {
		gain = change
	} else {
		loss = change * -1
	}

	return &RSIResponse{
		Gain: (values[0].(RSIResponse).Gain*(float64(conf.Length-1)) + gain) / float64(conf.Length),
		Loss: (values[0].(RSIResponse).Loss*(float64(conf.Length-1)) + loss) / float64(conf.Length),
		RSI:  100 - (100 / (1 + gain/loss)),
	}
}

func (conf *rsi) validateRSI(length int) error {
	if length-1 < conf.Length {
		return errors.New("candles length must bigger than indicator period length")
	}
	return nil
}
