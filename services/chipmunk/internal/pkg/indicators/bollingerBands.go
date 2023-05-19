package indicators

import (
	"errors"
	"github.com/google/uuid"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"math"
)

type bollingerBands struct {
	id uuid.UUID
	entity.BollingerBandsConfigs
	sma *movingAverage
}

func NewBollingerBands(id uuid.UUID, configs *entity.BollingerBandsConfigs) (*bollingerBands, error) {
	if err := validateBollingerBandsConfigs(configs); err != nil {
		return nil, err
	}
	return &bollingerBands{
		id:                    id,
		BollingerBandsConfigs: *configs,
		sma: &movingAverage{
			id: id,
			MovingAverageConfigs: entity.MovingAverageConfigs{
				Length: configs.Length,
				Source: configs.Source,
			},
		},
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
	_, err := conf.sma.sma(candles)
	if err != nil {
		return err
	}
	for i := conf.Length - 1; i < len(candles); i++ {
		conf.calculateBB(candles[1+i-conf.Length : i+1])
		//variance := float64(0)
		//ma := rateLimiters[i].MovingAverages[conf.id].Simple
		//for j := 1 + i - conf.Length; j <= i; j++ {
		//	sum := float64(0)
		//	switch conf.Source {
		//	case chipmunkApi.Source_Open:
		//		sum = rateLimiters[j].Open
		//	case chipmunkApi.Source_High:
		//		sum = rateLimiters[j].High
		//	case chipmunkApi.Source_Low:
		//		sum = rateLimiters[j].Low
		//	case chipmunkApi.Source_Close:
		//		sum = rateLimiters[j].Close
		//	case chipmunkApi.Source_OHLC4:
		//		sum = (rateLimiters[j].Open + rateLimiters[j].High + rateLimiters[j].Low + rateLimiters[j].Close) / 4
		//	case chipmunkApi.Source_HLC3:
		//		sum = (rateLimiters[j].Low + rateLimiters[j].High + rateLimiters[j].Close) / 3
		//	case chipmunkApi.Source_HL2:
		//		sum = (rateLimiters[j].Low + rateLimiters[j].High) / 2
		//	}
		//	variance += math.Pow(ma-sum, 2)
		//}
		//variance /= float64(conf.Length)
		//
		//if rateLimiters[i] == nil {
		//	log.Errorf("nil candle")
		//}
		//
		//rateLimiters[i].BollingerBands[conf.id] = &entity.BollingerBandsValue{
		//	UpperBand: ma + float64(conf.Deviation)*math.Sqrt(variance),
		//	LowerBand: ma - float64(conf.Deviation)*math.Sqrt(variance),
		//	MA:        ma,
		//}
	}
	return nil
}

func (conf *bollingerBands) Update(candles []*entity.Candle) {
	if len(candles) == 0 {
		return
	}
	first := candles[0]
	start := buffer.CandleBuffer.Before(first.MarketID.String(), first.ResolutionID.String(), first.Time, conf.Length-1)
	if len(start) == 0 {
		return
	}

	internalCandles := append(start, candles...)

	_, err := conf.sma.sma(internalCandles)
	if err != nil {
		log.WithError(err).Error("failed to calculate sma for bollinger bands")
		return
	}

	for i := conf.Length - 1; i < len(internalCandles); i++ {
		conf.calculateBB(candles[1+i-conf.Length : i+1])
		//variance := float64(0)
		//ma := internalCandles[i].MovingAverages[conf.id].Simple
		//for j := 1 + i - conf.Length; j <= i; j++ {
		//	sum := float64(0)
		//	switch conf.Source {
		//	case chipmunkApi.Source_Open:
		//		sum = rateLimiters[j].Open
		//	case chipmunkApi.Source_High:
		//		sum = rateLimiters[j].High
		//	case chipmunkApi.Source_Low:
		//		sum = rateLimiters[j].Low
		//	case chipmunkApi.Source_Close:
		//		sum = rateLimiters[j].Close
		//	case chipmunkApi.Source_OHLC4:
		//		sum = (rateLimiters[j].Open + rateLimiters[j].High + rateLimiters[j].Low + rateLimiters[j].Close) / 4
		//	case chipmunkApi.Source_HLC3:
		//		sum = (rateLimiters[j].Low + rateLimiters[j].High + rateLimiters[j].Close) / 3
		//	case chipmunkApi.Source_HL2:
		//		sum = (rateLimiters[j].Low + rateLimiters[j].High) / 2
		//	}
		//	variance += math.Pow(ma-sum, 2)
		//}
		//variance /= float64(conf.Length)
		//
		//rateLimiters[i].BollingerBands[conf.id] = &entity.BollingerBandsValue{
		//	UpperBand: ma + float64(conf.Deviation)*math.Sqrt(variance),
		//	LowerBand: ma - float64(conf.Deviation)*math.Sqrt(variance),
		//	MA:        ma,
		//}
	}
}

func (conf *bollingerBands) calculateBB(candles []*entity.Candle) {
	variance := float64(0)
	index := len(candles) - 1
	ma := candles[index].MovingAverages[conf.id].Simple
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

	candles[index].BollingerBands[conf.id] = &entity.BollingerBandsValue{
		UpperBand: ma + float64(conf.Deviation)*math.Sqrt(variance),
		LowerBand: ma - float64(conf.Deviation)*math.Sqrt(variance),
		MA:        ma,
	}
}

func (conf *bollingerBands) validateBollingerBand(length int) error {
	if length < conf.Length {
		return errors.New("length must be bigger than or equal to rateLimiters length")
	}
	if conf.Deviation < 1 {
		return errors.New("deviation value must be positive")
	}
	return nil
}

func validateBollingerBandsConfigs(indicator *entity.BollingerBandsConfigs) error {
	return nil
}
