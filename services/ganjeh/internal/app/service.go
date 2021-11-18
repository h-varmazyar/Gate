package app

import (
	"context"
	"github.com/mrNobody95/Gate/api"
	ganjehAPI "github.com/mrNobody95/Gate/services/ganjeh/api"
	"github.com/mrNobody95/Gate/services/ganjeh/internal/pkg/consul"
	"google.golang.org/grpc"
	"time"
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
	ganjehAPI.RegisterVariableServiceServer(server, s)
}

func (s *Service) Set(ctx context.Context, req *ganjehAPI.SetVariableRequest) (*ganjehAPI.Variable, error) {
	variable := &ganjehAPI.Variable{
		Key:       req.Key,
		Value:     req.Value,
		Namespace: req.Namespace,
		UpdatedAt: time.Now().Unix(),
	}
	if err := consul.Set(variable); err != nil {
		return nil, err
	}
	return variable, nil
}

func (s *Service) Get(ctx context.Context, req *ganjehAPI.GetVariableRequest) (*ganjehAPI.Variable, error) {
	return consul.Get(ctx, req.Namespace, req.Key)

}

func (s *Service) List(ctx context.Context, req *ganjehAPI.GetVariablesRequest) (*ganjehAPI.Variables, error) {
	return consul.GetList(ctx, req.Namespace)
}

func (s *Service) Delete(ctx context.Context, req *ganjehAPI.DeleteVariableRequest) (*api.Void, error) {
	err := consul.Delete(req.Namespace, req.Key)
	return new(api.Void), err
}
