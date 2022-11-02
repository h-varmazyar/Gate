package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	coreApi "github.com/h-varmazyar/Gate/services/core/api"
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/pkg/tgBotApi"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

const (
	ErrNotImplemented = "not implemented yet"
)

type Handler struct {
	brokerageService coreApi.BrokerageServiceClient
	marketService    chipmunkApi.MarketServiceClient
}

func NewInstance(configs *Configs) *Handler {
	handler := new(Handler)
	brokerageConn := grpcext.NewConnection(configs.BrokerageAddress)
	chipmunkConn := grpcext.NewConnection(configs.ChipmunkAddress)
	handler.brokerageService = coreApi.NewBrokerageServiceClient(brokerageConn)
	handler.marketService = chipmunkApi.NewMarketServiceClient(chipmunkConn)

	return handler
}

func (h *Handler) HandleMessage(ctx context.Context, msg *tgbotapi.Message) error {

	return nil
}

func (h *Handler) HandleCommand(ctx context.Context, msg *tgbotapi.Message) error {
	var err error
	log.Info(msg)
	switch msg.Command() {
	case CommandStart:
		err = h.startCommand(ctx, msg)
	case CmdBrokerageList:
		err = h.brokerageList(ctx, msg)
	case CmdUpdateMarkets:
		err = h.updateMarkets(ctx, msg)
	default:
		err = errors.New(ctx, codes.Unimplemented)
	}
	log.WithError(err).Error("failed to handle command")
	return err
}

func (h *Handler) HandleChannelPost(ctx context.Context, msg *tgbotapi.Message) error {
	log.Info(msg)
	return nil
}

func (h *Handler) HandleCallbackQuery(ctx context.Context, callbackQuery *tgbotapi.CallbackQuery) error {
	query, data, err := tgBotApi.ParseCallbackData(callbackQuery.Data)
	if err != nil {
		return err
	}
	switch query {
	case QueryStart:
		err = h.startCallback(ctx, callbackQuery, data)
	case QueryStop:
		err = h.stopCallback(ctx, callbackQuery, data)
	case QueryDelete:
		err = h.deleteCallback(ctx, callbackQuery, data)
	default:
		err = errors.New(ctx, codes.Unimplemented)
	}
	return err
}
