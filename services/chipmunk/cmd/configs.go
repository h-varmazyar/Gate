package main

import (
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/candles"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/markets"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/posts"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/resolutions"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/wallets"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
)

type Configs struct {
	ServiceName    string              `mapstructure:"service_name"`
	Version        string              `mapstructure:"version"`
	GRPCPort       uint16              `mapstructure:"grpc_port"`
	AMQPConfigs    *amqpext.Configs    `mapstructure:"amqp_configs"`
	BufferConfigs  *buffer.Configs     `mapstructure:"buffer_configs"`
	MarketsApp     markets.Configs     `mapstructure:"markets_app"`
	CandlesApp     candles.Configs     `mapstructure:"candles_app"`
	PostsApp       posts.Configs       `mapstructure:"posts_app"`
	ResolutionsApp resolutions.Configs `mapstructure:"resolutions_app"`
	WalletsApp     wallets.Configs     `mapstructure:"wallets_app"`
	DB             gormext.Configs     `mapstructure:"db"`
}
