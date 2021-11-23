package assets

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/errors"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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

func (s *Service) Add(ctx context.Context, brokerage *brokerageApi.Brokerage) (*brokerageApi.Brokerage, error) {
	br := new(repository.Brokerage)
	name := brokerage.Name.String()
	if name == fmt.Sprintf("%d", brokerage.Name) || name == brokerageApi.Names_UnknownBrokerage.String() {
		return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "invalid brokerage")
	}
	br.Name = repository.BrokerageName(brokerage.Name.String())
	if brokerage.Auth == nil {
		return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "authentication not set")
	}
	br.AuthType = brokerage.Auth.Type.String()
	switch br.AuthType {
	case api.AuthType_StaticToken.String():
		if brokerage.Auth.AccessID == "" {
			return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "access id not set")
		}
		br.AccessID = brokerage.Auth.AccessID
		if brokerage.Auth.SecretKey == "" {
			return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "secret key not set")
		}
		br.SecretKey = brokerage.Auth.SecretKey
	case api.AuthType_UsernamePassword.String():
		if brokerage.Auth.Username == "" {
			return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "username not set")
		}
		br.Username = brokerage.Auth.Username
		if brokerage.Auth.Password == "" {
			return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "password not set")
		}
		br.Password = brokerage.Auth.Password
	default:
		return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "invalid auth type")
	}
	if brokerage.Status == api.StatusType_Enable || brokerage.Status == api.StatusType_Disable {
		br.Status = brokerage.Status.String()
	} else {
		br.Status = api.StatusType_Disable.String()
	}
	if brokerage.ID == "" {
		br.ID = uuid.New()
	} else {
		var err error
		br.ID, err = uuid.Parse(brokerage.ID)
		if err != nil {
			return nil, err
		}
	}

	if err := repository.Brokerages.Create(br); err != nil {
		return nil, err
	}
	brokerage.ID = br.ID.String()
	return brokerage, nil
}

func (s *Service) Get(_ context.Context, req *brokerageApi.BrokerageIDReq) (*brokerageApi.GetBrokerage, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}
	brokerage, err := repository.Brokerages.ReturnByID(id)
	if err != nil {
		return nil, err
	}

	return &brokerageApi.GetBrokerage{
		ID:       brokerage.ID.String(),
		AuthType: brokerage.AuthType,
		Name:     string(brokerage.Name),
		Status:   brokerage.Status,
	}, err
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
		Status: api.StatusType(api.StatusType_value[brokerage.Status]),
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

func (s *Service) ChangeStatus(_ context.Context, req *api.StatusChangeRequest) (*api.Status, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}
	brokerage, err := repository.Brokerages.ReturnByID(id)
	if err != nil {
		return nil, err
	}
	switch brokerage.Status {
	case api.StatusType_Enable.String():
		brokerage.Status = api.StatusType_Disable.String()
	case api.StatusType_Disable.String():
		brokerage.Status = api.StatusType_Enable.String()
	}
	if err := repository.Brokerages.Update(brokerage); err != nil {
		return nil, err
	}
	return &api.Status{Status: api.StatusType(api.StatusType_value[brokerage.Status])}, nil
}
