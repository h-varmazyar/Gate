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
	OHLC(context.Context, *OHLCParams, Handler) ([]*api.Candle, error)
	MarketStatistics(context.Context, *MarketStatisticsParams, Handler) (*api.Candle, error)
	UpdateMarket(context.Context, Handler) ([]*repository.Market, error)
	NewOrder(context.Context, *NewOrderParams, Handler) (*brokerageApi.Order, error)
	CancelOrder(context.Context, *CancelOrderParams, Handler) (*brokerageApi.Order, error)
	OrderStatus(context.Context, *OrderStatusParams, Handler) (*brokerageApi.Order, error)
}
