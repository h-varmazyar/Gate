package calculator

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/errors"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	"google.golang.org/grpc/codes"
	"time"
)

type SMAValue struct {
	Value     *float64
	TimeFrame time.Time
}

type SMA struct {
	id            uuid.UUID
	Period        int
	Source        indicatorsAPI.Source
	lastValue     SMAValue
	periodCandles []*chipmunkAPI.Candle
}

func NewSMA(period int, source indicatorsAPI.Source) (*SMA, error) {
	return &SMA{
		id:     uuid.New(),
		Period: period,
		Source: source,
	}, nil
}

func (conf *SMA) Calculate(ctx context.Context, candles []*chipmunkAPI.Candle, values []*SMAValue) error {
	if len(candles) < conf.Period {
		return errors.New(ctx, codes.FailedPrecondition).AddDetails("invalid candle length")
	}

	firstNum := conf.getSmaNumber(candles[0])
	values[0] = &SMAValue{
		Value:     &firstNum,
		TimeFrame: time.Unix(candles[0].Time, 0),
	}

	for i := 1; i < len(candles); i++ {
		values[i] = &SMAValue{
			TimeFrame: time.Unix(candles[i].Time, 0),
		}

		newNum := conf.getSmaNumber(candles[i])
		var changeableNum float64
		if i >= conf.Period {
			changeableNum = conf.getSmaNumber(candles[i-conf.Period])
		}

		num := *values[i-1].Value + newNum - changeableNum
		values[i].Value = &num
	}
	for i := 0; i < conf.Period-1; i++ {
		values[i].Value = nil
	}

	conf.periodCandles = cloneCandles(candles[len(candles)-conf.Period:])
	if values[len(values)-1].Value != nil {
		num := *values[len(values)-1].Value
		conf.lastValue = SMAValue{
			Value:     &num,
			TimeFrame: values[len(values)-1].TimeFrame,
		}
	}
	return nil
}

func (conf *SMA) UpdateLast(_ context.Context, candle *chipmunkAPI.Candle, value *SMAValue) {
	var changeableCandle *chipmunkAPI.Candle
	if conf.lastValue.TimeFrame.Unix() == candle.Time {
		changeableCandle = conf.periodCandles[conf.Period-1]
	} else if conf.lastValue.TimeFrame.Unix() < candle.Time {
		changeableCandle = cloneCandle(conf.periodCandles[0])
		conf.periodCandles = conf.periodCandles[1:]
	}
	conf.periodCandles[conf.Period-1] = cloneCandle(candle)

	newNum := conf.getSmaNumber(candle)
	changeableNum := conf.getSmaNumber(changeableCandle)

	newValue := *conf.lastValue.Value + newNum - changeableNum
	conf.lastValue.Value = &newValue

	resp := newValue
	value = &SMAValue{
		Value:     &resp,
		TimeFrame: time.Unix(candle.Time, 0),
	}
}

func (conf *SMA) getSmaNumber(candle *chipmunkAPI.Candle) float64 {
	calculatorPeriod := float64(conf.Period)
	switch conf.Source {
	case indicatorsAPI.Source_CLOSE:
		return candle.Close / calculatorPeriod
	case indicatorsAPI.Source_OPEN:
		return candle.Open / calculatorPeriod
	case indicatorsAPI.Source_HIGH:
		return candle.High / calculatorPeriod
	case indicatorsAPI.Source_LOW:
		return candle.Low / calculatorPeriod
	case indicatorsAPI.Source_HL2:
		return (candle.High + candle.Low) / (2 * calculatorPeriod)
	case indicatorsAPI.Source_HLC3:
		return (candle.High + candle.Low + candle.Close) / (3 * calculatorPeriod)
	case indicatorsAPI.Source_OHLC4:
		return (candle.Open + candle.Close + candle.High + candle.Low) / (4 * calculatorPeriod)
	}

	return 0
}
