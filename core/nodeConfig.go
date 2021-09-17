package core

import (
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/strategies"
)

type NodeConfig struct {
	LoadMarketsOnline bool `yaml:"loadMarketsOnline"`
	Markets           []models.Market
	Resolutions       []models.Resolution
	Strategy          strategies.Strategy
	FakeTrading       bool `yaml:"fakeTrading"`
	EnableTrading     bool `yaml:"EnableTrading"`
}
