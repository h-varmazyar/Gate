package service

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Service struct {
	logger            *log.Logger
	configs           *Configs
	candleService     chipmunkApi.CandleServiceClient
	marketService     chipmunkApi.MarketServiceClient
	resolutionService chipmunkApi.ResolutionServiceClient
}

var (
	grpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, configs *Configs) *Service {
	if grpcService == nil {
		grpcService = new(Service)
		chipmunkConn := grpcext.NewConnection(configs.ChipmunkGrpcAddress)
		grpcService.candleService = chipmunkApi.NewCandleServiceClient(chipmunkConn)
		grpcService.marketService = chipmunkApi.NewMarketServiceClient(chipmunkConn)
		grpcService.resolutionService = chipmunkApi.NewResolutionServiceClient(chipmunkConn)
		grpcService.logger = logger
		grpcService.configs = configs
	}
	return grpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	coreApi.RegisterPlatformServiceServer(server, s)
}
