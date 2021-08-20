package indicators

import (
	"github.com/mrNobody95/Gate/models"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

func (conf *IndicatorConfig) CalculateCandleIndicators(candles []models.Candle, firstTime bool) {
	if firstTime {
		conf.Candles = append(conf.Candles, candles...)
		conf.calculateIndicators()
		for _, candle := range conf.Candles {
			if err := candle.Create(); err != nil {
				log.Errorf("creating candle failed:(brokerage %s - resolution: %s - symbol: %s - date: %s)",
					candle.Brokerage, candle.Resolution.Label, candle.Symbol, time.Unix(candle.Time, 0))
			}
		}
	} else {
		for _, candle := range candles {
			lastIndex := len(conf.Candles) - 1
			if candle.Time == conf.Candles[lastIndex].Time {
				conf.Candles[lastIndex] = candle
			} else {
				conf.Candles = append(conf.Candles, candle)
			}
			conf.updateIndicators()
			if len(conf.Candles)-1 == lastIndex {
				if err := conf.Candles[len(conf.Candles)-1].Update(); err != nil {
					log.Errorf("updating candle failed:(brokerage %s - resolution: %s - symbol: %s - date: %s)",
						candle.Brokerage, candle.Resolution.Label, candle.Symbol, time.Unix(candle.Time, 0))
				}
			} else {
				if err := conf.Candles[len(conf.Candles)-1].Create(); err != nil {
					log.Errorf("creating candle failed:(brokerage %s - resolution: %s - symbol: %s - date: %s)",
						candle.Brokerage, candle.Resolution.Label, candle.Symbol, time.Unix(candle.Time, 0))
				}
			}

		}
	}
}

func (conf *IndicatorConfig) calculateIndicators() {
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
		if err := conf.CalculateMACD(); err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		if err := conf.CalculatePSAR(); err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		if err := conf.CalculateRSI(); err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		if err := conf.CalculateStochastic(); err != nil {
			log.Error(err)
		}
	}(&wg)
	wg.Wait()
}

func (conf *IndicatorConfig) updateIndicators() {
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
		if err := conf.UpdateRSI(); err != nil {
			log.Error(err)
		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := conf.UpdateStochastic(); err != nil {
			log.Error(err)
		}
	}(&wg)

	wg.Wait()
}
