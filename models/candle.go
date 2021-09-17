package models

import (
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
	Symbol          Market     `gorm:"foreignKey:SymbolRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Resolution      Resolution `gorm:"foreignKey:ResolutionRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SymbolRefer     uint16
	ResolutionRefer uint
	Indicators      `gorm:"-"`
}

func (c *Candle) LoadLast() error {
	return db.Model(&Candle{}).
		Where("symbol_refer = ?", c.Symbol.Id).
		Where("resolution_refer = ?", c.Resolution.Id).
		Order("time ASC").
		Last(&c).Error
}

func (c *Candle) LoadList() ([]Candle, error) {
	var candles []Candle
	return candles, db.Model(&Candle{}).
		Where("symbol_refer = ?", c.Symbol.Id).
		Where("resolution_refer = ?", c.Resolution.Id).
		Where("time >= ?", c.Time).
		Order("time ASC").
		Limit(500).Find(&c).Error
}

func (c *Candle) Load() error {
	return db.Model(&Candle{}).
		Where("ID = ?", c.ID).
		Last(&c).Error
}

func (c *Candle) Create() error {
	return db.Create(&c).Error
}

func (c *Candle) Update() error {
	return db.Updates(&c).Error
}
