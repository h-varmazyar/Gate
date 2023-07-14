package service

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/errors"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	ipService "github.com/h-varmazyar/Gate/services/network/internal/app/IPs/service"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/requests"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Service struct {
	logger             *log.Logger
	rateLimiterManager *requests.Manager
	ipService          ipService.Service
}

var (
	GrpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, rateLimiterManager *requests.Manager) *Service {
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
		ip       *requests.IP
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
	request, err = requests.NewNetworkRequest(req, ip.ProxyURL)
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

func (s *Service) DoAsync(ctx context.Context, req *networkAPI.DoAsyncReq) (*networkAPI.DoAsyncResp, error) {
	s.logger.Infof("new async: %v", len(req.Requests))
	if req.CallbackQueue == "" {
		err := errors.New(ctx, codes.FailedPrecondition).AddDetails("callback queue can not be empty")
		s.logger.WithError(err).Errorf("invalid callback: %v", req.CallbackQueue)
		return nil, err
	}
	if req.ReferenceID == "" {
		err := errors.New(ctx, codes.FailedPrecondition).AddDetails("reference id must be declared")
		s.logger.WithError(err).Errorf("invalid reference: %v", req.ReferenceID)
		return nil, err
	}
	if len(req.Requests) == 0 {
		err := errors.New(ctx, codes.FailedPrecondition).AddDetails("empty requests not acceptable")
		s.logger.WithError(err).Errorf("zero request")
		return nil, err
	}

	totalPredictedInterval, err := s.rateLimiterManager.AddNewBucket(ctx, req.CallbackQueue, req.ReferenceID, req.Requests)
	if err != nil {
		s.logger.WithError(err).Errorf("failed to do async request for %v", req)
		return nil, err
	}

	return &networkAPI.DoAsyncResp{PredictedIntervalTime: totalPredictedInterval}, nil
}
