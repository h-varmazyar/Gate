package indicators

import (
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/models"
)

func (conf *Configuration) CalculateMACD(candles []models.Candle) error {
	if err := conf.validateMACD(len(candles)); err != nil {
		return err
	}
	slowEMA := Configuration{
		MovingAverageSource: conf.MacdSource,
		MovingAverageLength: conf.MacdSlowLength,
		MacdSlowLength:      conf.MacdSlowLength,
	}
	slowCandles := cloneCandles(candles)
	fastEMA := Configuration{
		MovingAverageSource: conf.MacdSource,
		MovingAverageLength: conf.MacdFastLength,
		MacdFastLength:      conf.MacdFastLength,
	}
	fastCandles := cloneCandles(candles)
	signalEMA := Configuration{
		MovingAverageSource: SourceClose,
		MovingAverageLength: conf.MacdSignalLength,
		MacdSignalLength:    conf.MacdSignalLength,
		From:                conf.MacdSlowLength - 1,
	}
	signalCandles := cloneCandles(candles)
	if err := slowEMA.CalculateEMA(slowCandles); err != nil {
		return err
	}
	if err := fastEMA.CalculateEMA(fastCandles); err != nil {
		return err
	}

	for i := conf.MacdSlowLength - 1; i < len(candles); i++ {
		candles[i].MACD.FastEMA = fastCandles[i].Exponential
		candles[i].MACD.SlowEMA = slowCandles[i].Exponential
		candles[i].MACD.MACD = fastCandles[i].Exponential - slowCandles[i].Exponential
		signalCandles[i].Close = candles[i].MACD.MACD
	}
	if err := signalEMA.CalculateEMA(signalCandles); err != nil {
		return err
	}
	for i := conf.MacdSlowLength + conf.MacdSignalLength - 2; i < len(candles); i++ {
		candles[i].MACD.Signal = signalCandles[i].Exponential
	}
	return nil
}

func (conf *Configuration) UpdateMACD(candles []models.Candle) {
	lastIndex := len(candles) - 1

	price := float64(0)
	fastFactor := 2 / float64(conf.MacdFastLength+1)
	slowFactor := 2 / float64(conf.MacdSlowLength+1)
	signalFactor := 2 / float64(conf.MacdSignalLength+1)
	switch conf.MovingAverageSource {
	case SourceClose:
		price = candles[lastIndex].Close
	case SourceOpen:
		price = candles[lastIndex].Open
	case SourceHigh:
		price = candles[lastIndex].High
	case SourceLow:
		price = candles[lastIndex].Low
	case SourceHL2:
		price = (candles[lastIndex].High + candles[lastIndex].Low) / 2
	case SourceHLC3:
		price = (candles[lastIndex].High + candles[lastIndex].Low + candles[lastIndex].Close) / 3
	case SourceOHLC4:
		price = (candles[lastIndex].Open + candles[lastIndex].Close + candles[lastIndex].High + candles[lastIndex].Low) / 4
	}

	candles[lastIndex].FastEMA = price*fastFactor + candles[lastIndex-1].FastEMA*(1-fastFactor)
	candles[lastIndex].SlowEMA = price*slowFactor + candles[lastIndex-1].SlowEMA*(1-slowFactor)
	candles[lastIndex].MACD.MACD = candles[lastIndex].FastEMA - candles[lastIndex].SlowEMA
	candles[lastIndex].Signal = candles[lastIndex].MACD.MACD*signalFactor + candles[lastIndex-1].Signal*(1-signalFactor)
}

func (conf *Configuration) validateMACD(length int) error {
	if conf.MacdSlowLength < conf.MacdFastLength {
		return errors.New("slow length must be bigger than fast length")
	}
	if length < conf.MacdSlowLength {
		return errors.New(fmt.Sprintf("candles length must be grater or equal than %d", conf.MacdSlowLength))
	}
	if conf.MacdSource == "" {
		return errors.New("macd source must be declared")
	}
	return nil
}
