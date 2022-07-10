package tgBotApi

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"google.golang.org/grpc/codes"
)

type UnImplementedHandler struct {
}

func (*UnImplementedHandler) HandleCommand(ctx context.Context, _ *tgbotapi.Message) error {
	return errors.New(ctx, codes.Unimplemented)
}

func (*UnImplementedHandler) HandleChannelPost(ctx context.Context, _ *tgbotapi.Message) error {
	return errors.New(ctx, codes.Unimplemented)
}

func (*UnImplementedHandler) HandleMessage(ctx context.Context, _ *tgbotapi.Message) error {
	return errors.New(ctx, codes.Unimplemented)
}
