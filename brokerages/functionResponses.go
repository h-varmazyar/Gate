package brokerages

import (
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/models/todo"
)

type OrderBookResponse struct {
	BasicResponse
	Symbol string
	Bids   []models.Order
	Asks   []models.Order
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
	Resolution   *models.Resolution
	Candles      []models.Candle
	Status       string
	Market       *models.Market
}

type UserInfoResponse struct {
	BasicResponse
	User        todo.User
	BankAccount []todo.BankAccount
}

type WalletListResponse struct {
	BasicResponse
	Wallets []models.Wallet
}

type WalletResponse struct {
	BasicResponse
	Wallet models.Wallet
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
	Order models.Order
}

type OrderListResponse struct {
	BasicResponse
	Orders []models.Order
}

type UpdateOrderStatusResponse struct {
	BasicResponse
	NewStatus models.OrderStatus
}

type MarketInfoResponse struct {
	BasicResponse
	Market models.Market
}

type MarketStatisticsResponse struct {
	BasicResponse
	Candle models.Candle
}
