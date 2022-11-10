package brokerages

import (
	"context"
	"github.com/h-varmazyar/Gate/services/core/internal/app/brokerages/repository"
	"github.com/h-varmazyar/Gate/services/core/internal/app/brokerages/service"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/db"
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
	return &App{
		Service: service.NewService(ctx, logger, configs.ServiceConfigs, repositoryInstance),
	}, nil
}
