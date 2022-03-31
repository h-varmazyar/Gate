package ohlc

import (
	"context"
	"errors"
	"fmt"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/indicators"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type worker struct {
	CandleService brokerageApi.CandleServiceClient
	Cancellations map[string]context.CancelFunc
}

type WorkerSettings struct {
	Context    context.Context
	Market     *brokerageApi.Market
	Resolution *brokerageApi.Resolution
	Indicators []indicators.Indicator
}

var (
	Worker *worker
)

func init() {
	Worker = new(worker)
	candleConnection := grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)
	Worker.CandleService = brokerageApi.NewCandleServiceClient(candleConnection)
	Worker.Cancellations = make(map[string]context.CancelFunc)
}

func (worker *worker) AddMarket(settings *WorkerSettings) {
	cancel, ok := worker.Cancellations[fmt.Sprintf("%d > %v", settings.Market.ID, settings.Resolution.ID)]
	if ok {
		cancel()
	}
	settings.Context, worker.Cancellations[fmt.Sprintf("%d > %d", settings.Market.ID, settings.Resolution.ID)] = context.WithCancel(context.Background())
	buffer.Markets.AddList(settings.Market.ID)
	go worker.run(settings)
}

func (worker *worker) CancelWorker(marketID, resolutionID uint32) error {
	fn, ok := worker.Cancellations[fmt.Sprintf("%d > %d", marketID, resolutionID)]
	if !ok {
		return errors.New("worker stopped before")
	}
	fn()
	delete(worker.Cancellations, fmt.Sprintf("%d > %d", marketID, resolutionID))
	buffer.Markets.RemoveList(marketID)
	return nil
}

func (worker *worker) run(settings *WorkerSettings) {
	if err := worker.loadPrimaryData(settings); err != nil {
		_ = worker.CancelWorker(settings.Market.ID, settings.Resolution.ID)
		log.WithError(err).Error("load primary failed")
		return
	}
	ticker := time.NewTicker(configs.Variables.OHLCWorkerHeartbeat)
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
			ticker.Stop()
			break LOOP
		case <-ticker.C:
			to := time.Now()
			if to.Sub(lastTime) <= time.Second {
				continue
			}
			candles := make([]*repository.Candle, 0)
			if candles, err = worker.downloadCandlesInfo(settings, lastTime.Unix(), to.Unix()); err != nil {
				time.Sleep(time.Minute)
				log.WithError(err).Error("get candle failed")
			} else {
				worker.calculateIndicators(settings, candles)
				for _, candle := range candles {
					buffer.Markets.Push(settings.Market.ID, candle)
				}
				lastTime = to
			}
		}
	}
}

func (worker *worker) loadPrimaryData(ws *WorkerSettings) error {
	totalCandles := make([]*repository.Candle, 0)
	end := false
	limit := 10000
	var from time.Time

	for i := 0; ; i += limit {
		list, err := repository.Candles.ReturnList(ws.Market.ID, ws.Resolution.ID, limit, i)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
			} else {
				log.WithError(err).Error("load primary candles failed")
				return err
			}
		}
		totalCandles = append(totalCandles, list...)
		if len(list) < limit {
			break
		}
	}

	if count := len(totalCandles); count == 0 {
		from = time.Unix(ws.Market.StartTime, 0)
	} else {
		from = totalCandles[count-1].Time
	}

	for !end {
		to := from.Add(time.Duration(1000*ws.Resolution.Duration) * time.Second)
		if to.After(time.Now()) {
			to = time.Now()
			end = true
		}
		if candles, err := worker.downloadCandlesInfo(ws, from.Unix(), to.Unix()); err != nil {
			log.WithError(err).Error("get candle failed")
			return err
		} else {
			from = to
			totalCandles = append(totalCandles, candles...)
		}
	}
	for _, candle := range totalCandles {
		candle.IndicatorValues = repository.NewIndicatorValues()
	}
	for _, indicator := range ws.Indicators {
		err := indicator.Calculate(totalCandles)
		if err != nil {
			return err
		}
	}
	for i := len(totalCandles) - configs.Variables.CandleBufferLength; i < len(totalCandles); i++ {
		buffer.Markets.Push(ws.Market.ID, totalCandles[i])
	}
	return nil
}

func (worker *worker) calculateIndicators(ws *WorkerSettings, candles []*repository.Candle) {
	for _, candle := range candles {
		candle.IndicatorValues = repository.NewIndicatorValues()
		for _, indicator := range ws.Indicators {
			switch indicator.GetType() {
			case indicators.RSI:
				data := buffer.Markets.GetLastNCandles(ws.Market.ID, 2)
				candle.RSIs[indicator.GetID()] = indicator.Update(data).RSI
			case indicators.Stochastic:
				data := buffer.Markets.GetLastNCandles(ws.Market.ID, indicator.GetLength())
				candle.Stochastics[indicator.GetID()] = indicator.Update(data).Stochastic
			case indicators.BollingerBands:
				data := buffer.Markets.GetLastNCandles(ws.Market.ID, indicator.GetLength())
				candle.BollingerBands[indicator.GetID()] = indicator.Update(data).BB
			case indicators.MovingAverage:
				data := buffer.Markets.GetLastNCandles(ws.Market.ID, 2)
				candle.MovingAverages[indicator.GetID()] = indicator.Update(data).MA
			}
		}
	}
}

func (worker *worker) downloadCandlesInfo(ws *WorkerSettings, from, to int64) ([]*repository.Candle, error) {
	response := make([]*repository.Candle, 0)
	c, err := worker.CandleService.OHLC(ws.Context, &brokerageApi.OhlcRequest{
		Resolution: ws.Resolution,
		Market:     ws.Market,
		From:       from,
		To:         to,
	})
	if err != nil {
		log.WithError(err).Error("failed to get candles")
		return nil, err
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
		response = append(response, tmp)
	}
	return response, nil
}
