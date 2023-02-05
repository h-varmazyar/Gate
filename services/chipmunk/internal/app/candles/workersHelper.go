package candles

import (
	"context"
	"github.com/google/uuid"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/workers"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	indicatorsPkg "github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/indicators"
)

func (app *App) initializeWorkers(ctx context.Context, configs *workers.Configs, candleBuffer *buffer.CandleBuffer, repositoryInstance repository.CandleRepository) (*workerHolder, error) {
	var err error
	holder := new(workerHolder)
	holder.candleReaderWorker, err = workers.NewCandleReaderWorker(ctx, repositoryInstance, configs, candleBuffer)
	if err != nil {
		app.logger.WithError(err).Error("failed to initialize primary data worker")
		return nil, err
	}

	holder.lastCandleWorker = workers.NewLastCandles(ctx, repositoryInstance, configs, app.logger, candleBuffer)

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

	//todo: must be pass copy of runner to each worker or not??
	holder.candleReaderWorker.Start(loadedIndicators)
	holder.lastCandleWorker.Start(runners, loadedIndicators)
	holder.missedCandlesWorker.Start(runners)
	holder.redundantRemoverWorker.Start(runners)

	return nil
}

func (app *App) prepareRunners(ctx context.Context, dependencies *AppDependencies, platforms []api.Platform) ([]*workers.Runner, error) {
	runners := make([]*workers.Runner, 0)
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
				runners = append(runners, &workers.Runner{
					Platform:   platform,
					Market:     market,
					Resolution: resolution,
				})
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
