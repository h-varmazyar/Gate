package brokerages

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/models"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type NobitexConfig struct {
	BrokerageConfig
	Token     string
	LongToken bool
}

func (n NobitexConfig) Validate() {

}

func (n NobitexConfig) Login(totp int) error {
	req := api.Request{
		Type:     api.POST,
		Endpoint: "https://api.nobitex.ir/auth/login/",
	}
	if totp > 0 {
		req.Headers = map[string][]string{"X-TOTP": {string(totp)}}
	}
	resp, err := req.Execute()
	if err != nil {
		return err
	}
	if resp.Code == 200 {
		respStr := struct {
			Key string `json:"key"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return err
		}
		n.Token = respStr.Key

		return nil
	} else {
		return errors.New(resp.Status)
	}
}

func (n NobitexConfig) OrderBook(symbol Symbol) (*api.OrderBookResponse, error) {
	req := api.Request{
		Type:     api.GET,
		Endpoint: "https://api.nobitex.ir/v2/orderbook/" + string(symbol),
	}

	resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if resp.Code == 200 {
		respStr := struct {
			Status string      `json:"status"`
			Bids   [][2]string `json:"bids"`
			Asks   [][2]string `json:"asks"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return nil, err
		}
		if respStr.Status == "ok" {
			orderBook := api.OrderBookResponse{
				Symbol: string(symbol),
				Bids:   make([]models.Order, len(respStr.Bids)),
				Asks:   make([]models.Order, len(respStr.Asks)),
			}
			for i, bid := range respStr.Bids {
				price, err := strconv.ParseFloat(bid[0], 64)
				if err != nil {
					return nil, err
				}
				orderBook.Bids[i].Price = price
				orderBook.Bids[i].Volume = bid[1]
			}
			for i, ask := range respStr.Asks {
				price, err := strconv.ParseFloat(ask[0], 64)
				if err != nil {
					return nil, err
				}
				orderBook.Asks[i].Price = price
				orderBook.Asks[i].Volume = ask[1]
			}
			return &orderBook, nil
		} else {
			return nil, errors.New("nobitex tesponse error")
		}
	} else {
		return nil, errors.New(resp.Status)
	}
}

func (n NobitexConfig) RecentTrades(symbol Symbol) (*api.RecentTradesResponse, error) {
	req := api.Request{
		Type:     api.GET,
		Endpoint: "https://api.nobitex.ir/v2/trades/" + string(symbol),
	}

	resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if resp.Code == 200 {
		respStr := struct {
			Status string `json:"status"`
			Trades []struct {
				Time   float64 `json:"time"`
				Price  string  `json:"price"`
				Volume string  `json:"volume"`
				Type   string  `json:"type"`
			} `json:"trades"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return nil, err
		}
		if respStr.Status == "ok" {
			recentTrade := api.RecentTradesResponse{
				Symbol: string(symbol),
				Trades: make([]models.Trade, len(respStr.Trades)),
			}
			for i, trade := range respStr.Trades {
				recentTrade.Trades[i].Time = trade.Time
				recentTrade.Trades[i].Price, _ = strconv.ParseFloat(trade.Price, 64)
				recentTrade.Trades[i].Volume, _ = strconv.ParseFloat(trade.Volume, 64)
				recentTrade.Trades[i].Type = models.OrderType(trade.Type)
			}
			return &recentTrade, nil
		} else {
			return nil, errors.New("nobitex tesponse error")
		}
	} else {
		return nil, errors.New(resp.Status)
	}
}

func (n NobitexConfig) MarketStats(destination, source string) (*api.MarketStatusResponse, error) {
	return nil, nil
}

func (n NobitexConfig) OHLC(symbol Symbol, resolution *models.Resolution, from, to float64) (*api.OHLCResponse, error) {
	req := api.Request{
		Type:     api.GET,
		Endpoint: "https://api.nobitex.ir/market/udf/history",
		Params:   map[string]interface{}{"symbol": symbol, "resolution": resolution.Value, "from": from, "to": to},
	}

	resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if resp.Code == 200 {
		respStr := struct {
			Status string    `json:"s"`
			Time   []float64 `json:"t"`
			Open   []string  `json:"o"`
			High   []string  `json:"h"`
			Low    []string  `json:"l"`
			Close  []string  `json:"c"`
			Volume []string  `json:"v"`
			Error  string    `json:"errmsg"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return nil, err
		}
		if respStr.Status == "ok" {
			ohlc := api.OHLCResponse{
				Symbol:     string(symbol),
				Resolution: resolution,
			}
			for i := 0; i < len(respStr.Time); i++ {
				ohlc.Candles[i].Time = respStr.Time[i]
				ohlc.Candles[i].Open, _ = strconv.ParseFloat(respStr.Open[i], 64)
				ohlc.Candles[i].High, _ = strconv.ParseFloat(respStr.High[i], 64)
				ohlc.Candles[i].Low, _ = strconv.ParseFloat(respStr.Low[i], 64)
				ohlc.Candles[i].Close, _ = strconv.ParseFloat(respStr.Close[i], 64)
				ohlc.Candles[i].Vol, _ = strconv.ParseFloat(respStr.Volume[i], 64)
			}
			return &ohlc, nil
		} else {
			return nil, errors.New(respStr.Error)
		}
	} else {
		return nil, errors.New(resp.Status)
	}
}

