package main

import "github.com/h-varmazyar/Gate/services/gateway/internal/app/chipmunk"
import "github.com/h-varmazyar/Gate/services/gateway/internal/app/eagle"
import "github.com/h-varmazyar/Gate/services/gateway/internal/app/core"
import "github.com/h-varmazyar/Gate/services/gateway/internal/app/telegramBot"

type Configs struct {
	ServiceName       string               `mapstructure:"service_name"`
	Version           string               `mapstructure:"version"`
	HttpPort          uint16               `mapstructure:"http_port"`
	ChipmunkRouter    *chipmunk.Configs    `mapstructure:"chipmunk_router"`
	CoreRouter        *core.Configs        `mapstructure:"core_router"`
	EagleRouter       *eagle.Configs       `mapstructure:"eagle_router"`
	TelegramBotRouter *telegramBot.Configs `mapstructure:"telegram_bot_router"`
}
