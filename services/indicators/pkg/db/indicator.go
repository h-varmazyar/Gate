package db

import (
	"context"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/entities"
)

type Indicators struct {
	db *DB
}

func (r Indicators) Create(ctx context.Context, indicator *entities.Indicator) error {
	return nil
}
