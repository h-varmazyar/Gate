package candles

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

type Candles interface {
	Save(*Candle) error
	ReturnLast(marketID, resolutionID uint32) (*Candle, error)
	ReturnList(marketID, resolutionID uint32, offset int) ([]*Candle, error)
}
