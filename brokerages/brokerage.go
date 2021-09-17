package brokerages

import (
	"time"
)

type BrokerageRequests interface {
	Validate() error
	//GetName() models.BrokerageName
	OHLC(MustImplementAsFunctionParameter) *OHLCResponse
	Login(MustImplementAsFunctionParameter) interface{}
	NewOrder(MustImplementAsFunctionParameter) interface{}
	UserInfo(MustImplementAsFunctionParameter) interface{}
	OrderList(MustImplementAsFunctionParameter) interface{}
	OrderBook(MustImplementAsFunctionParameter) interface{}
	MarketList(MustImplementAsFunctionParameter) interface{}
	WalletInfo(MustImplementAsFunctionParameter) interface{}
	WalletList(MustImplementAsFunctionParameter) interface{}
	OrderStatus(MustImplementAsFunctionParameter) interface{}
	MarketStats(MustImplementAsFunctionParameter) interface{}
	RecentTrades(MustImplementAsFunctionParameter) interface{}
	WalletBalance(MustImplementAsFunctionParameter) interface{}
	TransactionList(MustImplementAsFunctionParameter) interface{}
	UpdateOrderStatus(MustImplementAsFunctionParameter) interface{}
}

type ManagementFunctions interface {
	SubscribePeriodicOHLC(period time.Duration, endSignal chan bool)
}

type BasicResponse struct {
	Error error
}

type MustImplementAsFunctionParameter interface{}
