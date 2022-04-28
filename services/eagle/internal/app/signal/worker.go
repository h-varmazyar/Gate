package signal

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/errors"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/buffers"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/repository"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/strategies"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/strategies/automatedStrategy"
	"google.golang.org/grpc/codes"
)

type worker struct {
	strategy strategies.Strategy
	enable   bool
	stopSig  chan bool
	dataChan chan *eagleApi.CandleData
}

var (
	Worker *worker
)

func init() {
	Worker = new(worker)
}

func (w *worker) Start(ctx context.Context, strategy *repository.Strategy) error {
	if !w.enable {
		automated, err := automatedStrategy.NewAutomatedStrategy(ctx, strategy, 0.02, 0.02)
		if err != nil {
			return err
		}
		w.strategy = automated
		go w.dataListener()
		return nil
	} else {
		return errors.NewWithSlug(ctx, codes.FailedPrecondition, "worker_already_running")
	}
}

func (w *worker) Stop() {
	w.enable = false
	w.stopSig <- true
}

func (w *worker) dataListener() {
LOOP:
	for {
		select {
		case <-w.stopSig:
			break LOOP
		case data := <-w.dataChan:
			go w.handleData(data)
		}
	}
}

func (w *worker) handleData(data *eagleApi.CandleData) {
	marketID, err := uuid.Parse(data.MarketID)
	if err != nil {
		return
	}
	for _, candle := range data.Candles {
		buffers.Markets.Push(marketID, candle)
	}

	signalStrength := w.strategy.CheckForSignals(marketID, data.MarketName)

}
