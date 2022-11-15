package main

import (
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/services/core/internal/app/brokerages"
	"github.com/h-varmazyar/Gate/services/core/internal/app/functions"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages/coinex"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/db"
)

type Configs struct {
	ServiceName   string              `yaml:"service_name"`
	Version       string              `yaml:"version"`
	GRPCPort      uint16              `yaml:"grpc_port"`
	AMQPConfigs   *amqpext.Configs    `yaml:"amqp_configs"`
	CoinexConfigs *coinex.Configs     `yaml:"coinex_configs"`
	BrokeragesApp *brokerages.Configs `yaml:"brokerages_app"`
	FunctionsApp  *functions.Configs  `yaml:"functions_app"`
	DB            *db.Configs         `yaml:"db"`
}
