package nobitex

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/networkManager"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	brokerages.BrokerageConfig
	Token     string
	LongToken bool
}

type PeriodicOHLC struct {
	Config
	brokerages.Symbol
	*models.Resolution
	Response chan *OHLCResponse
}

func (config Config) Validate() error {
	return nil
}

func (config Config) GetName() brokerages.BrokerageName {
	return config.Name
}

func (config Config) Login(params LoginParams) *brokerages.BasicResponse {
	req := networkManager.Request{
		Type:     networkManager.POST,
		Endpoint: "https://api.nobitex.ir/auth/login/",
	}
	if params.Totp > 0 {
		req.Headers = map[string][]string{"X-TOTP": {strconv.Itoa(params.Totp)}}
	}
	resp, err := req.Execute()
	if err != nil {
		return &brokerages.BasicResponse{Error: err}
	}
	if resp.Code == 200 {
		respStr := struct {
			Key string `json:"key"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &brokerages.BasicResponse{Error: err}
		}
		config.Token = respStr.Key

		return nil
	} else {
		return &brokerages.BasicResponse{Error: errors.New(resp.Status)}
	}
}

func (config Config) OrderBook(params OrderBookParams) *OrderBookResponse {
	req := networkManager.Request{
		Type:     networkManager.GET,
		Endpoint: "https://api.nobitex.ir/v2/orderbook/" + string(params.Symbol),
	}

	resp, err := req.Execute()
	if err != nil {
		return &OrderBookResponse{
			BasicResponse: brokerages.BasicResponse{
				Error: err,
			},
		}
	}
	if resp.Code == 200 {
		respStr := struct {
			Status string      `json:"Status"`
			Bids   [][2]string `json:"bids"`
			Asks   [][2]string `json:"asks"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &OrderBookResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: err,
				},
			}
		}
		if respStr.Status == "ok" {
			orderBook := OrderBookResponse{
				Symbol: string(params.Symbol),
				Bids:   make([]models.Order, len(respStr.Bids)),
				Asks:   make([]models.Order, len(respStr.Asks)),
			}
			for i, bid := range respStr.Bids {
				price, err := strconv.ParseFloat(bid[0], 64)
				if err != nil {
					return &OrderBookResponse{
						BasicResponse: brokerages.BasicResponse{
							Error: err,
						},
					}
				}
				orderBook.Bids[i].Price = price
				orderBook.Bids[i].Volume = bid[1]
			}
			for i, ask := range respStr.Asks {
				price, err := strconv.ParseFloat(ask[0], 64)
				if err != nil {
					return &OrderBookResponse{
						BasicResponse: brokerages.BasicResponse{
							Error: err,
						},
					}
				}
				orderBook.Asks[i].Price = price
				orderBook.Asks[i].Volume = ask[1]
			}
			return &orderBook
		} else {
			return &OrderBookResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: errors.New("nobitex tesponse error"),
				},
			}
		}
	} else {
		return &OrderBookResponse{
			BasicResponse: brokerages.BasicResponse{
				Error: errors.New(resp.Status),
			},
		}
	}
}

