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

type IndicatorValues struct {
	BollingerBands map[uuid.UUID]*BollingerBandsValue
	MovingAverages map[uuid.UUID]*MovingAverageValue
	Stochastics    map[uuid.UUID]*StochasticValue
	RSIs           map[uuid.UUID]*RSIValue
}

func NewIndicatorValues() IndicatorValues {
	return IndicatorValues{
		BollingerBands: make(map[uuid.UUID]*BollingerBandsValue),
		MovingAverages: make(map[uuid.UUID]*MovingAverageValue),
		Stochastics:    make(map[uuid.UUID]*StochasticValue),
		RSIs:           make(map[uuid.UUID]*RSIValue),
	}
}

func (r *candlePostgresRepository) Save(candle *entity.Candle) error {
	item := new(entity.Candle)
	err := r.db.Model(new(entity.Candle)).
		Where("time = ?", candle.Time).
		Where("market_id = ?", candle.MarketID).
		Where("resolution_id = ?", candle.ResolutionID).First(item).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return r.db.Model(new(entity.Candle)).Create(candle).Error
		}
		return err
	}
	candle.ID = item.ID
	return r.db.Save(candle).Error
}

func (r *candlePostgresRepository) BulkInsert(candles []*entity.Candle) error {
	return r.db.CreateInBatches(candles, len(candles)).Error
}

func (r *candlePostgresRepository) ReturnLast(marketID, resolutionID uuid.UUID) (*entity.Candle, error) {
	item := new(entity.Candle)
	return item, r.db.Model(new(entity.Candle)).
		Where("market_id = ?", marketID).
		Where("resolution_id = ?", resolutionID).
		Order("time desc").First(item).Error
}

func (r *candlePostgresRepository) ReturnList(marketID, resolutionID uuid.UUID, limit, offset int) ([]*entity.Candle, error) {
	items := make([]*entity.Candle, 0)
	return items, r.db.Model(new(entity.Candle)).
		Where("market_id = ?", marketID).
		Where("resolution_id = ?", resolutionID).
		Order("time desc").Offset(offset).Limit(limit).Find(&items).Error
}
