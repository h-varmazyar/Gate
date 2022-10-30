package main

import (
	"github.com/h-varmazyar/Gate/services/core/internal/app/brokerages"
	"github.com/h-varmazyar/Gate/services/core/internal/app/functions"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/db"
)

type Configs struct {
	ServiceName   string              `yaml:"service_name"`
	Version       string              `yaml:"version"`
	GRPCPort      uint16              `yaml:"grpc_port"`
	BrokeragesApp *brokerages.Configs `yaml:"brokerages_app"`
	FunctionsApp  *functions.Configs  `yaml:"functions_app"`
	DB            *db.Configs         `yaml:"db"`
}
