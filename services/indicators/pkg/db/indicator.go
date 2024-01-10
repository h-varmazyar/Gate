package db

import (
	"context"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/entity"
	"gorm.io/gorm"
)

type Indicators struct {
	db *gorm.DB
}

func (r Indicators) Create(ctx context.Context) (*entity.Indicator, error) {
	return nil, nil
}
