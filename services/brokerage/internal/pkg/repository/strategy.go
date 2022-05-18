package repository

import (
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"gorm.io/gorm"
)

type Strategy struct {
	gormext.UniversalModel
	Name                     string
	Description              string
	MinDailyPercentageProfit int
	MinProfitPercentage      int
	MaxFund                  float64
	MaxFundPercentage        int
	Indicators               []*StrategyIndicator `gorm:"->;foreignkey:StrategyID;references:ID"`
}

type StrategyIndicator struct {
	StrategyID  uuid.UUID `gorm:"primary_key;type:uuid REFERENCES strategies(id)"`
	IndicatorID uuid.UUID `gorm:"primary_key;type:uuid"`
	Type        chipmunkApi.IndicatorType
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

func (r *StrategyRepository) List() ([]*Strategy, error) {
	strategies := make([]*Strategy, 0)
	return strategies, r.db.Model(new(Strategy)).Find(&strategies).Error
}
