package repositories

import (
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type AssetRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewAssetRepository(logger *log.Logger, db *gorm.DB) AssetRepository {
	return AssetRepository{
		db:     db,
		logger: logger,
	}
}

func (r AssetRepository) ReturnBySymbol(ctx context.Context, symbol string) (*models.Asset, error) {
	asset := new(models.Asset)
	return asset, r.db.WithContext(ctx).Model(&models.Asset{}).Where("symbol = ?", symbol).First(asset).Error
}

func (r AssetRepository) Create(ctx context.Context, asset *models.Asset) (*models.Asset, error) {
	if asset.Symbol == "" {
		asset.Symbol = asset.Name
	}
	return asset, r.db.WithContext(ctx).Create(asset).Error
}

func (r AssetRepository) Delete(ctx context.Context, assetID uint) error {
	return r.db.WithContext(ctx).Delete(&models.Asset{}, assetID).Error
}
