package repositories

import (
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type CandleRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

const candleListLimit = 100000

func NewCandleRepository(logger *log.Logger, db *gorm.DB) CandleRepository {
	return CandleRepository{
		db:     db,
		logger: logger,
	}
}

func (r CandleRepository) AllMarketCandles(ctx context.Context, marketID uint, offset int) ([]models.Candle, error) {
	candles := make([]models.Candle, 0)
	return candles, r.db.
		WithContext(ctx).
		Model(&models.Candle{}).
		Where("market_id = ?", marketID).
		Limit(candleListLimit).
		Offset(offset).
		Find(&candles).Error
}

func (r CandleRepository) DeleteMarketCandles(ctx context.Context, marketID uint) error {
	return r.db.WithContext(ctx).Where("market_id = ?", marketID).Delete(&models.Candle{}).Error
}

func (r CandleRepository) ReturnLast(ctx context.Context, marketID, resolutionID uint) (models.Candle, error) {
	var candle models.Candle
	return candle, r.db.
		WithContext(ctx).
		Model(&models.Candle{}).
		Where("market_id = ?", marketID).
		Where("resolution_id = ?", resolutionID).
		Order("time DESC").
		First(&candle).Error
}

func (r CandleRepository) BulkInsert(ctx context.Context, candles []models.Candle) error {
	return r.db.WithContext(ctx).Model(&models.Candle{}).CreateInBatches(&candles, 1000).Error
}
