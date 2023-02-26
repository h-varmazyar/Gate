package rateLimiters

import (
	"context"
	"github.com/h-varmazyar/Gate/services/network/internal/app/rateLimiters/repository"
	"github.com/h-varmazyar/Gate/services/network/internal/app/rateLimiters/service"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/db"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Service *service.Service
}

func NewApp(ctx context.Context, logger *log.Logger, db *db.DB) (*App, error) {
	repositoryInstance, err := repository.NewRepository(ctx, logger, db)
	if err != nil {
		return nil, err
	}
	s := &App{
		Service: service.NewService(ctx, logger, repositoryInstance),
	}
	return s, nil
}
