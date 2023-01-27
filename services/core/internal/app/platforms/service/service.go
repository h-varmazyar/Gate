package service

import (
	"context"
	api "github.com/h-varmazyar/Gate/api/proto"
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

func (s *Service) CollectMarketData(ctx context.Context, req *coreApi.PlatformCollectDataReq) (*api.Void, error) {
	markets, err := s.marketService.List(ctx, &chipmunkApi.MarketListReq{
		Platform: req.Platform,
	})
	if err != nil {
		return nil, err
	}

	s.logger.Infof("market len: %v", len(markets.Elements))
	resolutions, err := s.resolutionService.List(ctx, &chipmunkApi.ResolutionListReq{
		Platform: req.Platform,
	})
	_, err = s.candleService.StartWorkers(ctx, &chipmunkApi.CandleWorkerStartReq{
		Platform:    req.Platform,
		Markets:     markets,
		Resolutions: resolutions,
	})
	if err != nil {
		return nil, err
	}

	//_, err = s.marketService.StartWorker(ctx, &chipmunkApi.WorkerStartReq{
	//	Platform: req.Platform,
	//})
	//if err != nil {
	//	return nil, err
	//}

	return new(api.Void), nil
}
