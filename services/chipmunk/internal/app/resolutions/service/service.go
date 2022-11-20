package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/resolutions/repository"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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

func (s *Service) ReturnByID(ctx context.Context, _ *chipmunkApi.ResolutionReturnByIDReq) (*chipmunkApi.Resolution, error) {
	return nil, errors.New(ctx, codes.Unimplemented)
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
