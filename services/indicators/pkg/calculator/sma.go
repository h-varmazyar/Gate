package calculator

import (
	"github.com/google/uuid"
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	"time"
)

type SMAValue struct {
	Value     float64
	TimeFrame time.Time
}

type SMA struct {
	id     uuid.UUID
	values []SMAValue
	Period int
	Source indicatorsAPI.Source
}

func NewMovingAverage(id uuid.UUID, period int, source indicatorsAPI.Source) (*SMA, error) {
	return &SMA{
		id:     id,
		values: make([]SMAValue, 0),
		Period: period,
		Source: source,
	}, nil
}

func (conf SMA) GetType() indicatorsAPI.Type {
	return indicatorsAPI.Type_SMA
}

func (conf SMA) GetLength() int {
	return len(conf.values)
}

func (conf SMA) Values(from, to int) []SMAValue {
	if len(conf.values) < from {
		return nil
	}
	if to > len(conf.values) {
		return nil
	}

	return conf.values[from:to]
}

func (conf SMA) Calculate(candles []*chipmunkAPI.Candle) {
	calculatorPeriod := float64(conf.Period)
	for _, candle := range candles {
		value := SMAValue{
			Value:     0,
			TimeFrame: time.Unix(candle.Time, 0),
		}
		last := len(conf.values) - 1
		if last >= 0 {
			value.Value = conf.values[last].Value
		}
		if len(conf.values) > conf.Period {
			value.Value -= conf.values[last+1-conf.Period].Value
		}

		switch conf.Source {
		case indicatorsAPI.Source_CLOSE:
			value.Value += candle.Close / calculatorPeriod
		case indicatorsAPI.Source_OPEN:
			value.Value += candle.Open / calculatorPeriod
		case indicatorsAPI.Source_HIGH:
			value.Value += candle.High / calculatorPeriod
		case indicatorsAPI.Source_LOW:
			value.Value += candle.Low / calculatorPeriod
		case indicatorsAPI.Source_HL2:
			value.Value += (candle.High + candle.Low) / (2 * calculatorPeriod)
		case indicatorsAPI.Source_HLC3:
			value.Value += (candle.High + candle.Low + candle.Close) / (3 * calculatorPeriod)
		case indicatorsAPI.Source_OHLC4:
			value.Value += (candle.Open + candle.Close + candle.High + candle.Low) / (4 * calculatorPeriod)
		}
		conf.values = append(conf.values, value)
	}
}

