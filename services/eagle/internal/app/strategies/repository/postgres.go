package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type strategyPostgresRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewStrategyPostgresRepository(ctx context.Context, logger *log.Logger, db *gorm.DB) (StrategyRepository, error) {
	if db == nil {
		return nil, errors.New(ctx, codes.Internal).AddDetailF("invalid db instance")
	}
	return &strategyPostgresRepository{
		db:     db,
		logger: logger,
	}, nil
}

func (r *strategyPostgresRepository) Save(strategy *entity.Strategy) error {
	return r.db.Create(strategy).Error
}

func (r *strategyPostgresRepository) Return(strategyID uuid.UUID) (*entity.Strategy, error) {
	strategy := new(entity.Strategy)
	return strategy, r.db.Model(new(entity.Strategy)).Preload("Indicators").Where("id = ?", strategyID).Find(strategy).Error
}

func (r *strategyPostgresRepository) ReturnIndicators(strategyID uuid.UUID) ([]*entity.StrategyIndicator, error) {
	strategyIndicators := make([]*entity.StrategyIndicator, 0)
	return strategyIndicators, r.db.Model(new(entity.StrategyIndicator)).Where("strategy_id = ?", strategyID).Find(&strategyIndicators).Error
}

func (r *strategyPostgresRepository) List() ([]*entity.Strategy, error) {
	strategies := make([]*entity.Strategy, 0)
	return strategies, r.db.Model(new(entity.Strategy)).Find(&strategies).Error
}
