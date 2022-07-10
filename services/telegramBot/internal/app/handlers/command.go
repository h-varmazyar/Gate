package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/pkg/repository"
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/pkg/tgBotApi"
)

const (
	CommandStart = "start"
)

const (
	StartMsg = `welcome to Gate Bot. Gate will send you every signals that catch by robot`
)

func (*Handler) startCommand(ctx context.Context, msg *tgbotapi.Message) error {
	err := tgBotApi.SendMessage(ctx, tgBotApi.NewTextMessage(msg.Chat.ID, msg.MessageID, StartMsg, nil))
	if err != nil {
		return errors.Cast(ctx, err).AddDetailF("failed to send start message")
	}

	c := &repository.Client{
		TelegramAccountID: msg.From.ID,
	}
	if err = repository.Clients.Create(c); err != nil {
		return errors.Cast(ctx, err).AddDetailF("failed to create client")
	}

	return nil
}
