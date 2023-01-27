package candles

import (
	"context"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/workers"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Service *service.Service
}

type AppDependencies struct {
	ServiceDependencies *service.Dependencies
}

func NewApp(ctx context.Context, logger *log.Logger, db *db.DB, configs *Configs, appDependencies *AppDependencies) (*App, error) {
	repositoryInstance, err := repository.NewRepository(ctx, logger, db)
	if err != nil {
		return nil, err
	}

	candleBuffer := buffer.NewCandleBufferInstance(configs.BufferConfigs)
	primaryDataWorker, err := workers.NewPrimaryDataWorker(ctx, repositoryInstance, configs.WorkerConfigs, candleBuffer)
	if err != nil {
		logger.WithError(err).Error("failed to initialize primary data worker")
		return nil, err
	}

	lastCandleWorker := workers.NewLastCandles(ctx, repositoryInstance, configs.WorkerConfigs, logger, candleBuffer)
	if err != nil {
		logger.WithError(err).Error("failed to initialize primary data worker")
		return nil, err
	}

	missedCandleWorker := workers.NewMissedCandles(ctx, repositoryInstance, configs.WorkerConfigs, logger)
	if err != nil {
		logger.WithError(err).Error("failed to initialize primary data worker")
		return nil, err
	}

	redundantRemover := workers.NewRedundantRemover(ctx, repositoryInstance, configs.WorkerConfigs, logger)
	if err != nil {
		logger.WithError(err).Error("failed to initialize primary data worker")
		return nil, err
	}

	appDependencies.ServiceDependencies.Buffer = candleBuffer
	appDependencies.ServiceDependencies.PrimaryDataWorker = primaryDataWorker
	appDependencies.ServiceDependencies.LastCandleWorker = lastCandleWorker
	appDependencies.ServiceDependencies.MissedCandlesWorker = missedCandleWorker
	appDependencies.ServiceDependencies.RedundantRemoverWorker = redundantRemover
	return &App{
		Service: service.NewService(ctx, logger, configs.ServiceConfigs, repositoryInstance, appDependencies.ServiceDependencies),
	}, nil
}
