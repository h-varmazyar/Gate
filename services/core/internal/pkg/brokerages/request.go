package brokerages

import (
	"context"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
)

type Requests interface {
	AsyncOHLC(ctx context.Context, inputs *OHLCParams) (*networkAPI.Request, error)

	WalletList(context.Context, Handler) (*chipmunkApi.Wallets, error)
	OHLC(context.Context, *OHLCParams, Handler) ([]*chipmunkApi.Candle, error)
	MarketStatistics(context.Context, *MarketStatisticsParams, Handler) (*chipmunkApi.Candle, error)
	MarketList(context.Context, Handler) (*chipmunkApi.Markets, error)
	UpdateMarket(context.Context, Handler) ([]*chipmunkApi.Market, error)
	NewOrder(context.Context, *NewOrderParams, Handler) (*eagleApi.Order, error)
	CancelOrder(context.Context, *CancelOrderParams, Handler) (*eagleApi.Order, error)
	OrderStatus(context.Context, *OrderStatusParams, Handler) (*eagleApi.Order, error)
}