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
	Indicators            []*StrategyIndicator `->;gorm:"foreignkey:StrategyID;references:ID"`
}

// f93e0959-e55b-4afe-b233-8b30a244cdc8,RSI
// 7e5fd65a-9a02-43ae-81e2-e04b35f1e8b1,Stochastic
// 16e3591f-90f8-42d4-a38c-7b2d39a7de27,BollingerBands
// 513be79b-ee7b-4133-991a-05e7600c6a13,MovingAverage

// 10,1,2,2,ab28acd0-3517-483f-b3a1-7bd879fa85d0

type StrategyIndicator struct {
	StrategyID  uuid.UUID                 `gorm:"primary_key;type:uuid REFERENCES strategies(id)"`
	IndicatorID uuid.UUID                 `gorm:"primary_key;type:uuid"`
	Type        chipmunkApi.IndicatorType `gorm:"type:varchar(25);not null"`
}

type StrategyRepository struct {
	db *gorm.DB
}

func (r *StrategyRepository) Save(strategy *Strategy) error {
	return r.db.Create(strategy).Error
}

func (r *StrategyRepository) Return(strategyID uuid.UUID) (*Strategy, error) {
	strategy := new(Strategy)
	return strategy, r.db.Model(new(Strategy)).Preload("Indicators").Where("id = ?", strategyID).Find(strategy).Error
}

func (r *StrategyRepository) ReturnIndicators(strategyID uuid.UUID) ([]*StrategyIndicator, error) {
	strategyIndicators := make([]*StrategyIndicator, 0)
	return strategyIndicators, r.db.Model(new(StrategyIndicator)).Where("strategy_id = ?", strategyID).Find(&strategyIndicators).Error
}

func (r *StrategyRepository) List() ([]*Strategy, error) {
	strategies := make([]*Strategy, 0)
	return strategies, r.db.Model(new(Strategy)).Find(&strategies).Error
}
