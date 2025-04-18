package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	brokeragesService "github.com/h-varmazyar/Gate/services/core/internal/app/brokerages/service"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages/coinex"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api/proto"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
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

func (s *Service) AsyncOHLC(ctx context.Context, req *coreApi.AsyncOHLCReq) (*coreApi.AsyncOHLCResp, error) {
	var (
		err         error
		request     *networkAPI.Request
		referenceID = uuid.New()
		requests    = make([]*networkAPI.Request, 0)
		resp        *networkAPI.DoAsyncResp
	)

	for _, item := range req.Items {
		brokerageRequests := loadRequest(s.configs, &coreApi.Brokerage{Platform: req.Platform})

		from := time.Unix(item.From, 0)

		for end := false; !end; {
			to := from.Add(time.Duration(item.Resolution.Duration) * 999)
			if to.After(time.Now()) {
				to = time.Now()
				end = true
			}

			if to.Sub(from) < time.Duration(item.Resolution.Duration) {
				continue
			}

			params := &brokerages.OHLCParams{
				Resolution: item.Resolution,
				Market:     item.Market,
				From:       from,
				To:         to,
			}

			request, err = brokerageRequests.AsyncOHLC(ctx, params)
			if err != nil {
				s.logger.WithError(err).Error("failed to create async OHLC")
				return nil, err
			}
			request.Timeout = item.Timeout
			request.IssueTime = item.IssueTime
			request.ReferenceID = uuid.New().String()
			from = to.Add(time.Duration(item.Resolution.Duration))
			requests = append(requests, request)
		}
	}

	callbackQueue := ""
	switch req.Platform {
	case api.Platform_Coinex:
		callbackQueue = coinex.QueueOHLC
	}

	resp, err = s.requestService.DoAsync(ctx, &networkAPI.DoAsyncReq{
		Requests:      requests,
		CallbackQueue: callbackQueue,
		ReferenceID:   referenceID.String(),
	})
	if err != nil {
		s.logger.WithError(err).Error("failed to do async ohlc")
		return nil, err
	}

	return &coreApi.AsyncOHLCResp{
		LastRequestID:         referenceID.String(),
		PredictedIntervalTime: resp.PredictedIntervalTime,
	}, nil
}

func (s *Service) AllMarketStatistics(ctx context.Context, req *coreApi.AllMarketStatisticsReq) (*coreApi.AllMarketStatisticsResp, error) {
	brokerage := &coreApi.Brokerage{Platform: req.Platform}
	br := loadRequest(s.configs, brokerage)
	if br == nil {
		s.logger.Fatalf("nil br: %v", req.Platform)
	}
	request, err := br.AllMarketStatistics(ctx, new(brokerages.AllMarketStatisticsParams))
	if err != nil {
		return nil, err
	}

	networkResponse, err := s.doNetworkRequest(request)
	if err != nil {
		return nil, err
	}

	responseParser, err := loadResponse(s.configs, brokerage)
	if err != nil {
		s.logger.WithError(err).Error("failed to load response parser in all market statistics")
		return nil, err
	}

	response, err := responseParser.AllMarkerStatistics(ctx, networkResponse)
	if err != nil {
		s.logger.WithError(err).Error("failed to parse all market statistics response")
		return nil, err
	}

	return response, nil
}

func (s *Service) GetMarketInfo(ctx context.Context, req *coreApi.MarketInfoReq) (*coreApi.MarketInfo, error) {
	brokerage := &coreApi.Brokerage{Platform: req.Market.Platform}
	fmt.Println(brokerage)
	params := s.createMarketInfoParams(req)
	fmt.Println("params:", params)
	request, err := loadRequest(s.configs, brokerage).GetMarketInfo(ctx, params)
	if err != nil {
		return nil, err
	}

	networkResponse, err := s.doNetworkRequest(request)
	if err != nil {
		return nil, err
	}

	responseParser, err := loadResponse(s.configs, brokerage)
	if err != nil {
		s.logger.WithError(err).Error("failed to load response parser in all market statistics")
		return nil, err
	}

	response, err := responseParser.GetMarketInfo(ctx, networkResponse)
	if err != nil {
		s.logger.WithError(err).Error("failed to parse all market statistics response")
		return nil, err
	}

	return response, nil
}

