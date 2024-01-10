package service

import (
	"context"
	indicatorsAPI "github.com/h-varmazyar/Gate/services/indicators/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Service struct {
	log     *log.Logger
	configs Configs
}

var (
	grpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, configs Configs) *Service {
	if grpcService == nil {
		grpcService = new(Service)
		grpcService.log = logger
		grpcService.configs = configs
	}
	return grpcService
}

func (s Service) RegisterServer(server *grpc.Server) {
	indicatorsAPI.RegisterIndicatorServiceServer(server, s)
}

func (s Service) Register(ctx context.Context, req *indicatorsAPI.IndicatorRegisterReq) (*indicatorsAPI.Indicator, error) {

}

func (s Service) Values(ctx context.Context, req *indicatorsAPI.IndicatorValuesReq) (*indicatorsAPI.IndicatorValues, error) {

}
