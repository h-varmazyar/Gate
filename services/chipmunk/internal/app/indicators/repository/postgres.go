package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/errors"
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
