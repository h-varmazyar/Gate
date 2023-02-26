package main

import (
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/app"
)

type Configs struct {
	ServiceName    string          `mapstructure:"service_name"`
	Version        string          `mapstructure:"version"`
	GrpcPort       uint16          `mapstructure:"grpc_port"`
	ServiceConfigs *app.Configs    `mapstructure:"service_configs"`
	DB             gormext.Configs `mapstructure:"db"`
}
