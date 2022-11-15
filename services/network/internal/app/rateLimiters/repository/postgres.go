package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/errors"
	networkApi "github.com/h-varmazyar/Gate/services/network/api/proto"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type RateLimiterPostgres struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewRateLimiterPostgresRepository(ctx context.Context, logger *log.Logger, db *gorm.DB) (*RateLimiterPostgres, error) {
	if db == nil {
		return nil, errors.New(ctx, codes.Internal).AddDetailF("invalid db instance")
	}
	return &RateLimiterPostgres{
		db:     db,
		logger: logger,
	}, nil
}

func (repository *RateLimiterPostgres) Create(rateLimiter *entity.RateLimiter) error {
	return repository.db.Model(new(entity.RateLimiter)).Create(rateLimiter).Error
}

func (repository *RateLimiterPostgres) Return(id uuid.UUID) (*entity.RateLimiter, error) {
	rateLimiter := new(entity.RateLimiter)
	return rateLimiter, repository.db.Model(new(entity.RateLimiter)).Where("id = ?", id).First(rateLimiter).Error
}
func (repository *RateLimiterPostgres) List(Type networkApi.RateLimiterType) ([]*entity.RateLimiter, error) {
	rateLimiters := make([]*entity.RateLimiter, 0)
	db := repository.db.Model(new(entity.RateLimiter))
	if Type != networkApi.RateLimiter_All {
		db.Where("type = ?", Type)
	}

	return rateLimiters, db.Find(&rateLimiters).Error
}
