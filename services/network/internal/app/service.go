package app

import (
	"context"
	networkAPI "github.com/mrNobody95/Gate/services/network/api"
	"github.com/mrNobody95/Gate/services/network/internal/pkg/requests"
	"google.golang.org/grpc"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 10.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

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
