package nobitex

import (
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/models"
)

type LoginParams struct {
	brokerages.MustImplementAsFunctionParameter
	Totp int
}

type OrderBookParams struct {
	brokerages.MustImplementAsFunctionParameter
	Symbol brokerages.Symbol
}

type MarketStatusParams struct {
	brokerages.MustImplementAsFunctionParameter
	Destination string
	Source      string
}

type OHLCParams struct {
	brokerages.MustImplementAsFunctionParameter
	Resolution models.Resolution
	Symbol     brokerages.Symbol
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
	Order models.Order
}

type OrderStatusParams struct {
	brokerages.MustImplementAsFunctionParameter
	OrderId uint64
}

type OrderListParams struct {
	brokerages.MustImplementAsFunctionParameter
	Destination string
	WithDetails bool
	Status      models.OrderStatus
	Source      string
	Type        models.OrderType
}

type UpdateOrderStatusParams struct {
	brokerages.MustImplementAsFunctionParameter
	NewStatus models.OrderStatus
	OrderId   uint64
}
