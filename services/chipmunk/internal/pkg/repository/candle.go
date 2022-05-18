package repository

import (
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"gorm.io/gorm"
	"time"
)

type Candle struct {
	gormext.UniversalModel
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Time            time.Time
	Open            float64
	High            float64
	Low             float64
	Close           float64
	Volume          float64
	Amount          float64
	MarketID        uuid.UUID
	ResolutionID    uuid.UUID
	IndicatorValues `gorm:"-"`
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

type CandleRepository struct {
	db *gorm.DB
}

func (r *CandleRepository) Save(candle *Candle) error {
	item := new(Candle)
	err := r.db.Model(new(Candle)).
		Where("time = ?", candle.Time).
		Where("market_id = ?", candle.MarketID).
		Where("resolution_id = ?", candle.ResolutionID).First(item).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return r.db.Model(new(Candle)).Create(candle).Error
		}
		return err
	}
	candle.ID = item.ID
	return r.db.Save(candle).Error
}

func (r *CandleRepository) ReturnLast(marketID, resolutionID uuid.UUID) (*Candle, error) {
	item := new(Candle)
	return item, r.db.Model(new(Candle)).
		Where("market_id = ?", marketID).
		Where("resolution_id = ?", resolutionID).
		Order("time desc").First(item).Error
}

func (r *CandleRepository) ReturnList(marketID, resolutionID uuid.UUID, limit, offset int) ([]*Candle, error) {
	items := make([]*Candle, 0)
	return items, r.db.Model(new(Candle)).
		Where("market_id = ?", marketID).
		Where("resolution_id = ?", resolutionID).
		Order("time desc").Offset(offset).Limit(limit).Find(&items).Error
}
