package main

import (
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/gormext"
)

type Configs struct {
	ServiceName string           `mapstructure:"service_name"`
	Version     string           `mapstructure:"version"`
	GRPCPort    uint16           `mapstructure:"grpc_port"`
	AMQPConfigs *amqpext.Configs `mapstructure:"amqp_configs"`
	DB          gormext.Configs  `mapstructure:"db"`
}