func (n NobitexConfig) UserInfo() (*api.UserInfoResponse, error) {
	req := api.Request{
		Type:     api.GET,
		Endpoint: "https://api.nobitex.ir/users/profile",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", n.Token)}},
	}

	resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if resp.Code == 200 {
		respStr := struct {
			Status  string `json:"status"`
			Profile struct {
				FirstName    string `json:"firstName"`
				LastName     string `json:"lastName"`
				NationalCode string `json:"nationalCode"`
				Email        string `json:"email"`
				Username     string `json:"username"`
				Phone        string `json:"phone"`
				Mobile       string `json:"mobile"`
				City         string `json:"city"`
				BankCards    []struct {
					Number    string `json:"number"`
					Bank      string `json:"bank"`
					Owner     string `json:"owner"`
					Confirmed bool   `json:"confirmed"`
					Status    string `json:"status"`
				} `json:"bankCards"`
				BankAccounts []struct {
					Id        int    `json:"id"`
					Number    string `json:"number"`
					IBAN      string `json:"shaba"`
					Bank      string `json:"bank"`
					Owner     string `json:"owner"`
					Confirmed bool   `json:"confirmed"`
					Status    string `json:"status"`
				}
				Verifications struct {
					Email       bool `json:"email"`
					Phone       bool `json:"phone"`
					Mobile      bool `json:"mobile"`
					Identity    bool `json:"identity"`
					Selfie      bool `json:"selfie"`
					BankAccount bool `json:"bankAccount"`
					BankCard    bool `json:"bankCard"`
					Address     bool `json:"address"`
					City        bool `json:"city"`
				} `json:"verifications"`
				PendingVerifications struct {
					Email       bool `json:"email"`
					Phone       bool `json:"phone"`
					Mobile      bool `json:"mobile"`
					Identity    bool `json:"identity"`
					Selfie      bool `json:"selfie"`
					BankAccount bool `json:"bankAccount"`
					BankCard    bool `json:"bankCard"`
				} `json:"pendingVerifications"`
				Options struct {
					Fee               string `json:"fee"`
					FeeUsdt           string `json:"feeUsdt"`
					IsManualFee       bool   `json:"isManualFee"`
					TFA               bool   `json:"tfa"`
					SocialLoginEnable bool   `json:"socialLoginEnabled"`
				} `json:"options"`
				WithdrawEligible bool `json:"withdrawEligible"`
			} `json:"profile"`
			TradeStatus struct {
				MonthTradesTotal string `json:"monthTradesTotal"`
				MonthTradesCount int    `json:"monthTradesCount"`
			} `json:"tradeStatus"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return nil, err
		}
		if respStr.Status == "ok" {
			bankAccounts := make([]models.BankAccount, len(respStr.Profile.BankAccounts))
			for i, account := range respStr.Profile.BankAccounts {
				bankAccounts[i].AccountNumber = account.Number
				bankAccounts[i].OwnerName = account.Owner
				bankAccounts[i].BankName = account.Bank
				bankAccounts[i].Status = account.Status
				bankAccounts[i].IBAN = account.IBAN
			}
			userInfo := api.UserInfoResponse{
				User: models.User{
					FirstName: respStr.Profile.FirstName,
					LastName:  respStr.Profile.LastName,
					Email:     respStr.Profile.Email,
					Username:  respStr.Profile.Username,
					Phone:     respStr.Profile.Phone,
					Mobile:    respStr.Profile.Mobile,
					City:      respStr.Profile.City,
					IdCode:    respStr.Profile.NationalCode,
				},
				BankAccount: bankAccounts,
			}
			return &userInfo, nil
		} else {
			return nil, errors.New("get user profile error")
		}
	} else {
		return nil, errors.New(resp.Status)
	}
}

func (n NobitexConfig) WalletList() (*api.WalletsResponse, error) {
	req := api.Request{
		Type:     api.POST,
		Endpoint: "https://api.nobitex.ir/users/wallets/list",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", n.Token)}},
	}

	resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if resp.Code == 200 {
		respStr := struct {
			Status  string `json:"status"`
			Wallets []struct {
				RialBalanceSell float64 `json:"rialBalanceSell"`
				DepositAddress  string  `json:"depositAddress"`
				BlockedBalance  string  `json:"blockedBalance"`
				ActiveBalance   string  `json:"activeBalance"`
				RialBalance     float64 `json:"rialBalance"`
				Currency        string  `json:"currency"`
				Balance         string  `json:"balance"`
				User            string  `json:"user"`
				Id              int     `json:"id"`
			} `json:"wallets"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return nil, err
		}
		if respStr.Status == "ok" {
			wallets := make([]models.Wallet, len(respStr.Wallets))
			for i, wallet := range respStr.Wallets {
				wallets[i].ReferenceCurrencyBalance = wallet.RialBalance
				wallets[i].BlockedBalance = wallet.BlockedBalance
				wallets[i].ActiveBalance = wallet.ActiveBalance
				wallets[i].TotalBalance = wallet.Balance
				wallets[i].Currency = wallet.Currency
				wallets[i].ID = wallet.Id
			}

			return &api.WalletsResponse{Wallets: wallets}, nil
		} else {
			return nil, errors.New("get user profile error")
		}
	} else {
		return nil, errors.New(resp.Status)
	}
}