func (config Config) RecentTrades(params OrderBookParams) *RecentTradesResponse {
	req := networkManager.Request{
		Type:     networkManager.GET,
		Endpoint: "https://api.nobitex.ir/v2/trades/" + string(params.Symbol),
	}

	resp, err := req.Execute()
	if err != nil {
		return &RecentTradesResponse{
			BasicResponse: brokerages.BasicResponse{Error: err},
		}
	}
	if resp.Code == 200 {
		respStr := struct {
			Status string `json:"Status"`
			Trades []struct {
				Time   float64 `json:"time"`
				Price  string  `json:"price"`
				Volume string  `json:"volume"`
				Type   string  `json:"type"`
			} `json:"trades"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &RecentTradesResponse{
				BasicResponse: brokerages.BasicResponse{Error: err},
			}
		}
		if respStr.Status == "ok" {
			recentTrade := RecentTradesResponse{
				Symbol: string(params.Symbol),
				Trades: make([]models.Trade, len(respStr.Trades)),
			}
			for i, trade := range respStr.Trades {
				recentTrade.Trades[i].Time = trade.Time
				recentTrade.Trades[i].Price, _ = strconv.ParseFloat(trade.Price, 64)
				recentTrade.Trades[i].Volume, _ = strconv.ParseFloat(trade.Volume, 64)
				recentTrade.Trades[i].Type = models.OrderType(trade.Type)
			}
			return &recentTrade
		} else {
			return &RecentTradesResponse{
				BasicResponse: brokerages.BasicResponse{Error: errors.New("nobitex response error")},
			}
		}
	} else {
		return &RecentTradesResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New(resp.Status)},
		}
	}
}

func (config Config) MarketStats(params MarketStatusParams) *MarketStatusResponse {
	return &MarketStatusResponse{}
}

func (config Config) OHLC(params OHLCParams) *OHLCResponse {
	req := networkManager.Request{
		Type:     networkManager.GET,
		Endpoint: "https://api.nobitex.ir/market/udf/history",
		Params: map[string]interface{}{"symbol": params.Symbol,
			"resolution": params.Resolution.Value,
			"from":       params.From,
			"to":         params.To},
	}

	resp, err := req.Execute()
	if err != nil {
		return &OHLCResponse{
			BasicResponse: brokerages.BasicResponse{
				Error: err,
			},
		}
	}
	if resp.Code == 200 {
		respStr := struct {
			Status string   `json:"s"`
			Time   []int64  `json:"t"`
			Open   []string `json:"o"`
			High   []string `json:"h"`
			Low    []string `json:"l"`
			Close  []string `json:"c"`
			Volume []string `json:"v"`
			Error  string   `json:"errmsg"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &OHLCResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: err,
				},
			}
		}
		if respStr.Status == "ok" {
			ohlc := OHLCResponse{
				Symbol:     params.Symbol,
				Resolution: params.Resolution,
				Status:     respStr.Status,
			}
			for i := 0; i < len(respStr.Time); i++ {
				ohlc.Candles[i].Time = respStr.Time[i]
				ohlc.Candles[i].Open, _ = strconv.ParseFloat(respStr.Open[i], 64)
				ohlc.Candles[i].High, _ = strconv.ParseFloat(respStr.High[i], 64)
				ohlc.Candles[i].Low, _ = strconv.ParseFloat(respStr.Low[i], 64)
				ohlc.Candles[i].Close, _ = strconv.ParseFloat(respStr.Close[i], 64)
				ohlc.Candles[i].Vol, _ = strconv.ParseFloat(respStr.Volume[i], 64)
			}
			return &ohlc
		} else {
			return &OHLCResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: errors.New(respStr.Error),
				},
			}
		}
	} else {
		return &OHLCResponse{
			BasicResponse: brokerages.BasicResponse{
				Error: errors.New(resp.Status),
			},
		}
	}
}

func (config Config) UserInfo(brokerages.MustImplementAsFunctionParameter) *UserInfoResponse {
	req := networkManager.Request{
		Type:     networkManager.GET,
		Endpoint: "https://api.nobitex.ir/users/profile",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", config.Token)}},
	}

	resp, err := req.Execute()
	if err != nil {
		return &UserInfoResponse{
			BasicResponse: brokerages.BasicResponse{
				Error: err,
			},
		}
	}
	if resp.Code == 200 {
		respStr := struct {
			Status  string `json:"Status"`
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
					Status    string `json:"Status"`
				} `json:"bankCards"`
				BankAccounts []struct {
					Id        int    `json:"id"`
					Number    string `json:"number"`
					IBAN      string `json:"shaba"`
					Bank      string `json:"bank"`
					Owner     string `json:"owner"`
					Confirmed bool   `json:"confirmed"`
					Status    string `json:"Status"`
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
			return &UserInfoResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: err,
				},
			}
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
			userInfo := UserInfoResponse{
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
			return &userInfo
		} else {
			return &UserInfoResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: errors.New("get user profile error"),
				},
			}
		}
	} else {
		return &UserInfoResponse{
			BasicResponse: brokerages.BasicResponse{
				Error: errors.New(resp.Status),
			},
		}
	}
}

