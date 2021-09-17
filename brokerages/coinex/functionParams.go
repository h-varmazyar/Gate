package coinex

import (
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/models/todo"
)

type LoginParams struct {
	brokerages.MustImplementAsFunctionParameter
	Totp int
}

type OrderBookParams struct {
	brokerages.MustImplementAsFunctionParameter
	Symbol models.Market
}

type MarketStatusParams struct {
	brokerages.MustImplementAsFunctionParameter
	Destination string
	Source      string
}

type OHLCParams struct {
	brokerages.MustImplementAsFunctionParameter
	Resolution models.Resolution
	Symbol     models.Market
	From       int64
	To         int64
}

type WalletInfoParams struct {
	brokerages.MustImplementAsFunctionParameter
	WalletName string
}

type WalletBalanceParams struct {
	brokerages.MustImplementAsFunctionParameter
	Currency string
}

type TransactionListParams struct {
	brokerages.MustImplementAsFunctionParameter
	WalletID int
}

type NewOrderParams struct {
	brokerages.MustImplementAsFunctionParameter
	Order todo.Order
}

type OrderStatusParams struct {
	brokerages.MustImplementAsFunctionParameter
	OrderId uint64
}

type OrderListParams struct {
	brokerages.MustImplementAsFunctionParameter
	Destination string
	WithDetails bool
	Status      todo.OrderStatus
	Source      string
	Type        todo.OrderType
}

type UpdateOrderStatusParams struct {
	brokerages.MustImplementAsFunctionParameter
	NewStatus todo.OrderStatus
	OrderId   uint64
}
