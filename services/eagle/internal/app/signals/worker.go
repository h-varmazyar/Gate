package signals

import (
	"context"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/strategies"
	"time"
)

type worker struct {
	strategy   strategies.Strategy
	markets    []*chipmunkApi.Market
	running    bool
	cancelFunc func()
}

var (
	signalCheckWorker *worker
)

func SignalCheckWorkerInstance(strategy strategies.Strategy, markets []*chipmunkApi.Market) *worker {
	if signalCheckWorker == nil || !signalCheckWorker.running {
		signalCheckWorker = &worker{
			strategy: strategy,
			markets:  markets,
			running:  false,
		}
	}
	return signalCheckWorker
}

func StopSignalChecker() {
	if signalCheckWorker == nil {
		return
	}
	signalCheckWorker.running = false
	time.Sleep(time.Second)
	signalCheckWorker.cancelFunc()
	signalCheckWorker.strategy = nil
}

func (w *worker) Start() {
	w.running = true
	ctx, fn := context.WithCancel(context.Background())
	w.cancelFunc = fn
	for _, market := range w.markets {
		go w.strategy.CheckForSignals(ctx, market)
	}
}

func (w *worker) IsRunning() bool {
	if !w.running || w.cancelFunc == nil {
		return false
	}
	return true
}
