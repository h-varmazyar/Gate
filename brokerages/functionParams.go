package brokerages

import (
	"github.com/gofrs/uuid"
	"github.com/mrNobody95/Gate/models"
)

type LoginParams struct {
	Totp int
}

type OrderBookParams struct {
	Symbol models.Market
}

type MarketStatusParams struct {
	Destination string
	Source      string
}

type OHLCParams struct {
	Resolution *models.Resolution
	Market     *models.Market
	From       int64
	To         int64
}

type WalletInfoParams struct {
	WalletName string
}

type WalletBalanceParams struct {
	Currency string
}

type TransactionListParams struct {
	WalletID int
}

type NewOrderParams struct {
	OrderKind  models.OrderKind
	ClientUUID uuid.UUID
	BuyOrSell  models.OrderType
	Price      float64
	StopPrice  float64
	Market     models.Market
	Amount     float64
	Option     models.OrderOption
	HideOrder  bool
}

type CancelOrderParams struct {
	ServerOrderId int64
	Market        models.Market
	IsBuy         bool
	ClientUUID    uuid.UUID
	AllOrders     bool
}

type OrderStatusParams struct {
	ServerOrderId uint64
	Market        models.Market
	ClientUUID    uuid.UUID
}

type OrderListParams struct {
	WithDetails bool
	Status      models.OrderStatus
	Market      models.Market
	Type        models.OrderType
	IsBuy       models.OrderType
	ClientUUID  uuid.UUID
	Page        int
	Limit       int
	IsExecuted  bool
}

type UpdateOrderStatusParams struct {
	NewStatus models.OrderStatus
	OrderId   uint64
}

type MarketInfoParams struct {
	MarketName string
}
