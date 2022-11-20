package indicators

import (
	"errors"
	"github.com/google/uuid"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
)

type rsi struct {
	id uuid.UUID
	entity.RsiConfigs
}

func NewRSI(id uuid.UUID, configs *entity.RsiConfigs) (*rsi, error) {
	if err := validateRsiConfigs(configs); err != nil {
		return nil, err
	}
	return &rsi{
		id:         id,
		RsiConfigs: *configs,
	}, nil
}

func (conf *rsi) GetType() chipmunkApi.IndicatorType {
	return chipmunkApi.Indicator_RSI
}

func (conf *rsi) GetLength() int {
	return conf.Length
}

func (conf *rsi) Calculate(candles []*entity.Candle) error {
	if err := conf.validateRSI(len(candles)); err != nil {
		return err
	}

	{ //first RSI
		gain := float64(0)
		loss := float64(0)
		for i := 1; i <= conf.Length; i++ {
			if change := candles[i].Close - candles[i-1].Close; change > 0 {
				gain += change
			} else {
				loss += change
			}
		}
		loss *= -1
		gain = gain / float64(conf.Length)
		loss = loss / float64(conf.Length)
		rs := gain / loss
		candles[conf.Length].RSIs[conf.id] = &entity.RSIValue{
			Gain: gain,
			Loss: loss,
			RSI:  100 - (100 / (1 + rs)),
		}
	}

	for i := conf.Length + 1; i < len(candles); i++ {
		gain := float64(0)
		loss := float64(0)
		if change := candles[i].Close - candles[i-1].Close; change > 0 {
			gain = change
		} else {
			loss = change * -1
		}
		avgGain := (candles[i-1].RSIs[conf.id].Gain*float64(conf.Length-1) + gain) / float64(conf.Length)
		avgLoss := (candles[i-1].RSIs[conf.id].Loss*float64(conf.Length-1) + loss) / float64(conf.Length)
		rs := avgGain / avgLoss
		rsiValue := 100 - (100 / (1 + rs))

		candles[i].RSIs[conf.id] = &entity.RSIValue{
			Gain: avgGain,
			Loss: avgLoss,
			RSI:  rsiValue,
		}
	}
	return nil
}

func (conf *rsi) Update(candles []*entity.Candle) *entity.IndicatorValue {
	gain, loss := float64(0), float64(0)
	last := len(candles) - 1
	if last < 1 {
		return nil
	}
	if change := candles[last].Close - candles[last-1].Close; change > 0 {
		gain = change
	} else {
		loss = change
	}

	avgGain := (candles[last-1].RSIs[conf.id].Gain*float64(conf.Length-1) + gain) / float64(conf.Length)
	avgLoss := (candles[last-1].RSIs[conf.id].Loss*float64(conf.Length-1) - loss) / float64(conf.Length)
	rs := avgGain / avgLoss
	rsiValue := 100 - (100 / (1 + rs))

	return &entity.IndicatorValue{
		RSI: &entity.RSIValue{
			Gain: avgGain,
			Loss: avgLoss,
			RSI:  rsiValue,
		}}
}

func (conf *rsi) validateRSI(length int) error {
	if length-1 < conf.Length {
		return errors.New("candles length must bigger than indicators period length")
	}
	return nil
}

func validateRsiConfigs(indicator *entity.RsiConfigs) error {
	return nil
}