//
//func (conf SMA) sma(candles []*chipmunkAPI.Candle) ([]float64, error) {
//	//todo: return value must be delete
//	response := make([]float64, len(candles))
//	if err := conf.validateMA(len(candles)); err != nil {
//		return nil, err
//	}
//	for i := conf.Length - 1; i < len(candles); i++ {
//		sum := float64(0)
//		for _, innerCandle := range candles[i-conf.Length+1 : i+1] {
//			switch conf.Source {
//			case chipmunkApi.Source_Close:
//				sum += innerCandle.Close
//			case chipmunkApi.Source_Open:
//				sum += innerCandle.Open
//			case chipmunkApi.Source_High:
//				sum += innerCandle.High
//			case chipmunkApi.Source_Low:
//				sum += innerCandle.Low
//			case chipmunkApi.Source_HL2:
//				sum += (innerCandle.High + innerCandle.Low) / 2
//			case chipmunkApi.Source_HLC3:
//				sum += (innerCandle.High + innerCandle.Low + innerCandle.Close) / 3
//			case chipmunkApi.Source_OHLC4:
//				sum += (innerCandle.Open + innerCandle.Close + innerCandle.High + innerCandle.Low) / 4
//			}
//		}
//		if candles[i].MovingAverages[conf.id] == nil {
//			candles[i].MovingAverages[conf.id] = new(entity.MovingAverageValue)
//		}
//		candles[i].MovingAverages[conf.id].Simple = sum / float64(conf.Length)
//		response[i] = sum / float64(conf.Length)
//	}
//	return response, nil
//}
//
////func (conf *movingAverage) updateSMA(rateLimiters []*entity.Candle) float64 {
////	smaConf := movingAverage{
////		id:                   uuid.New(),
////		MovingAverageConfigs: conf.MovingAverageConfigs,
////	}
////	if sma, err := smaConf.sma(rateLimiters); err != nil {
////		return float64(0)
////	} else {
////		return sma[len(rateLimiters)-1]
////	}
////}
//
//func (conf SMA) Calculate(candles []*chipmunkApi.Candle) error {
//	//values := make([]*entity.MovingAverageValue, 0)
//	//var sma []float64
//	var err error
//	for i := 0; i < len(candles); i++ {
//		if candles[i].MovingAverages[conf.id] == nil {
//			candles[i].MovingAverages[conf.id] = new(entity.MovingAverageValue)
//		}
//	}
//	//if sma, err = conf.sma(rateLimiters, conf.id); err != nil {
//	if _, err = conf.sma(candles); err != nil {
//		log.WithError(err).Errorf("failed to calculate sma for moving average %v", conf.id)
//		return err
//	}
//	//for _, value := range sma {
//	//	values = append(values, &entity.MovingAverageValue{
//	//		Simple: value,
//	//	})
//	//}
//	candles[conf.Length-1].MovingAverages[conf.id].Exponential = candles[conf.Length-1].MovingAverages[conf.id].Simple
//
//	for i := conf.Length; i < len(candles); i++ {
//		conf.calculateEMA(candles[i], candles[i-1])
//	}
//
//	//for i := 0; i < len(rateLimiters); i++ {
//	//	rateLimiters[i].MovingAverages[conf.id] = &entity.MovingAverageValue{
//	//		Simple:      values[i].Simple,
//	//		Exponential: values[i].Exponential,
//	//	}
//	//}
//	return nil
//}
//
//func (conf *movingAverage) Update(candles []*entity.Candle) {
//	if len(candles) == 0 {
//		return
//	}
//	first := candles[0]
//	start := buffer.CandleBuffer.Before(first.MarketID.String(), first.ResolutionID.String(), first.Time, conf.Length-1)
//	if len(start) == 0 {
//		return
//	}
//
//	internalCandles := append(start, candles...)
//
//	_, err := conf.sma(internalCandles)
//	if err != nil {
//		log.WithError(err).Errorf("failed to calculate sma")
//		return
//	}
//
//	for i := conf.Length - 1; i < len(internalCandles); i++ {
//		conf.calculateEMA(candles[i-1], candles[i])
//	}
//}
//
//func (conf *movingAverage) calculateEMA(candle1, candle2 *entity.Candle) {
//	price := float64(0)
//	switch conf.Source {
//	case chipmunkApi.Source_Close:
//		price = candle2.Close
//	case chipmunkApi.Source_Open:
//		price = candle2.Open
//	case chipmunkApi.Source_High:
//		price = candle2.High
//	case chipmunkApi.Source_Low:
//		price = candle2.Low
//	case chipmunkApi.Source_HL2:
//		price = (candle2.High + candle2.Low) / 2
//	case chipmunkApi.Source_HLC3:
//		price = (candle2.High + candle2.Low + candle2.Close) / 3
//	case chipmunkApi.Source_OHLC4:
//		price = (candle2.Open + candle2.High + candle2.Low + candle2.Close) / 4
//	}
//	candle2.MovingAverages[conf.id].Exponential = price*conf.factor + candle1.MovingAverages[conf.id].Exponential*(1-conf.factor)
//}
//
//func (conf *movingAverage) validateMA(length int) error {
//	if length < conf.Length {
//		return errors.New(fmt.Sprintf("rateLimiters length must be grater or equal than %d", conf.Length))
//	}
//	switch conf.Source {
//	case chipmunkApi.Source_Close,
//		chipmunkApi.Source_Open,
//		chipmunkApi.Source_High,
//		chipmunkApi.Source_Low,
//		chipmunkApi.Source_HL2,
//		chipmunkApi.Source_HLC3,
//		chipmunkApi.Source_OHLC4,
//		chipmunkApi.Source_Custom:
//		return nil
//	default:
//		return errors.New(fmt.Sprintf("moving average source not valid: %s", conf.Source))
//	}
//}
