package entity

import (
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
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
}

type StrategyIndicator struct {
	StrategyID  uuid.UUID                 `gorm:"primary_key;type:uuid REFERENCES strategies(id)"`
	IndicatorID uuid.UUID                 `gorm:"primary_key;type:uuid"`
	Type        chipmunkApi.IndicatorType `gorm:"type:varchar(25);not null"`
}
