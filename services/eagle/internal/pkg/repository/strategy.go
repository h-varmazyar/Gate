package repository

import (
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"gorm.io/gorm"
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
	Indicators            []*StrategyIndicator `gorm:"->;foreignkey:StrategyID;references:ID"`
}

type StrategyIndicator struct {
	StrategyID  uuid.UUID                 `gorm:"primary_key;type:uuid REFERENCES strategies(id)"`
	IndicatorID uuid.UUID                 `gorm:"primary_key;type:uuid"`
	Type        chipmunkApi.IndicatorType `gorm:"type:varchar(25);not null"`
}

type StrategyRepository struct {
	db *gorm.DB
}

func (r *StrategyRepository) Save(strategy *Strategy) error {
	return r.db.Save(strategy).Error
}

func (r *StrategyRepository) Return(strategyID uuid.UUID) (*Strategy, error) {
	strategy := new(Strategy)
	return strategy, r.db.Model(new(Strategy)).Preload("Indicators").Where("id = ?", strategyID).Error
}

func (r *StrategyRepository) ReturnIndicators(strategyID uuid.UUID) ([]*StrategyIndicator, error) {
	strategyIndicators := make([]*StrategyIndicator, 0)
	return strategyIndicators, r.db.Model(new(StrategyIndicator)).Where("strategy_id = ?", strategyID).Find(&strategyIndicators).Error
}

func (r *StrategyRepository) List() ([]*Strategy, error) {
	strategies := make([]*Strategy, 0)
	return strategies, r.db.Model(new(Strategy)).Find(&strategies).Error
}
