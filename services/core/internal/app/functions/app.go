package functions

import (
	"context"
	brokeragesService "github.com/h-varmazyar/Gate/services/core/internal/app/brokerages/service"
	"github.com/h-varmazyar/Gate/services/core/internal/app/functions/service"
	"github.com/h-varmazyar/Gate/services/core/internal/app/functions/service/callbacks"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Service *service.Service
}

func NewApp(ctx context.Context, logger *log.Logger, configs *Configs, brService *brokeragesService.Service) (*App, error) {
	err := callbacks.ListenOHLCCallbacks()
	if err != nil {
		return nil, err
	}
	return &App{
		Service: service.NewService(ctx, logger, configs.ServiceConfigs, brService),
	}, nil
}
