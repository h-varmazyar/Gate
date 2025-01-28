package candleTicker

import (
	"github.com/h-varmazyar/Gate/services/gather/configs"
	candlesProducer "github.com/h-varmazyar/Gate/services/gather/internal/brokers/producer/candles"
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

type coinexAdapter interface {
	MarketsTicker(ctx context.Context) ([]Ticker, error)
}

type marketsRepo interface {
	All(ctx context.Context) ([]models.Market, error)
}

type Ticker struct {
	MarketName string
	LastPrice  float64
}

type Worker struct {
	started    bool
	ctx        context.Context
	cancelFunc context.CancelFunc

	logger         *log.Logger
	cfg            configs.WorkerTicker
	coinexAdapter  coinexAdapter
	candleProducer *candlesProducer.Producer
	marketIDMap    map[string]uint
	marketsRepo    marketsRepo
}

func NewWorker(
	logger *log.Logger,
	configs configs.WorkerTicker,
	coinexAdapter coinexAdapter,
	candlesProducer *candlesProducer.Producer,
	marketsRepo marketsRepo,
) *Worker {
	return &Worker{
		logger:         logger,
		cfg:            configs,
		coinexAdapter:  coinexAdapter,
		candleProducer: candlesProducer,
		ctx:            context.Background(),
		marketIDMap:    make(map[string]uint),
		marketsRepo:    marketsRepo,
	}
}

func (w *Worker) Start() error {
	if len(w.marketIDMap) == 0 {
		markets, err := w.marketsRepo.All(context.Background())
		if err != nil {
			return err
		}

		for _, market := range markets {
			w.marketIDMap[market.Name] = market.ID
		}
	}
	if !w.started {
		w.logger.Println("starting candle ticker")
		go w.run()
		w.started = true
	}
	return nil
}

func (w *Worker) run() {
	ticker := time.NewTicker(w.cfg.RunningInterval)
	for {
		select {
		case <-w.ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			tickers, err := w.coinexAdapter.MarketsTicker(w.ctx)
			if err != nil {
				w.logger.WithError(err).Error("Ticker error")
				continue
			}
			//w.logger.Infof("ticker running %v", tickers)
			for _, t := range tickers {
				payload := candlesProducer.TickerPayload{
					MarketID:  w.marketIDMap[t.MarketName],
					LastPrice: t.LastPrice,
				}
				if err = w.candleProducer.PublishTicker(payload); err != nil {
					w.logger.WithError(err).Error("failed to produce ticker")
					continue
				}
			}
		}
	}
}
