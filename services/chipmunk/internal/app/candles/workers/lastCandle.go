package workers

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	log "github.com/sirupsen/logrus"
	"time"
)

type LastCandles struct {
	db               repository.CandleRepository
	configs          *Configs
	functionsService coreApi.FunctionsServiceClient
	ctx              context.Context
	cancelFunc       context.CancelFunc
	logger           *log.Logger
	Started          bool
}

func NewLastCandles(_ context.Context, db repository.CandleRepository, configs *Configs, logger *log.Logger) *LastCandles {
	coreConn := grpcext.NewConnection(configs.CoreAddress)
	return &LastCandles{
		db:               db,
		logger:           logger,
		configs:          configs,
		functionsService: coreApi.NewFunctionsServiceClient(coreConn),
	}
}

func (w *LastCandles) Start(platformsPairs []*PlatformPairs) {
	if !w.Started {
		w.logger.Infof("starting last candle worker")
		w.ctx, w.cancelFunc = context.WithCancel(context.Background())

		go w.run(platformsPairs)
		w.Started = true
	}
}

func (w *LastCandles) Stop() {
	if w.Started {
		w.cancelFunc()
	}
}

func (w *LastCandles) run(platformsPairs []*PlatformPairs) {
	ticker := time.NewTicker(w.configs.LastCandlesInterval)
	for {
		select {
		case <-w.ctx.Done():
			w.logger.Infof("missed stopped")
			ticker.Stop()
			return
		case <-ticker.C:
			for _, platformPair := range platformsPairs {
				w.checkForLastCandle(platformPair)
			}
		}
	}
}

func (w *LastCandles) checkForLastCandle(platformPair *PlatformPairs) {
	items := make([]*coreApi.OHLCItem, 0)
	for _, pair := range platformPair.Pairs {
		last := buffer.CandleBuffer.Last(pair.Market.ID, pair.Resolution.ID)
		if last == nil {
			return
		}
		items = append(items, &coreApi.OHLCItem{
			Resolution: pair.Resolution,
			Market:     pair.Market,
			From:       last.Time.Unix(),
			To:         time.Now().Unix(),
			Timeout:    int64(w.configs.LastCandlesInterval),
			IssueTime:  time.Now().Unix(),
		})
	}

	_, err := w.functionsService.AsyncOHLC(context.Background(), &coreApi.AsyncOHLCReq{
		Items:    items,
		Platform: platformPair.Platform,
	})
	if err != nil {
		w.logger.WithError(err).Errorf("failed to create last candle request for %v", platformPair.Platform)
	}
}
