package repository

import (
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	"gorm.io/gorm"
)

type Indicator struct {
	gorm.Model
	StrategyRefer uint
	Name          brokerageApi.IndicatorNames
	Description   string
	Configs       []byte
}

type IndicatorRepository struct {
	db *gorm.DB
}

func (r *IndicatorRepository) Save(indicator *Indicator) error {
	return r.db.Save(indicator).Error
}
