package entity

import (
	"github.com/google/uuid"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	"gorm.io/gorm"
)

type Indicator struct {
	gorm.Model
	Type         indicatorsAPI.Type `gorm:"type:varchar(64);not null"`
	MarketId     uuid.UUID
	ResolutionId uuid.UUID
	Configs      *IndicatorConfigs `gorm:"embedded;embeddedPrefix:configs_"`
}

type IndicatorConfigs struct {
	RSI            *RsiConfigs            `gorm:"embedded;embeddedPrefix:rsi_"`
	Stochastic     *StochasticConfigs     `gorm:"embedded;embeddedPrefix:stochastic_"`
	SMA            *SMAConfigs            `gorm:"embedded;embeddedPrefix:sma_"`
	BollingerBands *BollingerBandsConfigs `gorm:"embedded;embeddedPrefix:bollinger_bands_"`
}

type RsiConfigs struct {
	Period int
	Source indicatorsAPI.Source
}

type StochasticConfigs struct {
	Period  int
	SmoothK int
	SmoothD int
}

type SMAConfigs struct {
	Period int
	Source indicatorsAPI.Source
}

type BollingerBandsConfigs struct {
	Period    int
	Deviation int
	Source    indicatorsAPI.Source
}
