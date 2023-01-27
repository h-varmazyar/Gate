package service

import (
	"context"
	"github.com/google/uuid"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api/proto"
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/strategies/repository"
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/strategies/workers"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Service struct {
	configs           *Configs
	logger            *log.Logger
	db                repository.StrategyRepository
	signalCheckWorker *workers.SignalCheckWorker
}

type Dependencies struct {
	SignalCheckWorker *workers.SignalCheckWorker
}

var (
	GrpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, configs *Configs, db repository.StrategyRepository, dependencies *Dependencies) *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		GrpcService.configs = configs
		GrpcService.logger = logger
		GrpcService.db = db
		GrpcService.signalCheckWorker = dependencies.SignalCheckWorker
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	eagleApi.RegisterStrategyServiceServer(server, s)
}

func (s *Service) Create(_ context.Context, req *eagleApi.CreateStrategyReq) (*eagleApi.Strategy, error) {
	strategy := new(entity.Strategy)
	mapper.Struct(req, strategy)
	if err := s.db.Save(strategy); err != nil {
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
	strategy, err := s.db.Return(strategyID)
	if err != nil {
		return nil, err
	}
	log.Warnf("ind after repo: %v", strategy.Indicators)
	response := new(eagleApi.Strategy)
	mapper.Struct(strategy, response)
	return response, nil
}

func (s *Service) List(_ context.Context, _ *api.Void) (*eagleApi.Strategies, error) {
	strategies, err := s.db.List()
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
	indicators, err := s.db.ReturnIndicators(strategyID)
	if err != nil {
		return nil, err
	}
	response := new(eagleApi.StrategyIndicators)
	mapper.Slice(indicators, &response.Elements)
	return response, nil
}

//func (s *Service) StartWorker(ctx context.Context, req *eagleApi.StrategyStartWorkerReq) (*api.Void, error) {
//	strategies, err := s.db.ReturnActives(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	for _, strategy := range strategies {
//		switch strategy.Type {
//		case eagleApi.StrategyType_Automated:
//			automated, err := automatedStrategy.NewAutomatedStrategy(strategy, s.configs.Automa)
//			if err != nil {
//				s.logger.WithError(err).Errorf("failed to create new instance of automated strategy")
//				return nil, err
//			}
//
//			s.signalCheckWorker.Start(automated, req.Markets.Elements, strategy.BrokerageID)
//		default:
//		}
//	}
//
//	return new(api.Void), nil
//}
//
//func (s *Service) StartSignalChecker(ctx context.Context, req *eagleApi.StrategySignalCheckStartReq) (*api.Void, error) {
//	if req.Markets == nil || len(req.Markets.Elements) == 0 {
//		err := errors.New(ctx, codes.FailedPrecondition)
//		s.logger.WithError(err).Errorf("invalid markets")
//		return nil, err
//	}
//	brokerageID, err := uuid.Parse(req.BrokerageID)
//	if err != nil {
//		s.logger.WithError(err).Errorf("failed to start worker. invalid brokerage id %v", req.BrokerageID)
//		return nil, err
//	}
//	strategyID, err := uuid.Parse(req.BrokerageID)
//	if err != nil {
//		s.logger.WithError(err).Errorf("failed to start worker. invalid strategy id %v", req.StrategyID)
//		return nil, err
//	}
//	strategy, err := s.db.Return(strategyID)
//	if err != nil {
//		s.logger.WithError(err).Errorf("failed to return strategy with id %v", req.StrategyID)
//	}
//
//	automated, err := automatedStrategy.NewAutomatedStrategy(strategy, req.WithTrading, s.configs.Automated)
//	if err != nil {
//		s.logger.WithError(err).Errorf("failed to create new instance of automated strategy for %v", req.StrategyID)
//		return nil, err
//	}
//
//	s.signalCheckWorker.Start(automated, req.Markets.Elements, brokerageID)
//	return new(api.Void), nil
//}
//
//func (s *Service) StopSignalChecker(_ context.Context, req *eagleApi.StrategySignalCheckStopReq) (*api.Void, error) {
//	brokerageID, err := uuid.Parse(req.BrokerageID)
//	if err != nil {
//		s.logger.WithError(err).Errorf("failed to stop worker. invalid brokerage id %v", req.BrokerageID)
//		return nil, err
//	}
//	s.signalCheckWorker.Stop(brokerageID)
//	return new(api.Void), nil
//}
