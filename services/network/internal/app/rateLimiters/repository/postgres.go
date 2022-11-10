package repository

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type rateLimiterPostgresRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewRateLimiterPostgresRepository(ctx context.Context, logger *log.Logger, db *gorm.DB) (*rateLimiterPostgresRepository, error) {
	if db == nil {
		return nil, errors.New(ctx, codes.Internal).AddDetailF("invalid db instance")
	}
	return &rateLimiterPostgresRepository{
		db:     db,
		logger: logger,
	}, nil
}

func (repository *rateLimiterPostgresRepository) Create(rateLimiter *entity.RateLimiter) error {
	return repository.db.Model(new(entity.RateLimiter)).Create(rateLimiter).Error
}
