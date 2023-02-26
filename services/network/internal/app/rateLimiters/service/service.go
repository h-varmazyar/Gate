package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	"github.com/h-varmazyar/Gate/services/network/internal/app/rateLimiters/repository"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Service struct {
	logger *log.Logger
	db     repository.RateLimiterRepository
}

var (
	GrpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, db repository.RateLimiterRepository) *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		GrpcService.logger = logger
		GrpcService.db = db
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	networkAPI.RegisterRateLimiterServiceServer(server, s)
}

func (s *Service) Create(_ context.Context, req *networkAPI.RateLimiterCreateReq) (*networkAPI.RateLimiter, error) {
	rateLimiter := new(entity.RateLimiter)
	mapper.Struct(req, rateLimiter)
	err := s.db.Create(rateLimiter)
	if err != nil {
		return nil, err
	}
	response := new(networkAPI.RateLimiter)
	mapper.Struct(rateLimiter, response)
	return response, nil
}

func (s *Service) Return(_ context.Context, req *networkAPI.RateLimiterReturnReq) (*networkAPI.RateLimiter, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}

	var rateLimiter *entity.RateLimiter
	rateLimiter, err = s.db.Return(id)
	if err != nil {
		return nil, err
	}
	response := new(networkAPI.RateLimiter)
	mapper.Struct(rateLimiter, response)
	return response, nil
}

func (s *Service) List(_ context.Context, req *networkAPI.RateLimiterListReq) (*networkAPI.RateLimiters, error) {
	rateLimiters, err := s.db.List(req.Type)
	if err != nil {
		return nil, err
	}
	response := new(networkAPI.RateLimiters)
	mapper.Slice(rateLimiters, &response.Elements)
	return response, nil
}
