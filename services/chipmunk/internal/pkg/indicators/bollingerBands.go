package indicators

import (
	"errors"
	"github.com/google/uuid"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"math"
)

type bollingerBands struct {
	id uuid.UUID
	entity.BollingerBandsConfigs
}

func NewBollingerBands(id uuid.UUID, configs *entity.BollingerBandsConfigs) (*bollingerBands, error) {
	if err := validateBollingerBandsConfigs(configs); err != nil {
		return nil, err
	}
	return &bollingerBands{
		id:                    id,
		BollingerBandsConfigs: *configs,
	}, nil
}

func (conf *bollingerBands) GetType() chipmunkApi.IndicatorType {
	return chipmunkApi.Indicator_BollingerBands
}

func (conf *bollingerBands) GetLength() int {
	return conf.Length
}

func (conf *bollingerBands) Calculate(candles []*entity.Candle) error {
	if err := conf.validateBollingerBand(len(candles)); err != nil {
		return err
	}
	cloned := cloneCandles(candles)
	smaConf := movingAverage{
		id: uuid.New(),
		MovingAverageConfigs: entity.MovingAverageConfigs{
			Length: conf.Length,
			Source: conf.Source,
		},
	}
	sma, err := smaConf.sma(cloned)
	if err != nil {
		return err
	}
	for i := conf.Length - 1; i < len(candles); i++ {
		variance := float64(0)
		ma := sma[i]
		for j := 1 + i - conf.Length; j <= i; j++ {
			sum := float64(0)
			switch conf.Source {
			case chipmunkApi.Source_Open:
				sum = candles[j].Open
			case chipmunkApi.Source_High:
				sum = candles[j].High
			case chipmunkApi.Source_Low:
				sum = candles[j].Low
			case chipmunkApi.Source_Close:
				sum = candles[j].Close
			case chipmunkApi.Source_OHLC4:
				sum = (candles[j].Open + candles[j].High + candles[j].Low + candles[j].Close) / 4
			case chipmunkApi.Source_HLC3:
				sum = (candles[j].Low + candles[j].High + candles[j].Close) / 3
			case chipmunkApi.Source_HL2:
				sum = (candles[j].Low + candles[j].High) / 2
			}
			variance += math.Pow(ma-sum, 2)
		}
		variance /= float64(conf.Length)

		if candles[i] == nil {
			log.Errorf("nil candle")
		}

		candles[i].BollingerBands[conf.id] = &entity.BollingerBandsValue{
			UpperBand: ma + float64(conf.Deviation)*math.Sqrt(variance),
			LowerBand: ma - float64(conf.Deviation)*math.Sqrt(variance),
			MA:        ma,
		}
	}
	return nil
}

func (conf *bollingerBands) Update(candles []*entity.Candle) *entity.IndicatorValue {
	smaConf := movingAverage{
		id: uuid.New(),
		MovingAverageConfigs: entity.MovingAverageConfigs{
			Length: conf.Length,
			Source: conf.Source,
		},
	}
	sma, err := smaConf.sma(candles)
	if err != nil {
		log.WithError(err).Error("failed to calculate sma for bollinger bands")
		return nil
	}
	variance := float64(0)
	ma := sma[len(candles)-1]
	for j := 0; j < len(candles); j++ {
		sum := float64(0)
		switch conf.Source {
		case chipmunkApi.Source_Open:
			sum = candles[j].Open
		case chipmunkApi.Source_High:
			sum = candles[j].High
		case chipmunkApi.Source_Low:
			sum = candles[j].Low
		case chipmunkApi.Source_Close:
			sum = candles[j].Close
		case chipmunkApi.Source_OHLC4:
			sum = (candles[j].Open + candles[j].High + candles[j].Low + candles[j].Close) / 4
		case chipmunkApi.Source_HLC3:
			sum = (candles[j].Low + candles[j].High + candles[j].Close) / 3
		case chipmunkApi.Source_HL2:
			sum = (candles[j].Low + candles[j].High) / 2
		}
		variance += math.Pow(ma-sum, 2)
	}
	variance /= float64(conf.Length)
	return &entity.IndicatorValue{
		BB: &entity.BollingerBandsValue{
			UpperBand: ma + float64(conf.Deviation)*math.Sqrt(variance),
			LowerBand: ma - float64(conf.Deviation)*math.Sqrt(variance),
			MA:        ma,
		},
	}
}

func (conf *bollingerBands) validateBollingerBand(length int) error {
	if length < conf.Length {
		return errors.New("length must be bigger than or equal to candles length")
	}
	if conf.Deviation < 1 {
		return errors.New("deviation value must be positive")
	}
	return nil
}

func validateBollingerBandsConfigs(indicator *entity.BollingerBandsConfigs) error {
	return nil
}
