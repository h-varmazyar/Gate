package candleTicker

import (
	"github.com/h-varmazyar/Gate/services/gather/configs"
	"github.com/h-varmazyar/Gate/services/gather/internal/brokers/producer"
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"sync"
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
	cancelFunc context.CancelFunc

	started        bool
	logger         *log.Logger
	cfg            configs.WorkerTicker
	coinexAdapter  coinexAdapter
	candleProducer *producer.Producer
	lock           sync.Mutex
	ctx            context.Context
	marketIDMap    map[string]uint
	marketsRepo    marketsRepo
}

func NewWorker(
	logger *log.Logger,
	configs configs.WorkerTicker,
	coinexAdapter coinexAdapter,
	candlesProducer *producer.Producer,
	marketsRepo marketsRepo,
) *Worker {
	return &Worker{
		logger:         logger,
		cfg:            configs,
		coinexAdapter:  coinexAdapter,
		candleProducer: candlesProducer,
		lock:           sync.Mutex{},
		ctx:            context.Background(),
		marketIDMap:    make(map[string]uint),
		marketsRepo:    marketsRepo,
	}
}

func (w *Worker) Run() error {
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

func (w *Worker) AttachMarket(_ context.Context, market models.Market) error {
	w.lock.Lock()
	w.marketIDMap[market.Name] = market.ID
	w.lock.Unlock()
	return nil
}

func (w *Worker) DetachMarket(_ context.Context, market models.Market) error {
	w.lock.Lock()
	delete(w.marketIDMap, market.Name)
	w.lock.Unlock()
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
				marketID, ok := w.marketIDMap[t.MarketName]
				if ok {
					payload := producer.TickerPayload{
						MarketID:  marketID,
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
}
