package indicators

import (
	"errors"
	"fmt"
)

func (conf *Configuration) CalculateMACD() error {
	if err := conf.validateMA(); err != nil {
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
		MovingAverageSource: conf.MacdSource,
		MovingAverageLength: conf.MacdSignalLength,
		MacdSignalLength:    conf.MacdSignalLength,
		Candles:             cloneCandles(conf.Candles),
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

	for i := conf.MacdSlowLength - 1; i < len(conf.Candles); i++ {
		indicatorLock.Lock()
		conf.Candles[i].MACD.MACD = fastEMA.Candles[i].MovingAverage.Exponential - slowEMA.Candles[i].MovingAverage.Exponential
		conf.Candles[i].MACD.Signal = signalEMA.Candles[i].MovingAverage.Exponential
		indicatorLock.Unlock()
	}
	return nil
}

func (conf *Configuration) UpdateMACD() error {
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
		MovingAverageSource: conf.MacdSource,
		MovingAverageLength: conf.MacdSignalLength,
		MacdSignalLength:    conf.MacdSignalLength,
		Candles:             cloneCandles(conf.Candles),
	}
	if err := slowEMA.UpdateEMA(); err != nil {
		return err
	}
	if err := fastEMA.UpdateEMA(); err != nil {
		return err
	}
	if err := signalEMA.UpdateEMA(); err != nil {
		return err
	}

	//color.HiYellow("lenght: %d", len(conf.Candles))
	//color.HiYellow("lenght: %d", len(fastEMA.Candles))
	//color.HiYellow("lenght: %d", len(signalEMA.Candles))
	//color.HiYellow("lenght: %d", len(slowEMA.Candles))
	indicatorLock.Lock()
	conf.Candles[len(conf.Candles)-1].MACD.MACD = fastEMA.Candles[len(fastEMA.Candles)-1].MovingAverage.Exponential - slowEMA.Candles[len(slowEMA.Candles)-1].MovingAverage.Exponential
	conf.Candles[len(conf.Candles)-1].MACD.Signal = signalEMA.Candles[len(signalEMA.Candles)-1].MovingAverage.Exponential
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
