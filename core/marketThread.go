package core

import (
	"github.com/fatih/color"
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/indicators"
	"github.com/mrNobody95/Gate/models"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type MarketThread struct {
	*Node
	Market                       models.Market
	StartFrom                    time.Time
	IndicatorConfigPerResolution map[uint]*indicators.Configuration
}

func (thread *MarketThread) CollectPrimaryData() {
	color.HiGreen("Collecting primary data for: %s", thread.Market.Value)
	wg := sync.WaitGroup{}
	wg.Add(len(thread.Resolutions))
	for _, resolution := range thread.Resolutions {
		go func(resolution models.Resolution) {
			defer wg.Done()
			conf, ok := thread.IndicatorConfigPerResolution[resolution.Id]
			if !ok {
				conf = indicators.DefaultConfig()
			}
			lastTime, err := conf.PreCalculation(thread.Market, resolution, thread.StartFrom, thread.Strategy.BufferedCandleCount)
			if err != nil {
				log.Errorf("fetch first candle failed: %v", err)
				return
			}
			count := (time.Now().Unix() - lastTime) / int64(resolution.Duration/time.Second)
			j := count / 500
			if count%500 != 0 {
				j++
			}
			for i := int64(1); i <= j; i++ {
				from := lastTime + 500*(i-1)*int64(resolution.Duration/time.Second)
				to := lastTime + 500*i*int64(resolution.Duration/time.Second)
				if err := thread.makeOHLCRequest(conf, resolution, from, to); err != nil {
					log.Errorf("ohlc request failed: %s", err.Error())
				}
			}
			//todo: check next line
			//thread.IndicatorConfigPerResolution[resolution.Id]=conf
		}(resolution)
	}
	wg.Wait()
}

func (thread *MarketThread) PeriodicOHLC() {
	color.HiGreen("Making Periodic ohlc for: %s", thread.Market.Value)
	for _, resolution := range thread.Resolutions {
		go func(resolution models.Resolution) {
			conf, ok := thread.IndicatorConfigPerResolution[resolution.Id]
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
				//thread.IndicatorConfigPerResolution[resolution.Id]=conf
				//data:=conf.Candles[len(conf.Candles)-1]
				//fmt.Printf("| %-7.3f | %-7.3f | %-7.3f | %-7.3f | %-7.3f |\n",
				//	data.Open, data.High, data.LowerBond, data.Close, data.Vol)
				thread.dataChannel <- conf.Candles[len(conf.Candles)-1]
				if thread.EnableTrading || thread.FakeTrading {
					thread.checkForSignals()
				}
				end := time.Now()
				idealTime := thread.Strategy.IndicatorUpdatePeriod - end.Sub(start)
				if idealTime > 0 {
					time.Sleep(idealTime)
				}
			}

		}(resolution)
	}
}

func (thread *MarketThread) makeOHLCRequest(conf *indicators.Configuration, resolution models.Resolution, from, to int64) error {
	response := thread.Requests.OHLC(brokerages.OHLCParams{
		Resolution: resolution,
		Market:     thread.Market,
		From:       from,
		To:         to,
	})
	if response.Error != nil {
		return response.Error
	}
	if len(response.Candles) == 0 {
		return nil
	}
	conf.CalculateIndicators(response.Candles, thread.Strategy.BufferedCandleCount)
	go func(symbol string, brokerage models.BrokerageName, candles []models.Candle) {
		for _, candle := range candles {
			if err := candle.CreateOrUpdate(); err != nil {
				log.Errorf("creating candle failed for %s in %s at %s",
					symbol, brokerage, candle.Time.Format("02/01/2006, 15:04:05"))
			}
		}
	}(thread.Market.Value, thread.Brokerage.Name, response.Candles)
	return nil
}

func (thread *MarketThread) checkForSignals() {
	for _, resolution := range thread.Resolutions {
		thread.IndicatorConfigPerResolution[resolution.Id].CheckIndicatorSignals()
		//todo: calculate lines and patterns
	}
}
