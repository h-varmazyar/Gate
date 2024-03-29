package workers

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"time"
)

type RedundantRemover struct {
	db           repository.CandleRepository
	configs      *Configs
	ctx          context.Context
	cancelFunc   context.CancelFunc
	logger       *log.Logger
	Started      bool
	removedCount int
}

func NewRedundantRemover(_ context.Context, db repository.CandleRepository, configs *Configs, logger *log.Logger) *RedundantRemover {
	return &RedundantRemover{
		db:      db,
		logger:  logger,
		configs: configs,
	}
}

func (w *RedundantRemover) Start(platformsPairs []*PlatformPairs) {
	if !w.Started {
		w.logger.Infof("starting redundant worker")
		w.ctx, w.cancelFunc = context.WithCancel(context.Background())
		go w.run(platformsPairs)
		w.Started = true
	}
}

func (w *RedundantRemover) Stop() {
	if w.Started {
		w.cancelFunc()
	}
}

func (w *RedundantRemover) run(platformsPairs []*PlatformPairs) {
	ticker := time.NewTicker(w.configs.RedundantRemoverInterval)
	for {
		select {
		case <-w.ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			w.removedCount = 0
			w.logger.Infof("prepare removed platform pairs")
			for _, platformPairs := range platformsPairs {
				if err := w.removeRedundantCandles(platformPairs); err != nil {
					w.logger.WithError(err).Error("failed to prepare remove redundant rateLimiters")
				}
			}
			w.logger.Infof("removed platform pairs: %v", w.removedCount)

		}
	}
}

//func (w *RedundantRemover) prepareMarkets(markets []*chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) error {
//	for _, market := range markets {
//		err := w.prepareResolutions(market, resolutions)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func (w *RedundantRemover) prepareResolutions(market *chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) error {
//	for _, resolution := range resolutions {
//		err := w.removeRedundantCandles(market, resolution)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

func (w *RedundantRemover) removeRedundantCandles(platformPairs *PlatformPairs) error {
	ids := make([]uuid.UUID, 0)
	for _, pair := range platformPairs.Pairs {
		resolutionID, err := uuid.Parse(pair.Resolution.ID)
		if err != nil {
			return err
		}
		marketID, err := uuid.Parse(pair.Market.ID)
		if err != nil {
			return err
		}
		candles, err := w.loadCandles(marketID, resolutionID)
		if err != nil {
			return err
		}
		for i := 1; i < len(candles); i++ {
			if candles[i-1].Time.Equal(candles[i].Time) {
				ids = append(ids, candles[i-1].ID)
				w.removedCount++
			}
		}
	}
	if len(ids) > 0 {
		w.logger.Infof("removing: %v", len(ids))
		for i := 0; i < len(ids); i += 1000 {
			end := i + 1000
			if end > len(ids) {
				end = len(ids)
			}
			if err := w.db.BulkHardDelete(ids[i:end]); err != nil {
				w.logger.WithError(err).Errorf("failed to delete rateLimiters")
			}
		}
	}

	return nil
}

func (w *RedundantRemover) loadCandles(marketID, resolutionID uuid.UUID) ([]*entity.Candle, error) {
	end := false
	limit := 1000000
	candles := make([]*entity.Candle, 0)
	for offset := 0; !end; offset += limit {
		list, err := w.db.List(marketID, resolutionID, limit, offset)
		if err != nil {
			w.logger.WithError(err).Errorf("failed to fetch candle list")
			return nil, err
		}
		if len(list) < limit {
			end = true
		}
		candles = append(candles, list...)
	}
	return candles, nil
}
