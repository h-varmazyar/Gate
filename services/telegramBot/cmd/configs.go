package main

import (
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/app"
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/pkg/repository"
)

type Configs struct {
	ServiceName    string              `yaml:"service_name" env:"SERVICE_NAME,required"`
	Version        string              `yaml:"version" env:"VERSION,required"`
	GrpcPort       uint16              `yaml:"grpc_port" env:"GRPC_PORT,required"`
	DBConfigs      *repository.Configs `yaml:"db_configs"`
	ServiceConfigs *app.Configs        `yaml:"service_configs"`
}
