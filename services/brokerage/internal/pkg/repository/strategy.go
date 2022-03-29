package repository

import (
	"gorm.io/gorm"
)

type Strategy struct {
	gorm.Model
	Name              string
	Description       string
	MinDailyProfit    float64
	MinProfitPerTrade float64
	MaxFundPerTrade   float64
	Indicators        []Indicator `gorm:"foreignKey:StrategyRefer"`
}

type StrategyRepository struct {
	db *gorm.DB
}

func (r *StrategyRepository) Save(strategy *Strategy) error {
	return r.db.Save(strategy).Error
}

func (r *StrategyRepository) Return(strategyID uint32) (*Strategy, error) {
	strategy := new(Strategy)
	return strategy, r.db.Model(new(Strategy)).Where("id = ?", strategyID).Error
}

func (r *StrategyRepository) List() ([]*Strategy, error) {
	strategies := make([]*Strategy, 0)
	return strategies, r.db.Model(new(Strategy)).Find(&strategies).Error
}
