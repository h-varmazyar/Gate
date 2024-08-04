package service

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/errors"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Service struct {
	logger *log.Logger
}

var (
	GrpcService *Service
)

type Dependencies struct {
}

func NewService(_ context.Context, logger *log.Logger, configs *Configs, dependencies *Dependencies) *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		GrpcService.logger = logger
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	chipmunkApi.RegisterPostServiceServer(server, s)
}

func (s *Service) List(ctx context.Context, req *chipmunkApi.PostListReq) (*chipmunkApi.Posts, error) {
	return nil, errors.New(ctx, codes.Unimplemented)
}
