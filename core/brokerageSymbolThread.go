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
	BufferedCandles map[uint][]models.Candle
	IndicatorConfig string
}

func (thread *BrokerageSymbolThread) CollectPrimaryData() {
	for _, resolution := range thread.Resolutions {
		go func(resolution models.Resolution) {
			conf := indicators.Configuration{
				Candles: make([]models.Candle, 0),
				//Length:  thread.Strategy.IndicatorCalcLength,
			}
			candle := models.Candle{
				Symbol:     thread.Symbol,
				Resolution: resolution,
			}
			if err := candle.LoadLast(); err != nil {
				if err.Error() == "record not found" {
					candle.Time = thread.StartFrom
				} else {
					log.Errorf("fetch first candle failed: %v", err)
					return
				}
			} else {
				conf.Candles = append(conf.Candles, candle)
			}
			count := (time.Now().Unix() - candle.Time.Unix()) / int64(resolution.Duration)

			for i := int64(0); i < count; i += 500 {
				from := candle.Time.Unix() + 500*i*int64(resolution.Duration)
				to := candle.Time.Unix() + 500*(i+1)*int64(resolution.Duration)
				if err := thread.makeOHLCRequest(&conf, resolution, from, to); err != nil {
					log.Errorf("ohlc request failed: %s", err.Error())
				}
			}
		}(resolution)
	}
}

func (thread *BrokerageSymbolThread) PeriodicOHLC() {
	for _, resolution := range thread.Resolutions {
		go func(resolution models.Resolution) {
			conf := indicators.Configuration{
				Candles: make([]models.Candle, 0),
				//Length:  thread.Strategy.IndicatorCalcLength,
			}
			candle := models.Candle{
				Symbol:     thread.Symbol,
				Resolution: resolution,
			}
			if err := candle.LoadLast(); err != nil {
				log.Errorf("fetch last candle failed: %v", err)
				return
			} else {
				conf.Candles = append(conf.Candles, candle)
			}
			if err := thread.makeOHLCRequest(&conf, resolution, candle.Time.Unix(), time.Now().Unix()); err != nil {
				log.Errorf("ohlc request failed: %s", err.Error())
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
	conf.CalculateIndicators(response.Candles)
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
