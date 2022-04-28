package brokerages

import (
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	"github.com/h-varmazyar/Gate/services/brokerage/internal/pkg/repository"
	"time"
)

type LoginParams struct {
	Totp int
}

type OrderBookParams struct {
	Symbol repository.Market
}

type MarketStatisticsParams struct {
	Destination string
	Source      string
	Market      string
}

type OHLCParams struct {
	Resolution *repository.Resolution
	Market     *repository.Market
	From       time.Time
	To         time.Time
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
	OrderModel brokerageApi.OrderModel
	ClientUUID string
	BuyOrSell  brokerageApi.OrderType
	Price      float64
	StopPrice  float64
	Market     *brokerageApi.Market
	Amount     float64
	Option     brokerageApi.OrderOption
	HideOrder  bool
}

type CancelOrderParams struct {
	ServerOrderId int64
	Market        *brokerageApi.Market
	Type          brokerageApi.OrderType
	ClientUUID    string
	AllOrders     bool
}

type OrderStatusParams struct {
	ServerOrderId int64
	Market        *brokerageApi.Market
	ClientUUID    string
}

//type OrderListParams struct {
//	WithDetails bool
//	Status      repository.OrderStatus
//	Market      repository.Market
//	Type        repository.OrderType
//	IsBuy       repository.OrderType
//	ClientUUID  uuid.UUID
//	Page        int
//	Limit       int
//	IsExecuted  bool
//}
//
//type UpdateOrderStatusParams struct {
//	NewStatus repository.OrderStatus
//	OrderId   uint64
//}
//
//type MarketInfoParams struct {
//	MarketName string
//}
