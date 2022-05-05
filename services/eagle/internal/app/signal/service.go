package signal

import (
	"context"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api"
	"github.com/h-varmazyar/Gate/services/eagle/configs"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/strategies/automatedStrategy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Service struct {
	StrategyService  brokerageApi.StrategyServiceClient
	MarketService    brokerageApi.MarketServiceClient
	IndicatorService chipmunkApi.IndicatorServiceClient
}

var (
	GrpcService *Service
)

func NewService() *Service {
	if GrpcService == nil {
		GrpcService = new(Service)

		brokerageConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)
		GrpcService.StrategyService = brokerageApi.NewStrategyServiceClient(brokerageConn)
		GrpcService.MarketService = brokerageApi.NewMarketServiceClient(brokerageConn)

		chipmunkConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Chipmunk)
		GrpcService.IndicatorService = chipmunkApi.NewIndicatorServiceClient(chipmunkConn)
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	eagleApi.RegisterSignalServiceServer(server, s)
}

func (s *Service) Start(ctx context.Context, req *eagleApi.SignalStartReq) (*api.Void, error) {
	if signalCheckWorker != nil && signalCheckWorker.IsRunning() {
		return nil, errors.New(ctx, codes.AlreadyExists)
	}
	strategy, err := s.StrategyService.Return(ctx, &brokerageApi.ReturnStrategyReq{
		ID: req.StrategyID,
	})
	if err != nil {
		return nil, err
	}

	markets, err := s.MarketService.List(ctx, &brokerageApi.MarketListRequest{BrokerageID: req.BrokerageID})
	if err != nil {
		return nil, err
	}

	automated, err := automatedStrategy.NewAutomatedStrategy(strategy, req.WithTrading)
	if err != nil {
		return nil, errors.Cast(ctx, err).AddDetailF("failed to start strategy")
	}

	signalCheckInstance := SignalCheckWorkerInstance(automated, markets.Markets)
	signalCheckInstance.Start()

	return new(api.Void), nil
}

func (s *Service) Stop(_ context.Context, _ *api.Void) (*api.Void, error) {
	StopSignalChecker()
	return new(api.Void), nil
}
