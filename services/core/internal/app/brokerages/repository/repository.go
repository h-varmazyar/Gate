package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/db"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
)

type BrokerageRepository interface {
	Create(brokerage *entity.Brokerage) error
	Delete(id uuid.UUID) error
	ReturnByID(id uuid.UUID) (*entity.Brokerage, error)
	List() ([]*entity.Brokerage, error)
	ChangeStatus(brokerageID uuid.UUID) error
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (BrokerageRepository, error) {
	if err := db.PostgresDB.AutoMigrate(new(entity.Brokerage)); err != nil {
		return nil, err
	}
	return NewBrokeragePostgresRepository(ctx, logger, db.PostgresDB)
}
