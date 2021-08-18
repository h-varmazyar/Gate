package nobitex

import (
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/models"
)

type OrderBookResponse struct {
	brokerages.BasicResponse
	Symbol string
	Bids   []models.Order
	Asks   []models.Order
}

type RecentTradesResponse struct {
	brokerages.BasicResponse
	Symbol string
	Trades []models.Trade
}

type MarketStatusResponse struct {
	brokerages.BasicResponse
	Symbol string
	Trades []models.Trade
}

type OHLCResponse struct {
	brokerages.BasicResponse
	ContinueLast uint
	Resolution   *models.Resolution
	Candles      []models.Candle
	Status       string
	Symbol       brokerages.Symbol
}

type UserInfoResponse struct {
	brokerages.BasicResponse
	User        models.User
	BankAccount []models.BankAccount
}

type WalletsResponse struct {
	brokerages.BasicResponse
	Wallets []models.Wallet
}

type WalletResponse struct {
	brokerages.BasicResponse
	Wallet models.Wallet
}

type BalanceResponse struct {
	brokerages.BasicResponse
	Symbol  string
	Balance string
}

type TransactionListResponse struct {
	brokerages.BasicResponse
	Transactions []models.Transaction
}

type OrderResponse struct {
	brokerages.BasicResponse
	Order models.Order
}

type OrderListResponse struct {
	brokerages.BasicResponse
	Orders []models.Order
}

type UpdateOrderStatusResponse struct {
	brokerages.BasicResponse
	NewStatus models.OrderStatus
}
