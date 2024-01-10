package workers

import (
	"context"
	"github.com/google/uuid"
	api "github.com/h-varmazyar/Gate/api/proto"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/strategies"
	log "github.com/sirupsen/logrus"
)

type SignalCheckWorker struct {
	cancelFunctions map[uuid.UUID]context.CancelFunc
	log             *log.Logger
}

var (
	signalCheckWorker *SignalCheckWorker
)

func SignalCheckWorkerInstance(log *log.Logger) *SignalCheckWorker {
	if signalCheckWorker == nil {
		signalCheckWorker = &SignalCheckWorker{
			cancelFunctions: make(map[uuid.UUID]context.CancelFunc),
			log:             log,
		}
	}
	return signalCheckWorker
}

func (w *SignalCheckWorker) Start(strategy strategies.Strategy, markets []*chipmunkApi.Market, brokerageID uuid.UUID) {
	ctx, fn := context.WithCancel(context.Background())
	w.cancelFunctions[brokerageID] = fn

	w.log.Infof("starting signal check for markets: %v", len(markets))

	for _, market := range markets {
		if market.Status == api.Status_Enable {
			go strategy.CheckForSignals(ctx, market)
		}
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
