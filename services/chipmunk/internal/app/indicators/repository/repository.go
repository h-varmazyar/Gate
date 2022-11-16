package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
)

type IndicatorRepository interface {
	Create(indicator *entity.Indicator) error
	Return(indicatorID uuid.UUID) (*entity.Indicator, error)
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (IndicatorRepository, error) {
	if err := db.PostgresDB.AutoMigrate(new(entity.Indicator)); err != nil {
		return nil, err
	}
	return NewIndicatorPostgresRepository(ctx, logger, db.PostgresDB)
}
