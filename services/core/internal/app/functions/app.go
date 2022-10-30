package functions

import (
	"context"
	brokeragesService "github.com/h-varmazyar/Gate/services/core/internal/app/brokerages/service"
	"github.com/h-varmazyar/Gate/services/core/internal/app/functions/service"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Service *service.Service
}

func NewApp(ctx context.Context, logger *log.Logger, configs *Configs, brService *brokeragesService.Service) (*App, error) {
	return &App{
		Service: service.NewService(ctx, logger, configs.ServiceConfigs, brService),
	}, nil
}