func (n NobitexConfig) WalletInfo(walletName string) (*api.WalletResponse, error) {
	req := api.Request{
		Type:     api.POST,
		Endpoint: "https://api.nobitex.ir/v2/wallets",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", n.Token)}},
		Params:   map[string]interface{}{"currencies": walletName},
	}

	resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if resp.Code == 200 {
		respStr := struct {
			Status  string `json:"status"`
			Wallets map[string]struct {
				Id      int    `json:"id"`
				Balance string `json:"balance"`
				Blocked string `json:"blocked"`
			}
		}{}

		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return nil, err
		}
		if respStr.Status == "ok" {
			return &api.WalletResponse{Wallet: models.Wallet{
				ID:             respStr.Wallets[strings.ToUpper(walletName)].Id,
				BlockedBalance: respStr.Wallets[strings.ToUpper(walletName)].Blocked,
				TotalBalance:   respStr.Wallets[strings.ToUpper(walletName)].Balance,
				Currency:       walletName,
			}}, nil
		} else {
			return nil, errors.New("get user profile error")
		}
	} else {
		return nil, errors.New(resp.Status)
	}
}

func (n NobitexConfig) WalletBalance(currency string) (*api.BalanceResponse, error) {
	req := api.Request{
		Type:     api.POST,
		Endpoint: "https://api.nobitex.ir/users/wallets/balance",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", n.Token)}},
		Params:   map[string]interface{}{"currency": currency},
	}

	resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if resp.Code == 200 {
		respStr := struct {
			Status  string `json:"status"`
			Balance string `json:"balance"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return nil, err
		}
		if respStr.Status == "ok" {
			return &api.BalanceResponse{
				Symbol:  currency,
				Balance: respStr.Balance,
			}, nil
		} else {
			return nil, errors.New("internal server error (wallet balance)")
		}
	} else {
		return nil, errors.New(resp.Status)
	}
}

func (n NobitexConfig) TransactionList(walletID int) (*api.TransactionListResponse, error) {
	req := api.Request{
		Type:     api.POST,
		Endpoint: "https://api.nobitex.ir/users/wallets/transactions/list",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", n.Token)}},
		Params:   map[string]interface{}{"wallet": walletID},
	}

	resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if resp.Code == 200 {
		respStr := struct {
			Status       string `json:"status"`
			Transactions []struct {
				Currency      string    `json:"currency"`
				CreatedAt     time.Time `json:"created_at"`
				CalculatedFee string    `json:"calculatedFee"`
				Id            uint      `json:"id"`
				Amount        string    `json:"amount"`
				Description   string    `json:"description"`
			} `json:"transactions"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return nil, err
		}
		if respStr.Status == "ok" {
			transactions := make([]models.Transaction, len(respStr.Transactions))
			for i, transaction := range respStr.Transactions {
				transactions[i] = models.Transaction{
					Model: gorm.Model{
						ID:        transaction.Id,
						CreatedAt: transaction.CreatedAt,
					},
					Volume:        transaction.Amount,
					Currency:      transaction.Currency,
					Description:   transaction.Description,
					CalculatedFee: transaction.CalculatedFee,
				}
			}
			return &api.TransactionListResponse{Transactions: transactions}, nil
		} else {
			return nil, errors.New("internal server error (wallet balance)")
		}
	} else {
		return nil, errors.New(resp.Status)
	}
}

