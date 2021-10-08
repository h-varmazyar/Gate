package core

import (
	"github.com/mrNobody95/Gate/indicators"
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/strategies"
)

type NodeConfig struct {
	Markets           []models.Market
	Strategy          strategies.Strategy
	Resolutions       []models.Resolution
	FakeTrading       bool                     `yaml:"fakeTrading"`
	EnableTrading     bool                     `yaml:"EnableTrading"`
	IndicatorConfigs  indicators.Configuration `yaml:"indicatorConfigs"`
	LoadMarketsOnline bool                     `yaml:"loadMarketsOnline"`
}
