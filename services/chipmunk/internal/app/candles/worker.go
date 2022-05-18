package candles

//import (
//	"context"
//	"errors"
//	"github.com/google/uuid"
//	"github.com/h-varmazyar/Gate/pkg/grpcext"
//	"github.com/h-varmazyar/Gate/pkg/mapper"
//	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
//	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
//	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
//	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
//	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/indicators"
//	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
//	log "github.com/sirupsen/logrus"
//	"gorm.io/gorm"
//	"time"
//)
//
//type worker struct {
//	functionsService brokerageApi.FunctionsServiceClient
//	Cancellations    map[uuid.UUID]context.CancelFunc
//}
//
//type WorkerSettings struct {
//	ctx        context.Context
//	Market     *chipmunkApi.Market
//	Resolution *chipmunkApi.Resolution
//	Indicators map[uuid.UUID]indicators.Indicator
//}
//
//var (
//	Worker *worker
//)
//
//func InitializeWorker() {
//	Worker = new(worker)
//	brokerageConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)
//	Worker.functionsService = brokerageApi.NewFunctionsServiceClient(brokerageConn)
//	Worker.Cancellations = make(map[uuid.UUID]context.CancelFunc)
//}
//
//func (worker *worker) AddMarket(settings *WorkerSettings) {
//	marketID, _ := uuid.Parse(settings.Market.ID)
//	cancel, ok := worker.Cancellations[marketID]
//	if ok {
//		cancel()
//	}
//	settings.ctx, worker.Cancellations[marketID] = context.WithCancel(context.Background())
//
//	buffer.Markets.AddList(marketID)
//	go worker.run(settings)
//}
//
//func (worker *worker) DeleteMarket(marketID uuid.UUID) error {
//	fn, ok := worker.Cancellations[marketID]
//	if !ok {
//		return errors.New("worker stopped before")
//	}
//	fn()
//	delete(worker.Cancellations, marketID)
//	buffer.Markets.RemoveList(marketID)
//	return nil
//}
//
//func (worker *worker) run(settings *WorkerSettings) {
//	marketID, _ := uuid.Parse(settings.Market.ID)
//	if err := worker.loadPrimaryData(settings); err != nil {
//		_ = worker.DeleteMarket(marketID)
//		log.WithError(err).Error("load primary failed")
//		return
//	}
//	ticker := time.NewTicker(configs.Variables.OHLCWorkerHeartbeat)
//	last, err := repository.Candles.ReturnLast(settings.Market.ID, settings.Resolution.ID)
//	if err != nil {
//		_ = worker.DeleteMarket(marketID)
//		log.WithError(err).Error("load last failed")
//		return
//	}
//	lastTime := last.Time
//LOOP:
//	for {
//		select {
//		case <-settings.ctx.Done():
//			ticker.Stop()
//			break LOOP
//		case <-ticker.C:
//			to := time.Now()
//			if to.Sub(lastTime) <= time.Second {
//				continue
//			}
//			candles := make([]*repository.Candle, 0)
//			if candles, err = worker.downloadCandlesInfo(settings, lastTime.Unix(), to.Unix()); err != nil {
//				time.Sleep(time.Minute)
//				log.WithError(err).Error("get candles failed")
//			} else {
//				worker.calculateIndicators(settings, candles)
//				for _, candle := range candles {
//					buffer.Markets.Push(marketID, candle)
//				}
//				lastTime = to
//			}
//		}
//	}
//}
//
//func (worker *worker) loadPrimaryData(ws *WorkerSettings) error {
//	totalCandles := make([]*repository.Candle, 0)
//	end := false
//	limit := 10000
//	var from time.Time
//
//	marketID, _ := uuid.Parse(ws.Market.ID)
//	resolutionID, _ := uuid.Parse(ws.Resolution.ID)
//
//	for i := 0; ; i += limit {
//		list, err := repository.Candles.ReturnList(marketID, resolutionID, limit, i)
//		if err != nil {
//			if err == gorm.ErrRecordNotFound {
//			} else {
//				log.WithError(err).Error("load primary candles failed")
//				return err
//			}
//		}
//		totalCandles = append(totalCandles, list...)
//		if len(list) < limit {
//			break
//		}
//	}
//
//	if count := len(totalCandles); count == 0 {
//		from = time.Unix(ws.Market.StartTime, 0)
//	} else {
//		from = totalCandles[count-1].Time
//	}
//
//	for !end {
//		to := from.Add(time.Duration(1000*ws.Resolution.Duration) * time.Second)
//		if to.After(time.Now()) {
//			to = time.Now()
//			end = true
//		}
//		if candles, err := worker.downloadCandlesInfo(ws, from.Unix(), to.Unix()); err != nil {
//			log.WithError(err).Error("get candles failed")
//			return err
//		} else {
//			from = to
//			totalCandles = append(totalCandles, candles...)
//		}
//	}
//	for _, candle := range totalCandles {
//		candle.IndicatorValues = repository.NewIndicatorValues()
//	}
//	for _, indicator := range ws.Indicators {
//		err := indicator.Calculate(totalCandles)
//		if err != nil {
//			return err
//		}
//	}
//	for i := len(totalCandles) - configs.Variables.CandleBufferLength; i < len(totalCandles); i++ {
//		buffer.Markets.Push(marketID, totalCandles[i])
//	}
//	return nil
//}
//
//func (worker *worker) calculateIndicators(ws *WorkerSettings, candles []*repository.Candle) {
//	marketID, _ := uuid.Parse(ws.Market.ID)
//	for _, candle := range candles {
//		candle.IndicatorValues = repository.NewIndicatorValues()
//		for id, indicator := range ws.Indicators {
//			switch indicator.GetType() {
//			case chipmunkApi.Indicator_RSI:
//				data := buffer.Markets.GetLastNCandles(marketID, 2)
//				candle.RSIs[id] = indicator.Update(data).RSI
//			case chipmunkApi.Indicator_Stochastic:
//				data := buffer.Markets.GetLastNCandles(marketID, indicator.GetLength())
//				candle.Stochastics[id] = indicator.Update(data).Stochastic
//			case chipmunkApi.Indicator_BollingerBands:
//				data := buffer.Markets.GetLastNCandles(marketID, indicator.GetLength())
//				candle.BollingerBands[id] = indicator.Update(data).BB
//			case chipmunkApi.Indicator_MovingAverage:
//				data := buffer.Markets.GetLastNCandles(marketID, 2)
//				candle.MovingAverages[id] = indicator.Update(data).MA
//			}
//		}
//	}
//}
//
//func (worker *worker) downloadCandlesInfo(ws *WorkerSettings, from, to int64) ([]*repository.Candle, error) {
//	marketID, _ := uuid.Parse(ws.Market.ID)
//	resolutionID, _ := uuid.Parse(ws.Resolution.ID)
//	response := make([]*repository.Candle, 0)
//	candles, err := worker.functionsService.OHLC(ws.ctx, &brokerageApi.OHLCReq{
//		Resolution: ws.Resolution,
//		Market:     ws.Market,
//		From:       from,
//		To:         to,
//	})
//	if err != nil {
//		log.WithError(err).Error("failed to get candles")
//		return nil, err
//	}
//	for _, candle := range candles.Elements {
//		tmp := new(repository.Candle)
//		mapper.Struct(candle, tmp)
//		tmp.MarketID = marketID
//		tmp.ResolutionID = resolutionID
//		err := repository.Candles.Save(tmp)
//		if err != nil {
//			log.WithError(err).Error("save candles failed")
//		}
//		response = append(response, tmp)
//	}
//	return response, nil
//}
