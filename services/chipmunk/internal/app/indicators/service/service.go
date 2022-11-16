package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/indicators/repository"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Service struct {
	logger  *log.Logger
	configs *Configs
	db      repository.IndicatorRepository
}

var (
	GrpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, configs *Configs, db repository.IndicatorRepository) *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		GrpcService.logger = logger
		GrpcService.configs = configs
		GrpcService.db = db
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	chipmunkApi.RegisterIndicatorServiceServer(server, s)
}

func (s *Service) Return(_ context.Context, req *chipmunkApi.IndicatorReturnReq) (*chipmunkApi.Indicator, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}

	indicator, err := s.db.Return(id)
	if err != nil {
		return nil, err
	}
	response := new(chipmunkApi.Indicator)
	mapper.Struct(indicator, response)
	return response, nil
}
