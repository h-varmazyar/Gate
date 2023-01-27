package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/db"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
)

type StrategyRepository interface {
	Save(strategy *entity.Strategy) error
	Return(strategyID uuid.UUID) (*entity.Strategy, error)
	ReturnActives(ctx context.Context) ([]*entity.Strategy, error)
	ReturnIndicators(strategyID uuid.UUID) ([]*entity.StrategyIndicator, error)
	List() ([]*entity.Strategy, error)
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (StrategyRepository, error) {
	if err := db.PostgresDB.AutoMigrate(new(entity.Strategy)); err != nil {
		return nil, err
	}
	if err := db.PostgresDB.AutoMigrate(new(entity.StrategyIndicator)); err != nil {
		return nil, err
	}
	return NewStrategyPostgresRepository(ctx, logger, db.PostgresDB)
}
