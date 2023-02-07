package candles

import (
	"context"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/workers"
	indicatorService "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/indicators/service"
	marketService "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets/service"
	resolutionService "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/resolutions/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Service *service.Service
	logger  *log.Logger
}

type AppDependencies struct {
	ServiceDependencies *service.Dependencies
	IndicatorService    *indicatorService.Service
	ResolutionService   *resolutionService.Service
	MarketService       *marketService.Service
}

type workerHolder struct {
	candleReaderWorker     *workers.CandleReader
	lastCandleWorker       *workers.LastCandles
	missedCandlesWorker    *workers.MissedCandles
	redundantRemoverWorker *workers.RedundantRemover
}

func NewApp(ctx context.Context, logger *log.Logger, db *db.DB, configs Configs, appDependencies *AppDependencies) (*App, error) {
	repositoryInstance, err := repository.NewRepository(ctx, logger, db)
	if err != nil {
		return nil, err
	}

	app := &App{
		logger: logger,
	}

	candleBuffer := buffer.NewCandleBufferInstance(configs.BufferConfigs)

	holder, err := app.initializeWorkers(ctx, configs.WorkerConfigs, candleBuffer, repositoryInstance)
	if err != nil {
		return nil, err
	}

	platforms := []api.Platform{api.Platform_Coinex}

	err = app.startWorkers(ctx, holder, appDependencies, platforms)

	appDependencies.ServiceDependencies.Buffer = candleBuffer
	app.Service = service.NewService(ctx, logger, configs.ServiceConfigs, repositoryInstance, appDependencies.ServiceDependencies)

	return app, nil
}
