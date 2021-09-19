package indicators

import (
	"github.com/mrNobody95/Gate/models"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

func (conf *Configuration) PreCalculation(market models.Market, resolution models.Resolution, startFrom time.Time, size int) (int64, error) {
	candle := models.Candle{
		Market:     market,
		Resolution: resolution,
		Time:       startFrom,
	}
	for {
		list, err := candle.LoadList()
		if err != nil {
			if err.Error() == "record not found" {
				return candle.Time.Unix(), nil
			} else {
				return 0, err
			}
		}
		if len(list) > 0 {
			conf.CalculateIndicators(list, size)
			candle.Time = list[len(list)-1].Time.Add(candle.Resolution.Duration)
		} else {
			return candle.Time.Unix(), nil
		}
	}
}

func (conf *Configuration) CalculateIndicators(candles []models.Candle, size int) {
	if len(conf.Candles) == 0 {
		conf.Candles = append(conf.Candles, candles...)
		conf.calculateIndicators()
	} else {
		for _, candle := range candles {
			if conf.Candles[len(conf.Candles)-1].Time == candle.Time {
				conf.Candles[len(conf.Candles)-1] = candle
			} else {
				conf.Candles = append(conf.Candles, candle)
			}
			conf.updateIndicators()
		}
	}
	if len(conf.Candles) > size {
		conf.Candles = conf.Candles[len(conf.Candles)-size:]
	}
}

func (conf *Configuration) calculateIndicators() {
	var wg sync.WaitGroup
	wg.Add(6)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := conf.CalculateADX(); err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := conf.CalculateBollingerBand(); err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := conf.CalculateMACD(); err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := conf.CalculatePSAR(); err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := conf.CalculateRSI(); err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := conf.CalculateStochastic(); err != nil {
			log.Error(err)
		}
	}(&wg)
	wg.Wait()
}

func (conf *Configuration) updateIndicators() {
	var wg sync.WaitGroup
	wg.Add(6)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		conf.UpdateADX()
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := conf.UpdateBollingerBand(); err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := conf.UpdateMACD(); err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := conf.UpdatePSAR(); err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		conf.UpdateRSI()
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		conf.UpdateStochastic()
	}(&wg)

	wg.Wait()
}

func (conf *Configuration) CheckIndicatorSignals() {
	//stochastic:=conf.Candles[len(conf.Candles)-1].Stochastic.SignalStrength()
	//bb:=conf.Candles[len(conf.Candles)-1].BollingerBand.SignalStrength()
	//rsi:=conf.Candles[len(conf.Candles)-1].RSI.SignalStrength()
	//atr:=conf.Candles[len(conf.Candles)-1].ATR.SignalStrength()
	//adx:=conf.Candles[len(conf.Candles)-1].ADX.SignalStrength()
	//pSar:=conf.Candles[len(conf.Candles)-1].ParabolicSAR.SignalStrength()
	//macd:=conf.Candles[len(conf.Candles)-1].MACD.SignalStrength()
	//ma:=conf.Candles[len(conf.Candles)-1].MovingAverage.SignalStrength()
}
