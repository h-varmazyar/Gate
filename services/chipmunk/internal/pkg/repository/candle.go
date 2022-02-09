package repository

import (
	"gorm.io/gorm"
	"time"
)

type Candle struct {
	ID           uint64 `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Time         time.Time
	Open         float64
	High         float64
	Low          float64
	Close        float64
	Volume       float64
	Amount       float64
	MarketID     uint32
	ResolutionID uint32
}

//
//type Candles interface {
//	Save(*Candle) error
//	ReturnLast(marketID, resolutionID uint32) (*Candle, error)
//	ReturnList(marketID, resolutionID uint32, offset int) ([]*Candle, error)
//}

type CandleRepository struct {
	DB *gorm.DB
}

func (r *CandleRepository) Save(candle *Candle) error {
	item := new(Candle)
	err := r.DB.Model(new(Candle)).
		Where("time = ?", candle.Time).
		Where("market_id = ?", candle.MarketID).
		Where("resolution_id = ?", candle.ResolutionID).First(item).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return r.DB.Model(new(Candle)).Create(candle).Error
		}
		return err
	}
	candle.ID = item.ID
	return r.DB.Save(candle).Error
}

func (r *CandleRepository) ReturnLast(marketID, resolutionID uint32) (*Candle, error) {
	item := new(Candle)
	return item, r.DB.Model(new(Candle)).
		Where("market_id = ?", marketID).
		Where("resolution_id = ?", resolutionID).
		Order("time desc").First(item).Error
}

func (r *CandleRepository) ReturnList(marketID, resolutionID uint32, offset int) ([]*Candle, error) {
	items := make([]*Candle, 0)
	return items, r.DB.Model(new(Candle)).
		Where("market_id = ?", marketID).
		Where("resolution_id = ?", resolutionID).
		Order("time desc").Offset(offset).Limit(1000).Find(items).Error
}
