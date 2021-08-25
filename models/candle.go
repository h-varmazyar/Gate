package models

import (
	"github.com/mrNobody95/Gate/models/todo"
	"github.com/mrNobody95/gorm"
	"time"
)

type Currency string

type Candle struct {
	ID              uint64 `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Time            time.Time
	Vol             float64
	Low             float64
	Open            float64
	High            float64
	Close           float64
	Symbol          Symbol     `gorm:"foreignKey:SymbolRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Resolution      Resolution `gorm:"foreignKey:ResolutionRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SymbolRefer     uint16
	ResolutionRefer uint
	todo.Indicators `gorm:"-"`
}

func (c *Candle) LoadLast() error {
	//return db.Model(&Candle{}).
	//	Preload("Resolutions").
	//	Where("brokerage LIKE ?", c.Brokerage).
	//	Where("symbol LIKE ?", c.Symbol).
	//	Where("value LIKE ?", c.Resolution.Value).
	//	Last(&c).Error
	return db.Model(&Candle{}).
		Where("symbol_refer = ?", c.Symbol.Id).
		Where("resolution_refer = ?", c.Resolution.Id).
		Order("time ASC").
		Last(&c).Error
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
