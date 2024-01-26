package calculator

import (
	"context"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/entity"
	"math"
)

type BollingerBands struct {
	id         uint
	Period     int
	Deviation  int
	Source     indicatorsAPI.Source
	Market     *chipmunkAPI.Market
	Resolution *chipmunkAPI.Resolution
	sma        *SMA
	//lastSMA       *indicatorsAPI.SMAValue
	periodCandles []*chipmunkAPI.Candle
}

func (conf *BollingerBands) GetMarket() *chipmunkAPI.Market {
	return conf.Market
}

func (conf *BollingerBands) GetResolution() *chipmunkAPI.Resolution {
	return conf.Resolution
}

func (conf *BollingerBands) GetId() uint {
	return conf.id
}

func NewBollingerBands(id uint, configs *entity.BollingerBandsConfigs, market *chipmunkAPI.Market, resolution *chipmunkAPI.Resolution) (*BollingerBands, error) {
	smaConfigs := &entity.SMAConfigs{
		Period: configs.Period,
		Source: configs.Source,
	}
	sma, err := NewSMA(id, smaConfigs, market, resolution)
	if err != nil {
		return nil, err
	}
	return &BollingerBands{
		id:        id,
		Period:    configs.Period,
		Deviation: configs.Deviation,
		Source:    configs.Source,
		sma:       sma,
	}, nil
}

func (conf *BollingerBands) Calculate(ctx context.Context, candles []*chipmunkAPI.Candle) (*indicatorsAPI.IndicatorValues, error) {
	smaValues, err := conf.sma.Calculate(ctx, candles)
	if err != nil {
		return nil, err
	}

	values := &indicatorsAPI.IndicatorValues{Values: make([]*indicatorsAPI.IndicatorValue, len(candles))}

	for i := 0; i < conf.Period-1; i++ {
		values.Values[i].Time = candles[i].Time
	}

	for i := conf.Period - 1; i < len(candles); i++ {
		values.Values[i].Time = candles[i].Time
		values.Values[i].Value = &indicatorsAPI.IndicatorValue_BollingerBands{
			BollingerBands: conf.calculateBB(candles[1+i-conf.Period:i+1], smaValues.GetValues()[i].GetSMA()),
		}
	}

	conf.periodCandles = cloneCandles(candles[len(candles)-conf.Period:])

	return values, nil
}

func (conf *BollingerBands) UpdateLast(ctx context.Context, candle *chipmunkAPI.Candle) *indicatorsAPI.IndicatorValue {
	if candle.Time > conf.periodCandles[len(conf.periodCandles)-1].Time {
		conf.periodCandles = conf.periodCandles[1:] //todo: must be check
	}
	conf.periodCandles[conf.Period-1] = cloneCandle(candle)

	sma := conf.sma.UpdateLast(ctx, candle)
	value := conf.calculateBB(conf.periodCandles, sma.GetSMA())

	return &indicatorsAPI.IndicatorValue{
		Time:  candle.Time,
		Value: &indicatorsAPI.IndicatorValue_BollingerBands{BollingerBands: value},
	}
}

func (conf *BollingerBands) calculateBB(candles []*chipmunkAPI.Candle, smaValue *indicatorsAPI.SMAValue) *indicatorsAPI.BollingerBandsValue {
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
		variance += math.Pow(smaValue.GetValue()-num, 2)
	}
	variance /= float64(len(candles))

	return &indicatorsAPI.BollingerBandsValue{
		UpperBand: smaValue.GetValue() + float64(conf.Deviation)*math.Sqrt(variance),
		LowerBand: smaValue.GetValue() - float64(conf.Deviation)*math.Sqrt(variance),
		MA:        smaValue.GetValue(),
	}
}