func (config Config) WalletList(brokerages.MustImplementAsFunctionParameter) *WalletsResponse {
	req := networkManager.Request{
		Type:     networkManager.POST,
		Endpoint: "https://api.nobitex.ir/users/wallets/list",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", config.Token)}},
	}

	resp, err := req.Execute()
	if err != nil {
		return &WalletsResponse{
			BasicResponse: brokerages.BasicResponse{
				Error: err,
			},
		}
	}
	if resp.Code == 200 {
		respStr := struct {
			Status  string `json:"Status"`
			Wallets []struct {
				RialBalanceSell float64 `json:"rialBalanceSell"`
				DepositAddress  string  `json:"depositAddress"`
				BlockedBalance  string  `json:"blockedBalance"`
				ActiveBalance   string  `json:"activeBalance"`
				RialBalance     float64 `json:"rialBalance"`
				Currency        string  `json:"Currency"`
				Balance         string  `json:"balance"`
				User            string  `json:"user"`
				Id              int     `json:"id"`
			} `json:"wallets"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &WalletsResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: err,
				},
			}
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

			return &WalletsResponse{Wallets: wallets}
		} else {
			return &WalletsResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: errors.New("get user profile error"),
				},
			}
		}
	} else {
		return &WalletsResponse{
			BasicResponse: brokerages.BasicResponse{
				Error: errors.New(resp.Status),
			},
		}
	}
}

func (config Config) WalletInfo(params WalletInfoParams) *WalletResponse {
	req := networkManager.Request{
		Type:     networkManager.POST,
		Endpoint: "https://api.nobitex.ir/v2/wallets",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", config.Token)}},
		Params:   map[string]interface{}{"currencies": params.WalletName},
	}

	resp, err := req.Execute()
	if err != nil {
		return &WalletResponse{
			BasicResponse: brokerages.BasicResponse{
				Error: err,
			},
		}
	}
	if resp.Code == 200 {
		respStr := struct {
			Status  string `json:"Status"`
			Wallets map[string]struct {
				Id      int    `json:"id"`
				Balance string `json:"balance"`
				Blocked string `json:"blocked"`
			}
		}{}

		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &WalletResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: err,
				},
			}
		}
		if respStr.Status == "ok" {
			return &WalletResponse{Wallet: models.Wallet{
				ID:             respStr.Wallets[strings.ToUpper(params.WalletName)].Id,
				BlockedBalance: respStr.Wallets[strings.ToUpper(params.WalletName)].Blocked,
				TotalBalance:   respStr.Wallets[strings.ToUpper(params.WalletName)].Balance,
				Currency:       params.WalletName,
			}}
		} else {
			return &WalletResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: errors.New("get user profile error"),
				},
			}
		}
	} else {
		return &WalletResponse{
			BasicResponse: brokerages.BasicResponse{
				Error: errors.New(resp.Status),
			},
		}
	}
}

func (config Config) WalletBalance(params WalletBalanceParams) *BalanceResponse {
	req := networkManager.Request{
		Type:     networkManager.POST,
		Endpoint: "https://api.nobitex.ir/users/wallets/balance",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", config.Token)}},
		Params:   map[string]interface{}{"currency": params.Currency},
	}

	resp, err := req.Execute()
	if err != nil {
		return &BalanceResponse{
			BasicResponse: brokerages.BasicResponse{Error: err},
		}
	}
	if resp.Code == 200 {
		respStr := struct {
			Status  string `json:"Status"`
			Balance string `json:"balance"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &BalanceResponse{
				BasicResponse: brokerages.BasicResponse{Error: err},
			}
		}
		if respStr.Status == "ok" {
			return &BalanceResponse{
				Symbol:  params.Currency,
				Balance: respStr.Balance,
			}
		} else {
			return &BalanceResponse{
				BasicResponse: brokerages.BasicResponse{Error: errors.New("internal server error (wallet balance)")},
			}
		}
	} else {
		return &BalanceResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New(resp.Status)},
		}
	}
}

