package functions

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	"github.com/h-varmazyar/Gate/services/brokerage/configs"
	"github.com/h-varmazyar/Gate/services/brokerage/internal/pkg/brokerages"
	"github.com/h-varmazyar/Gate/services/brokerage/internal/pkg/brokerages/coinex"
	"github.com/h-varmazyar/Gate/services/brokerage/internal/pkg/repository"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api"
	"google.golang.org/grpc"
	"time"
)

type Service struct {
	requestService networkAPI.RequestServiceClient
}

var (
	GrpcService *Service
)

func NewService() *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		networkConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Network)
		GrpcService.requestService = networkAPI.NewRequestServiceClient(networkConn)
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	brokerageApi.RegisterFunctionsServiceServer(server, s)
}

func (s *Service) OHLC(ctx context.Context, req *brokerageApi.OHLCReq) (*chipmunkApi.Candles, error) {
	brokerage, err := repository.Brokerages.ReturnEnable()
	if err != nil {
		return nil, err
	}
	resolution := new(repository.Resolution)
	mapper.Struct(req.Resolution, resolution)

	market := new(repository.Market)
	mapper.Struct(req.Market, market)
	inputs := &brokerages.OHLCParams{
		Resolution: resolution,
		Market:     market,
		From:       time.Unix(req.From, 0),
		To:         time.Unix(req.To, 0),
	}

	candles, err := loadBrokerage(brokerage).OHLC(ctx, inputs, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
		resp, err := s.requestService.Do(ctx, request)
		return resp, err
	})
	if err != nil {
		return nil, err
	}
	return &chipmunkApi.Candles{Elements: candles}, nil
}

func (s *Service) WalletsBalance(ctx context.Context, _ *api.Void) (*chipmunkApi.Wallets, error) {
	brokerage, err := repository.Brokerages.ReturnEnable()
	if err != nil {
		return nil, err
	}
	wallets, err := loadBrokerage(brokerage).WalletList(ctx, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
		resp, err := s.requestService.Do(ctx, request)
		return resp, err
	})
	if err != nil {
		return nil, err
	}
	return wallets, nil
}

func (s *Service) MarketStatistics(ctx context.Context, req *brokerageApi.MarketStatisticsReq) (*brokerageApi.MarketStatisticsResp, error) {
	brokerage, err := repository.Brokerages.ReturnEnable()
	if err != nil {
		return nil, err
	}
	params := &brokerages.MarketStatisticsParams{
		Market: req.MarketName,
	}
	statistics, err := loadBrokerage(brokerage).MarketStatistics(ctx, params, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
		resp, err := s.requestService.Do(ctx, request)
		return resp, err
	})
	if err != nil {
		return nil, err
	}
	resp := new(brokerageApi.MarketStatisticsResp)
	mapper.Struct(statistics, resp)
	return resp, nil
}

func (s *Service) NewOrder(ctx context.Context, req *brokerageApi.NewOrderReq) (*eagleApi.Order, error) {
	brokerage, err := repository.Brokerages.ReturnEnable()
	if err != nil {
		return nil, err
	}

	params := &brokerages.NewOrderParams{
		OrderModel: req.Model,
		ClientUUID: uuid.New().String(),
		BuyOrSell:  req.Type,
		Price:      req.Price,
		StopPrice:  req.StopPrice,
		Market:     req.Market,
		Amount:     req.Amount,
		Option:     req.Option,
		HideOrder:  req.ISHidden,
	}
	order, err := loadBrokerage(brokerage).NewOrder(ctx, params, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
		resp, err := s.requestService.Do(ctx, request)
		return resp, err
	})
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *Service) CancelOrder(ctx context.Context, req *brokerageApi.CancelOrderReq) (*eagleApi.Order, error) {
	brokerage, err := repository.Brokerages.ReturnEnable()
	if err != nil {
		return nil, err
	}

	params := &brokerages.CancelOrderParams{
		ServerOrderId: req.ServerOrderID,
		Market:        req.Market,
	}
	order, err := loadBrokerage(brokerage).CancelOrder(ctx, params, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
		resp, err := s.requestService.Do(ctx, request)
		return resp, err
	})
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *Service) OrderStatus(ctx context.Context, req *brokerageApi.OrderStatusReq) (*eagleApi.Order, error) {
	brokerage, err := repository.Brokerages.ReturnEnable()
	if err != nil {
		return nil, err
	}

	params := &brokerages.OrderStatusParams{
		ServerOrderId: req.ServerOrderID,
		Market:        req.Market,
	}
	order, err := loadBrokerage(brokerage).OrderStatus(ctx, params, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
		resp, err := s.requestService.Do(ctx, request)
		return resp, err
	})
	if err != nil {
		return nil, err
	}
	return order, nil
}

func loadBrokerage(brokerage *repository.Brokerage) brokerages.Brokerage {
	switch brokerage.Platform {
	case brokerageApi.Platform_Coinex:
		return new(coinex.Service)
	case brokerageApi.Platform_Nobitex:
		return nil
	}
	return nil
}
