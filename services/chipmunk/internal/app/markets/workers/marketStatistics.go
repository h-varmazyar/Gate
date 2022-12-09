package workers

import (
	"context"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	candles "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets/repository"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type StatisticsWorker struct {
	configs          *Configs
	functionsService coreApi.FunctionsServiceClient
	candlesService   *candles.Service
	cancelFunctions  map[api.Platform]context.CancelFunc
	lock             *sync.RWMutex
	db               repository.MarketRepository
}

type Runner struct {
	ctx      context.Context
	platform api.Platform
}

func NewStatisticsWorker(_ context.Context, configs *Configs, candlesService *candles.Service, db repository.MarketRepository) *StatisticsWorker {
	coreConn := grpcext.NewConnection(configs.CoreAddress)
	return &StatisticsWorker{
		configs:          configs,
		cancelFunctions:  make(map[api.Platform]context.CancelFunc),
		functionsService: coreApi.NewFunctionsServiceClient(coreConn),
		candlesService:   candlesService,
		lock:             new(sync.RWMutex),
		db:               db,
	}
}

func (w *StatisticsWorker) Start(platform api.Platform) {
	w.Stop(platform)
	cancelContext, cancelFunc := context.WithCancel(context.Background())
	runner := &Runner{
		ctx:      cancelContext,
		platform: platform,
	}
	w.lock.Lock()
	w.cancelFunctions[platform] = cancelFunc
	w.lock.Unlock()
	go w.run(runner)
}

func (w *StatisticsWorker) Stop(platform api.Platform) {
	w.lock.Lock()
	defer w.lock.Unlock()
	if fn, ok := w.cancelFunctions[platform]; ok {
		fn()
		delete(w.cancelFunctions, platform)
	}
}

func (w *StatisticsWorker) run(runner *Runner) {
	log.Infof("running market statistics worker for: %v", runner.platform)
	ticker := time.NewTicker(w.configs.MarketStatisticsWorkerInterval)

	for {
		select {
		case <-runner.ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			statistics, err := w.functionsService.AllMarketStatistics(runner.ctx, &coreApi.AllMarketStatisticsReq{
				Platform: runner.platform,
			})
			if err != nil {
				log.WithError(err).Info("market statistics failed")
				continue
			}

			tickers := make(map[string]*chipmunkApi.CandleBulkUpdateReqTicker)
			for marketName, data := range statistics.AllStatistics {
				market, err := w.db.Info(runner.platform, marketName)
				if err != nil {
					log.Infof("failed to fetch market %v", marketName)
					continue
				}
				tickers[market.ID.String()] = &chipmunkApi.CandleBulkUpdateReqTicker{
					Volume: data.Volume,
					Close:  data.Close,
					Open:   data.Open,
					High:   data.High,
					Low:    data.Low,
				}
			}

			bulkUpdateReq := &chipmunkApi.CandleBulkUpdateReq{
				Platform: runner.platform,
				Date:     statistics.Date,
				Tickers:  tickers,
			}
			_, err = w.candlesService.BulkUpdate(runner.ctx, bulkUpdateReq)
		}
	}
}
