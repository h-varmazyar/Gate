package candles

import (
	"context"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/workers"
	indicatorService "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/indicators/service"
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
	functionsService       coreApi.FunctionsServiceClient
	candleReaderWorker     *workers.CandleReader
	lastCandleWorker       *workers.LastCandles
	missedCandlesWorker    *workers.MissedCandles
	redundantRemoverWorker *workers.RedundantRemover
}

type AppDependencies struct {
	IndicatorService  *indicatorService.Service
	ResolutionService *resolutionService.Service
	MarketService     *marketService.Service
}

func NewApp(ctx context.Context, logger *log.Logger, db *db.DB, configs Configs, appDependencies *AppDependencies) (*App, error) {
	repositoryInstance, err := repository.NewRepository(ctx, logger, db)
	if err != nil {
		return nil, err
	}

	coreConn := grpcext.NewConnection(configs.WorkerConfigs.CoreAddress)
	app := &App{
		logger:           logger,
		db:               repositoryInstance,
		functionsService: coreApi.NewFunctionsServiceClient(coreConn),
	}

	err = app.initializeWorkers(ctx, configs.WorkerConfigs, repositoryInstance)
	if err != nil {
		return nil, err
	}

	platforms := []api.Platform{api.Platform_Coinex}

	err = app.startWorkers(ctx, appDependencies, platforms)
	if err != nil {
		return nil, err
	}

	app.Service = service.NewService(ctx, logger, configs.ServiceConfigs, repositoryInstance)

	return app, nil
}
