package assets

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/grpcext"
	"github.com/mrNobody95/Gate/pkg/mapper"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/brokerage/configs"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/repository"
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
* Date: 12.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type Service struct {
	BrokeragesService brokerageApi.BrokerageServiceClient
}

var (
	GrpcService *Service
)

func NewService(configs *configs.Configs) *Service {
	if GrpcService == nil {
		connection := grpcext.NewConnection(fmt.Sprintf("%d", configs.GrpcPort))
		GrpcService = new(Service)
		GrpcService.BrokeragesService = brokerageApi.NewBrokerageServiceClient(connection)
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	brokerageApi.RegisterMarketServiceServer(server, s)
}

func (s *Service) Set(_ context.Context, req *brokerageApi.Market) (*brokerageApi.Market, error) {
	market := new(repository.Market)
	mapper.Struct(req, market)
	if req.ID == "" {
		market.ID = uuid.New()
	} else {
		id, err := uuid.Parse(req.ID)
		if err != nil {
			return nil, err
		}
		market.ID = id
	}
	brokerageID, err := uuid.Parse(req.BrokerageID)
	if err != nil {
		return nil, err
	}
	market.BrokerageID = brokerageID
	//if req.Destination == nil {
	//	return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "destination not set")
	//}
	destinationID, err := uuid.Parse(req.DestinationID)
	if err != nil {
		return nil, err
	}
	market.DestinationID = destinationID
	//if req.Source == nil {
	//	return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "source not set")
	//}
	sourceID, err := uuid.Parse(req.SourceID)
	if err != nil {
		return nil, err
	}
	market.SourceID = sourceID
	if req.Status == api.StatusType_Enable || req.Status == api.StatusType_Disable {
		market.Status = req.Status.String()
	} else {
		market.Status = api.StatusType_Disable.String()
	}
	if err := repository.Markets.Update(market); err != nil {
		return nil, err
	}
	mapper.Struct(market, req)
	req.ID = market.ID.String()
	return req, nil
}

func (s *Service) Get(_ context.Context, req *brokerageApi.MarketRequest) (*brokerageApi.Market, error) {
	//brokerage, err := s.BrokeragesService.Get(ctx, &brokerageApi.BrokerageIDReq{ID: req.BrokerageID})
	//if err != nil {
	//	return nil, err
	//}
	info, err := repository.Markets.Info(req.BrokerageID, req.MarketName)
	if err != nil {
		return nil, err
	}
	response := new(brokerageApi.Market)
	//if req.WithUpdate {
	//	var br Interface.Brokerage
	//	br = new(Interface.Coinex)
	//	info, err = br.MarketInfo(func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
	//		return s.NetworkService.Do(ctx, request)
	//	})
	//	if err != nil {
	//		return nil, err
	//	}
	//	mapper.Struct(response, info)
	//	info.BrokerageID = req.BrokerageID
	//}
	mapper.Struct(info, response)
	switch info.Status {
	case api.StatusType_Disable.String():
		response.Status = api.StatusType_Disable
	case api.StatusType_Enable.String():
		response.Status = api.StatusType_Enable
	default:
		response.Status = api.StatusType_UnknownStatus
	}
	return response, nil
}

func (s *Service) List(_ context.Context, req *brokerageApi.MarketListRequest) (*brokerageApi.Markets, error) {
	response, err := repository.Markets.List(req.BrokerageID)
	if err != nil {
		return nil, err
	}
	markets := new(brokerageApi.Markets)
	mapper.Slice(response, &markets.Markets)
	for i, market := range response {
		switch market.Status {
		case api.StatusType_Disable.String():
			markets.Markets[i].Status = api.StatusType_Disable
		case api.StatusType_Enable.String():
			markets.Markets[i].Status = api.StatusType_Enable
		default:
			markets.Markets[i].Status = api.StatusType_UnknownStatus
		}
	}
	return markets, nil
}

func (s *Service) ChangeStatus(_ context.Context, req *api.StatusChangeRequest) (*api.Status, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}
	market, err := repository.Markets.ReturnByID(id)
	if err != nil {
		return nil, err
	}
	switch market.Status {
	case api.StatusType_Disable.String():
		market.Status = api.StatusType_Enable.String()
	default:
		market.Status = api.StatusType_Disable.String()
	}
	if err := repository.Markets.Update(market); err != nil {
		return nil, err
	}
	return &api.Status{Status: api.StatusType(api.StatusType_value[market.Status])}, nil
}
