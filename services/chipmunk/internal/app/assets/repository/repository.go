package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
)

type AssetRepository interface {
	ReturnByID(id uuid.UUID) (*entity.Asset, error)
	ReturnBySymbol(symbol string) (*entity.Asset, error)
	Create(asset *entity.Asset) error
	List(page int) ([]*entity.Asset, error)
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (AssetRepository, error) {
	if err := db.PostgresDB.AutoMigrate(new(entity.Asset)); err != nil {
		return nil, err
	}
	return NewAssetPostgresRepository(ctx, logger, db.PostgresDB)
}
