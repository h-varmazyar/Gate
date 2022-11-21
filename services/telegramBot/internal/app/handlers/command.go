package handlers

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/pkg/repository"
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/pkg/tgBotApi"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

const (
	CommandStart     = "start"
	CmdBrokerageList = "list"
	CmdUpdateMarkets = "update"
)

const (
	StartMsg = `welcome to The Gate Bot. Gate will send you every signals that catch by robot`
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

func (h *Handler) brokerageList(ctx context.Context, msg *tgbotapi.Message) error {
	brokerages, err := h.brokerageService.List(ctx, new(proto.Void))
	if err != nil {
		log.WithError(err).Error("failed to load brokerages")
		return err
	}

	brListItemTmp := `
%v- %s
platform: %v
status: %v
`

	for i, brokerage := range brokerages.Elements {
		statusText := QueryStart
		if brokerage.Status == proto.Status_Enable {
			statusText = QueryStop
		}

		text := fmt.Sprintf(brListItemTmp, i+1, brokerage.Title, brokerage.Platform, brokerage.Status)

		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgBotApi.NewCallbackDataButton(statusText, brokerage.ID),
				tgBotApi.NewCallbackDataButton(QueryDelete, brokerage.ID),
			),
		)
		err := tgBotApi.SendMessage(ctx, tgBotApi.NewTextMessage(msg.Chat.ID, 0, text, &tgBotApi.Keyboard{
			UseInline: true,
			Inline:    &keyboard,
		}))
		if err != nil {
			return errors.Cast(ctx, err).AddDetailF("failed to send start message")
		}
	}

	return nil
}

func (h *Handler) updateMarkets(ctx context.Context, msg *tgbotapi.Message) error {
	//brokerage, err := h.brokerageService.Enable(ctx, new(proto.Void))
	//if err != nil {
	//	return err
	//}
	//markets, err := h.marketService.Update(ctx, &chipmunkApi.MarketUpdateReq{BrokerageID: brokerage.ID})
	//if err != nil {
	//	return errors.Cast(ctx, err).AddDetailF("failed to update markets")
	//}
	//text := fmt.Sprintf("updated markets: %v", len(markets.Elements))
	//err = tgBotApi.SendMessage(ctx, tgBotApi.NewTextMessage(msg.Chat.ID, 0, text, nil))
	//if err != nil {
	//	return errors.Cast(ctx, err).AddDetailF("failed to send start message")
	//}
	//return nil
	return errors.New(ctx, codes.Unimplemented)
}
