package OHLC

import (
	"context"
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/pkg/mapper"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/buffer"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/indicators"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/repository"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type Worker struct {
	HeartbeatInterval time.Duration
	CandleService     brokerageApi.CandleServiceClient
	Cancellations     map[string]context.CancelFunc
}

type WorkerSettings struct {
	Context    context.Context
	Market     *brokerageApi.Market
	Resolution *brokerageApi.Resolution
	Indicators []indicators.Indicator
}

func (worker *Worker) AddMarket(settings *WorkerSettings) {
	worker.Cancellations[fmt.Sprintf("%d > %v", settings.Market.ID, settings.Resolution.ID)]()
	settings.Context, worker.Cancellations[fmt.Sprintf("%d > %d", settings.Market.ID, settings.Resolution.ID)] = context.WithCancel(context.Background())
	buffer.Candles.AddList(settings.Market.ID, settings.Resolution.ID)
	go worker.run(settings)
}

func (worker *Worker) CancelWorker(marketID, resolutionID uint32) error {
	fn, ok := worker.Cancellations[fmt.Sprintf("%d > %d", marketID, resolutionID)]
	if !ok {
		return errors.New("worker stopped before")
	}
	fn()
	delete(worker.Cancellations, fmt.Sprintf("%d > %d", marketID, resolutionID))
	buffer.Candles.RemoveList(marketID, resolutionID)
	return nil
}

func (worker *Worker) run(settings *WorkerSettings) {
	if err := worker.loadPrimaryData(settings); err != nil {
		_ = worker.CancelWorker(settings.Market.ID, settings.Resolution.ID)
		log.WithError(err).Error("load primary failed")
		return
	}
	ticker := time.NewTicker(worker.HeartbeatInterval)
	last, err := repository.Candles.ReturnLast(settings.Market.ID, settings.Resolution.ID)
	if err != nil {
		_ = worker.CancelWorker(settings.Market.ID, settings.Resolution.ID)
		log.WithError(err).Error("load last failed")
		return
	}
	lastTime := last.Time
LOOP:
	for {
		select {
		case <-settings.Context.Done():
			break LOOP
		case <-ticker.C:
			to := time.Now()
			if to.Sub(lastTime) <= time.Second {
				continue
			}
			if err := worker.getCandle(settings, lastTime.Unix(), to.Unix()); err != nil {
				time.Sleep(time.Minute)
				log.WithError(err).Error("get candle failed")
			} else {
				lastTime = to
			}
		}
	}
}

func (worker *Worker) loadPrimaryData(ws *WorkerSettings) error {
	last, err := repository.Candles.ReturnLast(ws.Market.ID, ws.Resolution.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			last = new(repository.Candle)
			last.Time = time.Unix(ws.Market.StartTime, 0)
		} else {
			log.WithError(err).Error("load last candle failed")
			return err
		}
	}
	from := last.Time
	end := false
	for !end {
		to := from.Add(time.Duration(1000*ws.Resolution.Duration) * time.Second)
		if to.After(time.Now()) {
			to = time.Now()
			end = true
		}
		if err := worker.getCandle(ws, from.Unix(), to.Unix()); err != nil {
			fmt.Println("from:", from, "to:", to, "duration:", ws.Resolution.Duration)
			log.WithError(err).Error("save candle failed")
			return err
		}
		from = to
	}
	return nil
}

func (worker *Worker) getCandle(ws *WorkerSettings, from, to int64) error {
	c, err := worker.CandleService.OHLC(ws.Context, &brokerageApi.OhlcRequest{
		Resolution: ws.Resolution,
		Market:     ws.Market,
		From:       from,
		To:         to,
	})
	if err != nil {
		return err
	}
	for _, candle := range c.Candles {
		tmp := new(repository.Candle)
		mapper.Struct(candle, tmp)
		tmp.MarketID = ws.Market.ID
		tmp.ResolutionID = ws.Resolution.ID
		err := repository.Candles.Save(tmp)
		if err != nil {
			log.WithError(err).Error("save candle failed")
		}
		indicatorsResponse := make([]buffer.IndicatorResp, len(ws.Indicators))
		for i, indicator := range ws.Indicators {
			indicatorsResponse[i] = buffer.IndicatorResp{
				ID:    indicator.GetID(),
				Value: indicator.Update(),
			}
		}
		buffer.Markets.Update(ws.Market.Name, tmp, indicatorsResponse...)
	}
	return nil
}
