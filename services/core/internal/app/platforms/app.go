package platforms

import (
	"context"
	"github.com/h-varmazyar/Gate/services/core/internal/app/platforms/service"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Service *service.Service
}

func NewApp(ctx context.Context, logger *log.Logger, configs *Configs) (*App, error) {
	return &App{
		Service: service.NewService(ctx, logger, configs.ServiceConfigs),
	}, nil
}
