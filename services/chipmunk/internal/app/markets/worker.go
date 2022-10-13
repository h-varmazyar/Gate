package markets

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/indicators"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type Worker struct {
	functionsService brokerageApi.FunctionsServiceClient
	Cancellations    map[uuid.UUID]context.CancelFunc
}

type WorkerSettings struct {
	ctx        context.Context
	Market     *repository.Market
	Resolution *repository.Resolution
	Indicators map[uuid.UUID]indicators.Indicator
}

var (
	worker *Worker
)

func InitializeWorker() {
	worker = new(Worker)
	brokerageConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)
	worker.functionsService = brokerageApi.NewFunctionsServiceClient(brokerageConn)
	worker.Cancellations = make(map[uuid.UUID]context.CancelFunc)
}

func (worker *Worker) AddMarket(settings *WorkerSettings) {
	cancel, ok := worker.Cancellations[settings.Market.ID]
	if ok {
		cancel()
	}
	settings.ctx, worker.Cancellations[settings.Market.ID] = context.WithCancel(context.Background())

	buffer.Markets.AddList(settings.Market.ID)
	log.Infof("add new market: %v", settings.Market.Name)
	go worker.run(settings)
}

func (worker *Worker) DeleteMarket(marketID uuid.UUID) error {
	log.Infof("deleting market: %v", marketID)
	fn, ok := worker.Cancellations[marketID]
	if !ok {
		return nil
	}
	fn()
	delete(worker.Cancellations, marketID)
	buffer.Markets.RemoveList(marketID)
	log.Infof("market deleted: %v", marketID)
	return nil
}

func (worker *Worker) run(settings *WorkerSettings) {
	log.Infof("runnig worker for %v", settings.Market.Name)
	if err := worker.loadPrimaryData(settings); err != nil {
		_ = worker.DeleteMarket(settings.Market.ID)
		log.WithError(err).Error("load primary failed")
		return
	}
	ticker := time.NewTicker(configs.Variables.OHLCWorkerHeartbeat)
	last, err := repository.Candles.ReturnLast(settings.Market.ID, settings.Resolution.ID)
	if err != nil {
		_ = worker.DeleteMarket(settings.Market.ID)
		log.WithError(err).Error("load last failed")
		return
	}
	lastTime := last.Time
LOOP:
	for {
		select {
		case <-settings.ctx.Done():
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
				log.WithError(err).Error("get candles failed")
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

func (worker *Worker) loadPrimaryData(ws *WorkerSettings) error {
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
		from = ws.Market.StartTime
	} else {
		from = totalCandles[count-1].Time
	}

	for !end {
		to := from.Add(1000 * ws.Resolution.Duration * time.Second)
		if to.After(time.Now()) {
			to = time.Now()
			end = true
		}
		if candles, err := worker.downloadCandlesInfo(ws, from.Unix(), to.Unix()); err != nil {
			log.WithError(err).Error("get candles failed")
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
	i := len(totalCandles) - configs.Variables.CandleBufferLength
	if i < 0 {
		i = 0
	}
	for ; i < len(totalCandles); i++ {
		buffer.Markets.Push(ws.Market.ID, totalCandles[i])
	}
	return nil
}

func (worker *Worker) calculateIndicators(ws *WorkerSettings, candles []*repository.Candle) {
	for _, candle := range candles {
		candle.IndicatorValues = repository.NewIndicatorValues()
		for id, indicator := range ws.Indicators {
			switch indicator.GetType() {
			case chipmunkApi.Indicator_RSI:
				data := buffer.Markets.GetLastNCandles(ws.Market.ID, 2)
				candle.RSIs[id] = indicator.Update(data).RSI
			case chipmunkApi.Indicator_Stochastic:
				data := buffer.Markets.GetLastNCandles(ws.Market.ID, indicator.GetLength())
				candle.Stochastics[id] = indicator.Update(data).Stochastic
			case chipmunkApi.Indicator_BollingerBands:
				data := buffer.Markets.GetLastNCandles(ws.Market.ID, indicator.GetLength())
				candle.BollingerBands[id] = indicator.Update(data).BB
			case chipmunkApi.Indicator_MovingAverage:
				data := buffer.Markets.GetLastNCandles(ws.Market.ID, 2)
				candle.MovingAverages[id] = indicator.Update(data).MA
			}
		}
	}
}

func (worker *Worker) downloadCandlesInfo(ws *WorkerSettings, from, to int64) ([]*repository.Candle, error) {
	response := make([]*repository.Candle, 0)
	resolution := new(chipmunkApi.Resolution)
	mapper.Struct(ws.Resolution, resolution)

	market := new(chipmunkApi.Market)
	mapper.Struct(ws.Market, market)
	candles, err := worker.functionsService.OHLC(ws.ctx, &brokerageApi.OHLCReq{
		Resolution: resolution,
		Market:     market,
		From:       from,
		To:         to,
	})
	if err != nil {
		log.WithError(err).Error("failed to get candles")
		return nil, err
	}
	for _, candle := range candles.Elements {
		tmp := new(repository.Candle)
		mapper.Struct(candle, tmp)
		tmp.MarketID = ws.Market.ID
		tmp.ResolutionID = ws.Resolution.ID
		err := repository.Candles.Save(tmp)
		if err != nil {
			log.WithError(err).Error("save candles failed")
		}
		response = append(response, tmp)
	}
	return response, nil
}
