package workers

import (
	"context"
	"github.com/google/uuid"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/strategies"
)

type SignalCheckWorker struct {
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

func (w *SignalCheckWorker) Start(strategy strategies.Strategy, markets []*chipmunkApi.Market, brokerageID uuid.UUID) {
	ctx, fn := context.WithCancel(context.Background())
	w.cancelFunctions[brokerageID] = fn
	for _, market := range markets {
		go strategy.CheckForSignals(ctx, market)
	}
}

func (w *SignalCheckWorker) Stop(brokerageID uuid.UUID) {
	fn, ok := w.cancelFunctions[brokerageID]
	if !ok {
		return
	}
	fn()
	delete(w.cancelFunctions, brokerageID)
}
