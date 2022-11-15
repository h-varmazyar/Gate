package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type IPPostgres struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewIPPostgresRepository(ctx context.Context, logger *log.Logger, db *gorm.DB) (*IPPostgres, error) {
	if db == nil {
		return nil, errors.New(ctx, codes.Internal).AddDetailF("invalid db instance")
	}
	return &IPPostgres{
		db:     db,
		logger: logger,
	}, nil
}

func (repository *IPPostgres) Create(ip *entity.IP) error {
	return repository.db.Model(new(entity.IP)).Create(ip).Error
}

func (repository *IPPostgres) Return(id uuid.UUID) (*entity.IP, error) {
	ip := new(entity.IP)
	return ip, repository.db.Model(new(entity.IP)).Where("id = ?", id).First(ip).Error
}
func (repository *IPPostgres) List() ([]*entity.IP, error) {
	ips := make([]*entity.IP, 0)
	return ips, repository.db.Model(new(entity.IP)).Find(&ips).Error
}
