package workers

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	indicatorsPkg "github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/indicators"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

func (w *LastCandles) Start(runners []*Runner, indicators []indicatorsPkg.Indicator) {
	if !w.Started {
		w.logger.Infof("starting last candle")
		w.ctx, w.cancelFunc = context.WithCancel(context.Background())

		w.preparePrimaryDataRequests(runners, indicators)

		go w.run(runners)
		w.Started = true
	}
}

func (w *LastCandles) Stop() {
	if w.Started {
		w.cancelFunc()
	}
}

func (w *LastCandles) preparePrimaryDataRequests(runners []*Runner, indicators []indicatorsPkg.Indicator) {
	for _, runner := range runners {
		from, err := w.prepareLocalCandles(runner, indicators)
		if err != nil {
			return
		}
		w.makePrimaryDataRequests(runner, from)
	}
}

func (w *LastCandles) prepareLocalCandles(runner *Runner, indicators []indicatorsPkg.Indicator) (time.Time, error) {
	marketID, err := uuid.Parse(runner.Market.ID)
	if err != nil {
		w.logger.WithError(err).Errorf("invalid market id %v", runner.Market)
		return time.Unix(0, 0), err
	}
	resolutionID, err := uuid.Parse(runner.Resolution.ID)
	if err != nil {
		w.logger.WithError(err).Errorf("invalid resolution id %v", runner.Resolution)
		return time.Unix(0, 0), err
	}
	var from time.Time
	candles, err := w.loadLocalCandles(marketID, resolutionID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			from = time.Unix(runner.Market.IssueDate, 0)
		} else {
			w.logger.WithError(err).Errorf("failed to load local candles for market %v in resolution %v", marketID, resolutionID)
			return time.Unix(0, 0), err
		}
	}

	for _, candle := range candles {
		candle.IndicatorValues = entity.NewIndicatorValues()
	}

	if len(candles) > 0 {
		if err = w.calculateIndicators(candles, indicators); err != nil {
			w.logger.WithError(err).Errorf("failed to calculate indicators for market %v in resolution %v", marketID, resolutionID)
			return time.Unix(0, 0), err
		}
		from = candles[len(candles)-1].Time.Add(time.Duration(runner.Resolution.Duration))

		for _, candle := range candles {
			buffer.CandleBuffer.Push(candle)
		}
	} else {
		from = time.Unix(runner.Market.IssueDate, 0)
	}
	return from, nil
}

func (w *LastCandles) loadLocalCandles(marketID, resolutionID uuid.UUID) ([]*entity.Candle, error) {
	end := false
	limit := 1000000
	candles := make([]*entity.Candle, 0)
	for offset := 0; !end; offset += limit {
		list, err := w.db.ReturnList(marketID, resolutionID, limit, offset)
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

func (w *LastCandles) calculateIndicators(candles []*entity.Candle, indicators []indicatorsPkg.Indicator) error {
	for _, indicator := range indicators {
		err := indicator.Calculate(candles)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *LastCandles) makePrimaryDataRequests(runner *Runner, from time.Time) {
	_, err := w.functionsService.AsyncOHLC(context.Background(), &coreApi.OHLCReq{
		Resolution: runner.Resolution,
		Market:     runner.Market,
		From:       from.Unix(),
		To:         time.Now().Unix(),
		Platform:   runner.Platform,
	})
	if err != nil {
		w.logger.WithError(err).Errorf("failed to create primary async OHLC request %v in resolution %v and Platform %v",
			runner.Market.Name, runner.Resolution.Duration, runner.Platform)
	}
}

func (w *LastCandles) run(runners []*Runner) {
	ticker := time.NewTicker(w.configs.LastCandlesInterval)
	for {
		select {
		case <-w.ctx.Done():
			w.logger.Infof("missed stopped")
			ticker.Stop()
			return
		case <-ticker.C:
			w.logger.Infof("running last: %v", len(runners))
			for _, runner := range runners {
				w.checkForLastCandle(runner)
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
