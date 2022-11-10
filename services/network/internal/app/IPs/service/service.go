package service

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/errors"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	"github.com/h-varmazyar/Gate/services/network/internal/app/IPs/repository"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Service struct {
	logger  *log.Logger
	configs *Configs
	db      repository.IPRepository
}

var (
	GrpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, configs *Configs, db repository.IPRepository) *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		GrpcService.logger = logger
		GrpcService.configs = configs
		GrpcService.db = db
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	networkAPI.RegisterIPServiceServer(server, s)
}

func (s *Service) Create(ctx context.Context, req *networkAPI.IPCreateReq) (*networkAPI.IP, error) {
	return nil, errors.New(ctx, codes.Unimplemented)
}

func (s *Service) Return(ctx context.Context, req *networkAPI.IPReturnReq) (*networkAPI.IP, error) {
	return nil, errors.New(ctx, codes.Unimplemented)
}

func (s *Service) List(ctx context.Context, req *networkAPI.IPListReq) (*networkAPI.IPs, error) {
	return nil, errors.New(ctx, codes.Unimplemented)
}
