package repository

import (
	"context"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
)

type CandleRepository interface {
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (CandleRepository, error) {
	if err := db.PostgresDB.AutoMigrate(new(entity.Candle)); err != nil {
		return nil, err
	}
	return NewCandlePostgresRepository(ctx, logger, db.PostgresDB)
}
