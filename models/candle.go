package models

import (
	"github.com/mrNobody95/Gate/models/todo"
	"gorm.io/gorm"
)

type Currency string

type Candle struct {
	gorm.Model
	Low             float64
	Vol             float64
	Time            int64
	Open            float64
	High            float64
	Close           float64
	Symbol          Symbol     `gorm:"foreignKey:SymbolRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Brokerage       Brokerage  `gorm:"foreignKey:BrokerageRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Resolution      Resolution `gorm:"foreignKey:ResolutionRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SymbolRefer     uint16
	BrokerageRefer  uint8
	ResolutionRefer uint
	todo.Indicators `gorm:"-"`
}

func (c *Candle) LoadLast() error {
	return db.Model(&Candle{}).
		Preload("Resolutions").
		Where("brokerage LIKE ?", c.Brokerage).
		Where("symbol LIKE ?", c.Symbol).
		Where("value LIKE ?", c.Resolution.Value).
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