func (config Config) TransactionList(params TransactionListParams) *TransactionListResponse {
	req := networkManager.Request{
		Type:     networkManager.POST,
		Endpoint: "https://api.nobitex.ir/users/wallets/transactions/list",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", config.Token)}},
		Params:   map[string]interface{}{"wallet": params.WalletID},
	}

	resp, err := req.Execute()
	if err != nil {
		return &TransactionListResponse{
			BasicResponse: brokerages.BasicResponse{Error: err},
		}
	}
	if resp.Code == 200 {
		respStr := struct {
			Status       string `json:"Status"`
			Transactions []struct {
				Currency      string    `json:"Currency"`
				CreatedAt     time.Time `json:"created_at"`
				CalculatedFee string    `json:"calculatedFee"`
				Id            uint      `json:"id"`
				Amount        string    `json:"amount"`
				Description   string    `json:"description"`
			} `json:"transactions"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &TransactionListResponse{
				BasicResponse: brokerages.BasicResponse{Error: err},
			}
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
			return &TransactionListResponse{Transactions: transactions}
		} else {
			return &TransactionListResponse{
				BasicResponse: brokerages.BasicResponse{Error: errors.New("internal server error (wallet balance)")},
			}
		}
	} else {
		return &TransactionListResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New(resp.Status)},
		}
	}
}

func (config Config) NewOrder(params NewOrderParams) *OrderResponse {
	body := make(map[string]interface{})
	body["price"] = params.Order.Price
	body["amount"] = params.Order.Volume
	body["type"] = params.Order.OrderType
	body["srcCurrency"] = params.Order.SourceCurrency
	body["destCurrency"] = params.Order.DestinationCurrency
	req := networkManager.Request{
		Type:     networkManager.POST,
		Endpoint: "https://api.nobitex.ir/market/orders/add",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", config.Token)}},
		Params:   body,
	}

	resp, err := req.Execute()
	if err != nil {
		return &OrderResponse{
			BasicResponse: brokerages.BasicResponse{Error: err},
		}
	}
	if resp.Code == 200 {
		respStr := struct {
			Status string `json:"Status"`
			Order  struct {
				Id         uint      `json:"id"`
				Fee        float64   `json:"fee"`
				Src        string    `json:"srcCurrency"`
				Dest       string    `json:"destCurrency"`
				Type       string    `json:"type"`
				User       string    `json:"user"`
				Price      string    `json:"price"`
				Amount     string    `json:"amount"`
				Status     string    `json:"Status"`
				Matched    string    `json:"matchedAmount"`
				Unmatched  string    `json:"unmatchedAmount"`
				CreatedAt  time.Time `json:"created_at"`
				TotalPrice string    `json:"totalPrice"`
			} `json:"transactions"`
			Message string `json:"message"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &OrderResponse{
				BasicResponse: brokerages.BasicResponse{Error: err},
			}
		}
		if respStr.Status == "ok" {
			price, err := strconv.ParseFloat(respStr.Order.Price, 64)
			if err != nil {
				return &OrderResponse{
					BasicResponse: brokerages.BasicResponse{Error: err},
				}
			}
			return &OrderResponse{
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
			}
		} else {
			return &OrderResponse{
				BasicResponse: brokerages.BasicResponse{Error: errors.New(respStr.Message)},
			}
		}
	} else {
		return &OrderResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New(resp.Status)},
		}
	}
}

