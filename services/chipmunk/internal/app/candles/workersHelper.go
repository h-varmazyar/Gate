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
	"time"
)

func (app *App) initializeWorkers(ctx context.Context, configs *workers.Configs, repositoryInstance repository.CandleRepository) error {
	var err error
	app.candleReaderWorker, err = workers.NewCandleReaderWorker(ctx, repositoryInstance, configs, app.logger)
	if err != nil {
		app.logger.WithError(err).Error("failed to initialize primary data worker")
		return err
	}

	app.lastCandleWorker = workers.NewLastCandles(ctx, repositoryInstance, configs, app.logger)

	app.missedCandlesWorker = workers.NewMissedCandles(ctx, repositoryInstance, configs, app.logger)

	app.redundantRemoverWorker = workers.NewRedundantRemover(ctx, repositoryInstance, configs, app.logger)

	return nil
}

func (app *App) startWorkers(ctx context.Context, dependencies *AppDependencies, platforms []api.Platform) error {
	indicators, err := dependencies.IndicatorService.List(ctx, &chipmunkApi.IndicatorListReq{Type: chipmunkApi.Indicator_All})
	if err != nil {
		app.logger.WithError(err).Error("failed to fetch all indicatorService")
		return err
	}

	loadedIndicators, err := app.loadIndicators(indicators.Elements)
	if err != nil {
		return err
	}

	pp := make([]*workers.PlatformPairs, 0)
	for _, platform := range platforms {
		pairs := new(workers.PlatformPairs)
		pairs, err = app.preparePlatformPairs(ctx, dependencies, platform)
		if err != nil {
			app.logger.WithError(err).Error("failed to prepare worker runners")
			return err
		}
		app.logger.Infof("loaded pairs: %v", len(pairs.Pairs))
		pp = append(pp, pairs)
	}

	app.candleReaderWorker.Start(loadedIndicators)

	go func() {
		predictedInterval := app.preparePrimaryDataRequests(pp, loadedIndicators)
		app.logger.Infof("predicted interval: %v", predictedInterval)
		time.Sleep(time.Duration(predictedInterval))
		app.lastCandleWorker.Start(pp)
		app.missedCandlesWorker.Start(pp)
		app.redundantRemoverWorker.Start(pp)
	}()

	return nil
}

func (app *App) preparePlatformPairs(ctx context.Context, dependencies *AppDependencies, platform api.Platform) (*workers.PlatformPairs, error) {
	var (
		err         error
		markets     *chipmunkApi.Markets
		resolutions *chipmunkApi.Resolutions
		pairs       = make([]*workers.Pair, 0)
	)
	markets, err = dependencies.MarketService.List(ctx, &chipmunkApi.MarketListReq{Platform: platform})
	if err != nil {
		app.logger.WithError(err).Errorf("failed to load markets of platform %v", platform)
		return nil, err
	}

	app.logger.Infof("market len: %v", len(markets.Elements))

	resolutions, err = dependencies.ResolutionService.List(ctx, &chipmunkApi.ResolutionListReq{Platform: platform})
	if err != nil {
		app.logger.WithError(err).Errorf("failed to load resolutions of platform %v", platform)
		return nil, err
	}

	app.logger.Infof("resolutions len: %v", len(resolutions.Elements))

	for _, market := range markets.Elements {
		for _, resolution := range resolutions.Elements {
			pairs = append(pairs, &workers.Pair{
				Market:     market,
				Resolution: resolution,
			})
		}
	}
	return &workers.PlatformPairs{
		Platform: platform,
		Pairs:    pairs,
	}, nil
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

func (app *App) preparePrimaryDataRequests(platformsPairs []*workers.PlatformPairs, indicators []indicatorsPkg.Indicator) int64 {
	app.logger.Infof("preparing primary candles")
	totalPredictedInterval := int64(0)
	for _, platformPairs := range platformsPairs {
		totalPredictedInterval += app.makePrimaryDataRequests(platformPairs, indicators)
	}
	return totalPredictedInterval
}

func (app *App) prepareLocalCandlesItem(pair *workers.Pair, indicators []indicatorsPkg.Indicator) (*coreApi.OHLCItem, error) {
	marketID, err := uuid.Parse(pair.Market.ID)
	if err != nil {
		app.logger.WithError(err).Errorf("invalid market id %v", marketID)
		return nil, err
	}
	resolutionID, err := uuid.Parse(pair.Resolution.ID)
	if err != nil {
		app.logger.WithError(err).Errorf("invalid resolution id %v", resolutionID)
		return nil, err
	}
	var from time.Time
	candles, err := app.loadLocalCandles(marketID, resolutionID)
	if err != nil {
		app.logger.WithError(err).Errorf("failed to load local candles for market %v in resolution %v", marketID, resolutionID)
		return nil, err
	}
	if len(candles) == 0 {
		from = time.Unix(pair.Market.IssueDate, 0)
	} else {
		from = candles[len(candles)-1].Time.Add(time.Duration(pair.Resolution.Duration))
	}

	for _, candle := range candles {
		candle.IndicatorValues = entity.NewIndicatorValues()
	}

	app.logger.Infof("loaded candles: %v", len(candles))

	if len(candles) > 0 {
		if err = app.calculateIndicators(candles, indicators); err != nil {
			app.logger.WithError(err).Errorf("failed to calculate indicators for market %v in resolution %v", marketID, resolutionID)
			return nil, err
		}
	}

	for _, candle := range candles {
		buffer.CandleBuffer.Push(candle)
	}

	item := &coreApi.OHLCItem{
		Resolution: pair.Resolution,
		Market:     pair.Market,
		From:       from.Unix(),
		To:         time.Now().Unix(),
		Timeout:    int64(time.Hour * 96),
		IssueTime:  time.Now().Unix(),
	}
	return item, nil
}

func (app *App) loadLocalCandles(marketID, resolutionID uuid.UUID) ([]*entity.Candle, error) {
	end := false
	limit := 1000000
	candles := make([]*entity.Candle, 0)
	for offset := 0; !end; offset += limit {
		list, err := app.db.List(marketID, resolutionID, limit, offset)
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

func (app *App) makePrimaryDataRequests(platformPairs *workers.PlatformPairs, indicators []indicatorsPkg.Indicator) int64 {
	predictedInterval := int64(0)
	for _, pair := range platformPairs.Pairs {
		app.logger.Infof("market: %v", pair.Market.Name)
		item, err := app.prepareLocalCandlesItem(pair, indicators)
		if err != nil {
			app.logger.WithError(err).Errorf("failed to prepare local candle item")
			continue
		}
		app.logger.Infof("item: %v - %v", time.Unix(item.From, 0), time.Unix(item.To, 0))

		asyncResp, err := app.functionsService.AsyncOHLC(context.Background(), &coreApi.AsyncOHLCReq{
			Items:    []*coreApi.OHLCItem{item},
			Platform: platformPairs.Platform,
		})
		if err != nil {
			app.logger.WithError(err).Errorf("failed to create primary async OHLC request for Platform %v", platformPairs.Platform)
			continue
		}
		app.logger.Infof("create new bulk request with id %v for %v. estimated execution time: %v", asyncResp.LastRequestID, platformPairs.Platform, time.Duration(asyncResp.PredictedIntervalTime))
		predictedInterval += asyncResp.PredictedIntervalTime
	}

	return predictedInterval
}
