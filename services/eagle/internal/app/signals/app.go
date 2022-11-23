package signals

import (
	"context"
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/signals/service"
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/signals/workers"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Service *service.Service
}

type AppDependencies struct {
	ServiceDependencies *service.Dependencies
}

func NewApp(ctx context.Context, logger *log.Logger, configs *Configs, dependencies *AppDependencies) (*App, error) {
	dependencies.ServiceDependencies.SignalCheckWorker = workers.SignalCheckWorkerInstance()
	return &App{
		Service: service.NewService(ctx, logger, configs.ServiceConfigs, dependencies.ServiceDependencies),
	}, nil
}
