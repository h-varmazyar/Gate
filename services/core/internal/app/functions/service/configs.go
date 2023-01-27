package service

import (
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages/coinex"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages/nobitex"
)

type Configs struct {
	NetworkGrpcAddress string           `yaml:"network_grpc_address"`
	Coinex             *coinex.Configs  `yaml:"coinex"`
	Nobitex            *nobitex.Configs `yaml:"nobitex"`
}
