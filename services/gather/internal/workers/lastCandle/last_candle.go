package lastCandle

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/errors"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/gather/configs"
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	"github.com/h-varmazyar/Gate/services/gather/internal/pkg/buffer"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"time"
)

type coreAdapter interface {
	OHLC(ctx context.Context, param OHLCParam) (*chipmunkApi.Candles, error)
}

type candlesRepo interface {
	ReturnLast(ctx context.Context, marketID, resolutionID uint) (models.Candle, error)
	BulkInsert(ctx context.Context, candles []models.Candle) error
}

type marketsRepo interface {
	All(ctx context.Context) ([]models.Market, error)
}

type resolutionsRepo interface {
	All(ctx context.Context) ([]models.Resolution, error)
}

type OHLCParam struct {
	MarketKey  string
	From       time.Time
	To         time.Time
	Timeout    time.Duration
	Resolution models.Resolution
}

type Worker struct {
	Started bool

	ctx             context.Context
	cancelFunc      context.CancelFunc
	logger          *log.Logger
	cfg             configs.WorkerLastCandle
	coreAdapter     coreAdapter
	candlesRepo     candlesRepo
	marketsRepo     marketsRepo
	resolutionsRepo resolutionsRepo
	pairs           []*pair
}

type pair struct {
	Market     models.Market
	Resolution models.Resolution
}

func NewWorker(
	logger *log.Logger,
	configs configs.WorkerLastCandle,
	coreAdapter coreAdapter,
	candlesRepo candlesRepo,
	marketsRepo marketsRepo,
	resolutionsRepo resolutionsRepo,
) *Worker {
	return &Worker{
		logger:          logger,
		cfg:             configs,
		coreAdapter:     coreAdapter,
		candlesRepo:     candlesRepo,
		marketsRepo:     marketsRepo,
		resolutionsRepo: resolutionsRepo,
		pairs:           []*pair{},
	}
}

func (w *Worker) Run() error {
	if !w.Started {
		markets, err := w.marketsRepo.All(context.Background())
		if err != nil {
			return err
		}

		resolutions, err := w.resolutionsRepo.All(context.Background())
		if err != nil {
			return err
		}

		for _, market := range markets {
			for _, resolution := range resolutions {
				w.pairs = append(w.pairs, &pair{
					Market:     market,
					Resolution: resolution,
				})
			}
		}

		w.logger.Infof("starting last candle worker(%v)", len(w.pairs))
		w.ctx, w.cancelFunc = context.WithCancel(context.Background())

		go w.run()
		w.Started = true
	}
	return nil
}

func (w *Worker) Stop() {
	if w.Started {
		w.cancelFunc()
	}
}

func (w *Worker) run() {
	for _, pair := range w.pairs {
		w.fillEmptyBuffer(pair)
	}

	w.logger.Infof("starting last candle worker loop")
	ticker := time.NewTicker(w.cfg.RunningInterval)
	for {
		select {
		case <-w.ctx.Done():
			w.logger.Infof("last candle stopped")
			ticker.Stop()
			return
		case <-ticker.C:
			w.logger.Infof("tickkkkk")
			for _, pair := range w.pairs {
				if err := w.checkForLastCandle(pair); err != nil {
					w.logger.WithError(err).Errorf("failed to check last candle for pair: %v:%v", pair.Market.ID, pair.Resolution.ID)
				}
			}
		}
	}
}

func (w *Worker) fillEmptyBuffer(p *pair) {
	last := buffer.CandleBuffer.Last(p.Market.ID, p.Resolution.ID)
	if last != nil {
		return
	}

	var err error
	candle, err := w.candlesRepo.ReturnLast(w.ctx, p.Market.ID, p.Resolution.ID)
	if err != nil {
		w.logger.WithError(err).Error("failed to get last candle")
		return
	}
	buffer.CandleBuffer.Push(&candle)
	return
}

func (w *Worker) checkForLastCandle(p *pair) error {
	last := buffer.CandleBuffer.Last(p.Market.ID, p.Resolution.ID)
	if last == nil {
		last = &models.Candle{
			Time: p.Market.IssueDate,
		}
	}

	resp, err := w.coreAdapter.OHLC(context.Background(), OHLCParam{
		MarketKey:  p.Market.Name,
		Resolution: p.Resolution,
		Timeout:    w.cfg.RunningInterval,
		To:         time.Now(),
		From:       last.Time,
	})
	if err != nil {
		return err
	}

	if resp.Elements == nil {
		return errors.New(w.ctx, codes.NotFound)
	}

	candles := make([]models.Candle, len(resp.Elements))
	for i, element := range resp.Elements {
		candle := models.Candle{
			Time:         time.Unix(element.Time, 0),
			Open:         element.Open,
			High:         element.High,
			Low:          element.Low,
			Close:        element.Close,
			Volume:       element.Volume,
			Amount:       element.Amount,
			MarketID:     p.Market.ID,
			ResolutionID: p.Resolution.ID,
		}

		buffer.CandleBuffer.Push(&candle)
		candles[i] = candle
	}
	if err = w.candlesRepo.BulkInsert(w.ctx, candles); err != nil {
		w.logger.WithError(err).Error("failed to insert candles")
		return err
	}

	//_, err := w.functionsService.AsyncOHLC(context.Background(), &coreApi.AsyncOHLCReq{
	//	Items:    items,
	//	Platform: platformPair.Platform,
	//})
	//if err != nil {
	//	w.logger.WithError(err).Errorf("failed to create last candle request for %v", platformPair.Platform)
	//}
	return nil
}
