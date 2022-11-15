package main

import (
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/assets"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
)

type Configs struct {
	ServiceName string           `yaml:"service_name"`
	Version     string           `yaml:"version"`
	GRPCPort    uint16           `yaml:"grpc_port"`
	AMQPConfigs *amqpext.Configs `yaml:"amqp_configs"`
	MarketsApp  *markets.Configs `yaml:"markets_app"`
	AssetsApp   *assets.Configs  `yaml:"assets_app"`
	DB          *db.Configs      `yaml:"db"`
}
