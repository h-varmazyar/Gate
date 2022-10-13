package wallets

import (
	"context"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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

func (s *Service) List(_ context.Context, _ *api.Void) (*chipmunkApi.Wallets, error) {
	return buffer.Wallets.FetchAll(), nil
}

func (s *Service) StartWorker(ctx context.Context, req *chipmunkApi.StartWorkerRequest) (*api.Void, error) {
	brokerage, err := s.brokerageService.Return(ctx, &brokerageApi.BrokerageReturnReq{ID: req.BrokerageID})
	if err != nil {
		return nil, err
	}
	if err := Worker.Start(brokerage); err != nil {
		return nil, err
	}
	return new(api.Void), nil
}

func (s *Service) StopWorker(_ context.Context, _ *api.Void) (*api.Void, error) {
	Worker.Stop()
	return new(api.Void), nil
}

func (s *Service) ReturnWallet(ctx context.Context, req *chipmunkApi.ReturnWalletReq) (*chipmunkApi.Wallet, error) {
	wallet := buffer.Wallets.FetchWallet(req.AssetName)
	if wallet != nil {
		return wallet, nil
	}
	return nil, errors.New(ctx, codes.NotFound).AddDetailF("wallet %v not found", req.AssetName)
}

func (s *Service) ReturnReference(ctx context.Context, req *chipmunkApi.ReturnReferenceReq) (*chipmunkApi.Reference, error) {
	wallet := buffer.Wallets.FetchReference(req.ReferenceName)
	if wallet != nil {
		return wallet, nil
	}
	return nil, errors.New(ctx, codes.NotFound).AddDetailF("wallet %v not found", req.ReferenceName)
}
