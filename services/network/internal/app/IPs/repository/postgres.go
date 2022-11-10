package repository

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type ipPostgresRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewIPPostgresRepository(ctx context.Context, logger *log.Logger, db *gorm.DB) (*ipPostgresRepository, error) {
	if db == nil {
		return nil, errors.New(ctx, codes.Internal).AddDetailF("invalid db instance")
	}
	return &ipPostgresRepository{
		db:     db,
		logger: logger,
	}, nil
}

func (repository *ipPostgresRepository) Create(ip *entity.IP) error {
	return repository.db.Model(new(entity.IP)).Create(ip).Error
}
