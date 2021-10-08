package models

import (
	"gorm.io/gorm"
	"time"
)

type Currency string

type Candle struct {
	ID              uint64 `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Time            time.Time
	Vol             float64
	Low             float64
	Open            float64
	High            float64
	Close           float64
	Market          *Market     `gorm:"foreignKey:MarketRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Resolution      *Resolution `gorm:"foreignKey:ResolutionRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MarketRefer     uint16
	ResolutionRefer uint
	Indicators      `gorm:"-"`
}

func (c *Candle) LoadLast() error {
	return db.Model(&Candle{}).
		Where("market_refer = ?", c.Market.Id).
		Where("resolution_refer = ?", c.Resolution.Id).
		Order("time ASC").
		Last(&c).Error
}

func LoadCandleList(marketId uint16, resolutionId uint, time int64) ([]Candle, error) {
	var candles []Candle
	return candles, db.Model(&Candle{}).
		Where("market_refer = ?", marketId).
		Where("resolution_refer = ?", resolutionId).
		Where("time > ?", time).
		Order("time ASC").
		Limit(500).Find(&candles).Error
}

func (c *Candle) Load() error {
	return db.Model(&Candle{}).
		Where("ID = ?", c.ID).
		Last(&c).Error
}

func (c *Candle) CreateOrUpdate() error {
	found := new(Candle)
	err := db.Model(&Candle{}).
		Where("market_refer = ?", c.Market.Id).
		Where("resolution_refer = ?", c.Resolution.Id).
		Where("time = ?", c.Time).
		First(found).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return db.Model(&Candle{}).Create(&c).Error
		}
		return err
	}
	return db.Model(&Candle{}).Where("id = ?", found.ID).Updates(&c).Error
}

func (c *Candle) Update() error {
	return db.Updates(&c).Error
}