func (n NobitexConfig) NewOrder(order models.Order) (*api.OrderResponse, error) {
	body := make(map[string]interface{})
	body["price"] = order.Price
	body["amount"] = order.Volume
	body["type"] = order.OrderType
	body["srcCurrency"] = order.SourceCurrency
	body["destCurrency"] = order.DestinationCurrency
	req := api.Request{
		Type:     api.POST,
		Endpoint: "https://api.nobitex.ir/market/orders/add",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", n.Token)}},
		Params:   body,
	}

	resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if resp.Code == 200 {
		respStr := struct {
			Status string `json:"status"`
			Order  struct {
				Id         uint      `json:"id"`
				Fee        float64   `json:"fee"`
				Src        string    `json:"srcCurrency"`
				Dest       string    `json:"destCurrency"`
				Type       string    `json:"type"`
				User       string    `json:"user"`
				Price      string    `json:"price"`
				Amount     string    `json:"amount"`
				Status     string    `json:"status"`
				Matched    string    `json:"matchedAmount"`
				Unmatched  string    `json:"unmatchedAmount"`
				CreatedAt  time.Time `json:"created_at"`
				TotalPrice string    `json:"totalPrice"`
			} `json:"transactions"`
			Message string `json:"message"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return nil, err
		}
		if respStr.Status == "ok" {
			price, err := strconv.ParseFloat(respStr.Order.Price, 64)
			if err != nil {
				return nil, err
			}
			return &api.OrderResponse{
				Order: models.Order{
					Model: gorm.Model{
						ID:        respStr.Order.Id,
						CreatedAt: respStr.Order.CreatedAt,
					},
					Fee:                 respStr.Order.Fee,
					User:                respStr.Order.User,
					Price:               price,
					Status:              models.OrderStatus(respStr.Order.Status),
					Volume:              respStr.Order.Amount,
					OrderType:           models.OrderType(respStr.Order.Type),
					TotalPrice:          respStr.Order.TotalPrice,
					MatchedVolume:       respStr.Order.Matched,
					SourceCurrency:      respStr.Order.Src,
					UnMatchedVolume:     respStr.Order.Unmatched,
					DestinationCurrency: respStr.Order.Dest,
				},
			}, nil
		} else {
			return nil, errors.New(respStr.Message)
		}
	} else {
		return nil, errors.New(resp.Status)
	}
}

func (n NobitexConfig) OrderStatus(orderId uint64) (*api.OrderResponse, error) {
	req := api.Request{
		Type:     api.POST,
		Endpoint: "https://api.nobitex.ir/market/orders/status",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", n.Token)}},
		Params:   map[string]interface{}{"id": orderId},
	}

	resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if resp.Code == 200 {
		respStr := struct {
			Status string `json:"status"`
			Order  struct {
				Id         uint      `json:"id"`
				Fee        float64   `json:"fee"`
				Src        string    `json:"srcCurrency"`
				Dest       string    `json:"destCurrency"`
				Type       string    `json:"type"`
				User       string    `json:"user"`
				Price      string    `json:"price"`
				Amount     string    `json:"amount"`
				Status     string    `json:"status"`
				Matched    string    `json:"matchedAmount"`
				Unmatched  string    `json:"unmatchedAmount"`
				CreatedAt  time.Time `json:"created_at"`
				TotalPrice string    `json:"totalPrice"`
			} `json:"transactions"`
			Message string `json:"message"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return nil, err
		}
		if respStr.Status == "ok" {
			price, err := strconv.ParseFloat(respStr.Order.Price, 64)
			if err != nil {
				return nil, err
			}
			return &api.OrderResponse{
				Order: models.Order{
					Model: gorm.Model{
						ID:        respStr.Order.Id,
						CreatedAt: respStr.Order.CreatedAt,
					},
					Fee:                 respStr.Order.Fee,
					User:                respStr.Order.User,
					Price:               price,
					Status:              models.OrderStatus(respStr.Order.Status),
					Volume:              respStr.Order.Amount,
					OrderType:           models.OrderType(respStr.Order.Type),
					TotalPrice:          respStr.Order.TotalPrice,
					MatchedVolume:       respStr.Order.Matched,
					SourceCurrency:      respStr.Order.Src,
					UnMatchedVolume:     respStr.Order.Unmatched,
					DestinationCurrency: respStr.Order.Dest,
				},
			}, nil
		} else {
			return nil, errors.New(respStr.Message)
		}
	} else {
		return nil, errors.New(resp.Status)
	}
}

