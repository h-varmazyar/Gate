package brokerages

import "time"

type Brokerage interface {
	//Validate() error
	//GetName() BrokerageName
	//Login(totp int) error
	//OrderBook(symbol Symbol) (*api.OrderBookResponse, error)
	//RecentTrades(symbol Symbol) (*api.RecentTradesResponse, error)
	//MarketStats(destination, source string) (*api.MarketStatusResponse, error)
	//OHLC(symbol Symbol, resolution *models.Resolution, from, to int64) (*api.OHLCResponse, error)
	//UserInfo() (*api.UserInfoResponse, error)
	//WalletList() (*api.WalletsResponse, error)
	//WalletInfo(walletName string) (*api.WalletResponse, error)
	//WalletBalance(currency string) (*api.BalanceResponse, error)
	//TransactionList(walletID int) (*api.TransactionListResponse, error)
	//NewOrder(order models.Order) (*api.OrderResponse, error)
	//OrderStatus(orderId uint64) (*api.OrderResponse, error)
	//OrderList(status models.OrderStatus, Type models.OrderType, source, destination string, withDetails bool) (*api.OrderListResponse, error)
	//UpdateOrderStatus(orderId uint64, newStatus models.OrderStatus) (*api.UpdateOrderStatusResponse, error)
	Validate() error
	GetName() BrokerageName
	OHLC(MustImplementAsFunctionParameter) *BasicResponse
	Login(MustImplementAsFunctionParameter) *BasicResponse
	NewOrder(MustImplementAsFunctionParameter) *BasicResponse
	UserInfo(MustImplementAsFunctionParameter) *BasicResponse
	OrderList(MustImplementAsFunctionParameter) *BasicResponse
	OrderBook(MustImplementAsFunctionParameter) *BasicResponse
	WalletInfo(MustImplementAsFunctionParameter) *BasicResponse
	WalletList(MustImplementAsFunctionParameter) *BasicResponse
	OrderStatus(MustImplementAsFunctionParameter) *BasicResponse
	MarketStats(MustImplementAsFunctionParameter) *BasicResponse
	RecentTrades(MustImplementAsFunctionParameter) *BasicResponse
	WalletBalance(MustImplementAsFunctionParameter) *BasicResponse
	TransactionList(MustImplementAsFunctionParameter) *BasicResponse
	UpdateOrderStatus(MustImplementAsFunctionParameter) *BasicResponse
}

type ManagementFunctions interface {
	SubscribePeriodicOHLC(period time.Duration, endSignal chan bool)
}

type BrokerageConfig struct {
	Name                 BrokerageName
	Username             string
	Password             string
	CandlePreRequestSize int
}

type BasicResponse struct {
	Error error
}

type MustImplementAsFunctionParameter struct{}

type Symbol string
type Currency string
type BrokerageName string

const (
	Binance BrokerageName = "binance"
	coinex  BrokerageName = "coinex"
	Nobitex BrokerageName = "nobitex"
)
