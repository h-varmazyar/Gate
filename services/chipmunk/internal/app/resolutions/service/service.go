package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/resolutions/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

type Service struct {
	db      repository.ResolutionRepository
	logger  *log.Logger
	configs *Configs
}

var (
	GrpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, configs *Configs, db repository.ResolutionRepository) *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		GrpcService.logger = logger
		GrpcService.configs = configs
		GrpcService.db = db
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	chipmunkApi.RegisterResolutionServiceServer(server, s)
}

func (s *Service) Set(_ context.Context, req *chipmunkApi.Resolution) (*proto.Void, error) {
	resolution := new(entity.Resolution)
	mapper.Struct(req, resolution)
	resolution.ID, _ = uuid.Parse(req.ID)
	if err := s.db.Set(resolution); err != nil {
		return nil, err
	}
	return new(proto.Void), nil
}

func (s *Service) ReturnByID(_ context.Context, req *chipmunkApi.ResolutionReturnByIDReq) (*chipmunkApi.Resolution, error) {
	resolutionID, err := uuid.Parse(req.ID)
	if err != nil {
		s.logger.WithError(err).Errorf("failed to parse resolution id: %v", req.ID)
		return nil, err
	}
	resolution, err := s.db.Return(resolutionID)
	if err != nil {
		s.logger.WithError(err).Errorf("failed to fetch resolution: %v", resolutionID)
		return nil, err
	}
	response := new(chipmunkApi.Resolution)
	mapper.Struct(resolution, response)
	return response, nil
}

func (s *Service) ReturnByDuration(_ context.Context, req *chipmunkApi.ResolutionReturnByDurationReq) (*chipmunkApi.Resolution, error) {
	resolution, err := s.db.ReturnByDuration(time.Duration(req.Duration), req.Platform)
	if err != nil {
		return nil, err
	}
	response := new(chipmunkApi.Resolution)
	mapper.Struct(resolution, response)
	return response, nil
}

func (s *Service) List(_ context.Context, req *chipmunkApi.ResolutionListReq) (*chipmunkApi.Resolutions, error) {
	resolutions, err := s.db.List(req.Platform)
	if err != nil {
		return nil, err
	}
	response := new(chipmunkApi.Resolutions)
	mapper.Slice(resolutions, &response.Elements)
	return response, nil
}
