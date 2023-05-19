package workers

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
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

func (w *MissedCandles) Start(runners []*Runner) {
	if !w.Started {
		w.logger.Infof("starting missed candle")
		w.ctx, w.cancelFunc = context.WithCancel(context.Background())
		go w.run(runners)
		w.Started = true
	}
}

func (w *MissedCandles) Stop() {
	if w.Started {
		w.cancelFunc()
	}
}

func (w *MissedCandles) run(runners []*Runner) {
	ticker := time.NewTicker(w.configs.MissedCandlesInterval)
	for {
		select {
		case <-w.ctx.Done():
			w.logger.Infof("missed stopped")
			ticker.Stop()
			return
		case <-ticker.C:
			w.logger.Infof("missed added: %v", time.Now())
			for _, runner := range runners {
				if err := w.checkForMissedCandles(runner); err != nil {
					w.logger.WithError(err).Error("failed to prepare missed rateLimiters")
				}
			}
		}
	}
}

//func (w *MissedCandles) prepareMarkets(markets []*chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) error {
//	for _, market := range markets {
//		err := w.prepareResolutions(market, resolutions)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func (w *MissedCandles) prepareResolutions(market *chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) error {
//	for _, resolution := range resolutions {
//		err := w.checkForMissedCandles(market, resolution)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

func (w *MissedCandles) checkForMissedCandles(runner *Runner) error {
	resolutionID, err := uuid.Parse(runner.Resolution.ID)
	if err != nil {
		return err
	}
	marketID, err := uuid.Parse(runner.Market.ID)
	if err != nil {
		return err
	}
	candles, err := w.db.ReturnList(marketID, resolutionID, 1000000000, 0)
	if err != nil {
		return err
	}

	for i := 1; i < len(candles); i++ {
		if candles[i-1].Time.Add(time.Duration(runner.Resolution.Duration)).Before(candles[i].Time) {
			from := candles[i-1].Time.Add(time.Duration(runner.Resolution.Duration))
			to := candles[i].Time.Add(time.Duration(runner.Resolution.Duration) * -1)
			if from.After(to) {
				continue
			}
			_, err := w.functionsService.AsyncOHLC(context.Background(), &coreApi.OHLCReq{
				Resolution: runner.Resolution,
				Market:     runner.Market,
				From:       from.Unix(),
				To:         to.Unix(),
				Platform:   runner.Platform,
			})
			if err != nil {
				w.logger.WithError(err).Errorf("failed to create missed async OHLC request %v in resolution %v and Platform %v",
					runner.Market.Name, runner.Resolution.Duration, runner.Platform)
			}
		}
	}
	return nil
}