func (s *Service) OHLC(ctx context.Context, req *coreApi.OHLCReq) (*chipmunkApi.Candles, error) {
	brokerage := &coreApi.Brokerage{Platform: req.Platform}
	from := time.Unix(req.Item.From, 0)

	allCandles := make([]*chipmunkApi.Candle, 0)

	for end := false; !end; {
		to := from.Add(time.Duration(req.Item.Resolution.Duration) * 999)
		if to.After(time.Unix(req.Item.To, 0)) {
			to = time.Unix(req.Item.To, 0)
			end = true
		} else if to.After(time.Now()) {
			to = time.Now()
			end = true
		}

		params := &brokerages.OHLCParams{
			Resolution: req.Item.Resolution,
			Market:     req.Item.Market,
			From:       from,
			To:         to,
		}
		candles, err := loadRequest(s.configs, brokerage).OHLC(ctx, params,
			func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
				resp, err := s.requestService.Do(ctx, request)
				return resp, err
			})
		if err != nil {
			s.logger.WithError(err).Error("failed to do OHLC")
			continue
		}
		from = to.Add(time.Duration(req.Item.Resolution.Duration))
		allCandles = append(allCandles, candles...)
	}
	return &chipmunkApi.Candles{Elements: allCandles}, nil
}

func (s *Service) WalletsBalance(ctx context.Context, req *coreApi.WalletsBalanceReq) (*chipmunkApi.Wallets, error) {
	brokerage, err := s.loadBrokerage(ctx, req.BrokerageID)
	if err != nil {
		return nil, err
	}
	request, err := loadRequest(s.configs, brokerage).WalletsBalance(ctx, nil)
	if err != nil {
		return nil, err
	}

	networkResponse, err := s.doNetworkRequest(request)
	if err != nil {
		return nil, err
	}

	responseParser, err := loadResponse(s.configs, brokerage)
	if err != nil {
		s.logger.WithError(err).Error("failed to load response parser in all market statistics")
		return nil, err
	}

	response, err := responseParser.WalletsBalance(ctx, networkResponse)
	if err != nil {
		s.logger.WithError(err).Error("failed to parse all market statistics response")
		return nil, err
	}

	return response, nil
}

func (s *Service) SingleMarketStatistics(ctx context.Context, req *coreApi.MarketStatisticsReq) (*coreApi.MarketStatistics, error) {
	s.logger.Infof("market statistics called: %v - %v", req.Platform, req.MarketName)
	params := &brokerages.MarketStatisticsParams{
		Market: req.MarketName,
	}
	brokerage := &coreApi.Brokerage{Platform: req.Platform}
	statistics, err := loadRequest(s.configs, brokerage).MarketStatistics(ctx, params, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
		resp, err := s.requestService.Do(ctx, request)
		s.logger.WithError(err).Error("failed to load market statistics request")
		return resp, err
	})
	if err != nil {
		s.logger.WithError(err).Error("market statistics network request failed")
		return nil, err
	}
	resp := new(coreApi.MarketStatistics)
	mapper.Struct(statistics, resp)
	return resp, nil
}

func (s *Service) MarketList(ctx context.Context, req *coreApi.MarketListReq) (*chipmunkApi.Markets, error) {
	brokerage := &coreApi.Brokerage{Platform: req.Platform}
	br := loadRequest(s.configs, brokerage)
	markets, err := br.MarketList(ctx, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
		resp, err := s.requestService.Do(ctx, request)
		return resp, err
	})
	if err != nil {
		s.logger.WithError(err).Errorf("failed to get market list")
		return nil, err
	}
	return markets, nil
}

func (s *Service) NewOrder(ctx context.Context, req *coreApi.NewOrderReq) (*eagleApi.Order, error) {
	brokerage, err := s.loadBrokerage(ctx, req.BrokerageID)
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
	brokerage, err := s.loadBrokerage(ctx, req.BrokerageID)
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
	brokerage, err := s.loadBrokerage(ctx, req.BrokerageID)
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
