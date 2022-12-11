package repository

import (
	"context"
	"github.com/google/uuid"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
)

type MarketRepository interface {
	Info(Platform api.Platform, marketName string) (*entity.Market, error)
	List(Platform api.Platform) ([]*entity.Market, error)
	ListBySource(Platform api.Platform, source string) ([]*entity.Market, error)
	ReturnByID(id uuid.UUID) (*entity.Market, error)
	ReturnByName(name string) (*entity.Market, error)
	SaveOrUpdate(market *entity.Market) error
	Delete(market *entity.Market) error
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (MarketRepository, error) {
	if err := db.PostgresDB.AutoMigrate(new(entity.Market)); err != nil {
		return nil, err
	}
	return NewMarketPostgresRepository(ctx, logger, db.PostgresDB)
}
