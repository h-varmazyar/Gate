package brokerages

import (
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/models"
)

type CoinexConfig struct {
}

func (c CoinexConfig) Login(totp int) error {
	return nil
}

func (c CoinexConfig) OrderBook(symbol Symbol) (*api.OrderBookResponse, error) {
	return nil, nil
}

func (c CoinexConfig) RecentTrades(symbol Symbol) (*api.RecentTradesResponse, error) {
	return nil, nil
}

func (c CoinexConfig) MarketStats(destination, source string) (*api.MarketStatusResponse, error) {
	return nil, nil
}

func (c CoinexConfig) OHLC(symbol Symbol, resolution *models.Resolution, from, to float64) (*api.OHLCResponse, error) {
	return nil, nil
}

func (c CoinexConfig) UserInfo() (*api.UserInfoResponse, error) {
	return nil, nil
}

func (c CoinexConfig) WalletList() (*api.WalletsResponse, error) {
	return nil, nil
}

func (c CoinexConfig) WalletInfo(walletName string) (*api.WalletResponse, error) {
	return nil, nil
}

func (c CoinexConfig) WalletBalance(currency string) (*api.BalanceResponse, error) {
	return nil, nil
}

func (c CoinexConfig) TransactionList(walletID int) (*api.TransactionListResponse, error) {
	return nil, nil
}

func (c CoinexConfig) NewOrder(order models.Order) (*api.OrderResponse, error) {
	return nil, nil
}

func (c CoinexConfig) OrderStatus(orderId uint64) (*api.OrderResponse, error) {
	return nil, nil
}

func (c CoinexConfig) OrderList(status models.OrderStatus, Type models.OrderType, source, destination string, withDetails bool) (*api.OrderListResponse, error) {
	return nil, nil
}

func (c CoinexConfig) UpdateOrderStatus(orderId uint64, newStatus models.OrderStatus) (*api.UpdateOrderStatusResponse, error) {
	return nil, nil
}
