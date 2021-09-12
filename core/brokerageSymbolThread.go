package core

import (
	"github.com/mrNobody95/Gate/brokerages/nobitex"
	"github.com/mrNobody95/Gate/indicators"
	"github.com/mrNobody95/Gate/models"
	log "github.com/sirupsen/logrus"
	"time"
)

type BrokerageSymbolThread struct {
	*Node
	Symbol          models.Symbol
	StartFrom       time.Time
	Resolutions     []models.Resolution
	SymbolConfig    map[uint]*indicators.Configuration
	IndicatorConfig string
}

func (thread *BrokerageSymbolThread) CollectPrimaryData() {
	for _, resolution := range thread.Resolutions {
		go func(resolution models.Resolution) {
			conf, ok := thread.SymbolConfig[resolution.Id]
			if !ok {
				conf = indicators.DefaultConfig()
			}
			lastTime, err := conf.PreCalculation(thread.Symbol, resolution, thread.StartFrom, thread.Strategy.CandleBufferLength)
			if err != nil {
				log.Errorf("fetch first candle failed: %v", err)
				return
			}
			count := (time.Now().Unix() - lastTime) / int64(resolution.Duration)

			for i := int64(0); i < count; i += 500 {
				from := lastTime + 500*i*int64(resolution.Duration)
				to := lastTime + 500*(i+1)*int64(resolution.Duration)
				if err := thread.makeOHLCRequest(conf, resolution, from, to); err != nil {
					log.Errorf("ohlc request failed: %s", err.Error())
				}
			}
			//todo: check next line
			//thread.SymbolConfig[resolution.Id]=conf
		}(resolution)
	}
}

func (thread *BrokerageSymbolThread) PeriodicOHLC() {
	for _, resolution := range thread.Resolutions {
		go func(resolution models.Resolution) {
			conf, ok := thread.SymbolConfig[resolution.Id]
			if !ok {
				conf = indicators.DefaultConfig()
			}
			for {
				start := time.Now()
				err := thread.makeOHLCRequest(conf, resolution, conf.Candles[len(conf.Candles)-1].Time.Unix(), time.Now().Unix())
				if err != nil {
					log.Errorf("ohlc request failed: %s", err.Error())
				}
				//todo: check next line
				//thread.SymbolConfig[resolution.Id]=conf
				thread.CheckForSignals()
				end := time.Now()
				idealTime := thread.Strategy.PeriodDuration - end.Sub(start)
				if idealTime > 0 {
					time.Sleep(idealTime)
				}
			}

		}(resolution)
	}
}

func (thread *BrokerageSymbolThread) makeOHLCRequest(conf *indicators.Configuration, resolution models.Resolution, from, to int64) error {
	log.Infof("make ohlc request for %s in %s with resolution %s from %v to %v",
		thread.Symbol.Value, thread.Brokerage.Name, resolution.Label,
		time.Unix(from, 0).Format("02/01/2006, 15:04:05"),
		time.Unix(to, 0).Format("02/01/2006, 15:04:05"))

	response := thread.Requests.OHLC(nobitex.OHLCParams{
		Resolution: resolution,
		Symbol:     thread.Symbol,
		From:       from,
		To:         to,
	})
	if response.Error != nil {
		return response.Error
	}
	conf.CalculateIndicators(response.Candles, thread.Strategy.CandleBufferLength)
	go func(symbol string, brokerage models.BrokerageName, candles []models.Candle) {
		for _, candle := range candles {
			if err := candle.Create(); err != nil {
				log.Errorf("creating candle failed for %s in %s at %s",
					symbol, brokerage, candle.Time.Format("02/01/2006, 15:04:05"))
			}
		}
	}(thread.Symbol.Value, thread.Brokerage.Name, response.Candles)
	return nil
}

func (thread *BrokerageSymbolThread) CheckForSignals() {
	for _, resolution := range thread.Resolutions {
		thread.SymbolConfig[resolution.Id].CheckIndicatorSignals()
		//todo: calculate lines and patterns
	}
}
