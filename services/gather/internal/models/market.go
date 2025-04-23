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
	Name                        string
	Status                      MarketStatus
	BaseCurrencyPrecision       uint8
	QuoteCurrencyPrecision      uint8
	QuoteCurrencyID             uint
	QuoteCurrency               *Asset `gorm:"foreignKey:QuoteCurrencyID"`
	BaseCurrencyID              uint
	BaseCurrency                *Asset `gorm:"foreignKey:BaseCurrencyID"`
	TakerFeeRate                float64
	MakerFeeRate                float64
	MinAmount                   float64
	IsAmmAvailable              bool
	IsMarginAvailable           bool
	IsPremarketTradingAvailable bool
	IssueDate                   time.Time
}

func (o Market) Table() string {
	return MarketTableName
}
