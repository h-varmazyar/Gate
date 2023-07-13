package workers

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	log "github.com/sirupsen/logrus"
	"time"
)

type MissedCandles struct {
	db               repository.CandleRepository
	configs          *Configs
	functionsService coreApi.FunctionsServiceClient
	ctx              context.Context
	cancelFunc       context.CancelFunc
	logger           *log.Logger
	Started          bool
}

func NewMissedCandles(_ context.Context, db repository.CandleRepository, configs *Configs, logger *log.Logger) *MissedCandles {
	coreConn := grpcext.NewConnection(configs.CoreAddress)
	return &MissedCandles{
		db:               db,
		logger:           logger,
		configs:          configs,
		functionsService: coreApi.NewFunctionsServiceClient(coreConn),
	}
}

func (w *MissedCandles) Start(platformsPairs []*PlatformPairs) {
	if !w.Started {
		w.logger.Infof("starting missed candle worker")
		w.ctx, w.cancelFunc = context.WithCancel(context.Background())
		go w.run(platformsPairs)
		w.Started = true
	}
}

func (w *MissedCandles) Stop() {
	if w.Started {
		w.cancelFunc()
	}
}

func (w *MissedCandles) run(platformsPairs []*PlatformPairs) {
	ticker := time.NewTicker(w.configs.MissedCandlesInterval)
	for {
		select {
		case <-w.ctx.Done():
			w.logger.Infof("missed stopped")
			ticker.Stop()
			return
		case <-ticker.C:
			w.logger.Infof("missed added: %v", time.Now())
			for _, platformPairs := range platformsPairs {
				if err := w.checkForMissedCandles(platformPairs); err != nil {
					w.logger.WithError(err).Error("failed to prepare missed rateLimiters")
				}
			}
		}
	}
}

func (w *MissedCandles) checkForMissedCandles(platformPairs *PlatformPairs) error {
	items := make([]*coreApi.OHLCItem, 0)
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
			if candles[i-1].Time.Add(time.Duration(pair.Resolution.Duration)).Before(candles[i].Time) {
				from := candles[i-1].Time.Add(time.Duration(pair.Resolution.Duration))
				to := candles[i].Time.Add(time.Duration(pair.Resolution.Duration) * -1)
				if from.After(to) {
					continue
				}
				item := &coreApi.OHLCItem{
					Resolution: pair.Resolution,
					Market:     pair.Market,
					From:       from.Unix(),
					To:         to.Unix(),
					Timeout:    int64(w.configs.MissedCandlesInterval),
					IssueTime:  time.Now().Unix(),
				}

				items = append(items, item)
			}
		}
	}
	_, err := w.functionsService.AsyncOHLC(context.Background(), &coreApi.AsyncOHLCReq{
		Items:    items,
		Platform: platformPairs.Platform,
	})
	if err != nil {
		w.logger.WithError(err).Errorf("failed to create missed OHLC request for Platform %v", platformPairs.Platform)
		return err
	}
	return nil
}

func (w *MissedCandles) loadCandles(marketID, resolutionID uuid.UUID) ([]*entity.Candle, error) {
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
