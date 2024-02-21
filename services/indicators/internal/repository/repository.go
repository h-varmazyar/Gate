package repository

import (
	"context"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/entities"
)

type Repository interface {
	Create(ctx context.Context, indicator *entities.Indicator) error
	List(ctx context.Context) ([]*entities.Indicator, error)
}
