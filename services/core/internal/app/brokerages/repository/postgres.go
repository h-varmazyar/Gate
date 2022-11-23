package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type brokeragePostgresRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewBrokeragePostgresRepository(ctx context.Context, logger *log.Logger, db *gorm.DB) (*brokeragePostgresRepository, error) {
	if db == nil {
		return nil, errors.New(ctx, codes.Internal).AddDetailF("invalid db instance")
	}
	return &brokeragePostgresRepository{
		db:     db,
		logger: logger,
	}, nil
}

func (repository *brokeragePostgresRepository) Create(brokerage *entity.Brokerage) error {
	return repository.db.Model(new(entity.Brokerage)).Create(brokerage).Error
}

func (repository *brokeragePostgresRepository) Delete(id uuid.UUID) error {
	return repository.db.Model(new(entity.Brokerage)).Where("id LIKE ?", id).Delete(new(entity.Brokerage)).Error
}

func (repository *brokeragePostgresRepository) ReturnByID(id uuid.UUID) (*entity.Brokerage, error) {
	brokerage := new(entity.Brokerage)
	return brokerage, repository.db.Where("id = ?", id).First(brokerage).Error
}

func (repository *brokeragePostgresRepository) List() ([]*entity.Brokerage, error) {
	brokerages := make([]*entity.Brokerage, 0)
	return brokerages, repository.db.Find(&brokerages).Error
}

func (repository *brokeragePostgresRepository) ChangeStatus(brokerageID uuid.UUID) error {
	status := ""
	if err := repository.db.Model(new(entity.Brokerage)).Where("id = ?", brokerageID).Select("status").Scan(&status).Error; err != nil {
		return err
	}
	newStatus := proto.Status_Enable
	if status == proto.Status_Enable.String() {
		newStatus = proto.Status_Disable
	}
	return repository.db.Model(new(entity.Brokerage)).Where("id = ?", brokerageID).Update("status", newStatus).Error
}
