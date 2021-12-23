package markets

import (
	"context"
	"github.com/google/uuid"
	"github.com/mrNobody95/Gate/pkg/mapper"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
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
	//brokerageID, err := uuid.Parse(req.BrokerageID)
	//if err != nil {
	//	return nil, err
	//}
	//market.BrokerageID = brokerageID
	//if req.Destination == nil {
	//	return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "destination not set")
	//}
	destinationID, err := uuid.Parse(req.Destination.ID)
	if err != nil {
		return nil, err
	}
	market.DestinationID = destinationID
	//if req.Source == nil {
	//	return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "source not set")
	//}
	sourceID, err := uuid.Parse(req.Source.ID)
	if err != nil {
		return nil, err
	}
	market.SourceID = sourceID
	market.Status = req.Status
	if err := repository.Markets.Update(market); err != nil {
		return nil, err
	}
	mapper.Struct(market, req)
	req.ID = market.ID.String()
	return req, nil
}

func (s *Service) Get(_ context.Context, req *brokerageApi.MarketRequest) (*brokerageApi.Market, error) {
	info, err := repository.Markets.Info(req.BrokerageName, req.MarketName)
	if err != nil {
		return nil, err
	}
	response := new(brokerageApi.Market)
	mapper.Struct(info, response)
	return response, nil
}

func (s *Service) List(_ context.Context, req *brokerageApi.MarketListRequest) (*brokerageApi.Markets, error) {
	response, err := repository.Markets.List(req.BrokerageName)
	if err != nil {
		return nil, err
	}
	markets := new(brokerageApi.Markets)
	mapper.Slice(response, &markets.Markets)
	return markets, nil
}
