package service

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/errors"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	ipService "github.com/h-varmazyar/Gate/services/network/internal/app/IPs/service"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/rateLimiter"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/requests"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Service struct {
	logger             *log.Logger
	rateLimiterManager *rateLimiter.Manager
	ipService          ipService.Service
}

var (
	GrpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, rateLimiterManager *rateLimiter.Manager) *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		GrpcService.logger = logger
		GrpcService.rateLimiterManager = rateLimiterManager
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	networkAPI.RegisterRequestServiceServer(server, s)
}

func (s *Service) Do(ctx context.Context, req *networkAPI.Request) (*networkAPI.Response, error) {
	response := new(networkAPI.Response)
	var err error
	if req.Type == networkAPI.Request_Async {
		err = s.rateLimiterManager.AddNewRequest(ctx, req)
	} else {
		response, err = s.handleSyncRequest(ctx, req)
	}
	if err != nil {
		s.logger.WithError(err).Errorf("failed to do network request for %v", req)
		return nil, err
	}

	return response, nil
}

func (s *Service) handleSyncRequest(ctx context.Context, req *networkAPI.Request) (*networkAPI.Response, error) {
	//todo: change next line
	ip, err := s.rateLimiterManager.GetRandomIP(ctx)
	if err != nil {
		if errors.Code(ctx, err) != codes.NotFound {
			s.logger.WithError(err).Errorf("failed to load IP")
			return nil, err
		}
	}
	request, err := requests.New(req, ip.ProxyURL)
	if err != nil {
		log.WithError(err).Errorf("failed to create network request")
		return nil, err
	}

	return request.Do()
}
