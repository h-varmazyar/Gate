package workers

import (
	"context"
	"github.com/google/uuid"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/strategies"
)

type SignalCheckWorker struct {
	//strategy   strategies.Strategy
	//markets    []*chipmunkApi.Market
	//running    bool
	cancelFunctions map[uuid.UUID]context.CancelFunc
}

var (
	signalCheckWorker *SignalCheckWorker
)

func SignalCheckWorkerInstance() *SignalCheckWorker {
	if signalCheckWorker == nil {
		signalCheckWorker = &SignalCheckWorker{
			cancelFunctions: make(map[uuid.UUID]context.CancelFunc),
		}
	}
	return signalCheckWorker
}

func StopSignalChecker(strategyID uuid.UUID) {
	//if signalCheckWorker == nil {
	//	return
	//}
	//signalCheckWorker.running = false
	//time.Sleep(time.Second)
	//signalCheckWorker.cancelFunc()
	//signalCheckWorker.strategy = nil
}

func (w *SignalCheckWorker) Start(strategy strategies.Strategy, markets []*chipmunkApi.Market) {
	//ctx, fn := context.WithCancel(context.Background())
	//w.cancelFunctions[strategy.GetID()] = fn
	//for _, market := range markets {
	//	go strategy.CheckForSignals(ctx, market)
	//}
}
