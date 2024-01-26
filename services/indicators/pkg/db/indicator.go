package db

import (
	"context"
	"github.com/h-varmazyar/Gate/services/indicators/internal/repository"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/entity"
)

type Indicators struct {
	db *DB
}

func NewIndicators(ctx context.Context, db *DB) repository.Repository {

}

func (r Indicators) Create(ctx context.Context) (*entity.Indicator, error) {
	return nil, nil
}
