package tgBotApi

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handlers interface {
	HandleMessage(ctx context.Context, msg *tgbotapi.Message) error
	HandleCommand(ctx context.Context, msg *tgbotapi.Message) error
	HandleChannelPost(ctx context.Context, msg *tgbotapi.Message) error
}
