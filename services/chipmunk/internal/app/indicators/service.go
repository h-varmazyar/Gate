package indicators

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
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
	chipmunkApi.RegisterIndicatorServiceServer(server, s)
}

func (s *Service) Return(_ context.Context, req *chipmunkApi.IndicatorReturnReq) (*chipmunkApi.Indicator, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}

	indicator, err := repository.Indicators.Return(id)
	if err != nil {
		return nil, err
	}
	response := new(chipmunkApi.Indicator)
	mapper.Struct(indicator, response)
	return response, nil
}
