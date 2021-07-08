package indicators

import (
	"errors"
	"fmt"
)

//type macdConfig struct {
//	Candles      []models.Candle
//	fastLength   int
//	slowLength   int
//	signalLength int
//	source       Source
//}

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

//func NewMacdConfig(fastLength, slowLength, signalLength int, source Source) *macdConfig {
//	return &macdConfig{
//		fastLength:   fastLength,
//		slowLength:   slowLength,
//		signalLength: signalLength,
//		source:       source,
//	}
//}

func (conf *IndicatorConfig) CalculateMACD() error {
	rangeCounter := conf.slowLength - 1
	if err := conf.validateMA(); err != nil {
		return err
	}
	slowEMA := IndicatorConfig{
		slowLength: conf.slowLength,
		source:     conf.source,
		Candles:    cloneCandles(conf.Candles[rangeCounter-conf.slowLength+1:]),
	}
	fastEMA := IndicatorConfig{
		fastLength: conf.fastLength,
		source:     conf.source,
		Candles:    cloneCandles(conf.Candles[rangeCounter-conf.fastLength+1:]),
	}
	signalEMA := IndicatorConfig{
		signalLength: conf.signalLength,
		source:       conf.source,
		Candles:      cloneCandles(conf.Candles[rangeCounter-conf.signalLength+1:]),
	}
	if err := slowEMA.CalculateEMA(); err != nil {
		return err
	}
	if err := fastEMA.CalculateEMA(); err != nil {
		return err
	}
	if err := signalEMA.CalculateEMA(); err != nil {
		return err
	}

	for i, candle := range conf.Candles[rangeCounter:] {
		candle.MACD.MACD = fastEMA.Candles[conf.fastLength+i-1].MovingAverage.Exponential - slowEMA.Candles[conf.slowLength+i-1].MovingAverage.Exponential
		candle.MACD.Signal = signalEMA.Candles[conf.signalLength+i-1].MovingAverage.Exponential
	}
	return nil
}

func (conf *IndicatorConfig) UpdateMACD() error {
	lastIndex := len(conf.Candles)

	slowEMA := IndicatorConfig{
		slowLength: conf.slowLength,
		source:     conf.source,
		Candles:    cloneCandles(conf.Candles[lastIndex-conf.slowLength:]),
	}
	fastEMA := IndicatorConfig{
		fastLength: conf.fastLength,
		source:     conf.source,
		Candles:    cloneCandles(conf.Candles[lastIndex-conf.fastLength:]),
	}
	signalEMA := IndicatorConfig{
		signalLength: conf.signalLength,
		source:       conf.source,
		Candles:      cloneCandles(conf.Candles[lastIndex-conf.signalLength:]),
	}
	if err := slowEMA.CalculateEMA(); err != nil {
		return err
	}
	if err := fastEMA.CalculateEMA(); err != nil {
		return err
	}
	if err := signalEMA.CalculateEMA(); err != nil {
		return err
	}

	indicatorLock.Lock()
	conf.Candles[lastIndex-1].MACD.MACD = fastEMA.Candles[conf.fastLength-1].MovingAverage.Exponential - slowEMA.Candles[conf.slowLength-1].MovingAverage.Exponential
	conf.Candles[lastIndex-1].MACD.Signal = signalEMA.Candles[conf.signalLength-1].MovingAverage.Exponential
	indicatorLock.Unlock()
	return nil
}

func (conf *IndicatorConfig) validateMACD() error {
	if conf.slowLength < conf.fastLength {
		return errors.New("slow length must be bigger than fast length")
	}
	if len(conf.Candles) < conf.slowLength {
		return errors.New(fmt.Sprintf("candles length must be grater or equal than %d", conf.slowLength))
	}
	return nil
}
