package resolutions

import (
	"context"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/errors"
	"github.com/mrNobody95/Gate/pkg/mapper"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/repository"
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
	brokerageApi.RegisterResolutionServiceServer(server, s)
}

func (s *Service) Set(_ context.Context, req *brokerageApi.Resolution) (*api.Void, error) {
	resolution := new(repository.Resolution)
	mapper.Struct(req, resolution)
	resolution.ID = uint(req.ID)
	if err := repository.Resolutions.Set(resolution); err != nil {
		return nil, err
	}
	return new(api.Void), nil
}

func (s *Service) GetByID(ctx context.Context, _ *brokerageApi.GetResolutionByIDRequest) (*brokerageApi.Resolution, error) {
	return nil, errors.New(ctx, codes.Unimplemented)
}

func (s *Service) GetByDuration(_ context.Context, req *brokerageApi.GetResolutionByDurationRequest) (*brokerageApi.Resolution, error) {
	resolution, err := repository.Resolutions.GetByDuration(time.Duration(req.Duration), req.BrokerageName)
	if err != nil {
		return nil, err
	}
	response := new(brokerageApi.Resolution)
	mapper.Struct(resolution, response)
	return response, nil
}

func (s *Service) List(_ context.Context, req *brokerageApi.GetResolutionListRequest) (*brokerageApi.Resolutions, error) {
	resolutions, err := repository.Resolutions.List(req.BrokerageName)
	if err != nil {
		return nil, err
	}
	response := new(brokerageApi.Resolutions)
	mapper.Slice(resolutions, &response.Resolutions)
	return response, nil
}
