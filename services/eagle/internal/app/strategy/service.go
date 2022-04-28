package strategy

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/repository"
	"google.golang.org/grpc"
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
	eagleApi.RegisterStrategyServiceServer(server, s)
}

func (s *Service) Create(ctx context.Context, req *eagleApi.CreateStrategyReq) (*eagleApi.Strategy, error) {
	strategy := new(repository.Strategy)
	mapper.Struct(req, strategy)
	if err := repository.Strategies.Save(strategy); err != nil {
		return nil, err
	}
	response := new(eagleApi.Strategy)
	mapper.Struct(strategy, response)
	return response, nil
}

func (s *Service) Return(ctx context.Context, req *eagleApi.ReturnStrategyReq) (*eagleApi.Strategy, error) {
	strategyID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}
	strategy, err := repository.Strategies.Return(strategyID)
	if err != nil {
		return nil, err
	}
	response := new(eagleApi.Strategy)
	mapper.Struct(strategy, response)
	return response, nil
}

func (s *Service) List(ctx context.Context, _ *api.Void) (*eagleApi.Strategies, error) {
	strategies, err := repository.Strategies.List()
	if err != nil {
		return nil, err
	}
	response := new(eagleApi.Strategies)
	mapper.Slice(strategies, &response.Elements)
	return response, nil
}
