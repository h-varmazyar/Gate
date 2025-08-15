package entities

import (
	"gorm.io/gorm"
)

type IndicatorSource string

const (
	IndicatorSourceOHLC4 = "OHLC4"
	IndicatorSourceClose = "CLOSE"
	IndicatorSourceOpen  = "OPEN"
	IndicatorSourceHigh  = "HIGH"
	IndicatorSourceHLC3  = "HLC3"
	IndicatorSourceLow   = "LOW"
	IndicatorSourceHL2   = "HL2"
)

type IndicatorType string

const (
	IndicatorTypeRSI            = "RSI"
	IndicatorTypeStochastic     = "STOCHASTIC"
	IndicatorTypeSMA            = "SMA"
	IndicatorTypeEMA            = "EMA"
	IndicatorTypeBollingerBands = "BOLLINGER_BANDS"
)

type Indicator struct {
	gorm.Model
	Type         IndicatorType `gorm:"type:varchar(16);not null"`
	MarketId     uint
	ResolutionId uint
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
	Source IndicatorSource
}

type StochasticConfigs struct {
	Period  int
	SmoothK int
	SmoothD int
}

type SMAConfigs struct {
	Period int
	Source IndicatorSource
}

type BollingerBandsConfigs struct {
	Period    int
	Deviation int
	Source    IndicatorSource
}
