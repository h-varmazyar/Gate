package models

import (
	"gorm.io/gorm"
	"strings"
	"time"
)

type Market struct {
	gorm.Model
	Id              uint16    `gorm:"primarykey"`
	Name            string    `gorm:"size:50"`
	Brokerage       Brokerage `gorm:"foreignKey:BrokerageRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StartTime       time.Time
	StartTimeString string `yaml:"startTime" gorm:"-"`
	BrokerageRefer  uint8
	TakerFeeRate    float64 `json:"taker_fee_rate"`
	MakerFeeRate    float64 `json:"maker_fee_rate"`
	PricingDecimal  int     `json:"pricing_decimal"`
	TradingDecimal  int     `json:"trading_decimal"`
	PricingName     string  `json:"pricing_name"`
	TradingName     string  `json:"trading_name"`
	MinAmount       float64 `json:"min_amount"`
	IsAMM           bool
	Source          *Asset `json:"source"`
	Destination     *Asset `json:"destination"`
}

func (market *Market) CreateOrLoad() error {
	err := db.
		Where("brokerage_refer = ?", market.BrokerageRefer).
		Where("name LIKE ?", strings.TrimSpace(market.Name)).
		First(&market).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return db.Model(&Market{}).Create(&market).Error
	}
	return err
}
