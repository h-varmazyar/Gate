package strategy

import (
	"context"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/mapper"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/repository"
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
	brokerageApi.RegisterStrategyServiceServer(server, s)
}

func (s *Service) Create(ctx context.Context, req *brokerageApi.CreateStrategyReq) (*brokerageApi.Strategy, error) {
	strategy := new(repository.Strategy)
	mapper.Struct(req, strategy)
	if err := repository.Strategies.Create(strategy); err != nil {
		return nil, err
	}
	response := new(brokerageApi.Strategy)
	mapper.Struct(strategy, response)
	return response, nil
}

func (s *Service) Return(ctx context.Context, req *brokerageApi.ReturnStrategyReq) (*brokerageApi.Strategy, error) {
	strategy, err := repository.Strategies.Return(req.ID)
	if err != nil {
		return nil, err
	}
	response := new(brokerageApi.Strategy)
	mapper.Struct(strategy, response)
	return response, nil
}

func (s *Service) List(ctx context.Context, _ *api.Void) (*brokerageApi.Strategies, error) {
	strategies, err := repository.Strategies.List()
	if err != nil {
		return nil, err
	}
	response := new(brokerageApi.Strategies)
	mapper.Slice(strategies, &response.Elements)
	return response, nil
}
