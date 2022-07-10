package app

import (
	"context"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/errors"
	telegramBotApi "github.com/h-varmazyar/Gate/services/telegramBot/api"
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/app/handlers"
	"github.com/h-varmazyar/Gate/services/telegramBot/internal/pkg/tgBotApi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Service struct {
}

var (
	GrpcService *Service
)

func NewService() *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	telegramBotApi.RegisterBotServiceServer(server, s)
}

func (s *Service) Start(_ context.Context, _ *api.Void) (*api.Void, error) {
	handler := new(handlers.Handler)
	if err := tgBotApi.Run(handler); err != nil {
		return nil, err
	}
	return new(api.Void), nil
}

func (s *Service) Stop(ctx context.Context, _ *api.Void) (*api.Void, error) {
	if err := tgBotApi.Stop(ctx); err != nil {
		return nil, err
	}
	return new(api.Void), nil
}

func (s *Service) SendMessage(ctx context.Context, req *telegramBotApi.Message) (*api.Void, error) {
	if !tgBotApi.IsRunning() {
		return nil, errors.New(ctx, codes.Aborted).AddDetails("bot is not running")
	}

	var (
		err error
		msg tgBotApi.Message
	)
	if msg, err = parseMessage(ctx, req); err != nil {
		return nil, err
	}

	if err = tgBotApi.SendMessage(ctx, msg); err != nil {
		return nil, err
	}
	return new(api.Void), nil
}

func parseMessage(_ context.Context, req *telegramBotApi.Message) (tgBotApi.Message, error) {
	msg := tgBotApi.NewTextMessage(req.ChatID, int(req.ReplyTo), req.Text, nil)
	return msg, nil
}
