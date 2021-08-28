package indicators

import (
	"github.com/mrNobody95/Gate/models"
	log "github.com/sirupsen/logrus"
	"sync"
)

func (conf *Configuration) CalculateIndicators(candles []models.Candle) {
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
