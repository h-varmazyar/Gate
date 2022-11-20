package service

import (
	"context"
	"github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/wallets/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/wallets/workers"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Service struct {
	brokerageService coreApi.BrokerageServiceClient
	logger           *log.Logger
	configs          *Configs
	buffer           *buffer.WalletBuffer
	worker           *workers.WalletCheck
}

type Dependencies struct {
	Buffer *buffer.WalletBuffer
	Worker *workers.WalletCheck
}

var (
	GrpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, configs *Configs, dependencies *Dependencies) *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		brokerageConnection := grpcext.NewConnection(configs.CoreAddress)
		GrpcService.brokerageService = coreApi.NewBrokerageServiceClient(brokerageConnection)
		GrpcService.logger = logger
		GrpcService.configs = configs
		GrpcService.buffer = dependencies.Buffer
		GrpcService.worker = dependencies.Worker
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	chipmunkApi.RegisterWalletsServiceServer(server, s)
}

func (s *Service) List(_ context.Context, _ *proto.Void) (*chipmunkApi.Wallets, error) {
	return s.buffer.FetchAll(), nil
}

func (s *Service) StartWorker(ctx context.Context, req *chipmunkApi.StartWorkerRequest) (*proto.Void, error) {
	brokerage, err := s.brokerageService.Return(ctx, &coreApi.BrokerageReturnReq{ID: req.BrokerageID})
	if err != nil {
		return nil, err
	}
	if err := s.worker.Start(brokerage); err != nil {
		return nil, err
	}
	return new(proto.Void), nil
}

func (s *Service) StopWorker(_ context.Context, _ *proto.Void) (*proto.Void, error) {
	s.worker.Stop()
	return new(proto.Void), nil
}

func (s *Service) ReturnWallet(ctx context.Context, req *chipmunkApi.ReturnWalletReq) (*chipmunkApi.Wallet, error) {
	wallet := s.buffer.FetchWallet(req.AssetName)
	if wallet != nil {
		return wallet, nil
	}
	return nil, errors.New(ctx, codes.NotFound).AddDetailF("wallet %v not found", req.AssetName)
}

func (s *Service) ReturnReference(ctx context.Context, req *chipmunkApi.ReturnReferenceReq) (*chipmunkApi.Reference, error) {
	wallet := s.buffer.FetchReference(req.ReferenceName)
	if wallet != nil {
		return wallet, nil
	}
	return nil, errors.New(ctx, codes.NotFound).AddDetailF("wallet %v not found", req.ReferenceName)
}
