package brokerages

import (
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/models"
)

type Brokerage interface {
	Validate() error
	Login(totp int) error
	OrderBook(symbol Symbol) (*api.OrderBookResponse, error)
	RecentTrades(symbol Symbol) (*api.RecentTradesResponse, error)
	MarketStats(destination, source string) (*api.MarketStatusResponse, error)
	OHLC(symbol Symbol, resolution *models.Resolution, from, to float64) (*api.OHLCResponse, error)
	UserInfo() (*api.UserInfoResponse, error)
	WalletList() (*api.WalletsResponse, error)
	WalletInfo(walletName string) (*api.WalletResponse, error)
	WalletBalance(currency string) (*api.BalanceResponse, error)
	TransactionList(walletID int) (*api.TransactionListResponse, error)
	NewOrder(order models.Order) (*api.OrderResponse, error)
	OrderStatus(orderId uint64) (*api.OrderResponse, error)
	OrderList(status models.OrderStatus, Type models.OrderType, source, destination string, withDetails bool) (*api.OrderListResponse, error)
	UpdateOrderStatus(orderId uint64, newStatus models.OrderStatus) (*api.UpdateOrderStatusResponse, error)
}

type BrokerageConfig struct {
	Name     string
	Username string
	Password string
}

type Symbol string
type Currency string

const (
	BTCIRT  Symbol = "BTCIRT"
	ETHIRT  Symbol = "ETHIRT"
	LTCIRT  Symbol = "LTCIRT"
	XRPIRT  Symbol = "XRPIRT"
	BCHIRT  Symbol = "BCHIRT"
	BNBIRT  Symbol = "BNBIRT"
	EOSIRT  Symbol = "EOSIRT"
	XLMIRT  Symbol = "XLMIRT"
	ETCIRT  Symbol = "ETCIRT"
	TRXIRT  Symbol = "TRXIRT"
	SDTIRT  Symbol = "SDTIRT"
	BTCUSDT Symbol = "BTCUSDT"
	ETHUSDT Symbol = "ETHUSDT"
	LTCUSDT Symbol = "LTCUSDT"
	XRPUSDT Symbol = "XRPUSDT"
	BCHUSDT Symbol = "BCHUSDT"
	BNBUSDT Symbol = "BNBUSDT"
	EOSUSDT Symbol = "EOSUSDT"
	XLMUSDT Symbol = "XLMUSDT"
	ETCUSDT Symbol = "ETCUSDT"
	TRXUSDT Symbol = "TRXUSDT"
)

const (
	BTC Currency = "btc"
	ETH Currency = "eth"
	LTC Currency = "ltc"
	XRP Currency = "xrp"
	BCH Currency = "bch"
	BNB Currency = "bnb"
	EOS Currency = "eos"
	XLM Currency = "xlm"
	ETC Currency = "etc"
	TRX Currency = "trx"
	SDT Currency = "sdt"
)
