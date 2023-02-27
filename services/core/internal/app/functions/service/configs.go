package service

import (
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages/coinex"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages/nobitex"
)

type Configs struct {
	NetworkGrpcAddress string           `mapstructure:"network_grpc_address"`
	Coinex             *coinex.Configs  `mapstructure:"coinex"`
	Nobitex            *nobitex.Configs `mapstructure:"nobitex"`
}
