package service

import "github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages/coinex"

type Configs struct {
	NetworkGrpcAddress string          `yaml:"network_grpc_address"`
	Coinex             *coinex.Configs `yaml:"coinex"`
}
