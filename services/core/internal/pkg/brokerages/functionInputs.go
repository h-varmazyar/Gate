package brokerages

import (
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/repository"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api"
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
	OrderModel eagleApi.OrderModel
	ClientUUID string
	BuyOrSell  eagleApi.OrderType
	Price      float64
	StopPrice  float64
	Market     *chipmunkApi.Market
	Amount     float64
	Option     eagleApi.OrderOption
	HideOrder  bool
}

type CancelOrderParams struct {
	ServerOrderId int64
	Market        *chipmunkApi.Market
	Type          eagleApi.OrderType
	ClientUUID    string
	AllOrders     bool
}

type OrderStatusParams struct {
	ServerOrderId int64
	Market        *chipmunkApi.Market
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
