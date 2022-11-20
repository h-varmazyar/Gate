package strategies

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/repository"
	log "github.com/sirupsen/logrus"
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

func (s *Service) Create(_ context.Context, req *eagleApi.CreateStrategyReq) (*eagleApi.Strategy, error) {
	strategy := new(repository.Strategy)
	mapper.Struct(req, strategy)
	if err := repository.Strategies.Save(strategy); err != nil {
		return nil, err
	}
	response := new(eagleApi.Strategy)
	mapper.Struct(strategy, response)
	return response, nil
}

func (s *Service) Return(_ context.Context, req *eagleApi.ReturnStrategyReq) (*eagleApi.Strategy, error) {
	strategyID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}
	strategy, err := repository.Strategies.Return(strategyID)
	if err != nil {
		return nil, err
	}
	log.Warnf("ind after repo: %v", strategy.Indicators)
	response := new(eagleApi.Strategy)
	mapper.Struct(strategy, response)
	return response, nil
}

func (s *Service) List(_ context.Context, _ *proto.Void) (*eagleApi.Strategies, error) {
	strategies, err := repository.Strategies.List()
	if err != nil {
		return nil, err
	}
	response := new(eagleApi.Strategies)
	mapper.Slice(strategies, &response.Elements)
	return response, nil
}

func (s *Service) Indicators(_ context.Context, req *eagleApi.StrategyIndicatorReq) (*eagleApi.StrategyIndicators, error) {
	strategyID, err := uuid.Parse(req.StrategyID)
	if err != nil {
		return nil, err
	}
	indicators, err := repository.Strategies.ReturnIndicators(strategyID)
	if err != nil {
		return nil, err
	}
	response := new(eagleApi.StrategyIndicators)
	mapper.Slice(indicators, &response.Elements)
	return response, nil
}
