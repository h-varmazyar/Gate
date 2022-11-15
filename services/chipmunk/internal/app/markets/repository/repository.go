package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
)

type MarketRepository interface {
	Info(brokerageID uuid.UUID, marketName string) (*entity.Market, error)
	List(brokerageID uuid.UUID) ([]*entity.Market, error)
	ListBySource(brokerageID uuid.UUID, source string) ([]*entity.Market, error)
	ReturnByID(id uuid.UUID) (*entity.Market, error)
	SaveOrUpdate(market *entity.Market) error
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (MarketRepository, error) {
	if err := db.PostgresDB.AutoMigrate(new(entity.Market)); err != nil {
		return nil, err
	}
	return NewMarketPostgresRepository(ctx, logger, db.PostgresDB)
}
