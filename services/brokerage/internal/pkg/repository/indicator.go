package repository

import (
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"gorm.io/gorm"
)

type Indicator struct {
	gorm.Model
	StrategyRefer uint
	Name          brokerageApi.IndicatorNames
	Configs       []byte
}
