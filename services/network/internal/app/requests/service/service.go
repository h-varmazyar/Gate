package service

import (
	"context"
	api "github.com/h-varmazyar/Gate/api/proto"
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
	var (
		err      error
		ip       *rateLimiter.IP
		request  *requests.Request
		response *networkAPI.Response
	)

	ip, err = s.rateLimiterManager.GetRandomIP(ctx)
	if err != nil {
		if errors.Code(ctx, err) != codes.NotFound {
			s.logger.WithError(err).Errorf("failed to load IP")
			return nil, err
		}
	}
	request, err = requests.New(req, ip.ProxyURL)
	if err != nil {
		log.WithError(err).Errorf("failed to create network request")
		return nil, err
	}

	response, err = request.Do()
	if err != nil {
		s.logger.WithError(err).Errorf("failed to do network request for %v", req)
		return nil, err
	}

	return response, nil
}

func (s *Service) DoAsync(ctx context.Context, req *networkAPI.DoAsyncReq) (*api.Void, error) {
	if req.CallbackQueue == "" {
		return nil, errors.New(ctx, codes.FailedPrecondition).AddDetails("callback queue can not be empty")
	}
	if req.ReferenceID == "" {
		return nil, errors.New(ctx, codes.FailedPrecondition).AddDetails("reference id must be declared")
	}
	if len(req.Requests) == 0 {
		return nil, errors.New(ctx, codes.FailedPrecondition).AddDetails("empty requests not acceptable")
	}

	request := &rateLimiter.Request{
		CallbackQueue: req.CallbackQueue,
		ReferenceID:   req.ReferenceID,
		Requests:      req.Requests,
	}
	err := s.rateLimiterManager.AddNewRequest(ctx, req.RateLimiterID, request)
	if err != nil {
		s.logger.WithError(err).Errorf("failed to do async request for %v", req)
		return nil, err
	}

	return new(api.Void), nil
}
