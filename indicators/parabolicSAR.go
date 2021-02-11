package indicators

import (
	"errors"
	"github.com/mrNobody95/Gate/models"
	"math"
)

type pSarConfig struct {
	*basicConfig
	acceleration       float64
	maxAcceleration    float64
	accelerationFactor float64
	startAcceleration  float64
	extremePoint       float64
	trend              Trend
}

type Trend int

const (
	Long  Trend = 1
	Short Trend = -1
)

func NewParabolicSARConfig(length int, acceleration, maxAcceleration, startAF float64) *pSarConfig {
	return &pSarConfig{
		basicConfig: &basicConfig{
			Length: length,
		},
		acceleration:       acceleration,
		maxAcceleration:    maxAcceleration,
		startAcceleration:  startAF,
		accelerationFactor: startAF,
	}
}

func (conf *pSarConfig) CalculatePSAR(candles []models.Candle, appendCandles bool) error {
	var rangeCounter int
	if appendCandles {
		rangeCounter = len(conf.Candles) - 1
		conf.Candles = append(conf.Candles, candles...)
	} else {
		conf.Candles = candles
		rangeCounter = 1
		if conf.Candles[1].High >= conf.Candles[0].High || conf.Candles[0].Low <= conf.Candles[1].Low {
			conf.trend = Long
			conf.Candles[1].ParabolicSAR.SAR = conf.Candles[0].Low
			conf.extremePoint = conf.Candles[0].High
		} else {
			conf.trend = Short
			conf.extremePoint = conf.Candles[0].Low
		}
		conf.Candles[1].ParabolicSAR.Trend = conf.trend
	}
	if err := conf.validate(); err != nil {
		return err
	}

	for i, candle := range conf.Candles[rangeCounter : len(conf.Candles)-1] {
		nextSar := candle.ParabolicSAR.SAR
		if conf.trend == Long {
			if candle.High > conf.extremePoint {
				conf.extremePoint = candle.High
				conf.accelerationFactor = math.Min(conf.maxAcceleration, conf.acceleration+conf.accelerationFactor)
			}
			tmpSar := nextSar + conf.accelerationFactor*(conf.extremePoint-nextSar)
			nextSar = math.Min(math.Min(conf.Candles[rangeCounter+i-1].Low, candle.Low), tmpSar)
			if nextSar > conf.Candles[rangeCounter+i+1].Low {
				conf.trend = Short
				nextSar = conf.extremePoint
				conf.extremePoint = conf.Candles[rangeCounter+i+1].Low
				conf.accelerationFactor = conf.startAcceleration
			}
		} else if conf.trend == Short {
			if candle.Low < conf.extremePoint {
				conf.extremePoint = candle.Low
				conf.accelerationFactor = math.Min(conf.maxAcceleration, conf.acceleration+conf.accelerationFactor)
			}
			tmpSar := nextSar + conf.accelerationFactor*(conf.extremePoint-nextSar)
			nextSar = math.Max(math.Max(conf.Candles[rangeCounter+i-1].High, candle.High), tmpSar)
			if nextSar < conf.Candles[rangeCounter+i+1].High {
				conf.trend = Long
				nextSar = conf.extremePoint
				conf.extremePoint = conf.Candles[rangeCounter+i+1].High
				conf.accelerationFactor = conf.startAcceleration
			}
		}
		conf.Candles[rangeCounter+i+1].ParabolicSAR.SAR = nextSar
		conf.Candles[rangeCounter+i+1].ParabolicSAR.Trend = conf.trend
		conf.Candles[rangeCounter+i+1].ParabolicSAR.TrendFlipped = candle.ParabolicSAR.Trend != conf.trend
	}
	return nil
}

func (conf *pSarConfig) validate() error {
	if len(conf.Candles) < 2 {
		return errors.New("candle length must be more than 2")
	}
	return nil
}
