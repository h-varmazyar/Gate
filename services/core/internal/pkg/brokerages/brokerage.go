package brokerages

import (
	"context"
	api "github.com/h-varmazyar/Gate/api/proto"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
)

type Handler func(ctx context.Context, request *networkAPI.Request) (*networkAPI.Response, error)

//type Brokerage interface {
//	WalletList(context.Context, Handler) (*chipmunkApi.Wallets, error)
//	OHLC(context.Context, *OHLCParams, Handler) ([]*chipmunkApi.Candle, error)
//	MarketStatistics(context.Context, *MarketStatisticsParams, Handler) (*chipmunkApi.Candle, error)
//	MarketList(context.Context, Handler) (*chipmunkApi.Markets, error)
//	UpdateMarket(context.Context, Handler) ([]*chipmunkApi.Market, error)
//	NewOrder(context.Context, *NewOrderParams, Handler) (*eagleApi.Order, error)
//	CancelOrder(context.Context, *CancelOrderParams, Handler) (*eagleApi.Order, error)
//	OrderStatus(context.Context, *OrderStatusParams, Handler) (*eagleApi.Order, error)
//}

type Metadata struct {
	Method       string
	Platform     api.Platform
	MarketID     string
	ResolutionID string
}
