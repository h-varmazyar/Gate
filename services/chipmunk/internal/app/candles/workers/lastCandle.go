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

func (w *LastCandles) Start(runners map[string]*Runner) {
	if !w.Started {
		w.logger.Infof("starting last candle worker")
		w.ctx, w.cancelFunc = context.WithCancel(context.Background())

		go w.run(runners)
		w.Started = true
	}
}

func (w *LastCandles) Stop() {
	if w.Started {
		w.cancelFunc()
	}
}

func (w *LastCandles) run(runners map[string]*Runner) {
	ticker := time.NewTicker(w.configs.LastCandlesInterval)
	for {
		select {
		case <-w.ctx.Done():
			w.logger.Infof("missed stopped")
			ticker.Stop()
			return
		case <-ticker.C:
			for _, runner := range runners {
				if runner.IsPrimaryCandlesLoaded {
					w.checkForLastCandle(runner)
				}
			}
		}
	}
}

//func (w *LastCandles) prepareMarkets(markets []*chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) {
//	for _, market := range markets {
//		w.prepareResolutions(market, resolutions)
//	}
//}
//
//func (w *LastCandles) prepareResolutions(market *chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) {
//	for _, resolution := range resolutions {
//		w.checkForLastCandle(market, resolution)
//	}
//}

func (w *LastCandles) checkForLastCandle(runner *Runner) {
	last := buffer.CandleBuffer.Last(runner.Market.ID, runner.Resolution.ID)
	if last == nil {
		return
	}
	_, err := w.functionsService.AsyncOHLC(context.Background(), &coreApi.OHLCReq{
		Resolution: runner.Resolution,
		Market:     runner.Market,
		From:       last.Time.Unix(),
		To:         time.Now().Unix(),
		Platform:   runner.Market.Platform,
		Timeout:    int64(time.Second),
		IssueTime:  time.Now().Unix(),
	})
	if err != nil {
		w.logger.WithError(err).Errorf("failed to create missed async OHLC request %v in resolution %v and Platform %v",
			runner.Market.Name, runner.Resolution.Duration, runner.Platform)
	}
}
