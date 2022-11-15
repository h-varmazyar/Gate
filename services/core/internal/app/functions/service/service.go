package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	brokeragesService "github.com/h-varmazyar/Gate/services/core/internal/app/brokerages/service"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Service struct {
	requestService   networkAPI.RequestServiceClient
	brokerageService *brokeragesService.Service
	logger           *log.Logger
	configs          *Configs
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
		grpcService.configs = configs
	}
	return grpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	coreApi.RegisterFunctionsServiceServer(server, s)
}

func (s *Service) OHLC(ctx context.Context, req *coreApi.OHLCReq) (*chipmunkApi.Candles, error) {
	brokerage, err := s.loadBrokerage(ctx, req.BrokerageID)
	if err != nil {
		return nil, err
	}

	candles, err := loadRequest(s.configs, brokerage).OHLC(ctx, s.createOHLCParams(req),
		func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
			resp, err := s.requestService.Do(ctx, request)
			return resp, err
		})
	if err != nil {
		return nil, err
	}
	return &chipmunkApi.Candles{Elements: candles}, nil
}

func (s *Service) AsyncOHLC(ctx context.Context, req *coreApi.OHLCReq) (*api.Void, error) {
	brokerage, err := s.loadBrokerage(ctx, req.BrokerageID)
	if err != nil {
		return nil, err
	}

	request, err := loadRequest(s.configs, brokerage).AsyncOHLC(ctx, s.createOHLCParams(req))
	if err != nil {
		return nil, err
	}

	request.Type = networkAPI.Request_Async

	s.doAsyncRequest(request)

	return new(api.Void), nil
}

func (s *Service) WalletsBalance(ctx context.Context, _ *api.Void) (*chipmunkApi.Wallets, error) {
	brokerage, err := s.brokerageService.Enable(ctx, new(api.Void))
	if err != nil {
		return nil, err
	}
	wallets, err := loadRequest(s.configs, brokerage).WalletList(ctx, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
		resp, err := s.requestService.Do(ctx, request)
		return resp, err
	})
	if err != nil {
		return nil, err
	}
	return wallets, nil
}

func (s *Service) MarketStatistics(ctx context.Context, req *coreApi.MarketStatisticsReq) (*coreApi.MarketStatisticsResp, error) {
	brokerage, err := s.brokerageService.Enable(ctx, new(api.Void))
	if err != nil {
		return nil, err
	}
	params := &brokerages.MarketStatisticsParams{
		Market: req.MarketName,
	}
	statistics, err := loadRequest(s.configs, brokerage).MarketStatistics(ctx, params, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
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
	brokerage, err := s.brokerageService.Enable(ctx, new(api.Void))
	if err != nil {
		return nil, err
	}
	markets, err := loadRequest(s.configs, brokerage).MarketList(ctx, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
		resp, err := s.requestService.Do(ctx, request)
		return resp, err
	})
	if err != nil {
		return nil, err
	}
	return markets, nil
}

func (s *Service) NewOrder(ctx context.Context, req *coreApi.NewOrderReq) (*eagleApi.Order, error) {
	brokerage, err := s.brokerageService.Enable(ctx, new(api.Void))
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
	order, err := loadRequest(s.configs, brokerage).NewOrder(ctx, params, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
		resp, err := s.requestService.Do(ctx, request)
		return resp, err
	})
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *Service) CancelOrder(ctx context.Context, req *coreApi.CancelOrderReq) (*eagleApi.Order, error) {
	brokerage, err := s.brokerageService.Enable(ctx, new(api.Void))
	if err != nil {
		return nil, err
	}

	params := &brokerages.CancelOrderParams{
		ServerOrderId: req.ServerOrderID,
		Market:        req.Market,
	}
	order, err := loadRequest(s.configs, brokerage).CancelOrder(ctx, params, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
		resp, err := s.requestService.Do(ctx, request)
		return resp, err
	})
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *Service) OrderStatus(ctx context.Context, req *coreApi.OrderStatusReq) (*eagleApi.Order, error) {
	brokerage, err := s.brokerageService.Enable(ctx, new(api.Void))
	if err != nil {
		return nil, err
	}

	params := &brokerages.OrderStatusParams{
		ServerOrderId: req.ServerOrderID,
		Market:        req.Market,
	}
	order, err := loadRequest(s.configs, brokerage).OrderStatus(ctx, params, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
		resp, err := s.requestService.Do(ctx, request)
		return resp, err
	})
	if err != nil {
		return nil, err
	}
	return order, nil
}
