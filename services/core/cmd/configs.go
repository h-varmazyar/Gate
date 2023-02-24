package main

import (
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"github.com/h-varmazyar/Gate/services/core/internal/app/brokerages"
	"github.com/h-varmazyar/Gate/services/core/internal/app/functions"
	"github.com/h-varmazyar/Gate/services/core/internal/app/platforms"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages/coinex"
)

type Configs struct {
	ServiceName   string              `mapstructure:"service_name"`
	Version       string              `mapstructure:"version"`
	GRPCPort      uint16              `mapstructure:"grpc_port"`
	AMQPConfigs   *amqpext.Configs    `mapstructure:"amqp_configs"`
	CoinexConfigs *coinex.Configs     `mapstructure:"coinex_configs"`
	BrokeragesApp *brokerages.Configs `mapstructure:"brokerages_app"`
	PlatformsApp  *platforms.Configs  `mapstructure:"platforms_app"`
	FunctionsApp  *functions.Configs  `mapstructure:"functions_app"`
	DB            gormext.Configs     `mapstructure:"db"`
}