func (n NobitexConfig) OrderList(status models.OrderStatus, Type models.OrderType, source, destination string, withDetails bool) (*api.OrderListResponse, error) {
	body := make(map[string]interface{})
	if status != "" {
		body["status"] = status
	}
	if Type != "" {
		body["type"] = Type
	}
	if source != "" {
		body["srcCurrency"] = source
	} else {
		return nil, errors.New("please specify source currency")
	}
	if destination != "" {
		body["dstCurrency"] = destination
	} else {
		return nil, errors.New("please specify destination currency")
	}
	if withDetails {
		body["details"] = 2
	} else {
		body["details"] = 1
	}
	req := api.Request{
		Type:     api.POST,
		Endpoint: "https://api.nobitex.ir/market/orders/status",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", n.Token)}},
		Params:   map[string]interface{}{"id": body},
	}

	resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if resp.Code == 200 {
		respStr := struct {
			Status string `json:"status"`
			Orders []struct {
				Id           uint      `json:"id"`
				Fee          float64   `json:"fee"`
				Src          string    `json:"srcCurrency"`
				Dest         string    `json:"destCurrency"`
				Type         string    `json:"type"`
				User         string    `json:"user"`
				Price        string    `json:"price"`
				Amount       string    `json:"amount"`
				Status       string    `json:"status"`
				Matched      string    `json:"matchedAmount"`
				Unmatched    string    `json:"unmatchedAmount"`
				CreatedAt    time.Time `json:"created_at"`
				TotalPrice   string    `json:"totalPrice"`
				AveragePrice string    `json:"averagePrice"`
			} `json:"transactions"`
			Message string `json:"message"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return nil, err
		}
		if respStr.Status == "ok" {
			orders := make([]models.Order, len(respStr.Orders))
			for i, order := range respStr.Orders {
				price, err := strconv.ParseFloat(order.Price, 64)
				if err != nil {
					continue
				}
				orders[i] = models.Order{
					Model: gorm.Model{
						ID:        order.Id,
						CreatedAt: order.CreatedAt,
					},
					Fee:                 order.Fee,
					User:                order.User,
					Price:               price,
					Status:              models.OrderStatus(order.Status),
					Volume:              order.Amount,
					OrderType:           models.OrderType(order.Type),
					TotalPrice:          order.TotalPrice,
					MatchedVolume:       order.Matched,
					SourceCurrency:      order.Src,
					UnMatchedVolume:     order.Unmatched,
					DestinationCurrency: order.Dest,
				}
			}
			return &api.OrderListResponse{Orders: orders}, nil
		} else {
			return nil, errors.New(respStr.Message)
		}
	} else {
		return nil, errors.New(resp.Status)
	}
}

func (n NobitexConfig) UpdateOrderStatus(orderId uint64, newStatus models.OrderStatus) (*api.UpdateOrderStatusResponse, error) {
	body := make(map[string]interface{})
	if orderId < 0 {
		body["order"] = orderId
	} else {
		return nil, errors.New("please specify order id currency")
	}
	if newStatus != "" {
		body["status"] = newStatus
	} else {
		return nil, errors.New("please specify new status currency")
	}
	req := api.Request{
		Type:     api.POST,
		Endpoint: "https://api.nobitex.ir/market/orders/update-status",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", n.Token)}},
		Params:   body,
	}

	resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	if resp.Code == 200 {
		respStr := struct {
			Status        string             `json:"status"`
			Message       string             `json:"message"`
			UpdatedStatus models.OrderStatus `json:"updatedStatus"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return nil, err
		}
		if respStr.Status == "ok" {
			return &api.UpdateOrderStatusResponse{NewStatus: respStr.UpdatedStatus}, nil
		} else {
			return nil, errors.New(respStr.Message)
		}
	} else {
		return nil, errors.New(resp.Status)
	}
}
