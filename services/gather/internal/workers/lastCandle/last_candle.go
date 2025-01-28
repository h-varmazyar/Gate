package lastCandle

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/errors"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/gather/configs"
	candlesProducer "github.com/h-varmazyar/Gate/services/gather/internal/brokers/producer/candles"
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	"github.com/h-varmazyar/Gate/services/gather/internal/pkg/buffer"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"sync"
	"time"
)

type coreAdapter interface {
	OHLC(ctx context.Context, param OHLCParam) (*chipmunkApi.Candles, error)
}

type coinexAdapter interface {
	OHLC(ctx context.Context, market models.Market, resolution models.Resolution, from, to time.Time) ([]models.Candle, error)
}

type candlesRepo interface {
	ReturnLast(ctx context.Context, marketID, resolutionID uint) (models.Candle, error)
	BulkInsert(ctx context.Context, candles []models.Candle) error
	Update(ctx context.Context, candles models.Candle) error
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
	coinexAdapter   coinexAdapter
	candlesRepo     candlesRepo
	marketsRepo     marketsRepo
	resolutionsRepo resolutionsRepo
	pairs           []*pair
	lock            sync.Mutex
	candleProducer  *candlesProducer.Producer
}

type pair struct {
	Market     models.Market
	Resolution models.Resolution
}

func NewWorker(
	logger *log.Logger,
	configs configs.WorkerLastCandle,
	coreAdapter coreAdapter,
	coinexAdapter coinexAdapter,
	candlesRepo candlesRepo,
	marketsRepo marketsRepo,
	resolutionsRepo resolutionsRepo,
	candlesProducer *candlesProducer.Producer,
) *Worker {
	return &Worker{
		logger:          logger,
		cfg:             configs,
		coreAdapter:     coreAdapter,
		coinexAdapter:   coinexAdapter,
		candlesRepo:     candlesRepo,
		marketsRepo:     marketsRepo,
		resolutionsRepo: resolutionsRepo,
		pairs:           []*pair{},
		lock:            sync.Mutex{},
		candleProducer:  candlesProducer,
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

func (w *Worker) AttachMarket(ctx context.Context, market models.Market) error {
	resolutions, err := w.resolutionsRepo.All(ctx)
	if err != nil {
		return err
	}

	w.lock.Lock()
	for _, resolution := range resolutions {
		w.pairs = append(w.pairs, &pair{
			Market:     market,
			Resolution: resolution,
		})
	}
	w.lock.Unlock()
	return nil
}

func (w *Worker) DetachMarket(_ context.Context, market models.Market) error {
	w.lock.Lock()
	for i, _ := range w.pairs {
		if w.pairs[i].Market.ID == market.ID {
			w.pairs = append(w.pairs[:i], w.pairs[i+1:]...)
			i--
		}
	}
	w.lock.Unlock()
	return nil
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
			wg := &sync.WaitGroup{}
			wg.Add(len(w.pairs))
			eachPeriodDuration := time.Duration(int64(w.cfg.RunningInterval) / int64(len(w.pairs)))
			w.logger.Infof("last candle duration: %v - %v", w.cfg.RunningInterval, time.Duration(eachPeriodDuration))
			totalStart := time.Now()
			for _, p := range w.pairs {
				start := time.Now()
				go func(p *pair) {
					len, err := w.checkForLastCandle(p)
					if err != nil {
						w.logger.WithError(err)
					} else {
						w.logger.Infof("len %v: %v", p.Market.ID, len)
					}
					wg.Done()
				}(p)
				diff := time.Now().Sub(start)
				if diff < eachPeriodDuration {
					time.Sleep(eachPeriodDuration - diff)
				}
			}
			wg.Wait()
			w.logger.Infof("done one period: %v - %v", w.cfg.RunningInterval, time.Now().Sub(totalStart))
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

func (w *Worker) checkForLastCandle(p *pair) (int, error) {
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
		return 0, err
	}

	if resp.Elements == nil {
		return 0, errors.New(w.ctx, codes.NotFound)
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
		return 0, err
	}

	payload := candlesProducer.CandlePayload{
		MarketID:     p.Market.ID,
		ResolutionID: p.Resolution.ID,
		Candles: []candlesProducer.Candle{
			{
				Timestamp: candles[len(candles)-1].Time.Unix(),
				Open:      candles[len(candles)-1].Open,
				High:      candles[len(candles)-1].High,
				Low:       candles[len(candles)-1].Low,
				Close:     candles[len(candles)-1].Close,
				Volume:    candles[len(candles)-1].Volume,
			},
		},
	}
	if err = w.candleProducer.PublishCandleUpdates(payload); err != nil {
		w.logger.WithError(err).Error("failed to produce candle")
		return 0, err
	}

	//_, err := w.functionsService.AsyncOHLC(context.Background(), &coreApi.AsyncOHLCReq{
	//	Items:    items,
	//	Platform: platformPair.Platform,
	//})
	//if err != nil {
	//	w.logger.WithError(err).Errorf("failed to create last candle request for %v", platformPair.Platform)
	//}
	return len(candles), nil
}

//func (w *Worker) checkForLastCandle(p *pair) (int, error) {
//	last := buffer.CandleBuffer.Last(p.Market.ID, p.Resolution.ID)
//	if last == nil {
//		last = &models.Candle{
//			Time: p.Market.IssueDate,
//		}
//	}
//
//	candles, err := w.coinexAdapter.OHLC(context.Background(), p.Market, p.Resolution, last.Time, time.Now())
//	if err != nil {
//		return 0, err
//	}
//
//	for _, candle := range candles {
//		buffer.CandleBuffer.Push(&candle)
//	}
//	if candles[0].Time.Equal(last.Time) {
//		if err = w.candlesRepo.Update(context.Background(), candles[0]); err != nil {
//			return 0, err
//		}
//		candles = append(candles[:0], candles[1:]...)
//	}
//
//	if len(candles) > 0 {
//		if err = w.candlesRepo.BulkInsert(w.ctx, candles); err != nil {
//			w.logger.WithError(err).Error("failed to insert candles")
//			return 0, err
//		}
//	}
//	return len(candles), nil
//}
