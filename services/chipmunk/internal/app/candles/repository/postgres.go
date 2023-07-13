package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type candlePostgresRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewCandlePostgresRepository(ctx context.Context, logger *log.Logger, db *gorm.DB) (CandleRepository, error) {
	if db == nil {
		return nil, errors.New(ctx, codes.Internal).AddDetailF("invalid db instance")
	}
	return &candlePostgresRepository{
		db:     db,
		logger: logger,
	}, nil
}

func (r *candlePostgresRepository) Save(candle *entity.Candle) error {
	item := new(entity.Candle)
	err := r.db.Model(new(entity.Candle)).
		Where("time = ?", candle.Time).
		Where("market_id = ?", candle.MarketID).
		Where("resolution_id = ?", candle.ResolutionID).First(item).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Infof("created: %v - %v", candle.MarketID, candle.ResolutionID)
			return r.db.Model(new(entity.Candle)).Create(candle).Error
		}
		return err
	}
	candle.ID = item.ID
	return r.db.Updates(candle).Error
}

func (r *candlePostgresRepository) HardDelete(candle *entity.Candle) error {
	err := r.db.Model(new(entity.Candle)).Where("id = ?", candle.ID).Unscoped().Delete(candle).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *candlePostgresRepository) BulkHardDelete(candleIDs []uuid.UUID) error {
	return r.db.Unscoped().Where("id in ?", candleIDs).Delete(new(entity.Candle)).Error
}

func (r *candlePostgresRepository) BulkInsert(candles []*entity.Candle) error {
	return r.db.CreateInBatches(candles, 1000).Error
}

func (r *candlePostgresRepository) ReturnLast(marketID, resolutionID uuid.UUID) (*entity.Candle, error) {
	item := new(entity.Candle)
	return item, r.db.Model(new(entity.Candle)).
		Where("market_id = ?", marketID).
		Where("resolution_id = ?", resolutionID).
		Order("time desc").First(item).Error
}

func (r *candlePostgresRepository) Count(marketID, resolutionID uuid.UUID) (int64, error) {
	count := int64(0)
	return count, r.db.Model(new(entity.Candle)).
		Where("market_id = ?", marketID).
		Where("resolution_id = ?", resolutionID).
		Count(&count).Error
}

func (r *candlePostgresRepository) List(marketID, resolutionID uuid.UUID, limit, offset int) ([]*entity.Candle, error) {
	items := make([]*entity.Candle, 0)
	return items, r.db.Model(new(entity.Candle)).
		Where("market_id = ?", marketID).
		Where("resolution_id = ?", resolutionID).
		Order("time asc").Offset(offset).Limit(limit).Find(&items).Error
}
