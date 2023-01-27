package workers

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
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
	buffer           *buffer.CandleBuffer
	Started          bool
}

func NewLastCandles(_ context.Context, db repository.CandleRepository, configs *Configs, logger *log.Logger, buffer *buffer.CandleBuffer) *LastCandles {
	coreConn := grpcext.NewConnection(configs.CoreAddress)
	return &LastCandles{
		db:               db,
		logger:           logger,
		configs:          configs,
		buffer:           buffer,
		functionsService: coreApi.NewFunctionsServiceClient(coreConn),
	}
}

func (w *LastCandles) Start(markets []*chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) {
	if !w.Started {
		w.logger.Infof("starting last candle")
		w.ctx, w.cancelFunc = context.WithCancel(context.Background())
		go w.run(markets, resolutions)
		w.Started = true
	}
}

func (w *LastCandles) Stop() {
	if w.Started {
		w.cancelFunc()
	}
}

func (w *LastCandles) run(markets []*chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) {
	ticker := time.NewTicker(w.configs.LastCandlesInterval)
	for {
		select {
		case <-w.ctx.Done():
			w.logger.Infof("missed stopped")
			ticker.Stop()
			return
		case <-ticker.C:
			w.logger.Infof("running last: %v - %v", len(markets), len(resolutions))
			w.prepareMarkets(markets, resolutions)
		}
	}
}

func (w *LastCandles) prepareMarkets(markets []*chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) {
	for _, market := range markets {
		w.prepareResolutions(market, resolutions)
	}
}

func (w *LastCandles) prepareResolutions(market *chipmunkApi.Market, resolutions []*chipmunkApi.Resolution) {
	for _, resolution := range resolutions {
		w.checkForLastCandle(market, resolution)
	}
}

func (w *LastCandles) checkForLastCandle(market *chipmunkApi.Market, resolution *chipmunkApi.Resolution) {
	last := w.buffer.ReturnCandles(market.ID, resolution.ID, 1)
	if len(last) == 0 {
		return
	}
	_, err := w.functionsService.AsyncOHLC(context.Background(), &coreApi.OHLCReq{
		Resolution: resolution,
		Market:     market,
		From:       last[0].Time.Unix(),
		To:         time.Now().Unix(),
		Platform:   market.Platform,
		Timeout:    int64(time.Second),
		IssueTime:  time.Now().Unix(),
	})
	if err != nil {
		w.logger.WithError(err).Errorf("failed to create missed async OHLC request %v in resolution %v and Platform %v", market.Name, resolution.Duration, market.Platform)
	}
}
