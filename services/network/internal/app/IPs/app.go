package IPs

import (
	"context"
	"github.com/h-varmazyar/Gate/services/network/internal/app/IPs/repository"
	"github.com/h-varmazyar/Gate/services/network/internal/app/IPs/service"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/db"
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
	s := &App{
		Service: service.NewService(ctx, logger, configs.ServiceConfigs, repositoryInstance),
	}
	return s, nil
}