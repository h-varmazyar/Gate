package models

import (
	"time"
)

type Currency string

type Candle struct {
	ID              uint32 `gorm:"primarykey"`
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
	//FromDb          bool `gorm:"-"`
	Indicators `gorm:"-"`
}

func (c *Candle) LoadLast() error {
	return db.Model(&Candle{}).
		Where("market_refer = ?", c.MarketRefer).
		Where("resolution_refer = ?", c.ResolutionRefer).
		Order("time ASC").
		Last(&c).Error
}

func LoadCandleList(marketId uint16, resolutionId uint, lastTime time.Time, limit int) ([]Candle, error) {
	var candles []Candle
	return candles, db.Model(&Candle{}).
		Where("market_refer = ?", marketId).
		Where("resolution_refer = ?", resolutionId).
		Where("time > ?", lastTime).
		Order("time ASC").
		Limit(limit).Find(&candles).Error
}

func (c *Candle) Load() error {
	return db.Model(&Candle{}).
		Where("ID = ?", c.ID).
		Last(&c).Error
}

func (c *Candle) CreateOrUpdate() error {
	count := int64(0)
	db.Model(&Candle{}).Where("id = ?", c.ID).Count(&count)
	if count == 0 {
		return db.Model(&Candle{}).Create(c).Error
	}
	return db.Save(c).Error
}

func (c *Candle) Update() error {
	return db.Updates(&c).Error
}
