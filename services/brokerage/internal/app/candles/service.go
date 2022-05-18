package candles

//import (
//	"context"
//	"fmt"
//	"github.com/h-varmazyar/Gate/api"
//	"github.com/h-varmazyar/Gate/pkg/grpcext"
//	"github.com/h-varmazyar/Gate/pkg/mapper"
//	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
//	"github.com/h-varmazyar/Gate/services/brokerage/configs"
//	"github.com/h-varmazyar/Gate/services/brokerage/internal/pkg/brokerages"
//	"github.com/h-varmazyar/Gate/services/brokerage/internal/pkg/brokerages/coinex"
//	"github.com/h-varmazyar/Gate/services/brokerage/internal/pkg/repository"
//	networkAPI "github.com/h-varmazyar/Gate/services/network/api"
//	"google.golang.org/grpc"
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
//	brokerageApi.RegisterCandleServiceServer(server, s)
//}
//
//func (s *Service) OHLC(ctx context.Context, req *brokerageApi.OhlcRequest) (*api.Candles, error) {
//	br := new(coinex.Service)
//	resolution := new(repository.Resolution)
//	mapper.Struct(req.Resolution, resolution)
//
//	market := new(repository.Market)
//	mapper.Struct(req.Market, market)
//	inputs := brokerages.OHLCParams{
//		Resolution: resolution,
//		Market:     market,
//		From:       time.Unix(req.From, 0),
//		To:         time.Unix(req.To, 0),
//	}
//	candles, err := br.OHLC(ctx, inputs, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
//		resp, err := s.networkService.Do(ctx, request)
//		return resp, err
//	})
//	if err != nil {
//		return nil, err
//	}
//	return &api.Candles{Candles: candles}, nil
//}
