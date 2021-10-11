package core

import (
	"github.com/fatih/color"
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/indicators"
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/storage"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type MarketThread struct {
	*Node
	Market          *models.Market
	StartFrom       time.Time
	Resolution      *models.Resolution
	IndicatorConfig *indicators.Configuration
	CandlePool      *storage.CandlePool
}

func (thread *MarketThread) CollectPrimaryData() error {
	color.HiGreen("Collecting primary data for: %s", thread.Market.Name)

	if thread.IndicatorConfig == nil {
		thread.IndicatorConfig = indicators.DefaultConfig()
	}
	lastTime := thread.StartFrom
	var list []models.Candle
	for {
		time.Sleep(time.Second)
		if tmpList, err := models.LoadCandleList(thread.Market.Id, thread.PivotResolution.Id, lastTime); err != nil {
			if err == gorm.ErrRecordNotFound {
				break
			} else {
				return err
			}
		} else if len(tmpList) > 0 {
			list = append(list, tmpList...)
			lastTime = tmpList[len(tmpList)-1].Time
		} else {
			break
		}
	}
	for i := 0; i < len(list); i++ {
		list[i].FromDb = true
	}

	count := (time.Now().Unix() - lastTime.Unix()) / int64(thread.PivotResolution.Duration/time.Second)
	j := count / 500
	if count%500 != 0 {
		j++
	}
	for i := int64(1); i <= j; i++ {
		from := lastTime.Unix() + 500*(i-1)*int64(thread.PivotResolution.Duration/time.Second)
		to := lastTime.Unix() + 500*i*int64(thread.PivotResolution.Duration/time.Second)
		if candles, err := thread.makeOHLCRequest(thread.PivotResolution, from, to); err != nil {
			return err
		} else {
			if len(list) > 0 && list[len(list)-1].Time == candles[0].Time {
				candles[0].ID = list[len(list)-1].ID
				list = append(list[:len(list)-1], candles...)
			} else {
				list = append(list, candles...)
			}
		}
	}
	if err := thread.IndicatorConfig.CalculateIndicators(list); err != nil {
		return err
	}
	if err := thread.CandlePool.ImportNewCandles(list); err != nil {
		return err
	}
	return nil
}

func (thread *MarketThread) PeriodicOHLC() {
	color.HiGreen("Making Periodic ohlc for: %s", thread.Market.Name)
	for {
		start := time.Now()
		if candles, err := thread.makeOHLCRequest(thread.PivotResolution, thread.CandlePool.GetLastCandle().Time.Unix(), time.Now().Unix()); err != nil {
			log.Errorf("ohlc request failed: %s", err.Error())
		} else {
			for _, candle := range candles {
				if poolErr := thread.CandlePool.UpdateLastCandle(candle); poolErr != nil {
					log.WithError(poolErr).Errorf("update pool failed for market %s in timeframe %s",
						candle.Market.Name, candle.Resolution.Label)
					continue
				}
				thread.IndicatorConfig.UpdateIndicators(thread.CandlePool)
			}
		}
		thread.dataChannel <- *thread.CandlePool.GetLastCandle()
		if thread.EnableTrading || thread.FakeTrading {
			thread.checkForSignals()
		}
		end := time.Now()
		idealTime := thread.Strategy.IndicatorUpdatePeriod - end.Sub(start)
		if idealTime > 0 {
			time.Sleep(idealTime)
		}
	}
}

func (thread *MarketThread) makeOHLCRequest(resolution *models.Resolution, from, to int64) ([]models.Candle, error) {
	response := thread.Requests.OHLC(brokerages.OHLCParams{
		Resolution: resolution,
		Market:     thread.Market,
		From:       from,
		To:         to,
	})
	if response.Error != nil {
		return nil, response.Error
	}
	return response.Candles, nil
}

func (thread *MarketThread) checkForSignals() {
}
