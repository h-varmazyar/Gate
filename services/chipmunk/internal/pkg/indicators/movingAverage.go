package indicators

//
//import (
//	"errors"
//	"fmt"
//	"github.com/google/uuid"
//	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
//	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
//	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
//	log "github.com/sirupsen/logrus"
//)
//
//type movingAverage struct {
//	id uuid.UUID
//	entity.MovingAverageConfigs
//	factor float64
//}
//
//func NewMovingAverage(id uuid.UUID, configs *entity.MovingAverageConfigs) (*movingAverage, error) {
//	return &movingAverage{
//		id:                   id,
//		MovingAverageConfigs: *configs,
//		factor:               2 / float64(configs.Length+1),
//	}, nil
//}
//
//func (conf *movingAverage) GetType() chipmunkApi.IndicatorType {
//	return chipmunkApi.Indicator_MovingAverage
//}
//
//func (conf *movingAverage) GetLength() int {
//	return conf.Length
//}
//
//func (conf *movingAverage) sma(candles []*entity.Candle) ([]float64, error) {
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
////func (conf *movingAverage) updateSMA(rateLimiters []*entities.Candle) float64 {
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
//func (conf *movingAverage) Calculate(candles []*entity.Candle) error {
//	//values := make([]*entities.MovingAverageValue, 0)
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
//	//	values = append(values, &entities.MovingAverageValue{
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
//	//	rateLimiters[i].MovingAverages[conf.id] = &entities.MovingAverageValue{
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
