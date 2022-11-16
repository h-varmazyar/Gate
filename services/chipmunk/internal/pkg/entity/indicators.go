package entity

import (
	"github.com/h-varmazyar/Gate/pkg/gormext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
)

type Indicator struct {
	gormext.UniversalModel
	Type    chipmunkApi.IndicatorType `gorm:"type:varchar(25);not null"`
	Configs *IndicatorConfigs         `gorm:"embedded;embeddedPrefix:configs_"`
}

type IndicatorConfigs struct {
	RSI            *RsiConfigs            `gorm:"embedded;embeddedPrefix:rsi_"`
	Stochastic     *StochasticConfigs     `gorm:"embedded;embeddedPrefix:stochastic_"`
	MovingAverage  *MovingAverageConfigs  `gorm:"embedded;embeddedPrefix:moving_average_"`
	BollingerBands *BollingerBandsConfigs `gorm:"embedded;embeddedPrefix:bollinger_bands_"`
}

type RsiConfigs struct {
	Length int
}

type StochasticConfigs struct {
	Length  int
	SmoothK int
	SmoothD int
}

type MovingAverageConfigs struct {
	Length int
	Source chipmunkApi.Source
}

type BollingerBandsConfigs struct {
	Length    int
	Deviation int
	Source    chipmunkApi.Source
}

type BollingerBandsValue struct {
	UpperBand float64
	LowerBand float64
	MA        float64
}

type MovingAverageValue struct {
	Simple      float64
	Exponential float64
}

type StochasticValue struct {
	IndexK float64
	IndexD float64
	FastK  float64
}

type RSIValue struct {
	Gain float64
	Loss float64
	RSI  float64
}

type IndicatorValue struct {
	BB         *BollingerBandsValue
	MA         *MovingAverageValue
	Stochastic *StochasticValue
	RSI        *RSIValue
}
