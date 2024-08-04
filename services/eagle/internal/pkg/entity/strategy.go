package entity

import (
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api/proto"
	"github.com/lib/pq"
)

type Strategy struct {
	gormext.UniversalModel
	Name                  string
	Description           string
	MinDailyProfitRate    float64
	MinProfitPerTradeRate float64
	MaxFundPerTrade       float64
	MaxFundPerTradeRate   float64
	WorkingResolutionID   uuid.UUID
	Indicators            []*StrategyIndicator `->;gorm:"foreignkey:StrategyID;references:ID"`
	IsActive              bool
	MarketIDs             pq.StringArray `gorm:"type:varchar(20000)[]"`
	Type                  eagleApi.StrategyType
	BrokerageID           uuid.UUID
	WithTrading           bool
}

type StrategyIndicator struct {
	StrategyID  uuid.UUID `gorm:"primary_key;type:uuid REFERENCES strategies(id)"`
	IndicatorID uuid.UUID `gorm:"primary_key;type:uuid"`
	//Type        chipmunkApi.IndicatorType `gorm:"type:varchar(25);not null"`
}
