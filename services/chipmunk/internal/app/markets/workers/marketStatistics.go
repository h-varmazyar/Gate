package workers

import (
	"context"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	candles "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/service"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	"sync"
	"time"
)

type StatisticsWorker struct {
	configs          *Configs
	functionsService coreApi.FunctionsServiceClient
	candlesService   *candles.Service
	cancelFunctions  map[api.Platform]context.CancelFunc
	lock             *sync.RWMutex
}

type Runner struct {
	ctx      context.Context
	platform api.Platform
}

func NewStatisticsWorker(_ context.Context, configs *Configs, candlesService *candles.Service) *StatisticsWorker {
	coreConn := grpcext.NewConnection(configs.CoreAddress)
	return &StatisticsWorker{
		configs:          configs,
		cancelFunctions:  make(map[api.Platform]context.CancelFunc),
		functionsService: coreApi.NewFunctionsServiceClient(coreConn),
		candlesService:   candlesService,
		lock:             new(sync.RWMutex),
	}
}

func (w *StatisticsWorker) Start(ctx context.Context, platform api.Platform) {
	w.Stop(platform)
	cancelContext, cancelFunc := context.WithCancel(ctx)
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
	ticker := time.NewTicker(w.configs.MarketStatisticsWorkerInterval)
LOOP:
	for {
		select {
		case <-runner.ctx.Done():
			ticker.Stop()
			break LOOP
		case <-ticker.C:
			statistics, err := w.functionsService.AllMarketStatistics(runner.ctx, new(coreApi.AllMarketStatisticsReq))
			if err != nil {
				continue
			}

			tickers := make(map[string]*chipmunkApi.CandleBulkUpdateReqTicker)
			for market, data := range statistics.AllStatistics {
				tickers[market] = &chipmunkApi.CandleBulkUpdateReqTicker{
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
