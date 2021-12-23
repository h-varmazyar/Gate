package brokerages

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/errors"
	"github.com/mrNobody95/Gate/pkg/mapper"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"strconv"
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
* Date: 12.11.21
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
	brokerageApi.RegisterBrokerageServiceServer(server, s)
}

func (s *Service) Create(ctx context.Context, brokerage *brokerageApi.CreateBrokerageReq) (*brokerageApi.Brokerage, error) {
	//todo: validation on auth and other fields
	br := new(repository.Brokerage)
	if _, ok := brokerageApi.Names_value[brokerage.Name.String()]; !ok {
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "wrong_name")
	} else {
		br.Name = brokerage.Name.String()
	}
	if _, ok := api.AuthType_value[brokerage.Auth.Type.String()]; !ok {
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "wrong_auth_type")
	} else {
		br.AuthType = brokerage.Auth.Type.String()
	}
	br.Status = api.Status_Disable.String()
	br.ID = uuid.New()
	fmt.Println("before res")
	resID, err := strconv.Atoi(brokerage.ResolutionID)
	if err != nil {
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "invalid_resolution")
	}
	br.ResolutionID = uint(resID)
	mapper.Struct(brokerage, br)

	for _, market := range brokerage.Markets.Markets {
		id, err := uuid.Parse(market.ID)
		if err != nil {
			return nil, err
		}
		tmp := new(repository.Market)
		tmp.ID = id
		br.Markets = append(br.Markets, tmp)
	}
	if err := repository.Brokerages.Create(br); err != nil {
		return nil, err
	}
	response := new(brokerageApi.Brokerage)
	mapper.Struct(br, response)
	return response, nil
}

func (s *Service) List(_ context.Context, _ *api.Void) (*brokerageApi.Brokerages, error) {
	bb, err := repository.Brokerages.List()
	if err != nil {
		return nil, err
	}

	response := new(brokerageApi.Brokerages)
	mapper.Slice(bb, &response.Brokerages)
	return response, err
}

func (s *Service) Get(_ context.Context, req *brokerageApi.BrokerageIDReq) (*brokerageApi.Brokerage, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}
	brokerage, err := repository.Brokerages.ReturnByID(id)
	if err != nil {
		return nil, err
	}
	response := new(brokerageApi.Brokerage)
	response.Name = brokerageApi.Names(brokerageApi.Names_value[brokerage.Name])
	mapper.Struct(brokerage, response)
	mapper.Slice(brokerage.Markets, &response.Markets.Markets)
	return response, err
}

func (s *Service) GetInternal(_ context.Context, req *brokerageApi.BrokerageIDReq) (*brokerageApi.Brokerage, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}
	brokerage, err := repository.Brokerages.ReturnByID(id)
	if err != nil {
		return nil, err
	}

	return &brokerageApi.Brokerage{
		ID: brokerage.ID.String(),
		Auth: &api.Auth{
			Type:      api.AuthType(api.AuthType_value[brokerage.AuthType]),
			Username:  brokerage.Username,
			Password:  brokerage.Password,
			AccessID:  brokerage.AccessID,
			SecretKey: brokerage.SecretKey,
		},
		Name:   brokerageApi.Names(brokerageApi.Names_value[string(brokerage.Name)]),
		Status: api.Status(api.Status_value[brokerage.Status]),
	}, err
}

func (s *Service) Delete(_ context.Context, req *brokerageApi.BrokerageIDReq) (*api.Void, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}
	if err := repository.Brokerages.Delete(id); err != nil {
		return nil, err
	}
	return new(api.Void), err
}

func (s *Service) ChangeStatus(_ context.Context, req *api.StatusChangeRequest) (*brokerageApi.BrokerageStatus, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}
	brokerage, err := repository.Brokerages.ReturnByID(id)
	if err != nil {
		return nil, err
	}
	switch brokerage.Status {
	case api.Status_Enable.String():
		brokerage.Status = api.Status_Disable.String()
	case api.Status_Disable.String():
		brokerage.Status = api.Status_Enable.String()
	}
	if err := repository.Brokerages.Update(brokerage); err != nil {
		return nil, err
	}
	return &brokerageApi.BrokerageStatus{Status: api.Status(api.Status_value[brokerage.Status])}, nil
}
