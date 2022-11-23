package strategies

import (
	"context"
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/strategies/repository"
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/strategies/service"
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/strategies/workers"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/db"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Service *service.Service
}

func NewApp(ctx context.Context, logger *log.Logger, db *db.DB, configs *Configs) (*App, error) {
	repositoryInstance, err := repository.NewRepository(ctx, logger, db)
	if err != nil {
		return nil, err
	}

	dependencies := &service.Dependencies{
		SignalCheckWorker: workers.SignalCheckWorkerInstance(),
	}
	return &App{
		Service: service.NewService(ctx, logger, configs.ServiceConfigs, repositoryInstance, dependencies),
	}, nil
}
