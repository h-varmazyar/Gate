package indicators

import (
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/storage"
	log "github.com/sirupsen/logrus"
	"sync"
)

//func (conf *Configuration) PreCalculation(market *models.Market, resolution *models.Resolution, startFrom time.Time, size int) (int64, error) {
//
//}
//
//func (conf *Configuration) CalculateIndicators(candles []models.Candle) {
//	if len(conf.Candles) == 0 {
//		conf.Candles = append(conf.Candles, candles...)
//		conf.calculateIndicators()
//	} else {
//		for _, candle := range candles {
//			if conf.Candles[len(conf.Candles)-1].Time == candle.Time {
//				conf.Candles[len(conf.Candles)-1] = candle
//			} else {
//				conf.Candles = append(conf.Candles, candle)
//			}
//			conf.updateIndicators()
//		}
//	}
//	if len(conf.Candles) > size {
//		conf.Candles = conf.Candles[len(conf.Candles)-size:]
//	}
//}

func (conf *Configuration) CalculateIndicators(candles []models.Candle) error {
	if err := conf.CalculateBollingerBand(candles); err != nil {
		return err
	}
	if err := conf.CalculateMACD(candles); err != nil {
		return err
	}
	if err := conf.CalculatePSAR(candles); err != nil {
		return err
	}
	if err := conf.CalculateRSI(candles); err != nil {
		return err
	}
	if err := conf.CalculateStochastic(candles); err != nil {
		return err
	}
	return nil
}

func (conf *Configuration) UpdateIndicators(pool *storage.CandlePool) {
	var wg sync.WaitGroup
	wg.Add(5)
	go func(wg *sync.WaitGroup, pool *storage.CandlePool) {
		defer wg.Done()
		candles := pool.GetLastNCandle(conf.BollingerLength)
		if err := conf.UpdateBollingerBand(candles); err != nil {
			log.Error(err)
		}
		if err := pool.UpdateLastCandle(candles[len(candles)-1]); err != nil {
			log.WithError(err).Error("update bollinger bands failed")
		}
	}(&wg, pool)
	go func(wg *sync.WaitGroup, pool *storage.CandlePool) {
		defer wg.Done()
		candles := pool.GetLastNCandle(2)
		conf.UpdateMACD(candles)
		if err := pool.UpdateLastCandle(candles[1]); err != nil {
			log.Error(err)
		}
	}(&wg, pool)
	go func(wg *sync.WaitGroup, pool *storage.CandlePool) {
		defer wg.Done()
		candles := pool.GetLastNCandle(3)
		conf.UpdatePSAR(candles)
		if err := pool.UpdateLastCandle(candles[1]); err != nil {
			log.Error(err)
		}
	}(&wg, pool)
	go func(wg *sync.WaitGroup, pool *storage.CandlePool) {
		defer wg.Done()
		candles := pool.GetLastNCandle(2)
		conf.UpdateRSI(candles)
		if err := pool.UpdateLastCandle(candles[1]); err != nil {
			log.Error(err)
		}
	}(&wg, pool)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		candles := pool.GetLastNCandle(conf.StochasticLength)
		conf.UpdateStochastic(candles)
		if err := pool.UpdateLastCandle(candles[1]); err != nil {
			log.Error(err)
		}
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
