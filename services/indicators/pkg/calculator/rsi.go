package calculator

import (
	"context"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/entities"
)

type RSI struct {
	id            uint
	Period        int
	Source        entities.IndicatorSource
	Market        *chipmunkAPI.Market
	Resolution    *chipmunkAPI.Resolution
	lastValue     *indicatorsAPI.RSIValue
	lastGain      float64
	lastLoss      float64
	lastTimeframe int64
	lastCandle    *chipmunkAPI.Candle
}

func NewRSI(id uint, configs *entities.RsiConfigs, market *chipmunkAPI.Market, resolution *chipmunkAPI.Resolution) (*RSI, error) {
	return &RSI{
		id:         id,
		Period:     configs.Period,
		Source:     configs.Source,
		Market:     market,
		Resolution: resolution,
	}, nil
}

func (conf *RSI) GetMarket() *chipmunkAPI.Market {
	return conf.Market
}

func (conf *RSI) GetResolution() *chipmunkAPI.Resolution {
	return conf.Resolution
}

func (conf *RSI) GetId() uint {
	return conf.id
}

func (conf *RSI) Calculate(_ context.Context, candles []*chipmunkAPI.Candle) (*indicatorsAPI.IndicatorValues, error) {
	values := &indicatorsAPI.IndicatorValues{
		Values: make([]*indicatorsAPI.IndicatorValue, len(candles)),
	}

	//first RSI
	{
		for i := 1; i <= conf.Period; i++ {
			if change := candles[i].Close - candles[i-1].Close; change > 0 {
				conf.lastGain += change
			} else {
				conf.lastLoss += change
			}
		}
		conf.lastLoss *= -1
		conf.lastGain = conf.lastGain / float64(conf.Period)
		conf.lastLoss = conf.lastLoss / float64(conf.Period)
		rs := conf.lastGain / conf.lastLoss
		rsiValue := 100 - (100 / (1 + rs))
		//gain:      gain,
		//	loss:      loss,
		//	Value:     &rsiValue,
		//	TimeFrame: time.Unix(candles[conf.Period].Time, 0),
		values.Values[conf.Period].Time = candles[conf.Period].Time
		values.Values[conf.Period].Value = &indicatorsAPI.IndicatorValue_RSI{
			RSI: &indicatorsAPI.RSIValue{
				Rsi: rsiValue,
			},
		}

		conf.lastValue = &indicatorsAPI.RSIValue{Rsi: rsiValue}
	}

	for i := 0; i < conf.Period; i++ {
		values.Values[i].Time = candles[i].Time
		values.Values[i].Value = nil
	}

	for i := conf.Period + 1; i < len(candles); i++ {
		conf.lastCandle = cloneCandle(candles[i-1])
		value := conf.calculateRSIValue(candles[i])

		values.Values[i].Time = candles[i].Time
		values.Values[i].Value = &indicatorsAPI.IndicatorValue_RSI{RSI: &indicatorsAPI.RSIValue{Rsi: value}}
	}
	return values, nil
}

func (conf *RSI) UpdateLast(_ context.Context, candle *chipmunkAPI.Candle) *indicatorsAPI.IndicatorValue {
	value := conf.calculateRSIValue(candle)

	conf.lastCandle = cloneCandle(candle)

	return &indicatorsAPI.IndicatorValue{
		Time:  candle.Time,
		Value: &indicatorsAPI.IndicatorValue_RSI{RSI: &indicatorsAPI.RSIValue{Rsi: value}},
	}
}

func (conf *RSI) calculateRSIValue(candle *chipmunkAPI.Candle) float64 {
	gain, loss := float64(0), float64(0)
	if change := candle.Close - conf.lastCandle.Close; change > 0 {
		gain = change
	} else {
		loss = change
	}
	conf.lastGain = (conf.lastGain*float64(conf.Period-1) + gain) / float64(conf.Period)
	conf.lastLoss = (conf.lastLoss*float64(conf.Period-1) - loss) / float64(conf.Period)
	rs := conf.lastGain / conf.lastLoss
	rsiValue := 100 - (100 / (1 + rs))

	conf.lastValue = &indicatorsAPI.RSIValue{Rsi: rsiValue}

	return rsiValue
}
