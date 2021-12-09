package signals

import (
	"context"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/errors"
	eagleApi "github.com/mrNobody95/Gate/services/eagle/api"
	"github.com/mrNobody95/Gate/services/eagle/internal/pkg/strategies/automatedStrategy"
	"github.com/mrNobody95/Gate/services/eagle/internal/pkg/workers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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
* Date: 09.12.21
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
	eagleApi.RegisterSignalsServiceServer(server, s)
}

func (s *Service) NewWorker(ctx context.Context, req *eagleApi.AddSignalsWorkerRequest) (*api.Void, error) {
	settings := new(workers.SignalsSettings)
	if req.Market == nil {
		return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "market is nil")
	}
	if req.Resolution == nil {
		return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "resolution is nil")
	}
	settings.Market = req.Market
	settings.Resolution = req.Resolution
	settings.Strategy = &automatedStrategy.Automated{MinGainPercentage: 1}
	workers.SignalsWorker.AddMarket(settings)
	return &api.Void{}, nil
}

func (s *Service) CancelWorker(_ context.Context, req *eagleApi.CancelSignalsWorkerRequest) (*api.Void, error) {
	return new(api.Void), workers.SignalsWorker.CancelWorker(req.MarketID, req.ResolutionID)
}
