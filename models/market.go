package models

import (
	"gorm.io/gorm"
	"strings"
	"time"
)

type Market struct {
	gorm.Model
	Id               uint16    `gorm:"primarykey"`
	Name             string    `gorm:"size:50"`
	Brokerage        Brokerage `gorm:"foreignKey:BrokerageRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StartTime        time.Time
	StartTimeString  string `yaml:"startTime" gorm:"-"`
	BrokerageRefer   uint8
	TakerFeeRate     float64 `json:"taker_fee_rate"`
	MakerFeeRate     float64 `json:"maker_fee_rate"`
	PricingDecimal   int     `json:"pricing_decimal"`
	TradingDecimal   int     `json:"trading_decimal"`
	MinAmount        float64 `json:"min_amount"`
	IsAMM            bool
	Source           *Asset `gorm:"foreignKey:SourceRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Destination      *Asset `gorm:"foreignKey:DestinationRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SourceRefer      uint
	DestinationRefer uint
}

func (market *Market) CreateOrLoad() error {
	err := db.Model(&Market{}).
		Where("brokerage_refer = ?", market.BrokerageRefer).
		Where("name LIKE ?", strings.TrimSpace(market.Name)).
		First(market).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return db.Model(&Market{}).Create(market).Error
	}
	return err
}

func GetBrokerageMarkets(brokerageId uint8) ([]Market, error) {
	markets := make([]Market, 0)
	return markets, db.Model(&Market{}).Where("brokerage_refer = ?", brokerageId).Limit(10).Offset(20).Find(&markets).Error
}
