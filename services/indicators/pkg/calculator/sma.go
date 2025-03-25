package calculator

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/errors"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/entities"
	"google.golang.org/grpc/codes"
	"sync"
	"time"
)

type SMA struct {
	base
	Period        int
	Source        entities.IndicatorSource
	lastValue     indicatorsAPI.SMAValue
	lastTimeframe time.Time
	periodCandles []Candle
}

func NewSMA(id uint, configs *entities.SMAConfigs, market *chipmunkAPI.Market, resolution *chipmunkAPI.Resolution) (*SMA, error) {
	return &SMA{
		base: base{
			lock:       sync.Mutex{},
			id:         id,
			Market:     market,
			Resolution: resolution,
		},
		Period: configs.Period,
		Source: configs.Source,
	}, nil
}

func (conf *SMA) GetMarket() *chipmunkAPI.Market {
	return conf.Market
}

func (conf *SMA) GetResolution() *chipmunkAPI.Resolution {
	return conf.Resolution
}

func (conf *SMA) GetId() uint {
	return conf.id
}

func (conf *SMA) Calculate(ctx context.Context, candles []Candle) (*indicatorsAPI.IndicatorValues, error) {
	if len(candles) < conf.Period {
		return nil, errors.New(ctx, codes.FailedPrecondition).AddDetails("invalid candle length")
	}

	values := &indicatorsAPI.IndicatorValues{
		Values: make([]*indicatorsAPI.IndicatorValue, len(candles)),
	}

	values.Values[0].Time = time.Now().Unix()

	baseNum := conf.getSmaNumber(candles[0])
	values.Values[0].Value = &indicatorsAPI.IndicatorValue_SMA{SMA: &indicatorsAPI.SMAValue{Value: baseNum}}
	values.Values[0].Time = candles[0].Time.Unix()

	for i := 1; i < len(candles); i++ {

		newNum := conf.getSmaNumber(candles[i])
		var changeableNum float64
		if i >= conf.Period {
			changeableNum = conf.getSmaNumber(candles[i-conf.Period])
		}

		baseNum = baseNum + newNum - changeableNum
		values.Values[i].Value = &indicatorsAPI.IndicatorValue_SMA{SMA: &indicatorsAPI.SMAValue{Value: baseNum}}
		values.Values[i].Time = candles[i].Time.Unix()
	}
	for i := 0; i < conf.Period-1; i++ {
		values.Values[i].Value = nil
	}

	//conf.periodCandles = cloneCandles(candles[len(candles)-conf.Period:])
	conf.periodCandles = candles[len(candles)-conf.Period:]
	conf.lastValue = indicatorsAPI.SMAValue{
		Value: baseNum,
	}
	conf.lastTimeframe = candles[len(candles)-1].Time
	return values, nil
}

func (conf *SMA) UpdateLast(_ context.Context, candle Candle) *indicatorsAPI.IndicatorValue {
	conf.lock.Lock()
	defer conf.lock.Unlock()
	var changeableCandle Candle
	if conf.lastTimeframe == candle.Time {
		changeableCandle = conf.periodCandles[conf.Period-1]
	} else if conf.lastTimeframe.Before(candle.Time) {
		//changeableCandle = cloneCandle(conf.periodCandles[0])
		changeableCandle = conf.periodCandles[0]
		conf.periodCandles = conf.periodCandles[1:] //todo: must be check
	}
	//conf.periodCandles[conf.Period-1] = cloneCandle(candle)
	conf.periodCandles[conf.Period-1] = candle

	newNum := conf.getSmaNumber(candle)
	changeableNum := conf.getSmaNumber(changeableCandle)

	v := conf.lastValue.Value + newNum - changeableNum
	conf.lastValue.Value = v

	return &indicatorsAPI.IndicatorValue{
		Time:  candle.Time.Unix(),
		Value: &indicatorsAPI.IndicatorValue_SMA{SMA: &indicatorsAPI.SMAValue{Value: v}},
	}
}

func (conf *SMA) getSmaNumber(candle Candle) float64 {
	calculatorPeriod := float64(conf.Period)
	switch conf.Source {
	case entities.IndicatorSourceClose:
		return candle.Close / calculatorPeriod
	case entities.IndicatorSourceOpen:
		return candle.Open / calculatorPeriod
	case entities.IndicatorSourceHigh:
		return candle.High / calculatorPeriod
	case entities.IndicatorSourceLow:
		return candle.Low / calculatorPeriod
	case entities.IndicatorSourceHL2:
		return (candle.High + candle.Low) / (2 * calculatorPeriod)
	case entities.IndicatorSourceHLC3:
		return (candle.High + candle.Low + candle.Close) / (3 * calculatorPeriod)
	case entities.IndicatorSourceOHLC4:
		return (candle.Open + candle.Close + candle.High + candle.Low) / (4 * calculatorPeriod)
	}

	return 0
}
