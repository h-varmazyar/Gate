package app

import (
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/app/handlers"
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/pkg/tgBotApi"
)

type Configs struct {
	HandlerConfigs *handlers.Configs `yaml:"handler_configs"`
	BotConfigs     *tgBotApi.Configs `yaml:"bot_configs"`
}
