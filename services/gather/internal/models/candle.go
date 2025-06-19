package models

import (
	"gorm.io/gorm"
	"time"
)

const CandleTableName = "candles"

type Candle struct {
	gorm.Model
	Time         time.Time
	Open         float64
	High         float64
	Low          float64
	Close        float64
	Volume       float64
	Amount       float64
	MarketID     uint        `gorm:"index"`
	Market       *Market     `gorm:"foreignKey:MarketID"`
	ResolutionID uint        `gorm:"index"`
	Resolution   *Resolution `gorm:"foreignKey:ResolutionID"`
}

func (o Candle) Table() string {
	return CandleTableName
}
