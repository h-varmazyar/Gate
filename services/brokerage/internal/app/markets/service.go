package markets

//import (
//	"context"
//	"fmt"
//	"github.com/h-varmazyar/Gate/api"
//	"github.com/h-varmazyar/Gate/pkg/errors"
//	"github.com/h-varmazyar/Gate/pkg/grpcext"
//	"github.com/h-varmazyar/Gate/pkg/mapper"
//	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
//	"github.com/h-varmazyar/Gate/services/brokerage/configs"
//	"github.com/h-varmazyar/Gate/services/brokerage/internal/pkg/brokerages"
//	"github.com/h-varmazyar/Gate/services/brokerage/internal/pkg/brokerages/coinex"
//	"github.com/h-varmazyar/Gate/services/brokerage/internal/pkg/repository"
//	networkAPI "github.com/h-varmazyar/Gate/services/network/api"
//	log "github.com/sirupsen/logrus"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/codes"
//	"gorm.io/gorm"
//	"time"
//)
//
//type Service struct {
//	networkService networkAPI.RequestServiceClient
//}
//
//var (
//	GrpcService *Service
//)
//
//func NewService(configs *configs.Configs) *Service {
//	if GrpcService == nil {
//		GrpcService = new(Service)
//		networkConnection := grpcext.NewConnection(fmt.Sprintf(":%d", configs.NetworkGrpcPort))
//		GrpcService.networkService = networkAPI.NewRequestServiceClient(networkConnection)
//	}
//	return GrpcService
//}
//
//func (s *Service) RegisterServer(server *grpc.Server) {
//	brokerageApi.RegisterMarketServiceServer(server, s)
//}
//
//func (s *Service) Set(_ context.Context, req *brokerageApi.Market) (*brokerageApi.Market, error) {
//	market := new(repository.Market)
//	mapper.Struct(req, market)
//	market.DestinationID = uint(req.Destination.ID)
//	market.SourceID = uint(req.Source.ID)
//	market.Status = req.Status
//	if err := repository.Markets.SaveOrUpdate(market); err != nil {
//		return nil, err
//	}
//	mapper.Struct(market, req)
//	req.ID = uint32(market.ID)
//	return req, nil
//}
//
//func (s *Service) Get(_ context.Context, req *brokerageApi.MarketRequest) (*brokerageApi.Market, error) {
//	info, err := repository.Markets.Info(req.BrokerageName, req.MarketName)
//	if err != nil {
//		return nil, err
//	}
//	response := new(brokerageApi.Market)
//	mapper.Struct(info, response)
//	return response, nil
//}
//
//func (s *Service) List(_ context.Context, req *brokerageApi.MarketListRequest) (*brokerageApi.Markets, error) {
//	response, err := repository.Markets.List(req.BrokerageName)
//	if err != nil {
//		return nil, err
//	}
//	markets := new(brokerageApi.Markets)
//	mapper.Slice(response, &markets.Markets)
//	return markets, nil
//}
//
//func (s *Service) UpdateMarkets(ctx context.Context, req *brokerageApi.UpdateMarketsReq) (*api.Void, error) {
//	switch req.BrokerageName {
//	case brokerageApi.Names_Coinex.String():
//		br := new(coinex.Service)
//		markets, err := br.UpdateMarket(ctx, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
//			resp, err := s.networkService.Do(ctx, request)
//			return resp, err
//		})
//		if err != nil {
//			return nil, err
//		}
//		for _, market := range markets {
//			source, err := repository.Assets.ReturnBySymbol(market.SourceName)
//			if err != nil {
//				if err == gorm.ErrRecordNotFound {
//					asset := new(repository.Asset)
//					asset.Name = market.SourceName
//					asset.Symbol = market.SourceName
//					asset.IssueDate = time.Now()
//					asset, err = repository.Assets.Create(asset)
//					if err != nil {
//						log.WithError(err).WithField("source_name", market.SourceName).Error("failed to create assets")
//						continue
//					}
//				}
//				log.WithError(err).WithField("source_name", market.SourceName).Error("failed to get assets")
//				continue
//			}
//			market.SourceID = source.ID
//			destination, err := repository.Assets.ReturnBySymbol(market.DestinationName)
//			if err != nil {
//				if err == gorm.ErrRecordNotFound {
//					asset := new(repository.Asset)
//					asset.Name = market.DestinationName
//					asset.Symbol = market.DestinationName
//					asset.IssueDate = time.Now()
//					asset, err = repository.Assets.Create(asset)
//					if err != nil {
//						log.WithError(err).WithField("destination_name", market.DestinationName).Error("failed to create assets")
//						continue
//					}
//				}
//				log.WithError(err).WithField("destination_name", market.DestinationName).Error("failed to get assets")
//				continue
//			}
//			market.DestinationID = destination.ID
//			err = repository.Markets.SaveOrUpdate(market)
//			if err != nil {
//				log.WithError(err).Error("failed to update markets")
//				continue
//			}
//		}
//	case brokerageApi.Names_Nobitex.String():
//		return nil, errors.NewWithSlug(ctx, codes.Unimplemented, "not_supported")
//	default:
//		return nil, errors.NewWithSlug(ctx, codes.Unimplemented, "brokerage_not_supported")
//	}
//	return new(api.Void), nil
//}
//
//func (s *Service) MarketStatistics(ctx context.Context, req *brokerageApi.MarketStatisticsRequest) (*api.Candle, error) {
//	br := new(coinex.Service)
//	params := brokerages.MarketStatisticsParams{
//		Market: req.MarketName,
//	}
//	return br.MarketStatistics(ctx, params, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
//		resp, err := s.networkService.Do(ctx, request)
//		return resp, err
//	})
//}
//
//func (s *Service) ReturnBySource(_ context.Context, req *brokerageApi.MarketListBySourceRequest) (*brokerageApi.Markets, error) {
//	response := new(brokerageApi.Markets)
//	if markets, err := repository.Markets.ListBySource(req.BrokerageName, req.Source); err != nil {
//		return nil, err
//	} else {
//		list := make([]*brokerageApi.Market, 0)
//		mapper.Slice(markets, &list)
//		response.Markets = list
//		return response, nil
//	}
//}
