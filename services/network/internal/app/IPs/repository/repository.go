package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/db"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
)

type IPRepository interface {
	Create(ip *entity.IP) error
	Return(id uuid.UUID) (*entity.IP, error)
	List() ([]*entity.IP, error)
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (IPRepository, error) {
	if err := db.PostgresDB.AutoMigrate(new(entity.IP)); err != nil {
		return nil, err
	}
	return NewIPPostgresRepository(ctx, logger, db.PostgresDB)
}