package brokerages

import (
	"github.com/mrNobody95/Gate/models"
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
	Trades []models.Trade
}

type MarketStatusResponse struct {
	BasicResponse
	Symbol string
	Trades []models.Trade
}

type OHLCResponse struct {
	BasicResponse
	ContinueLast uint
	Resolution   models.Resolution
	Candles      []models.Candle
	Status       string
	Symbol       Symbol
}

type UserInfoResponse struct {
	BasicResponse
	User        models.User
	BankAccount []models.BankAccount
}

type WalletsResponse struct {
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
	Transactions []models.Transaction
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
