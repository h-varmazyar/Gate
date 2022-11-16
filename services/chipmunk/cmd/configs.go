package main

import (
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/assets"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/indicators"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/resolutions"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
)

type Configs struct {
	ServiceName    string               `yaml:"service_name"`
	Version        string               `yaml:"version"`
	GRPCPort       uint16               `yaml:"grpc_port"`
	AMQPConfigs    *amqpext.Configs     `yaml:"amqp_configs"`
	AssetsApp      *assets.Configs      `yaml:"assets_app"`
	MarketsApp     *markets.Configs     `yaml:"markets_app"`
	CandlesApp     *candles.Configs     `yaml:"candles_app"`
	IndicatorsApp  *indicators.Configs  `yaml:"indicators_app"`
	ResolutionsApp *resolutions.Configs `yaml:"resolutions_app"`
	DB             *db.Configs          `yaml:"db"`
}
