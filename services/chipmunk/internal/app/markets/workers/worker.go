package workers

//
//import (
//	"context"
//	"github.com/google/uuid"
//	"github.com/h-varmazyar/Gate/pkg/grpcext"
//	"github.com/h-varmazyar/Gate/pkg/mapper"
//	"github.com/h-varmazyar/Gate/services/Dolphin/configs"
//	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
//	rateLimiters "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/rateLimiters/service"
//	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
//	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
//	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/indicators"
//	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
//	log "github.com/sirupsen/logrus"
//	"gorm.io/gorm"
//	"time"
//)
//
//type PrimaryDataWorker struct {
//	functionsService coreApi.FunctionsServiceClient
//	Cancellations    map[uuid.UUID]context.CancelFunc
//	configs          *Configs
//	candlesService   *rateLimiters.Service
//}
//
//type WorkerSettings struct {
//	ctx context.Context
//	//Market      *entity.Market
//	Resolution  *entity.Resolution
//	Indicators  map[uuid.UUID]indicators.Indicator
//	BrokerageID uuid.UUID
//}
//
//func InitializeWorker(_ context.Context, configs *Configs, candlesService *rateLimiters.Service) *PrimaryDataWorker {
//	worker := new(PrimaryDataWorker)
//	brokerageConn := grpcext.NewConnection(configs.CoreAddress)
//	worker.functionsService = coreApi.NewFunctionsServiceClient(brokerageConn)
//	worker.Cancellations = make(map[uuid.UUID]context.CancelFunc)
//	worker.configs = configs
//	worker.candlesService = candlesService
//
//	return worker
//}
//
//func (worker *PrimaryDataWorker) AddMarket(settings *WorkerSettings) {
//	cancel, ok := worker.Cancellations[settings.Market.ID]
//	if ok {
//		cancel()
//	}
//	settings.ctx, worker.Cancellations[settings.Market.ID] = context.WithCancel(context.Background())
//
//	buffer.Markets.AddList(settings.Market.ID)
//	go worker.run(settings)
//}
//
//func (worker *PrimaryDataWorker) DeleteMarket(marketID uuid.UUID) error {
//	log.Infof("deleting market: %v", marketID)
//	fn, ok := worker.Cancellations[marketID]
//	if !ok {
//		return nil
//	}
//	fn()
//	delete(worker.Cancellations, marketID)
//	buffer.Markets.RemoveList(marketID)
//	log.Infof("market deleted: %v", marketID)
//	return nil
//}
//
//func (worker *PrimaryDataWorker) run(settings *WorkerSettings) {
//	time.Sleep(time.Second)
//	log.Infof("runnig workers for %v", settings.Market.Name)
//	if err := worker.loadPrimaryData(settings); err != nil {
//		_ = worker.DeleteMarket(settings.Market.ID)
//		log.WithError(err).Error("load primary failed")
//		return
//	}
//	ticker := time.NewTicker(worker.configs.OHLCWorkerHeartbeat)
//	last, err := worker.candlesService.ReturnLast(settings.Market.ID, settings.Resolution.ID)
//	if err != nil {
//		_ = worker.DeleteMarket(settings.Market.ID)
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
//			rateLimiters := make([]*entity.Candle, 0)
//			if rateLimiters, err = worker.downloadCandlesInfo(settings, lastTime.Unix(), to.Unix()); err != nil {
//				log.WithError(err).Error("get rateLimiters failed")
//			} else {
//				worker.calculateIndicators(settings, rateLimiters)
//				for _, candle := range rateLimiters {
//					buffer.Markets.Push(settings.Market.ID, candle)
//				}
//				lastTime = to
//			}
//		}
//	}
//}
//
//func (worker *PrimaryDataWorker) loadPrimaryData(ws *WorkerSettings) error {
//	totalCandles := make([]*entity.Candle, 0)
//	end := false
//	limit := 10000
//	var from time.Time
//
//	for i := 0; ; i += limit {
//		list, err := repository.Candles.ReturnList(ws.Market.ID, ws.Resolution.ID, limit, i)
//		if err != nil && err != gorm.ErrRecordNotFound {
//			log.WithError(err).Error("load primary rateLimiters failed")
//			return err
//		}
//		totalCandles = append(totalCandles, list...)
//		if len(list) < limit {
//			break
//		}
//	}
//
//	if len(totalCandles) == 0 {
//		from = ws.Market.StartTime
//	} else {
//		from = totalCandles[len(totalCandles)-1].Time
//	}
//
//	for !end {
//		to := from.Add(1000 * ws.Resolution.Duration * time.Second)
//		if to.After(time.Now()) {
//			to = time.Now()
//			end = true
//		}
//		if rateLimiters, err := worker.downloadCandlesInfo(ws, from.Unix(), to.Unix()); err != nil {
//			log.WithError(err).Error("get rateLimiters failed")
//			return err
//		} else {
//			from = to
//			totalCandles = append(totalCandles, rateLimiters...)
//		}
//	}
//	for _, candle := range totalCandles {
//		candle.IndicatorValues = entity.NewIndicatorValues()
//	}
//	for _, indicator := range ws.Indicators {
//		err := indicator.Calculate(totalCandles)
//		if err != nil {
//			return err
//		}
//	}
//	i := len(totalCandles) - configs.Variables.CandleBufferLength
//	if i < 0 {
//		i = 0
//	}
//	for ; i < len(totalCandles); i++ {
//		buffer.Markets.Push(ws.Market.ID, totalCandles[i])
//	}
//	return nil
//}
//
//func (worker *PrimaryDataWorker) calculateIndicators(ws *WorkerSettings, rateLimiters []*entity.Candle) {
//	for _, candle := range rateLimiters {
//		candle.IndicatorValues = entity.NewIndicatorValues()
//		for id, indicator := range ws.Indicators {
//			switch indicator.GetType() {
//			case chipmunkApi.Indicator_RSI:
//				data := buffer.Markets.GetLastNCandles(ws.Market.ID, 2)
//				candle.RSIs[id] = indicator.Update(data).RSI
//			case chipmunkApi.Indicator_Stochastic:
//				data := buffer.Markets.GetLastNCandles(ws.Market.ID, indicator.GetLength())
//				candle.Stochastics[id] = indicator.Update(data).Stochastic
//			case chipmunkApi.Indicator_BollingerBands:
//				data := buffer.Markets.GetLastNCandles(ws.Market.ID, indicator.GetLength())
//				candle.BollingerBands[id] = indicator.Update(data).BB
//			case chipmunkApi.Indicator_MovingAverage:
//				data := buffer.Markets.GetLastNCandles(ws.Market.ID, 2)
//				candle.MovingAverages[id] = indicator.Update(data).MA
//			}
//		}
//	}
//}
//
//func (worker *PrimaryDataWorker) downloadCandlesInfo(ws *WorkerSettings, from, to int64) ([]*entity.Candle, error) {
//	resolution := new(chipmunkApi.Resolution)
//	mapper.Struct(ws.Resolution, resolution)
//
//	market := new(chipmunkApi.Market)
//	mapper.Struct(ws.Market, market)
//	rateLimiters, err := worker.functionsService.OHLC(ws.ctx, &coreApi.OHLCReq{
//		Resolution:  resolution,
//		Market:      market,
//		From:        from,
//		To:          to,
//		BrokerageID: ws.BrokerageID.String(),
//	})
//	if err != nil {
//		log.WithError(err).Error("failed to get rateLimiters")
//		return nil, err
//	}
//	localCandles := make([]*entity.Candle, 0)
//	for _, candle := range rateLimiters.Elements {
//		tmp := new(entity.Candle)
//		mapper.Struct(candle, tmp)
//		tmp.MarketID = ws.Market.ID
//		tmp.ResolutionID = ws.Resolution.ID
//		localCandles = append(localCandles, tmp)
//	}
//	if err = repository.Candles.BulkInsert(localCandles); err != nil {
//		return nil, err
//	}
//	return localCandles, nil
//}
