package candles

import (
	"context"
	"fmt"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/grpcext"
	"github.com/mrNobody95/Gate/pkg/mapper"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/brokerage/configs"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/brokerages"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/brokerages/coinex"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/repository"
	networkAPI "github.com/mrNobody95/Gate/services/network/api"
	"google.golang.org/grpc"
	"time"
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
	networkService networkAPI.RequestServiceClient
}

var (
	GrpcService *Service
)

func NewService(configs *configs.Configs) *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		networkConnection := grpcext.NewConnection(fmt.Sprintf(":%d", configs.NetworkGrpcPort))
		GrpcService.networkService = networkAPI.NewRequestServiceClient(networkConnection)
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	brokerageApi.RegisterCandleServiceServer(server, s)
}

func (s *Service) OHLC(ctx context.Context, req *brokerageApi.OhlcRequest) (*api.Candles, error) {
	br := new(coinex.Service)
	resolution := new(repository.Resolution)
	mapper.Struct(req.Resolution, resolution)

	market := new(repository.Market)
	mapper.Struct(req.Market, market)
	inputs := brokerages.OHLCParams{
		Resolution: resolution,
		Market:     market,
		From:       time.Unix(req.From, 0),
		To:         time.Unix(req.To, 0),
	}
	candles, err := br.OHLC(ctx, inputs, func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error) {
		resp, err := s.networkService.Do(ctx, request)
		return resp, err
	})
	if err != nil {
		fmt.Println("after net failed")
		return nil, err
	}
	return &api.Candles{Candles: candles}, nil
}
