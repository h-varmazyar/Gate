package brokerages

import (
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/models/todo"
)

type OrderBookResponse struct {
	BasicResponse
	Symbol string
	Bids   []todo.Order
	Asks   []todo.Order
}

type RecentTradesResponse struct {
	BasicResponse
	Symbol string
	Trades []todo.Trade
}

type MarketListResponse struct {
	BasicResponse
	Markets []models.Market
}

type MarketStatusResponse struct {
	BasicResponse
	Symbol string
	Trades []todo.Trade
}

type OHLCResponse struct {
	BasicResponse
	ContinueLast uint
	Resolution   models.Resolution
	Candles      []models.Candle
	Status       string
	Symbol       models.Market
}

type UserInfoResponse struct {
	BasicResponse
	User        todo.User
	BankAccount []todo.BankAccount
}

type WalletsResponse struct {
	BasicResponse
	Wallets []todo.Wallet
}

type WalletResponse struct {
	BasicResponse
	Wallet todo.Wallet
}

type BalanceResponse struct {
	BasicResponse
	Symbol  string
	Balance string
}

type TransactionListResponse struct {
	BasicResponse
	Transactions []todo.Transaction
}

type OrderResponse struct {
	BasicResponse
	Order todo.Order
}

type OrderListResponse struct {
	BasicResponse
	Orders []todo.Order
}

type UpdateOrderStatusResponse struct {
	BasicResponse
	NewStatus todo.OrderStatus
}
