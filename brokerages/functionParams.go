package brokerages

import (
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/models/todo"
)

type LoginParams struct {
	MustImplementAsFunctionParameter
	Totp int
}

type OrderBookParams struct {
	MustImplementAsFunctionParameter
	Symbol models.Market
}

type MarketStatusParams struct {
	MustImplementAsFunctionParameter
	Destination string
	Source      string
}

type OHLCParams struct {
	MustImplementAsFunctionParameter
	Resolution models.Resolution
	Market     models.Market
	From       int64
	To         int64
}

type WalletInfoParams struct {
	MustImplementAsFunctionParameter
	WalletName string
}

type WalletBalanceParams struct {
	MustImplementAsFunctionParameter
	Currency string
}

type TransactionListParams struct {
	MustImplementAsFunctionParameter
	WalletID int
}

type NewOrderParams struct {
	MustImplementAsFunctionParameter
	Order todo.Order
}

type OrderStatusParams struct {
	MustImplementAsFunctionParameter
	OrderId uint64
}

type OrderListParams struct {
	MustImplementAsFunctionParameter
	Destination string
	WithDetails bool
	Status      todo.OrderStatus
	Source      string
	Type        todo.OrderType
}

type UpdateOrderStatusParams struct {
	MustImplementAsFunctionParameter
	NewStatus todo.OrderStatus
	OrderId   uint64
}
