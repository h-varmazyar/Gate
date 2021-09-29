package indicators

import (
	"errors"
	"fmt"
)

func (conf *Configuration) CalculateMACD() error {
	if err := conf.validateMACD(); err != nil {
		return err
	}
	slowEMA := Configuration{
		MovingAverageSource: conf.MacdSource,
		MovingAverageLength: conf.MacdSlowLength,
		MacdSlowLength:      conf.MacdSlowLength,
		Candles:             cloneCandles(conf.Candles),
	}
	fastEMA := Configuration{
		MovingAverageSource: conf.MacdSource,
		MovingAverageLength: conf.MacdFastLength,
		MacdFastLength:      conf.MacdFastLength,
		Candles:             cloneCandles(conf.Candles),
	}
	signalEMA := Configuration{
		MovingAverageSource: SourceClose,
		MovingAverageLength: conf.MacdSignalLength,
		MacdSignalLength:    conf.MacdSignalLength,
		Candles:             cloneCandles(conf.Candles),
		From:                conf.MacdSlowLength - 1,
	}
	if err := slowEMA.CalculateEMA(); err != nil {
		return err
	}
	if err := fastEMA.CalculateEMA(); err != nil {
		return err
	}

	indicatorLock.Lock()
	for i := conf.MacdSlowLength - 1; i < len(conf.Candles); i++ {
		conf.Candles[i].MACD.FastEMA = fastEMA.Candles[i].Exponential
		conf.Candles[i].MACD.SlowEMA = slowEMA.Candles[i].Exponential
		conf.Candles[i].MACD.MACD = fastEMA.Candles[i].Exponential - slowEMA.Candles[i].Exponential
		signalEMA.Candles[i].Close = conf.Candles[i].MACD.MACD
	}
	indicatorLock.Unlock()
	if err := signalEMA.CalculateEMA(); err != nil {
		return err
	}
	indicatorLock.Lock()
	for i := conf.MacdSlowLength + conf.MacdSignalLength - 2; i < len(conf.Candles); i++ {
		conf.Candles[i].MACD.Signal = signalEMA.Candles[i].Exponential
	}
	indicatorLock.Unlock()
	return nil
}

func (conf *Configuration) UpdateMACD() error {
	lastIndex := len(conf.Candles) - 1

	price := float64(0)
	fastFactor := 2 / float64(conf.MacdFastLength+1)
	slowFactor := 2 / float64(conf.MacdSlowLength+1)
	signalFactor := 2 / float64(conf.MacdSignalLength+1)
	switch conf.MovingAverageSource {
	case SourceClose:
		price = conf.Candles[lastIndex].Close
	case SourceOpen:
		price = conf.Candles[lastIndex].Open
	case SourceHigh:
		price = conf.Candles[lastIndex].High
	case SourceLow:
		price = conf.Candles[lastIndex].Low
	case SourceHL2:
		price = (conf.Candles[lastIndex].High + conf.Candles[lastIndex].Low) / 2
	case SourceHLC3:
		price = (conf.Candles[lastIndex].High + conf.Candles[lastIndex].Low + conf.Candles[lastIndex].Close) / 3
	case SourceOHLC4:
		price = (conf.Candles[lastIndex].Open + conf.Candles[lastIndex].Close + conf.Candles[lastIndex].High + conf.Candles[lastIndex].Low) / 4
	}

	indicatorLock.Lock()
	conf.Candles[lastIndex].FastEMA = price*fastFactor + conf.Candles[lastIndex-1].FastEMA*(1-fastFactor)
	conf.Candles[lastIndex].SlowEMA = price*slowFactor + conf.Candles[lastIndex-1].SlowEMA*(1-slowFactor)
	conf.Candles[lastIndex].MACD.MACD = conf.Candles[lastIndex].FastEMA - conf.Candles[lastIndex].SlowEMA
	conf.Candles[lastIndex].Signal = conf.Candles[lastIndex].MACD.MACD*signalFactor + conf.Candles[lastIndex-1].Signal*(1-signalFactor)
	indicatorLock.Unlock()
	return nil
}

func (conf *Configuration) validateMACD() error {
	if conf.MacdSlowLength < conf.MacdFastLength {
		return errors.New("slow length must be bigger than fast length")
	}
	if len(conf.Candles) < conf.MacdSlowLength {
		return errors.New(fmt.Sprintf("candles length must be grater or equal than %d", conf.MacdSlowLength))
	}
	if conf.MacdSource == "" {
		return errors.New("macd source must be declared")
	}
	return nil
}
