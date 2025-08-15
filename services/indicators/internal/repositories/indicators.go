package repository

import (
	"github.com/h-varmazyar/Gate/services/indicators/pkg/entities"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type IndicatorRepository struct {
	db *gorm.DB
}

func NewIndicatorRepository(db *gorm.DB) *IndicatorRepository {
	return &IndicatorRepository{db: db}
}

func (r *IndicatorRepository) All(ctx context.Context) ([]*entities.Indicator, error) {
	indicators := make([]*entities.Indicator, 0)
	return indicators, r.db.WithContext(ctx).Find(&indicators).Error
}
