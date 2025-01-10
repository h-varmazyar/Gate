package models

import (
	"gorm.io/gorm"
	"time"
)

type MarketStatus string

const (
	MarketStatusEnable  MarketStatus = "enable"
	MarketStatusDisable MarketStatus = "disable"
)

const MarketTableName = "markets"

type Market struct {
	gorm.Model
	PricingDecimal float64
	TradingDecimal float64
	TakerFeeRate   float64
	MakerFeeRate   float64
	DestinationID  uint
	Destination    *Asset `gorm:"foreignKey:DestinationID"`
	IssueDate      time.Time
	MinAmount      float64
	SourceID       uint
	Source         *Asset `gorm:"foreignKey:SourceID"`
	IsAMM          bool
	Name           string
	Status         MarketStatus
	//SourceName      string `gorm:"-"`
	//DestinationName string `gorm:"-"`
}

func (o Market) Table() string {
	return MarketTableName
}
