package repositories

import (
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type MarketRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewMarketRepository(logger *log.Logger, db *gorm.DB) MarketRepository {
	return MarketRepository{
		db:     db,
		logger: logger,
	}
}

func (r MarketRepository) All(ctx context.Context) ([]models.Market, error) {
	markets := make([]models.Market, 0)
	return markets, r.db.WithContext(ctx).Model(&models.Market{}).Find(&markets).Error
}

func (r MarketRepository) Delete(ctx context.Context, marketID uint) error {
	return r.db.WithContext(ctx).Delete(&models.Market{}, marketID).Error
}

func (r MarketRepository) MarketCount(ctx context.Context, assetID uint) (int64, error) {
	var count int64
	return count, r.db.
		WithContext(ctx).
		Model(&models.Market{}).
		Where("source_id = ?", assetID).
		Or("destination_id = ?", assetID).
		Count(&count).Error
}

func (r MarketRepository) Create(ctx context.Context, market models.Market) (models.Market, error) {
	return market, r.db.WithContext(ctx).Create(&market).Error
}
