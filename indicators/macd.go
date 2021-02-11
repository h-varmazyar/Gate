package indicators

import (
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/models"
)

type macdConfig struct {
	Candles      []models.Candle
	fastLength   int
	slowLength   int
	signalLength int
	source       Source
	//lastSlowValue   float64
	//lastFastValue   float64
	//lastSignalValue float64
}

type Source string

const (
	SourceOpen  = "open"
	SourceClose = "close"
	SourceHigh  = "high"
	SourceLow   = "low"
	SourceHL2   = "hl2"
	SourceHLC3  = "hlc3"
	SourceOHLC4 = "ohlc4"
)

func NewMacdConfig(fastLength, slowLength, signalLength int, source Source) *macdConfig {
	return &macdConfig{
		fastLength:   fastLength,
		slowLength:   slowLength,
		signalLength: signalLength,
		source:       source,
	}
}

func (conf *macdConfig) CalculateMacd(candles []models.Candle, appendCandles bool) error {
	var rangeCounter int
	if appendCandles {
		rangeCounter = len(conf.Candles)
		conf.Candles = append(conf.Candles, candles...)
	} else {
		conf.Candles = candles
		rangeCounter = conf.slowLength - 1
	}
	if err := conf.validate(); err != nil {
		return err
	}
	slowEMA := NewMovingAverageConfig(conf.slowLength, conf.source)
	fastEMA := NewMovingAverageConfig(conf.fastLength, conf.source)
	signalEMA := NewMovingAverageConfig(conf.signalLength, conf.source)
	if err := slowEMA.CalculateExponential(cloneCandles(conf.Candles[rangeCounter-conf.slowLength+1:]), false); err != nil {
		return err
	}
	if err := fastEMA.CalculateExponential(cloneCandles(conf.Candles[rangeCounter-conf.fastLength+1:]), false); err != nil {
		return err
	}
	if err := signalEMA.CalculateExponential(cloneCandles(conf.Candles[rangeCounter-conf.signalLength+1:]), false); err != nil {
		return err
	}

	for i, candle := range conf.Candles[rangeCounter:] {
		candle.MACD.MACD = fastEMA.Candles[conf.fastLength+i-1].MovingAverage.Exponential - slowEMA.Candles[conf.slowLength+i-1].MovingAverage.Exponential
		candle.MACD.Signal = signalEMA.Candles[conf.signalLength+i-1].MovingAverage.Exponential
	}

	//conf.lastSlowValue = slowEMA.Candles[len(slowEMA.Candles)-1].MovingAverage.Exponential
	//conf.lastFastValue = fastEMA.Candles[len(fastEMA.Candles)-1].MovingAverage.Exponential
	//conf.lastSignalValue = signalEMA.Candles[len(signalEMA.Candles)-1].MovingAverage.Exponential

	return nil
}

func (conf *macdConfig) validate() error {
	if conf.slowLength < conf.fastLength {
		return errors.New("slow length must be bigger than fast length")
	}
	if len(conf.Candles) < conf.slowLength {
		return errors.New(fmt.Sprintf("candles length must be grater or equal than %d", conf.slowLength))
	}
	return nil
}

func cloneCandles(candles []models.Candle) []models.Candle {
	return append([]models.Candle{}, candles...)
}
