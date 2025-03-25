package calculator

import (
	"context"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/entities"
	"math"
	"sync"
)

type BollingerBands struct {
	base
	Period        int
	Deviation     int
	Source        entities.IndicatorSource
	sma           *SMA
	periodCandles []Candle
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

func NewBollingerBands(id uint, configs *entities.BollingerBandsConfigs, market *chipmunkAPI.Market, resolution *chipmunkAPI.Resolution) (*BollingerBands, error) {
	smaConfigs := &entities.SMAConfigs{
		Period: configs.Period,
		Source: configs.Source,
	}
	sma, err := NewSMA(id, smaConfigs, market, resolution)
	if err != nil {
		return nil, err
	}
	return &BollingerBands{
		base: base{
			lock:       sync.Mutex{},
			id:         id,
			Market:     market,
			Resolution: resolution,
		},
		Period:    configs.Period,
		Deviation: configs.Deviation,
		Source:    configs.Source,
		sma:       sma,
	}, nil
}

func (conf *BollingerBands) Calculate(ctx context.Context, candles []Candle) (*indicatorsAPI.IndicatorValues, error) {
	smaValues, err := conf.sma.Calculate(ctx, candles)
	if err != nil {
		return nil, err
	}

	values := &indicatorsAPI.IndicatorValues{Values: make([]*indicatorsAPI.IndicatorValue, len(candles))}

	for i := 0; i < conf.Period-1; i++ {
		values.Values[i].Time = candles[i].Time.Unix()
	}

	for i := conf.Period - 1; i < len(candles); i++ {
		values.Values[i].Time = candles[i].Time.Unix()
		values.Values[i].Value = &indicatorsAPI.IndicatorValue_BollingerBands{
			BollingerBands: conf.calculateBB(candles[1+i-conf.Period:i+1], smaValues.GetValues()[i].GetSMA()),
		}
	}

	//conf.periodCandles = cloneCandles(candles[len(candles)-conf.Period:])
	conf.periodCandles = candles[len(candles)-conf.Period:]

	return values, nil
}

func (conf *BollingerBands) UpdateLast(ctx context.Context, candle Candle) *indicatorsAPI.IndicatorValue {
	conf.lock.Lock()
	defer conf.lock.Unlock()
	if candle.Time.After(conf.periodCandles[len(conf.periodCandles)-1].Time) {
		conf.periodCandles = conf.periodCandles[1:] //todo: must be check
	}
	//conf.periodCandles[conf.Period-1] = cloneCandle(candle)
	conf.periodCandles[conf.Period-1] = candle

	sma := conf.sma.UpdateLast(ctx, candle)
	value := conf.calculateBB(conf.periodCandles, sma.GetSMA())

	return &indicatorsAPI.IndicatorValue{
		Time:  candle.Time.Unix(),
		Value: &indicatorsAPI.IndicatorValue_BollingerBands{BollingerBands: value},
	}
}

func (conf *BollingerBands) calculateBB(candles []Candle, smaValue *indicatorsAPI.SMAValue) *indicatorsAPI.BollingerBandsValue {
	variance := float64(0)
	for j := 0; j < len(candles); j++ {
		num := float64(0)
		switch conf.Source {
		case entities.IndicatorSourceOpen:
			num = candles[j].Open
		case entities.IndicatorSourceHigh:
			num = candles[j].High
		case entities.IndicatorSourceLow:
			num = candles[j].Low
		case entities.IndicatorSourceClose:
			num = candles[j].Close
		case entities.IndicatorSourceOHLC4:
			num = (candles[j].Open + candles[j].High + candles[j].Low + candles[j].Close) / 4
		case entities.IndicatorSourceHLC3:
			num = (candles[j].Low + candles[j].High + candles[j].Close) / 3
		case entities.IndicatorSourceHL2:
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
