package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/pkg/tgBotApi"
	log "github.com/sirupsen/logrus"
)

const (
	ErrNotImplemented = "not implemented yet"
)

type Handler struct {
	*tgBotApi.UnImplementedHandler
}

func (h *Handler) HandleMessage(ctx context.Context, msg *tgbotapi.Message) error {
	var err error
	switch msg.Command() {
	case CommandStart:
		err = h.startCommand(ctx, msg)
	}
	log.WithError(err).Error("failed to handle command")
	return nil
}

func (h *Handler) HandleCommand(ctx context.Context, msg *tgbotapi.Message) error {
	log.Info(msg)
	return nil
}

func (h *Handler) HandleChannelPost(ctx context.Context, msg *tgbotapi.Message) error {
	log.Info(msg)
	return nil
}
