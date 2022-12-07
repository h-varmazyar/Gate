package workers

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
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

func (w *MissedCandles) Start(markets []*chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) {
	if !w.Started {
		w.logger.Infof("starting missed candle")
		w.ctx, w.cancelFunc = context.WithCancel(context.Background())
		go w.run(markets, resolutions)
		w.Started = true
	}
}

func (w *MissedCandles) Stop() {
	if w.Started {
		w.cancelFunc()
	}
}

func (w *MissedCandles) run(markets []*chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) {
	ticker := time.NewTicker(w.configs.MissedCandlesInterval)
	for {
		select {
		case <-w.ctx.Done():
			w.logger.Infof("missed stopped")
			ticker.Stop()
			return
		case <-ticker.C:
			w.logger.Infof("missed added: %v", time.Now())
			if err := w.prepareMarkets(markets, resolutions); err != nil {
				w.logger.WithError(err).Error("failed to prepare missed candles")
			}
		}
	}
}

func (w *MissedCandles) prepareMarkets(markets []*chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) error {
	for _, market := range markets {
		err := w.prepareResolutions(market, resolutions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *MissedCandles) prepareResolutions(market *chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) error {
	for _, resolution := range resolutions {
		err := w.checkForMissedCandles(market, resolution)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *MissedCandles) checkForMissedCandles(market *chipmunkApi.Market, resolution *chipmunkApi.Resolution) error {
	resolutionID, err := uuid.Parse(resolution.ID)
	if err != nil {
		return err
	}
	marketID, err := uuid.Parse(market.ID)
	if err != nil {
		return err
	}
	candles, err := w.db.ReturnList(marketID, resolutionID, 1000000000, 0)
	if err != nil {
		return err
	}

	for i := 1; i < len(candles); i++ {
		if candles[i-1].Time.Add(time.Duration(resolution.Duration)).Before(candles[i].Time) {
			from := candles[i-1].Time.Add(time.Duration(resolution.Duration))
			to := candles[i].Time.Add(time.Duration(resolution.Duration) * -1)
			if from.After(to) {
				continue
			}
			_, err := w.functionsService.AsyncOHLC(context.Background(), &coreApi.OHLCReq{
				Resolution: resolution,
				Market:     market,
				From:       from.Unix(),
				To:         to.Unix(),
				Platform:   market.Platform,
			})
			if err != nil {
				w.logger.WithError(err).Errorf("failed to create async OHLC request for marker %v in resolution %v and Platform %v", market.Name, resolution.Duration, market.Platform)
			}
		}
	}
	return nil
}