func (config Config) OrderStatus(params OrderStatusParams) *OrderResponse {
	req := networkManager.Request{
		Type:     networkManager.POST,
		Endpoint: "https://api.nobitex.ir/market/orders/Status",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", config.Token)}},
		Params:   map[string]interface{}{"id": params.OrderId},
	}

	resp, err := req.Execute()
	if err != nil {
		return &OrderResponse{
			BasicResponse: brokerages.BasicResponse{Error: err},
		}
	}
	if resp.Code == 200 {
		respStr := struct {
			Status string `json:"Status"`
			Order  struct {
				Id         uint      `json:"id"`
				Fee        float64   `json:"fee"`
				Src        string    `json:"srcCurrency"`
				Dest       string    `json:"destCurrency"`
				Type       string    `json:"type"`
				User       string    `json:"user"`
				Price      string    `json:"price"`
				Amount     string    `json:"amount"`
				Status     string    `json:"Status"`
				Matched    string    `json:"matchedAmount"`
				Unmatched  string    `json:"unmatchedAmount"`
				CreatedAt  time.Time `json:"created_at"`
				TotalPrice string    `json:"totalPrice"`
			} `json:"transactions"`
			Message string `json:"message"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &OrderResponse{
				BasicResponse: brokerages.BasicResponse{Error: err},
			}
		}
		if respStr.Status == "ok" {
			price, err := strconv.ParseFloat(respStr.Order.Price, 64)
			if err != nil {
				return &OrderResponse{
					BasicResponse: brokerages.BasicResponse{Error: err},
				}
			}
			return &OrderResponse{
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
			}
		} else {
			return &OrderResponse{
				BasicResponse: brokerages.BasicResponse{Error: errors.New(respStr.Message)},
			}
		}
	} else {
		return &OrderResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New(resp.Status)},
		}
	}
}

func (config Config) OrderList(params OrderListParams) *OrderListResponse {
	body := make(map[string]interface{})
	if params.Status != "" {
		body["status"] = params.Status
	}
	if params.Type != "" {
		body["type"] = params.Type
	}
	if params.Source != "" {
		body["srcCurrency"] = params.Source
	} else {
		return &OrderListResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New("please specify Source Currency")},
		}
	}
	if params.Destination != "" {
		body["dstCurrency"] = params.Destination
	} else {
		return &OrderListResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New("please specify Destination Currency")},
		}
	}
	if params.WithDetails {
		body["details"] = 2
	} else {
		body["details"] = 1
	}
	req := networkManager.Request{
		Type:     networkManager.POST,
		Endpoint: "https://api.nobitex.ir/market/orders/Status",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", config.Token)}},
		Params:   map[string]interface{}{"id": body},
	}

	resp, err := req.Execute()
	if err != nil {
		return &OrderListResponse{
			BasicResponse: brokerages.BasicResponse{Error: err},
		}
	}
	if resp.Code == 200 {
		respStr := struct {
			Status string `json:"Status"`
			Orders []struct {
				Id           uint      `json:"id"`
				Fee          float64   `json:"fee"`
				Src          string    `json:"srcCurrency"`
				Dest         string    `json:"destCurrency"`
				Type         string    `json:"type"`
				User         string    `json:"user"`
				Price        string    `json:"price"`
				Amount       string    `json:"amount"`
				Status       string    `json:"Status"`
				Matched      string    `json:"matchedAmount"`
				Unmatched    string    `json:"unmatchedAmount"`
				CreatedAt    time.Time `json:"created_at"`
				TotalPrice   string    `json:"totalPrice"`
				AveragePrice string    `json:"averagePrice"`
			} `json:"transactions"`
			Message string `json:"message"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &OrderListResponse{
				BasicResponse: brokerages.BasicResponse{Error: err},
			}
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
			return &OrderListResponse{Orders: orders}
		} else {
			return &OrderListResponse{
				BasicResponse: brokerages.BasicResponse{Error: errors.New(respStr.Message)},
			}
		}
	} else {
		return &OrderListResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New(resp.Status)},
		}
	}
}

func (config Config) UpdateOrderStatus(params UpdateOrderStatusParams) *UpdateOrderStatusResponse {
	body := make(map[string]interface{})
	if params.OrderId < 0 {
		body["Order"] = params.OrderId
	} else {
		return &UpdateOrderStatusResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New("please specify Order id Currency")},
		}
	}
	if params.NewStatus != "" {
		body["Status"] = params.NewStatus
	} else {
		return &UpdateOrderStatusResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New("please specify new Status Currency")},
		}
	}
	req := networkManager.Request{
		Type:     networkManager.POST,
		Endpoint: "https://api.nobitex.ir/market/orders/update-Status",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", config.Token)}},
		Params:   body,
	}

	resp, err := req.Execute()
	if err != nil {
		return &UpdateOrderStatusResponse{
			BasicResponse: brokerages.BasicResponse{Error: err},
		}
	}
	if resp.Code == 200 {
		respStr := struct {
			Status        string             `json:"Status"`
			Message       string             `json:"message"`
			UpdatedStatus models.OrderStatus `json:"updatedStatus"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &UpdateOrderStatusResponse{
				BasicResponse: brokerages.BasicResponse{Error: err},
			}
		}
		if respStr.Status == "ok" {
			return &UpdateOrderStatusResponse{NewStatus: respStr.UpdatedStatus}
		} else {
			return &UpdateOrderStatusResponse{
				BasicResponse: brokerages.BasicResponse{Error: errors.New(respStr.Message)},
			}
		}
	} else {
		return &UpdateOrderStatusResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New(resp.Status)},
		}
	}
}

func (periodic PeriodicOHLC) SubscribePeriodicOHLC(period time.Duration, endSignal chan bool) {
	ticker := time.NewTicker(period)
	go func() {
		for {
			select {
			case <-endSignal:
				return
			case _ = <-ticker.C:
				now := time.Now().Unix()
				periodic.Response <- periodic.Config.OHLC(OHLCParams{
					Resolution: periodic.Resolution,
					Symbol:     periodic.Symbol,
					From:       now,
					To:         now,
				})

			}
		}
	}()
}
