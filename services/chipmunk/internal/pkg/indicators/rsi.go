package indicators

//
//import (
//	"errors"
//	"github.com/google/uuid"
//	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
//	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
//	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
//)
//
//type rsi struct {
//	id uuid.UUID
//	entity.RsiConfigs
//}
//
//func NewRSI(id uuid.UUID, configs *entity.RsiConfigs) (*rsi, error) {
//	return &rsi{
//		id:         id,
//		RsiConfigs: *configs,
//	}, nil
//}
//
//func (conf *rsi) GetType() chipmunkApi.IndicatorType {
//	return chipmunkApi.Indicator_RSI
//}
//
//func (conf *rsi) GetLength() int {
//	return conf.Length
//}
//
//func (conf *rsi) Calculate(candles []*entity.Candle) error {
//	if err := conf.validateRSI(len(candles)); err != nil {
//		return err
//	}
//
//	{ //first RSI
//		gain := float64(0)
//		loss := float64(0)
//		for i := 1; i <= conf.Length; i++ {
//			if change := candles[i].Close - candles[i-1].Close; change > 0 {
//				gain += change
//			} else {
//				loss += change
//			}
//		}
//		loss *= -1
//		gain = gain / float64(conf.Length)
//		loss = loss / float64(conf.Length)
//		rs := gain / loss
//		candles[conf.Length].RSIs[conf.id] = &entity.RSIValue{
//			Gain: gain,
//			Loss: loss,
//			RSI:  100 - (100 / (1 + rs)),
//		}
//	}
//
//	for i := conf.Length + 1; i < len(candles); i++ {
//		conf.calculateRSIValue(candles[i-1], candles[i])
//	}
//	return nil
//}
//
//func (conf *rsi) Update(candles []*entity.Candle) {
//	if len(candles) == 0 {
//		return
//	}
//
//	start := buffer.CandleBuffer.Before(candles[0].MarketID.String(), candles[0].ResolutionID.String(), candles[0].Time, 1)
//
//	if len(start) == 0 {
//		return
//	}
//
//	internalCandles := append(start, candles...)
//
//	for i := 1; i < len(internalCandles); i++ {
//		conf.calculateRSIValue(internalCandles[i-1], internalCandles[i])
//	}
//}
//
//func (conf *rsi) validateRSI(length int) error {
//	if length-1 < conf.Length {
//		return errors.New("rateLimiters length must bigger than indicators period length")
//	}
//	return nil
//}
//
//func (conf *rsi) calculateRSIValue(candle1, candle2 *entity.Candle) {
//	gain, loss := float64(0), float64(0)
//	if change := candle2.Close - candle1.Close; change > 0 {
//		gain = change
//	} else {
//		loss = change
//	}
//	avgGain := (candle1.RSIs[conf.id].Gain*float64(conf.Length-1) + gain) / float64(conf.Length)
//	avgLoss := (candle1.RSIs[conf.id].Loss*float64(conf.Length-1) - loss) / float64(conf.Length)
//	rs := avgGain / avgLoss
//	rsiValue := 100 - (100 / (1 + rs))
//
//	candle2.RSIs[conf.id] = &entity.RSIValue{
//		Gain: avgGain,
//		Loss: avgLoss,
//		RSI:  rsiValue,
//	}
//}
