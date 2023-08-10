package markets

import (
	"context"
	api "github.com/h-varmazyar/Gate/api/proto"
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
	ServiceDependencies *service.Dependencies
}

func NewApp(ctx context.Context, logger *log.Logger, db *db.DB, configs Configs, dependencies *AppDependencies) (*App, error) {
	repositoryInstance, err := repository.NewRepository(ctx, logger, db)
	if err != nil {
		return nil, err
	}

	updateMarketWorker := workers.NewUpdateMarketWorker(repositoryInstance, logger, dependencies.ServiceDependencies.AssetsService, configs.WorkerConfigs.CoreAddress)
	_, err = updateMarketWorker.UpdateFromPlatform(ctx, api.Platform_Coinex)
	if err != nil {
		return nil, err
	}
	
	return &App{
		Service: service.NewService(ctx, logger, configs.ServiceConfigs, repositoryInstance, dependencies.ServiceDependencies),
	}, nil
}
