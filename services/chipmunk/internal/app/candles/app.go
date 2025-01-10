package candles

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/workers"
	marketService "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets/service"
	resolutionService "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/resolutions/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Service                *service.Service
	logger                 *log.Logger
	db                     repository.CandleRepository
	candleReaderWorker     *workers.CandleReader
	lastCandleWorker       *workers.LastCandles
	missedCandlesWorker    *workers.MissedCandles
	redundantRemoverWorker *workers.RedundantRemover
}

type AppDependencies struct {
	ResolutionService *resolutionService.Service
	MarketService     *marketService.Service
}

func NewApp(ctx context.Context, logger *log.Logger, db *db.DB, configs Configs, appDependencies *AppDependencies) (*App, error) {
	repositoryInstance, err := repository.NewRepository(ctx, logger, db)
	if err != nil {
		return nil, err
	}

	app := &App{
		logger: logger,
		db:     repositoryInstance,
	}

	err = app.startWorkers(ctx, configs.WorkerConfigs, repositoryInstance, appDependencies)
	if err != nil {
		return nil, err
	}

	app.Service = service.NewService(ctx, logger, configs.ServiceConfigs, repositoryInstance)

	return app, nil
}

func (app *App) startWorkers(ctx context.Context, configs *workers.Configs, db repository.CandleRepository, dependencies *AppDependencies) error {
	log.Infof("starting candle app workers...")
	pp := preparePlatformPairs(ctx, dependencies.MarketService, dependencies.ResolutionService)
	nextRequestTrigger := make(chan bool, 10)
	candleReaderWorker, err := workers.NewCandleReaderWorker(ctx, db, configs, app.logger, nextRequestTrigger)
	if err != nil {
		return err
	}
	candleReaderWorker.Start()

	if configs.DataWarmupMood {
		nextRequestTrigger <- true
		coreConn := grpcext.NewConnection(configs.CoreAddress)
		functionsService := coreApi.NewFunctionsServiceClient(coreConn)
		preparePrimaryDataRequests(nextRequestTrigger, db, pp, functionsService)
		log.Infof("total warmup requests sent")
	}

	if configs.DataCorrectionMode {
		workers.NewMissedCandles(ctx, db, configs, app.logger).Start(pp)
		workers.NewRedundantRemover(ctx, db, configs, app.logger).Start(pp)
	}

	if configs.NormalDataGathering {
		workers.NewLastCandles(ctx, db, configs, app.logger).Start(pp)
	}

	return nil
}
