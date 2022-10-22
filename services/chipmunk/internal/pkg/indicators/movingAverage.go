package indicators

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
	log "github.com/sirupsen/logrus"
)

type movingAverage struct {
	id uuid.UUID
	repository.MovingAverageConfigs
}

func NewMovingAverage(id uuid.UUID, configs *repository.MovingAverageConfigs) (*movingAverage, error) {
	return &movingAverage{
		id:                   id,
		MovingAverageConfigs: *configs,
	}, nil
}

func (conf *movingAverage) GetType() chipmunkApi.IndicatorType {
	return chipmunkApi.Indicator_MovingAverage
}

func (conf *movingAverage) GetLength() int {
	return conf.Length
}

func (conf *movingAverage) sma(candles []*repository.Candle) ([]float64, error) {
	response := make([]float64, len(candles))
	if err := conf.validateMA(len(candles)); err != nil {
		return nil, err
	}
	for i := conf.Length - 1; i < len(candles); i++ {
		sum := float64(0)
		for _, innerCandle := range candles[i-conf.Length+1 : i+1] {
			switch conf.Source {
			case chipmunkApi.Source_Close:
				sum += innerCandle.Close
			case chipmunkApi.Source_Open:
				sum += innerCandle.Open
			case chipmunkApi.Source_High:
				sum += innerCandle.High
			case chipmunkApi.Source_Low:
				sum += innerCandle.Low
			case chipmunkApi.Source_HL2:
				sum += (innerCandle.High + innerCandle.Low) / 2
			case chipmunkApi.Source_HLC3:
				sum += (innerCandle.High + innerCandle.Low + innerCandle.Close) / 3
			case chipmunkApi.Source_OHLC4:
				sum += (innerCandle.Open + innerCandle.Close + innerCandle.High + innerCandle.Low) / 4
			}
		}
		response[i] = sum / float64(conf.Length)
		//log.Infof("sma %v is %v", i, response[i])
	}
	return response, nil
}

func (conf *movingAverage) updateSMA(candles []*repository.Candle) float64 {
	smaConf := movingAverage{
		id:                   uuid.New(),
		MovingAverageConfigs: conf.MovingAverageConfigs,
	}
	if sma, err := smaConf.sma(candles); err != nil {
		return float64(0)
	} else {
		return sma[len(candles)-1]
	}
}

func (conf *movingAverage) Calculate(candles []*repository.Candle) error {
	values := make([]*repository.MovingAverageValue, 0)
	var sma []float64
	var err error
	if sma, err = conf.sma(candles); err != nil {
		log.WithError(err).Errorf("failed to calculate sma for moving average %v", conf.id)
		return err
	}
	for _, value := range sma {
		values = append(values, &repository.MovingAverageValue{
			Simple: value,
		})
	}
	values[conf.Length-1].Exponential = values[conf.Length-1].Simple

	factor := 2 / float64(conf.Length+1)
	for i := conf.Length; i < len(candles); i++ {
		price := float64(0)
		switch conf.Source {
		case chipmunkApi.Source_Close:
			price = candles[i].Close
		case chipmunkApi.Source_Open:
			price = candles[i].Open
		case chipmunkApi.Source_High:
			price = candles[i].High
		case chipmunkApi.Source_Low:
			price = candles[i].Low
		case chipmunkApi.Source_HL2:
			price = (candles[i].High + candles[i].Low) / 2
		case chipmunkApi.Source_HLC3:
			price = (candles[i].High + candles[i].Low + candles[i].Close) / 3
		case chipmunkApi.Source_OHLC4:
			price = (candles[i].Open + candles[i].Close + candles[i].High + candles[i].Low) / 4
		}
		values[i].Exponential = price*factor + values[i-1].Exponential*(1-factor)
	}

	for i := 0; i < len(candles); i++ {
		candles[i].MovingAverages[conf.id] = &repository.MovingAverageValue{
			Simple:      values[i].Simple,
			Exponential: values[i].Exponential,
		}
	}
	return nil
}

func (conf *movingAverage) Update(candles []*repository.Candle) *repository.IndicatorValue {
	//candles := buffer.Markets.GetLastNCandles(conf.MarketName, 2)
	//values := buffer.Markets.GetLastNIndicatorValue(conf.MarketName, conf.GetID(), 2)

	i := len(candles) - 1
	price := float64(0)
	factor := 2 / float64(conf.Length+1)

	switch conf.Source {
	case chipmunkApi.Source_Close:
		price = candles[i].Close
	case chipmunkApi.Source_Open:
		price = candles[i].Open
	case chipmunkApi.Source_High:
		price = candles[i].High
	case chipmunkApi.Source_Low:
		price = candles[i].Low
	case chipmunkApi.Source_HL2:
		price = (candles[i].High + candles[i].Low) / 2
	case chipmunkApi.Source_HLC3:
		price = (candles[i].High + candles[i].Low + candles[i].Close) / 3
	case chipmunkApi.Source_OHLC4:
		price = (candles[i].Open + candles[i].Close + candles[i].High + candles[i].Low) / 4
	}
	return &repository.IndicatorValue{
		MA: &repository.MovingAverageValue{
			Exponential: price*factor + candles[i-1].MovingAverages[conf.id].Exponential*(1-factor),
			Simple:      conf.updateSMA(candles),
		},
	}
}

func (conf *movingAverage) validateMA(length int) error {
	if length < conf.Length {
		return errors.New(fmt.Sprintf("candles length must be grater or equal than %d", conf.Length))
	}
	switch conf.Source {
	case chipmunkApi.Source_Close,
		chipmunkApi.Source_Open,
		chipmunkApi.Source_High,
		chipmunkApi.Source_Low,
		chipmunkApi.Source_HL2,
		chipmunkApi.Source_HLC3,
		chipmunkApi.Source_OHLC4,
		chipmunkApi.Source_Custom:
		return nil
	default:
		return errors.New(fmt.Sprintf("moving average source not valid: %s", conf.Source))
	}
}
