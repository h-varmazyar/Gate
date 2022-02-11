package ohlc

import (
	"context"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/errors"
	"github.com/mrNobody95/Gate/pkg/mapper"
	chipmunkApi "github.com/mrNobody95/Gate/services/chipmunk/api"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/buffer"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/indicators"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/repository"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/workers"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/pkg/workers/OHLC"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

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
	chipmunkApi.RegisterOhlcServiceServer(server, s)
}

func (s *Service) AddMarket(ctx context.Context, req *chipmunkApi.AddMarketRequest) (*api.Void, error) {
	settings := new(OHLC.OhlcWorkerSettings)
	if req.Market == nil {
		return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "market is nil")
	}
	if req.Resolution == nil {
		return nil, errors.NewWithSlug(ctx, codes.InvalidArgument, "resolution is nil")
	}
	settings.Market = req.Market
	settings.Resolution = req.Resolution
	settings.Indicators = make([]indicators.Indicator, 0)
	workers.OHLCWorker.AddMarket(settings)
	return &api.Void{}, nil
}

func (s *Service) ReturnCandles(_ context.Context, req *chipmunkApi.BufferedCandlesRequest) (*api.Candles, error) {
	response := make([]*api.Candle, 0)
	for i := 0; ; i += 1000 {
		list, err := repository.Candles.ReturnList(req.MarketID, req.ResolutionID, i)
		if err != nil {
			return nil, err
		}
		tmp := make([]*api.Candle, 0)
		mapper.Slice(list, &tmp)
		response = append(response, tmp...)
		if len(list) < 1000 {
			break
		}
	}
	return &api.Candles{Candles: response}, nil
}

func (s *Service) ReturnLastCandle(_ context.Context, req *chipmunkApi.BufferedCandlesRequest) (*api.Candle, error) {
	candle := buffer.Candles.Last(req.MarketID, req.ResolutionID)
	response := new(api.Candle)
	mapper.Struct(candle, &response)
	return response, nil
}

func (s *Service) CancelWorker(_ context.Context, req *chipmunkApi.CancelWorkerRequest) (*api.Void, error) {
	return new(api.Void), workers.OHLCWorker.CancelWorker(req.MarketID, req.ResolutionID)
}
