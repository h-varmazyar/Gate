package workers

import (
	"context"
	"github.com/google/uuid"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	log "github.com/sirupsen/logrus"
	"time"
)

type RedundantRemover struct {
	db           repository.CandleRepository
	configs      *Configs
	ctx          context.Context
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

func (w *RedundantRemover) Start(markets []*chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) {
	if !w.Started {
		w.logger.Infof("starting redundant candles")
		w.ctx = context.Background()
		go w.run(markets, resolutions)
		w.Started = true
	}
}

func (w *RedundantRemover) run(markets []*chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) {
	ticker := time.NewTicker(w.configs.RedundantRemoverInterval)
	for {
		select {
		case <-w.ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			w.removedCount = 0
			w.logger.Infof("prepare removed candles")
			if err := w.prepareMarkets(markets, resolutions); err != nil {
				log.WithError(err).Error("failed to prepare missed candles")
			}
			w.logger.Infof("removed candles: %v", w.removedCount)
		}
	}
}

func (w *RedundantRemover) prepareMarkets(markets []*chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) error {
	for _, market := range markets {
		err := w.prepareResolutions(market, resolutions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *RedundantRemover) prepareResolutions(market *chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) error {
	for _, resolution := range resolutions {
		err := w.removeRedundantCandles(market, resolution)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *RedundantRemover) removeRedundantCandles(market *chipmunkApi.Market, resolution *chipmunkApi.Resolution) error {
	resolutionID, err := uuid.Parse(resolution.ID)
	if err != nil {
		return err
	}
	marketID, err := uuid.Parse(market.ID)
	if err != nil {
		return err
	}
	candles, err := w.db.ReturnList(marketID, resolutionID, 1000000, 0)
	if err != nil {
		return err
	}

	for i := 1; i < len(candles); i++ {
		if candles[i-1].Time.Equal(candles[i].Time) {
			if err := w.db.HardDelete(candles[i-1]); err != nil {
				w.logger.WithError(err).Errorf("failed to delete candle %v", candles[i-1].ID)
			}
			w.removedCount++
		}
	}
	return nil
}
