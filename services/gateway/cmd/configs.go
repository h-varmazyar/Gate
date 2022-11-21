package main

import "github.com/h-varmazyar/Gate/services/gateway/internal/app/chipmunk"
import "github.com/h-varmazyar/Gate/services/gateway/internal/app/eagle"
import "github.com/h-varmazyar/Gate/services/gateway/internal/app/core"
import "github.com/h-varmazyar/Gate/services/gateway/internal/app/telegramBot"

type Configs struct {
	ServiceName       string               `yaml:"service_name"`
	Version           string               `yaml:"version"`
	HttpPort          uint16               `yaml:"http_port"`
	ChipmunkRouter    *chipmunk.Configs    `yaml:"chipmunk_router"`
	CoreRouter        *core.Configs        `yaml:"core_router"`
	EagleRouter       *eagle.Configs       `yaml:"eagle_router"`
	TelegramBotRouter *telegramBot.Configs `yaml:"telegram_bot_router"`
}
