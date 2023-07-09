package candles

import (
	"context"
	"github.com/google/uuid"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/workers"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	indicatorsPkg "github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/indicators"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	"gorm.io/gorm"
	"time"
)

func (app *App) initializeWorkers(ctx context.Context, configs *workers.Configs, repositoryInstance repository.CandleRepository) (*workerHolder, error) {
	var err error
	holder := new(workerHolder)
	holder.candleReaderWorker, err = workers.NewCandleReaderWorker(ctx, repositoryInstance, configs, app.logger)
	if err != nil {
		app.logger.WithError(err).Error("failed to initialize primary data worker")
		return nil, err
	}

	holder.lastCandleWorker = workers.NewLastCandles(ctx, repositoryInstance, configs, app.logger)

	holder.missedCandlesWorker = workers.NewMissedCandles(ctx, repositoryInstance, configs, app.logger)

	holder.redundantRemoverWorker = workers.NewRedundantRemover(ctx, repositoryInstance, configs, app.logger)

	return holder, nil
}

func (app *App) startWorkers(ctx context.Context, holder *workerHolder, dependencies *AppDependencies, platforms []api.Platform) error {
	indicators, err := dependencies.IndicatorService.List(ctx, &chipmunkApi.IndicatorListReq{Type: chipmunkApi.Indicator_All})
	if err != nil {
		app.logger.WithError(err).Error("failed to fetch all indicatorService")
		return err
	}

	loadedIndicators, err := app.loadIndicators(indicators.Elements)
	if err != nil {
		return err
	}

	runners, err := app.prepareRunners(ctx, dependencies, platforms)
	if err != nil {
		app.logger.WithError(err).Error("failed to prepare worker runners")
		return err
	}

	app.preparePrimaryDataRequests(runners, loadedIndicators)

	holder.candleReaderWorker.Start(runners, loadedIndicators)
	holder.lastCandleWorker.Start(runners)
	holder.missedCandlesWorker.Start(runners)
	holder.redundantRemoverWorker.Start(runners)

	return nil
}

func (app *App) prepareRunners(ctx context.Context, dependencies *AppDependencies, platforms []api.Platform) (map[string]*workers.Runner, error) {
	runners := make(map[string]*workers.Runner)
	for _, platform := range platforms {
		markets, err := dependencies.MarketService.List(ctx, &chipmunkApi.MarketListReq{Platform: platform})
		if err != nil {
			app.logger.WithError(err).Errorf("failed to load markets of platform %v", platform)
			return nil, err
		}

		resolutions, err := dependencies.ResolutionService.List(ctx, &chipmunkApi.ResolutionListReq{Platform: platform})
		if err != nil {
			app.logger.WithError(err).Errorf("failed to load resolutions of platform %v", platform)
			return nil, err
		}
		for _, market := range markets.Elements {
			for _, resolution := range resolutions.Elements {
				runners[market.ID] = &workers.Runner{
					Platform:   platform,
					Market:     market,
					Resolution: resolution,
				}
			}
		}
	}
	return runners, nil
}

func (app *App) loadIndicators(indicators []*chipmunkApi.Indicator) ([]indicatorsPkg.Indicator, error) {
	response := make([]indicatorsPkg.Indicator, 0)
	for _, i := range indicators {
		var indicatorCalculator indicatorsPkg.Indicator
		var err error
		indicator := new(entity.Indicator)
		mapper.Struct(i, indicator)
		indicator.ID, err = uuid.Parse(i.ID)
		if err != nil {
			return nil, err
		}
		switch indicator.Type {
		case chipmunkApi.Indicator_RSI:
			indicatorCalculator, err = indicatorsPkg.NewRSI(indicator.ID, indicator.Configs.RSI)
		case chipmunkApi.Indicator_Stochastic:
			indicatorCalculator, err = indicatorsPkg.NewStochastic(indicator.ID, indicator.Configs.Stochastic)
		case chipmunkApi.Indicator_MovingAverage:
			indicatorCalculator, err = indicatorsPkg.NewMovingAverage(indicator.ID, indicator.Configs.MovingAverage)
		case chipmunkApi.Indicator_BollingerBands:
			indicatorCalculator, err = indicatorsPkg.NewBollingerBands(indicator.ID, indicator.Configs.BollingerBands)
		}
		if err != nil {
			return nil, err
		}
		response = append(response, indicatorCalculator)
	}
	return response, nil
}

func (app *App) preparePrimaryDataRequests(runners map[string]*workers.Runner, indicators []indicatorsPkg.Indicator) {
	for _, runner := range runners {
		from, err := app.prepareLocalCandles(runner, indicators)
		if err != nil {
			return
		}
		app.makePrimaryDataRequests(runner, from)
	}
}

func (app *App) prepareLocalCandles(runner *workers.Runner, indicators []indicatorsPkg.Indicator) (time.Time, error) {
	marketID, err := uuid.Parse(runner.Market.ID)
	if err != nil {
		app.logger.WithError(err).Errorf("invalid market id %v", runner.Market)
		return time.Unix(0, 0), err
	}
	resolutionID, err := uuid.Parse(runner.Resolution.ID)
	if err != nil {
		app.logger.WithError(err).Errorf("invalid resolution id %v", runner.Resolution)
		return time.Unix(0, 0), err
	}
	var from time.Time
	candles, err := app.loadLocalCandles(marketID, resolutionID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			from = time.Unix(runner.Market.IssueDate, 0)
		} else {
			app.logger.WithError(err).Errorf("failed to load local rateLimiters for market %v in resolution %v", marketID, resolutionID)
			return time.Unix(0, 0), err
		}
	}

	for _, candle := range candles {
		candle.IndicatorValues = entity.NewIndicatorValues()
	}

	if len(candles) > 0 {
		if err = app.calculateIndicators(candles, indicators); err != nil {
			app.logger.WithError(err).Errorf("failed to calculate indicators for market %v in resolution %v", marketID, resolutionID)
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

func (app *App) loadLocalCandles(marketID, resolutionID uuid.UUID) ([]*entity.Candle, error) {
	end := false
	limit := 1000000
	candles := make([]*entity.Candle, 0)
	for offset := 0; !end; offset += limit {
		list, err := app.db.ReturnList(marketID, resolutionID, limit, offset)
		if err != nil {
			app.logger.WithError(err).Errorf("failed to fetch candle list")
			return nil, err
		}
		if len(list) < limit {
			end = true
		}
		candles = append(candles, list...)
	}
	return candles, nil
}

func (app *App) calculateIndicators(candles []*entity.Candle, indicators []indicatorsPkg.Indicator) error {
	for _, indicator := range indicators {
		err := indicator.Calculate(candles)
		if err != nil {
			return err
		}
	}
	return nil
}

func (app *App) makePrimaryDataRequests(runner *workers.Runner, from time.Time) {
	asyncResp, err := app.functionsService.AsyncOHLC(context.Background(), &coreApi.OHLCReq{
		Resolution: runner.Resolution,
		Market:     runner.Market,
		From:       from.Unix(),
		To:         time.Now().Unix(),
		Platform:   runner.Platform,
	})
	if err != nil {
		app.logger.WithError(err).Errorf("failed to create primary async OHLC request %v in resolution %v and Platform %v",
			runner.Market.Name, runner.Resolution.Duration, runner.Platform)
	}
	runner.LastEventID = asyncResp.LastRequestID
}
