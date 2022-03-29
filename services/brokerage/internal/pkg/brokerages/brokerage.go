package brokerages

import (
	"context"
	"github.com/h-varmazyar/Gate/api"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	"github.com/h-varmazyar/Gate/services/brokerage/internal/pkg/repository"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api"
)

type Handler func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error)

type Brokerage interface {
	WalletList(context.Context, Handler) (*brokerageApi.Wallets, error)
	OHLC(context.Context, OHLCParams, Handler) ([]*api.Candle, error)
	MarketStatistics(context.Context, MarketStatisticsParams, Handler) (*api.Candle, error)
	UpdateMarket(ctx context.Context, handler Handler) ([]*repository.Market, error)
}
