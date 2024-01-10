package main

import (
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"github.com/h-varmazyar/Gate/services/indicators/internal"
)

type Configs struct {
	ServiceName string           `mapstructure:"service_name"`
	Version     string           `mapstructure:"version"`
	GRPCPort    uint16           `mapstructure:"grpc_port"`
	AppConfigs  internal.Configs `mapstructure:"app_configs"`
	DB          gormext.Configs  `mapstructure:"db"`
}
