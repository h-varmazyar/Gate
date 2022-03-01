package repository

import (
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"gorm.io/gorm"
)

type Strategy struct {
	gorm.Model
	Indicators struct {
		Name    brokerageApi.IndicatorNames
		Configs []byte
	}
}

type StrategyRepository struct {
	db *gorm.DB
}

func (r *StrategyRepository) Create(strategy *Strategy) error {
	return r.db.Model(new(Strategy)).Create(strategy).Error
}

func (r *StrategyRepository) Return(strategyID uint32) (*Strategy, error) {
	strategy := new(Strategy)
	return strategy, r.db.Model(new(Strategy)).Where("id = ?", strategyID).Error
}
