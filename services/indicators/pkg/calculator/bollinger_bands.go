package calculator

import (
	"context"
	"github.com/google/uuid"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	"math"
)

type BollingerBands struct {
	id            uuid.UUID
	Period        int
	Deviation     int
	Source        indicatorsAPI.Source
	sma           *SMA
	lastSMA       *SMAValue
	periodCandles []*chipmunkAPI.Candle
}

type BollingerBandsValue struct {
	UpperBand float64
	LowerBand float64
	MA        float64
}

func NewBollingerBands(period, deviation int, source indicatorsAPI.Source) (*BollingerBands, error) {
	sma, err := NewSMA(period, source)
	if err != nil {
		return nil, err
	}
	return &BollingerBands{
		id:        uuid.New(),
		Period:    period,
		Deviation: deviation,
		Source:    source,
		sma:       sma,
	}, nil
}

func (conf *BollingerBands) Calculate(ctx context.Context, candles []*chipmunkAPI.Candle, values []*BollingerBandsValue) error {
	smaValues := make([]*SMAValue, len(candles))
	err := conf.sma.Calculate(ctx, candles, smaValues)
	if err != nil {
		return err
	}

	for i := conf.Period - 1; i < len(candles); i++ {
		values[i] = conf.calculateBB(candles[1+i-conf.Period:i+1], smaValues[i])
	}

	lastValue := *smaValues[len(smaValues)-1].Value
	conf.lastSMA = &SMAValue{
		Value:     &lastValue,
		TimeFrame: smaValues[len(smaValues)-1].TimeFrame,
	}
	conf.periodCandles = cloneCandles(candles[len(candles)-conf.Period:])

	return nil
}

func (conf *BollingerBands) UpdateLast(ctx context.Context, candle *chipmunkAPI.Candle, value *BollingerBandsValue) {
	if candle.Time > conf.periodCandles[len(conf.periodCandles)-1].Time {
		conf.periodCandles = conf.periodCandles[1:] //todo: must be check
	}
	conf.periodCandles[conf.Period-1] = cloneCandle(candle)

	conf.sma.UpdateLast(ctx, candle, conf.lastSMA)
	value = conf.calculateBB(conf.periodCandles, conf.lastSMA)
}

func (conf *BollingerBands) calculateBB(candles []*chipmunkAPI.Candle, smaValue *SMAValue) *BollingerBandsValue {
	variance := float64(0)
	for j := 0; j < len(candles); j++ {
		num := float64(0)
		switch conf.Source {
		case indicatorsAPI.Source_OPEN:
			num = candles[j].Open
		case indicatorsAPI.Source_HIGH:
			num = candles[j].High
		case indicatorsAPI.Source_LOW:
			num = candles[j].Low
		case indicatorsAPI.Source_CLOSE:
			num = candles[j].Close
		case indicatorsAPI.Source_OHLC4:
			num = (candles[j].Open + candles[j].High + candles[j].Low + candles[j].Close) / 4
		case indicatorsAPI.Source_HLC3:
			num = (candles[j].Low + candles[j].High + candles[j].Close) / 3
		case indicatorsAPI.Source_HL2:
			num = (candles[j].Low + candles[j].High) / 2
		}
		variance += math.Pow(*smaValue.Value-num, 2)
	}
	variance /= float64(len(candles))

	return &BollingerBandsValue{
		UpperBand: *smaValue.Value + float64(conf.Deviation)*math.Sqrt(variance),
		LowerBand: *smaValue.Value - float64(conf.Deviation)*math.Sqrt(variance),
		MA:        *smaValue.Value,
	}
}
