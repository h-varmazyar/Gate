package service

import (
	"context"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/rateLimiter"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/requests"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Service struct {
	logger             *log.Logger
	configs            *Configs
	rateLimiterManager *rateLimiter.Manager
}

var (
	GrpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, configs *Configs, rateLimiterManager *rateLimiter.Manager) *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		GrpcService.logger = logger
		GrpcService.configs = configs
		GrpcService.rateLimiterManager = rateLimiterManager
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	networkAPI.RegisterRequestServiceServer(server, s)
}

func (s *Service) Do(ctx context.Context, req *networkAPI.Request) (*networkAPI.Response, error) {
	log.Infof("new request: %v", req.Endpoint)

	response := new(networkAPI.Response)
	var err error
	if req.Type == networkAPI.Request_Async {
		err = s.rateLimiterManager.AddNewRequest(ctx, req)
	} else {
		response, err = s.handleSyncRequest(ctx, req)
	}
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *Service) handleSyncRequest(_ context.Context, req *networkAPI.Request) (*networkAPI.Response, error) {
	request, err := requests.New(req, nil)
	if err != nil {
		log.WithError(err).Errorf("failed to create network request")
		return nil, err
	}

	return request.Do()
}
