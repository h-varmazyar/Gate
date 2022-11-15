package resolutions

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
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
	chipmunkApi.RegisterResolutionServiceServer(server, s)
}

func (s *Service) Set(_ context.Context, req *chipmunkApi.Resolution) (*api.Void, error) {
	resolution := new(repository.Resolution)
	mapper.Struct(req, resolution)
	resolution.ID, _ = uuid.Parse(req.ID)
	if err := repository.Resolutions.Set(resolution); err != nil {
		return nil, err
	}
	return new(api.Void), nil
}

func (s *Service) GetByID(ctx context.Context, _ *chipmunkApi.GetResolutionByIDRequest) (*chipmunkApi.Resolution, error) {
	return nil, errors.New(ctx, codes.Unimplemented)
}

func (s *Service) GetByDuration(_ context.Context, req *chipmunkApi.GetResolutionByDurationRequest) (*chipmunkApi.Resolution, error) {
	resolution, err := repository.Resolutions.GetByDuration(time.Duration(req.Duration), req.BrokerageName)
	if err != nil {
		return nil, err
	}
	response := new(chipmunkApi.Resolution)
	mapper.Struct(resolution, response)
	return response, nil
}

func (s *Service) List(_ context.Context, req *chipmunkApi.GetResolutionListRequest) (*chipmunkApi.Resolutions, error) {
	resolutions, err := repository.Resolutions.List(req.BrokerageName)
	if err != nil {
		return nil, err
	}
	response := new(chipmunkApi.Resolutions)
	mapper.Slice(resolutions, &response.Elements)
	return response, nil
}
