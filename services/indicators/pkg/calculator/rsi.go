package calculator

import (
	"context"
	"github.com/google/uuid"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	"time"
)

type RSIValue struct {
	gain      float64
	loss      float64
	Value     *float64
	TimeFrame time.Time
}

type RSI struct {
	id         uuid.UUID
	Period     int
	Source     indicatorsAPI.Source
	lastValue  RSIValue
	lastCandle *chipmunkAPI.Candle
}

func NewRSI(period int, source indicatorsAPI.Source) (*RSI, error) {
	return &RSI{
		id:     uuid.New(),
		Period: period,
		Source: source,
	}, nil
}

func (conf *RSI) Calculate(_ context.Context, candles []*chipmunkAPI.Candle, values []*RSIValue) error {
	{ //first RSI
		gain := float64(0)
		loss := float64(0)
		for i := 1; i <= conf.Period; i++ {
			if change := candles[i].Close - candles[i-1].Close; change > 0 {
				gain += change
			} else {
				loss += change
			}
		}
		loss *= -1
		gain = gain / float64(conf.Period)
		loss = loss / float64(conf.Period)
		rs := gain / loss
		rsiValue := 100 - (100 / (1 + rs))
		values[conf.Period] = &RSIValue{
			gain:      gain,
			loss:      loss,
			Value:     &rsiValue,
			TimeFrame: time.Unix(candles[conf.Period].Time, 0),
		}
		conf.lastValue = RSIValue{
			gain:      gain,
			loss:      loss,
			Value:     &rsiValue,
			TimeFrame: time.Unix(candles[conf.Period].Time, 0),
		}
	}

	for i := 0; i < conf.Period; i++ {
		values[i] = &RSIValue{
			gain:      0,
			loss:      0,
			Value:     nil,
			TimeFrame: time.Unix(candles[i].Time, 0),
		}
	}

	for i := conf.Period + 1; i < len(candles); i++ {
		conf.lastCandle = cloneCandle(candles[i-1])
		values[i] = conf.calculateRSIValue(candles[i])

		lastValue := values[i]
		conf.lastValue = *lastValue
	}
	return nil
}

func (conf *RSI) UpdateLast(_ context.Context, candle *chipmunkAPI.Candle, value *RSIValue) {
	value = conf.calculateRSIValue(candle)
	lastValue := value

	conf.lastCandle = cloneCandle(candle)
	conf.lastValue = *lastValue
}

func (conf *RSI) calculateRSIValue(candle *chipmunkAPI.Candle) *RSIValue {
	gain, loss := float64(0), float64(0)
	if change := candle.Close - conf.lastCandle.Close; change > 0 {
		gain = change
	} else {
		loss = change
	}
	avgGain := (conf.lastValue.gain*float64(conf.Period-1) + gain) / float64(conf.Period)
	avgLoss := (conf.lastValue.loss*float64(conf.Period-1) - loss) / float64(conf.Period)
	rs := avgGain / avgLoss
	rsiValue := 100 - (100 / (1 + rs))

	return &RSIValue{
		gain:  avgGain,
		loss:  avgLoss,
		Value: &rsiValue,
	}
}
