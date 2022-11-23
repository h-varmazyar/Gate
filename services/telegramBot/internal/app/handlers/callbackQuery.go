package handlers

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/errors"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/pkg/tgBotApi"
	"google.golang.org/grpc/codes"
)

const (
	QueryStart  = "start"
	QueryStop   = "stop"
	QueryDelete = "delete"
)

func (h *Handler) startCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, data interface{}) error {
	brokerageID, err := uuid.Parse(data.(string))
	if err != nil {
		return err
	}
	var brokerage *coreApi.Brokerage
	if brokerage, err = h.brokerageService.Start(ctx, &coreApi.BrokerageStartReq{
		ID:          brokerageID.String(),
		WithTrading: true,
	}); err != nil {
		return err
	}
	startMsg := fmt.Sprintf("core %v started successfully", brokerage.Title)
	return tgBotApi.SendMessage(ctx, tgBotApi.NewTextMessage(callback.Message.Chat.ID, callback.Message.MessageID, startMsg, nil))
}

func (h *Handler) stopCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, data interface{}) error {
	brokerageID, err := uuid.Parse(data.(string))
	if err != nil {
		return err
	}
	var brokerage *coreApi.Brokerage
	if brokerage, err = h.brokerageService.Stop(ctx, &coreApi.BrokerageStopReq{ID: brokerageID.String()}); err != nil {
		return err
	}
	startMsg := fmt.Sprintf("core %v stopped successfully", brokerage.Title)
	return tgBotApi.SendMessage(ctx, tgBotApi.NewTextMessage(callback.Message.Chat.ID, callback.Message.MessageID, startMsg, nil))
}

func (h *Handler) deleteCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, data interface{}) error {
	return errors.New(ctx, codes.Unimplemented)
}
