package main

import (
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/strategies"
)

type Configs struct {
	ServiceName   string              `mapstructure:"service_name"`
	Version       string              `mapstructure:"version"`
	GRPCPort      uint16              `mapstructure:"grpc_port"`
	StrategiesApp *strategies.Configs `mapstructure:"strategies_app"`
	DB            gormext.Configs     `mapstructure:"db"`
}
