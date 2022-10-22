package app

import (
	"context"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/requests"
	"google.golang.org/grpc"
)

type Service struct {
}

var (
	GrpcService *Service
)

func NewService() *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	networkAPI.RegisterRequestServiceServer(server, s)
}

func (s *Service) Do(ctx context.Context, req *networkAPI.Request) (*networkAPI.Response, error) {
	request := requests.New(req.Type, req.Endpoint)

	request.AddHeaders(req.Headers)
	switch req.Type {
	case networkAPI.Type_POST:
		request.AddBody(req.Params)
	case networkAPI.Type_GET:
		request.AddQueryParams(req.Params)
	}
	return request.Do()
}
