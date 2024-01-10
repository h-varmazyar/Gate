package repository

import (
	"context"
	"github.com/h-varmazyar/Gate/services/indicators/pkg/entity"
)

type Repository interface {
	Create(ctx context.Context) (*entity.Indicator, error)
}
