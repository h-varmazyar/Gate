package repository

import (
	"context"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/db"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
)

type RateLimiterRepository interface {
	Create(brokerage *entity.RateLimiter) error
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (RateLimiterRepository, error) {
	if err := db.PostgresDB.AutoMigrate(new(entity.RateLimiter)); err != nil {
		return nil, err
	}
	return NewRateLimiterPostgresRepository(ctx, logger, db.PostgresDB)
}
