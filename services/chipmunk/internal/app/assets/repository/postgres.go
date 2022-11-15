package repository

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type assetPostgresRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewAssetPostgresRepository(ctx context.Context, logger *log.Logger, db *gorm.DB) (AssetRepository, error) {
	if db == nil {
		return nil, errors.New(ctx, codes.Internal).AddDetailF("invalid db instance")
	}
	return &assetPostgresRepository{
		db:     db,
		logger: logger,
	}, nil
}

func (repository *assetPostgresRepository) ReturnBySymbol(symbol string) (*entity.Asset, error) {
	asset := new(entity.Asset)
	return asset, repository.db.Model(&entity.Asset{}).Where("symbol = ?", symbol).First(asset).Error
}

func (repository *assetPostgresRepository) Create(asset *entity.Asset) error {
	return repository.db.Model(&entity.Asset{}).Create(asset).Error
}

func (repository *assetPostgresRepository) Set(asset *entity.Asset) (*entity.Asset, error) {
	return asset, repository.db.Model(&entity.Asset{}).Save(asset).Error
}

func (repository *assetPostgresRepository) List(page int) ([]*entity.Asset, error) {
	assets := make([]*entity.Asset, 0)
	return assets, repository.db.Model(&entity.Asset{}).Offset((page - 1) * 25).Limit(25).Find(&assets).Error
}
