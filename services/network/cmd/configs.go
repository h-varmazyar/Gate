package main

import (
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/services/network/internal/app/IPs"
	"github.com/h-varmazyar/Gate/services/network/internal/app/rateLimiters"
	"github.com/h-varmazyar/Gate/services/network/internal/app/requests"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/db"
)

type Configs struct {
	ServiceName     string                `yaml:"service_name"`
	Version         string                `yaml:"version"`
	GRPCPort        uint16                `yaml:"grpc_port"`
	DB              *db.Configs           `yaml:"db"`
	AMQPConfigs     *amqpext.Configs      `yaml:"amqp_configs"`
	RequestsApp     *requests.Configs     `yaml:"requests_app"`
	IPsApp          *IPs.Configs          `yaml:"ips_app"`
	RateLimitersApp *rateLimiters.Configs `yaml:"rate_limiters_app"`
}
