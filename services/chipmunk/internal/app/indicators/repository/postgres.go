package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/errors"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type indicatorPostgresRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewIndicatorPostgresRepository(ctx context.Context, logger *log.Logger, db *gorm.DB) (IndicatorRepository, error) {
	if db == nil {
		return nil, errors.New(ctx, codes.Internal).AddDetailF("invalid db instance")
	}
	return &indicatorPostgresRepository{
		db:     db,
		logger: logger,
	}, nil
}

func (r *indicatorPostgresRepository) Create(indicator *entity.Indicator) error {
	return r.db.Save(indicator).Error
}

func (r *indicatorPostgresRepository) Return(indicatorID uuid.UUID) (*entity.Indicator, error) {
	indicator := new(entity.Indicator)
	return indicator, r.db.Model(new(entity.Indicator)).Where("id = ?", indicatorID).First(indicator).Error
}

func (r *indicatorPostgresRepository) List(_ context.Context, indicatorType chipmunkApi.IndicatorType) ([]*entity.Indicator, error) {
	indicators := make([]*entity.Indicator, 0)
	tx := r.db.Model(new(entity.Indicator))
	if indicatorType != chipmunkApi.Indicator_All {
		tx.Where("type = ?", indicatorType)
	}
	err := tx.Find(&indicators).Error
	if err != nil {
		r.logger.WithError(err).Error("failed to load indicator list")
		return nil, err
	}

	return indicators, nil
}
