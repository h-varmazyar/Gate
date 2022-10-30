package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	coreApi "github.com/h-varmazyar/Gate/services/core/api"
	"github.com/h-varmazyar/Gate/services/core/internal/app/brokerages/repository"
	brokeragesService "github.com/h-varmazyar/Gate/services/core/internal/app/brokerages/service"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages/coinex"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/entity"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

type Service struct {
	requestService   networkAPI.RequestServiceClient
	brokerageService *brokeragesService.Service
	logger           *log.Logger
}

var (
	grpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, configs *Configs, brService *brokeragesService.Service) *Service {
	if grpcService == nil {
		grpcService = new(Service)
		networkConn := grpcext.NewConnection(configs.NetworkGrpcAddress)
		grpcService.requestService = networkAPI.NewRequestServiceClient(networkConn)
		grpcService.brokerageService = brService
		grpcService.logger = logger
	}
	return grpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	coreApi.RegisterFunctionsServiceServer(server, s)
}

func (s *Service) OHLC(ctx context.Context, req *coreApi.OHLCReq) (*chipmunkApi.Candles, error) {
	brokerageID, err := uuid.Parse(req.BrokerageID)
	if err != nil {
		return nil, err
	}
	brokerage, err := s.db.ReturnByID(brokerageID)
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
	brokerage, err := s.db.ReturnEnable()
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

func (s *Service) MarketStatistics(ctx context.Context, req *coreApi.MarketStatisticsReq) (*coreApi.MarketStatisticsResp, error) {
	brokerage, err := s.db.ReturnEnable()
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
	resp := new(coreApi.MarketStatisticsResp)
	mapper.Struct(statistics, resp)
	return resp, nil
}

func (s *Service) MarketList(ctx context.Context, req *coreApi.MarketListReq) (*chipmunkApi.Markets, error) {
	brokerage, err := s.db.ReturnEnable()
	if err != nil {
		return nil, err
	}
	markets, err := loadBrokerage(brokerage).MarketList(ctx, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
		resp, err := s.requestService.Do(ctx, request)
		return resp, err
	})
	if err != nil {
		return nil, err
	}
	return markets, nil
}

func (s *Service) NewOrder(ctx context.Context, req *coreApi.NewOrderReq) (*eagleApi.Order, error) {
	brokerage, err := s.db.ReturnEnable()
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

func (s *Service) CancelOrder(ctx context.Context, req *coreApi.CancelOrderReq) (*eagleApi.Order, error) {
	brokerage, err := s.db.ReturnEnable()
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

func (s *Service) OrderStatus(ctx context.Context, req *coreApi.OrderStatusReq) (*eagleApi.Order, error) {
	brokerage, err := s.db.ReturnEnable()
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

func loadBrokerage(brokerage *entity.Brokerage) brokerages.Brokerage {
	switch brokerage.Platform {
	case brokerageApi.Platform_Coinex:
		coinexInstance := new(coinex.Service)
		coinexInstance.Auth = &api.Auth{
			Type:      api.AuthType_StaticToken,
			AccessID:  brokerage.AccessID,
			SecretKey: brokerage.SecretKey,
		}
		return coinexInstance
	case brokerageApi.Platform_Nobitex:
		return nil
	}
	return nil
}
