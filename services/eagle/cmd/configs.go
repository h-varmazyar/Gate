package main

import (
	signals "github.com/h-varmazyar/Gate/services/eagle/internal/app/signals"
	strategies "github.com/h-varmazyar/Gate/services/eagle/internal/app/strategies"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/db"
)

type Configs struct {
	ServiceName   string              `yaml:"service_name"`
	Version       string              `yaml:"version"`
	GRPCPort      uint16              `yaml:"grpc_port"`
	StrategiesApp *strategies.Configs `yaml:"strategies_app"`
	SignalsApp    *signals.Configs    `yaml:"signals_app"`
	DB            *db.Configs         `yaml:"db"`
}
