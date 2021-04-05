package models

import "gorm.io/gorm"

type Candle struct {
	gorm.Model
	Open  float64
	High  float64
	Low   float64
	Close float64
	Vol   float64
	Time  float64
	Indicators
}
