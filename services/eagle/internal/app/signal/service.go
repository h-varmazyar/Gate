package signal

import (
	"context"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api"
	"github.com/h-varmazyar/Gate/services/eagle/configs"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Service struct {
	strategyService eagleApi.StrategyServiceClient
}

var (
	GrpcService *Service
)

func NewService() *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		brokerageConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)
		GrpcService.strategyService = eagleApi.NewStrategyServiceClient(brokerageConn)
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	eagleApi.RegisterSignalServiceServer(server, s)
}

func (s *Service) StartWorker(ctx context.Context, req *eagleApi.StartWorkerReq) (*api.Void, error) {
	strategy, err := s.strategyService.Return(ctx, &eagleApi.ReturnStrategyReq{ID: req.StrategyID})
	if err != nil {
		return nil, err
	}
	localStrategy := new(repository.Strategy)
	mapper.Struct(strategy, localStrategy)
	err = Worker.Start(ctx, localStrategy)
	if err != nil {
		return nil, err
	}
	return new(api.Void), nil
}

func (s *Service) StopWorker(ctx context.Context, _ *api.Void) (*api.Void, error) {
	if !Worker.enable {
		return nil, errors.NewWithSlug(ctx, codes.Aborted, "worker_stopped_before")
	}
	Worker.Stop()
	return nil, nil
}

func (s *Service) AddCandleData(ctx context.Context, req *eagleApi.CandleData) (*api.Void, error) {
	if !Worker.enable {
		return nil, errors.NewWithSlug(ctx, codes.Aborted, "worker_stopped_before")
	}
	Worker.dataChan <- req
	return nil, nil
}
