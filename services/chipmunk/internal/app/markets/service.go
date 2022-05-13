package markets

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Service struct {
	networkService networkAPI.RequestServiceClient
}

var (
	GrpcService *Service
)

func NewService() *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		networkConnection := grpcext.NewConnection(configs.Variables.GrpcAddresses.Network)
		GrpcService.networkService = networkAPI.NewRequestServiceClient(networkConnection)
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	chipmunkApi.RegisterMarketServiceServer(server, s)
}

func (s *Service) Set(ctx context.Context, req *chipmunkApi.Market) (*chipmunkApi.Market, error) {
	market := new(repository.Market)
	mapper.Struct(req, market)
	var err error
	if req.Destination == nil {
		return nil, errors.New(ctx, codes.FailedPrecondition).AddDetailF("destination not found")
	}
	market.DestinationID, err = uuid.Parse(req.Destination.ID)
	if err != nil {
		return nil, errors.Cast(ctx, err).AddDetailF("invalid destination id %v", req.Destination.ID)
	}
	if req.Source == nil {
		return nil, errors.New(ctx, codes.FailedPrecondition).AddDetailF("source not found")
	}
	market.SourceID, err = uuid.Parse(req.Source.ID)
	if err != nil {
		return nil, errors.Cast(ctx, err).AddDetailF("invalid source id %v", req.Source.ID)
	}
	market.Status = req.Status
	if err := repository.Markets.SaveOrUpdate(market); err != nil {
		return nil, err
	}
	mapper.Struct(market, req)
	req.ID = market.ID.String()
	return req, nil
}

func (s *Service) Return(_ context.Context, req *chipmunkApi.ReturnMarketRequest) (*chipmunkApi.Market, error) {
	marketID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}
	market, err := repository.Markets.ReturnByID(marketID)
	if err != nil {
		return nil, err
	}
	response := new(chipmunkApi.Market)
	mapper.Struct(market, response)
	return response, nil
}

func (s *Service) Get(_ context.Context, req *chipmunkApi.MarketRequest) (*chipmunkApi.Market, error) {
	brID, err := uuid.Parse(req.BrokerageID)
	if err != nil {
		return nil, err
	}
	info, err := repository.Markets.Info(brID, req.MarketName)
	if err != nil {
		return nil, err
	}
	response := new(chipmunkApi.Market)
	mapper.Struct(info, response)
	return response, nil
}

func (s *Service) List(_ context.Context, req *chipmunkApi.MarketListRequest) (*chipmunkApi.Markets, error) {
	brID, err := uuid.Parse(req.BrokerageID)
	if err != nil {
		return nil, err
	}
	response, err := repository.Markets.List(brID)
	if err != nil {
		return nil, err
	}
	markets := new(chipmunkApi.Markets)
	mapper.Slice(response, &markets.Elements)
	return markets, nil
}

func (s *Service) UpdateMarkets(ctx context.Context, req *chipmunkApi.UpdateMarketsReq) (*api.Void, error) {
	//switch req.BrokerageName {
	//case brokerageApi.Names_Coinex.String():
	//	br := new(coinex.Service)
	//	markets, err := br.UpdateMarket(ctx, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
	//		resp, err := s.networkService.Do(ctx, request)
	//		return resp, err
	//	})
	//	if err != nil {
	//		return nil, err
	//	}
	//	for _, markets := range markets {
	//		source, err := repository.Assets.ReturnBySymbol(markets.SourceName)
	//		if err != nil {
	//			if err == gorm.ErrRecordNotFound {
	//				assets := new(repository.Asset)
	//				assets.Name = markets.SourceName
	//				assets.Symbol = markets.SourceName
	//				assets.IssueDate = time.Now()
	//				assets, err = repository.Assets.Create(assets)
	//				if err != nil {
	//					log.WithError(err).WithField("source_name", markets.SourceName).Error("failed to create assets")
	//					continue
	//				}
	//			}
	//			log.WithError(err).WithField("source_name", markets.SourceName).Error("failed to get assets")
	//			continue
	//		}
	//		markets.SourceID = source.ID
	//		destination, err := repository.Assets.ReturnBySymbol(markets.DestinationName)
	//		if err != nil {
	//			if err == gorm.ErrRecordNotFound {
	//				assets := new(repository.Asset)
	//				assets.Name = markets.DestinationName
	//				assets.Symbol = markets.DestinationName
	//				assets.IssueDate = time.Now()
	//				assets, err = repository.Assets.Create(assets)
	//				if err != nil {
	//					log.WithError(err).WithField("destination_name", markets.DestinationName).Error("failed to create assets")
	//					continue
	//				}
	//			}
	//			log.WithError(err).WithField("destination_name", markets.DestinationName).Error("failed to get assets")
	//			continue
	//		}
	//		markets.DestinationID = destination.ID
	//		err = repository.Markets.SaveOrUpdate(markets)
	//		if err != nil {
	//			log.WithError(err).Error("failed to update markets")
	//			continue
	//		}
	//	}
	//case chipmunkApi.nobi.String():
	//	return nil, errors.NewWithSlug(ctx, codes.Unimplemented, "not_supported")
	//default:
	//	return nil, errors.NewWithSlug(ctx, codes.Unimplemented, "brokerage_not_supported")
	//}
	//return new(api.Void), nil
	return nil, errors.New(ctx, codes.Unimplemented)
}

func (s *Service) ReturnBySource(_ context.Context, req *chipmunkApi.MarketListBySourceRequest) (*chipmunkApi.Markets, error) {
	brID, err := uuid.Parse(req.BrokerageID)
	if err != nil {
		return nil, err
	}
	response := new(chipmunkApi.Markets)
	if markets, err := repository.Markets.ListBySource(brID, req.Source); err != nil {
		return nil, err
	} else {
		list := make([]*chipmunkApi.Market, 0)
		mapper.Slice(markets, &list)
		response.Elements = list
		return response, nil
	}
}
