package signals

import (
	"context"
	"fmt"
	"github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api"
	"github.com/h-varmazyar/Gate/services/eagle/configs"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/strategies/automatedStrategy"
	"google.golang.org/grpc"
)

type Service struct {
	strategyService  eagleApi.StrategyServiceClient
	marketService    chipmunkApi.MarketServiceClient
	indicatorService chipmunkApi.IndicatorServiceClient
}

var (
	GrpcService *Service
)

func NewService() *Service {
	if GrpcService == nil {
		GrpcService = new(Service)

		eagleConn := grpcext.NewConnection(fmt.Sprintf(":%v", configs.Variables.GrpcPort))
		GrpcService.strategyService = eagleApi.NewStrategyServiceClient(eagleConn)

		chipmunkConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Chipmunk)
		GrpcService.marketService = chipmunkApi.NewMarketServiceClient(chipmunkConn)
		GrpcService.indicatorService = chipmunkApi.NewIndicatorServiceClient(chipmunkConn)
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	eagleApi.RegisterSignalServiceServer(server, s)
}

func (s *Service) Start(ctx context.Context, req *eagleApi.SignalStartReq) (*proto.Void, error) {
	if signalCheckWorker != nil && signalCheckWorker.IsRunning() {
		return new(proto.Void), nil
	}
	strategy, err := s.strategyService.Return(ctx, &eagleApi.ReturnStrategyReq{
		ID: req.StrategyID,
	})
	if err != nil {
		return nil, err
	}

	markets, err := s.marketService.List(ctx, &chipmunkApi.MarketListRequest{BrokerageID: req.BrokerageID})
	if err != nil {
		return nil, err
	}

	automated, err := automatedStrategy.NewAutomatedStrategy(strategy, req.WithTrading)
	if err != nil {
		return nil, errors.Cast(ctx, err).AddDetailF("failed to start strategy")
	}

	signalCheckInstance := SignalCheckWorkerInstance(automated, markets.Elements)
	signalCheckInstance.Start()

	return new(proto.Void), nil
}

func (s *Service) Stop(_ context.Context, _ *proto.Void) (*proto.Void, error) {
	StopSignalChecker()
	return new(proto.Void), nil
}
