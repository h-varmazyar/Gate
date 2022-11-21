package service

import (
	"context"
	"github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api/proto"
	"github.com/h-varmazyar/Gate/services/eagle/internal/app/signals/workers"
	strategies "github.com/h-varmazyar/Gate/services/eagle/internal/app/strategies/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Service struct {
	strategyService   *strategies.Service
	marketService     chipmunkApi.MarketServiceClient
	indicatorService  chipmunkApi.IndicatorServiceClient
	signalCheckWorker *workers.SignalCheckWorker
	configs           *Configs
	logger            *log.Logger
}

type Dependencies struct {
	StrategyService   *strategies.Service
	SignalCheckWorker *workers.SignalCheckWorker
}

var (
	GrpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, configs *Configs, dependencies *Dependencies) *Service {
	if GrpcService == nil {
		GrpcService = new(Service)

		chipmunkConn := grpcext.NewConnection(configs.ChipmunkAddress)
		GrpcService.marketService = chipmunkApi.NewMarketServiceClient(chipmunkConn)
		GrpcService.indicatorService = chipmunkApi.NewIndicatorServiceClient(chipmunkConn)
		GrpcService.strategyService = dependencies.StrategyService
		GrpcService.signalCheckWorker = dependencies.SignalCheckWorker
		GrpcService.configs = configs
		GrpcService.logger = logger
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	eagleApi.RegisterSignalServiceServer(server, s)
}

func (s *Service) Start(ctx context.Context, req *eagleApi.SignalStartReq) (*proto.Void, error) {
	//if s.signalCheckWorker != nil && s.signalCheckWorker.IsRunning() {
	//	return new(proto.Void), nil
	//}
	//strategy, err := s.strategyService.Return(ctx, &eagleApi.ReturnStrategyReq{
	//	ID: req.StrategyID,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//markets, err := s.marketService.List(ctx, &chipmunkApi.MarketListReq{Platform: req.Platform})
	//if err != nil {
	//	return nil, err
	//}
	//
	//automated, err := automatedStrategy.NewAutomatedStrategy(strategy, req.WithTrading)
	//if err != nil {
	//	return nil, errors.Cast(ctx, err).AddDetailF("failed to start strategy")
	//}
	//
	//s.signalCheckWorker.Start(automated, markets.Elements)

	return new(proto.Void), nil
}

func (s *Service) Stop(_ context.Context, _ *proto.Void) (*proto.Void, error) {
	//workers.StopSignalChecker()
	return new(proto.Void), nil
}
