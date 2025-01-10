package models

import (
	"gorm.io/gorm"
	"time"
)

const CandleTableName = "candles"

//CREATE INDEX idx_candles_market_resolution_time ON public.candles (market_id, resolution_id, time DESC);

type Candle struct {
	gorm.Model
	Time         time.Time
	Open         float64
	High         float64
	Low          float64
	Close        float64
	Volume       float64
	Amount       float64
	MarketID     uint `gorm:"index"`
	ResolutionID uint `gorm:"index"`
}

func (o Candle) Table() string {
	return CandleTableName
}
