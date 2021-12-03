package resolutions

import (
	"context"
	"github.com/google/uuid"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/errors"
	"github.com/mrNobody95/Gate/pkg/mapper"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 12.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

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

func (s *Service) Set(ctx context.Context, req *api.Resolution) (*api.Void, error) {
	resolution := new(repository.Resolution)
	mapper.Struct(req, resolution)
	if req.BrokerageID == "" {
		return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "brokerage id not set")
	} else if reqID, err := uuid.Parse(req.BrokerageID); err != nil {
		return nil, err
	} else {
		resolution.BrokerageID = reqID
	}
	if req.ID == "" {
		resolution.ID = uuid.New()
	} else if reqID, err := uuid.Parse(req.ID); err != nil {
		return nil, err
	} else {
		resolution.ID = reqID
	}
	if err := repository.Resolutions.Set(resolution); err != nil {
		return nil, err
	}
	return new(api.Void), nil
}

func (s *Service) GetByID(ctx context.Context, _ *brokerageApi.GetResolutionByIDRequest) (*api.Resolution, error) {
	return nil, errors.New(ctx, codes.Unimplemented)
}

func (s *Service) GetByDuration(_ context.Context, req *brokerageApi.GetResolutionByDurationRequest) (*api.Resolution, error) {
	id, err := uuid.Parse(req.BrokerageID)
	if err != nil {
		return nil, err
	}
	resolution, err := repository.Resolutions.GetByDuration(time.Duration(req.Duration), id)
	if err != nil {
		return nil, err
	}
	response := new(api.Resolution)
	mapper.Struct(resolution, response)
	return response, nil
}

func (s *Service) List(_ context.Context, req *brokerageApi.GetResolutionListRequest) (*api.Resolutions, error) {
	id, err := uuid.Parse(req.BrokerageID)
	if err != nil {
		return nil, err
	}
	resolutions, err := repository.Resolutions.List(id)
	if err != nil {
		return nil, err
	}
	response := new(api.Resolutions)
	mapper.Slice(resolutions, response.Resolutions)
	return response, nil
}
