package markets

import (
	"context"
	candles "github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets/service"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets/workers"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Service *service.Service
}

type AppDependencies struct {
	CandlesService      *candles.Service
	ServiceDependencies *service.Dependencies
}

func NewApp(ctx context.Context, logger *log.Logger, db *db.DB, configs *Configs, dependencies *AppDependencies) (*App, error) {
	repositoryInstance, err := repository.NewRepository(ctx, logger, db)
	if err != nil {
		return nil, err
	}
	//worker := workers.InitializeWorker(ctx, configs.WorkerConfigs, dependencies.CandlesService)
	statisticsWorker := workers.NewStatisticsWorker(ctx, configs.WorkerConfigs, dependencies.CandlesService)
	//dependencies.ServiceDependencies.PrimaryDataWorker = worker
	dependencies.ServiceDependencies.StatisticsWorker = statisticsWorker
	return &App{
		Service: service.NewService(ctx, logger, configs.ServiceConfigs, repositoryInstance, dependencies.ServiceDependencies),
	}, nil
}
