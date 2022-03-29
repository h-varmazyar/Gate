package wallets

import (
	"context"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api"
	"google.golang.org/grpc"
)

type Service struct {
	brokerageService brokerageApi.BrokerageServiceClient
	networkService   networkAPI.RequestServiceClient
}

var (
	GrpcService *Service
)

func NewService() *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		brokerageConnection := grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)
		networkConnection := grpcext.NewConnection(configs.Variables.GrpcAddresses.Network)
		GrpcService.brokerageService = brokerageApi.NewBrokerageServiceClient(brokerageConnection)
		GrpcService.networkService = networkAPI.NewRequestServiceClient(networkConnection)
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	chipmunkApi.RegisterWalletsServiceServer(server, s)
}

func (s *Service) List(_ context.Context, _ *chipmunkApi.WalletListRequest) (*brokerageApi.Wallets, error) {
	return buffer.Wallets.FetchAll(), nil
}

func (s *Service) StartWorker(ctx context.Context, req *chipmunkApi.StartWorkerRequest) (*api.Void, error) {
	brokerage, err := s.brokerageService.Get(ctx, &brokerageApi.BrokerageIDReq{ID: req.BrokerageID})
	if err != nil {
		return nil, err
	}
	if err := Worker.Start(brokerage); err != nil {
		return nil, err
	}
	return new(api.Void), nil
}

func (s *Service) CancelWorker(ctx context.Context, _ *api.Void) (*api.Void, error) {
	Worker.Stop()
	return new(api.Void), nil
}
