package internal

import (
	"context"
	"github.com/h-varmazyar/Gate/services/indicators/internal/service"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Service *service.Service
}

type AppDependencies struct {
}

func NewApp(ctx context.Context, logger *log.Logger, configs Configs, dependencies *AppDependencies) (*App, error) {
	return &App{
		Service: service.NewService(ctx, logger, configs.ServiceConfigs),
	}, nil
}
