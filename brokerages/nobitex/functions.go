package nobitex

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/models/todo"
	"github.com/mrNobody95/Gate/networkManager"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	models.Brokerage
	Token     string
	LongToken bool
}

type PeriodicOHLC struct {
	models.Symbol
	models.Resolution
	Response chan *brokerages.OHLCResponse
}

func (config Config) Validate() error {
	return nil
}

func (config Config) GetName() models.BrokerageName {
	return config.Name
}

func (config Config) Login(input brokerages.MustImplementAsFunctionParameter) interface{} {
	params := input.(LoginParams)
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

func (config Config) OrderBook(input brokerages.MustImplementAsFunctionParameter) interface{} {
	params := input.(OrderBookParams)
	req := networkManager.Request{
		Type:     networkManager.GET,
		Endpoint: "https://api.nobitex.ir/v2/orderbook/" + params.Symbol.Value,
	}

	resp, err := req.Execute()
	if err != nil {
		return &brokerages.OrderBookResponse{
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
			return &brokerages.OrderBookResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: err,
				},
			}
		}
		if respStr.Status == "ok" {
			orderBook := brokerages.OrderBookResponse{
				Symbol: params.Symbol.Value,
				Bids:   make([]todo.Order, len(respStr.Bids)),
				Asks:   make([]todo.Order, len(respStr.Asks)),
			}
			for i, bid := range respStr.Bids {
				price, err := strconv.ParseFloat(bid[0], 64)
				if err != nil {
					return &brokerages.OrderBookResponse{
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
					return &brokerages.OrderBookResponse{
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
			return &brokerages.OrderBookResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: errors.New("nobitex tesponse error"),
				},
			}
		}
	} else {
		return &brokerages.OrderBookResponse{
			BasicResponse: brokerages.BasicResponse{
				Error: errors.New(resp.Status),
			},
		}
	}
}

func (config Config) RecentTrades(input brokerages.MustImplementAsFunctionParameter) interface{} {
	params := input.(OrderBookParams)
	req := networkManager.Request{
		Type:     networkManager.GET,
		Endpoint: "https://api.nobitex.ir/v2/trades/" + params.Symbol.Value,
	}

	resp, err := req.Execute()
	if err != nil {
		return &brokerages.RecentTradesResponse{
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
			return &brokerages.RecentTradesResponse{
				BasicResponse: brokerages.BasicResponse{Error: err},
			}
		}
		if respStr.Status == "ok" {
			recentTrade := brokerages.RecentTradesResponse{
				Symbol: params.Symbol.Value,
				Trades: make([]todo.Trade, len(respStr.Trades)),
			}
			for i, trade := range respStr.Trades {
				recentTrade.Trades[i].Time = trade.Time
				recentTrade.Trades[i].Price, _ = strconv.ParseFloat(trade.Price, 64)
				recentTrade.Trades[i].Volume, _ = strconv.ParseFloat(trade.Volume, 64)
				recentTrade.Trades[i].Type = todo.OrderType(trade.Type)
			}
			return &recentTrade
		} else {
			return &brokerages.RecentTradesResponse{
				BasicResponse: brokerages.BasicResponse{Error: errors.New("nobitex response error")},
			}
		}
	} else {
		return &brokerages.RecentTradesResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New(resp.Status)},
		}
	}
}

func (config Config) MarketStats(input brokerages.MustImplementAsFunctionParameter) interface{} {
	return &brokerages.MarketStatusResponse{}
}

func (config Config) OHLC(input brokerages.MustImplementAsFunctionParameter) *brokerages.OHLCResponse {
	params := input.(OHLCParams)
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
		return &brokerages.OHLCResponse{
			BasicResponse: brokerages.BasicResponse{
				Error: err,
			},
		}
	}
	if resp.Code == 200 {
		respStr := struct {
			Status string    `json:"s"`
			Time   []int64   `json:"t"`
			Open   []float64 `json:"o"`
			High   []float64 `json:"h"`
			Low    []float64 `json:"l"`
			Close  []float64 `json:"c"`
			Volume []float64 `json:"v"`
			Error  string    `json:"errmsg"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &brokerages.OHLCResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: err,
				},
			}
		}
		if respStr.Status == "ok" {
			ohlc := brokerages.OHLCResponse{
				Symbol:     params.Symbol,
				Resolution: params.Resolution,
				Status:     respStr.Status,
			}
			ohlc.Candles = make([]models.Candle, len(respStr.Time))
			fmt.Println("candle:", len(respStr.Time))
			for i := 0; i < len(respStr.Time); i++ {
				ohlc.Candles[i].Time = time.Unix(respStr.Time[i], 0)
				ohlc.Candles[i].Open = respStr.Open[i]
				ohlc.Candles[i].High = respStr.High[i]
				ohlc.Candles[i].Low = respStr.Low[i]
				ohlc.Candles[i].Close = respStr.Close[i]
				ohlc.Candles[i].Vol = respStr.Volume[i]
				ohlc.Candles[i].Symbol = params.Symbol

			}
			return &ohlc
		} else {
			return &brokerages.OHLCResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: errors.New(respStr.Error),
				},
			}
		}
	} else {
		return &brokerages.OHLCResponse{
			BasicResponse: brokerages.BasicResponse{
				Error: errors.New(resp.Status),
			},
		}
	}
}

func (config Config) UserInfo(brokerages.MustImplementAsFunctionParameter) interface{} {
	req := networkManager.Request{
		Type:     networkManager.GET,
		Endpoint: "https://api.nobitex.ir/users/profile",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", config.Token)}},
	}

	resp, err := req.Execute()
	if err != nil {
		return &brokerages.UserInfoResponse{
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
			return &brokerages.UserInfoResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: err,
				},
			}
		}
		if respStr.Status == "ok" {
			bankAccounts := make([]todo.BankAccount, len(respStr.Profile.BankAccounts))
			for i, account := range respStr.Profile.BankAccounts {
				bankAccounts[i].AccountNumber = account.Number
				bankAccounts[i].OwnerName = account.Owner
				bankAccounts[i].BankName = account.Bank
				bankAccounts[i].Status = account.Status
				bankAccounts[i].IBAN = account.IBAN
			}
			userInfo := brokerages.UserInfoResponse{
				User: todo.User{
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
			return &brokerages.UserInfoResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: errors.New("get user profile error"),
				},
			}
		}
	} else {
		return &brokerages.UserInfoResponse{
			BasicResponse: brokerages.BasicResponse{
				Error: errors.New(resp.Status),
			},
		}
	}
}

func (config Config) WalletList(brokerages.MustImplementAsFunctionParameter) interface{} {
	req := networkManager.Request{
		Type:     networkManager.POST,
		Endpoint: "https://api.nobitex.ir/users/wallets/list",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", config.Token)}},
	}

	resp, err := req.Execute()
	if err != nil {
		return &brokerages.WalletsResponse{
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
			return &brokerages.WalletsResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: err,
				},
			}
		}
		if respStr.Status == "ok" {
			wallets := make([]todo.Wallet, len(respStr.Wallets))
			for i, wallet := range respStr.Wallets {
				wallets[i].ReferenceCurrencyBalance = wallet.RialBalance
				wallets[i].BlockedBalance = wallet.BlockedBalance
				wallets[i].ActiveBalance = wallet.ActiveBalance
				wallets[i].TotalBalance = wallet.Balance
				wallets[i].Currency = wallet.Currency
				wallets[i].ID = wallet.Id
			}

			return &brokerages.WalletsResponse{Wallets: wallets}
		} else {
			return &brokerages.WalletsResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: errors.New("get user profile error"),
				},
			}
		}
	} else {
		return &brokerages.WalletsResponse{
			BasicResponse: brokerages.BasicResponse{
				Error: errors.New(resp.Status),
			},
		}
	}
}

func (config Config) WalletInfo(input brokerages.MustImplementAsFunctionParameter) interface{} {
	params := input.(WalletInfoParams)
	req := networkManager.Request{
		Type:     networkManager.POST,
		Endpoint: "https://api.nobitex.ir/v2/wallets",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", config.Token)}},
		Params:   map[string]interface{}{"currencies": params.WalletName},
	}

	resp, err := req.Execute()
	if err != nil {
		return &brokerages.WalletResponse{
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
			return &brokerages.WalletResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: err,
				},
			}
		}
		if respStr.Status == "ok" {
			return &brokerages.WalletResponse{Wallet: todo.Wallet{
				ID:             respStr.Wallets[strings.ToUpper(params.WalletName)].Id,
				BlockedBalance: respStr.Wallets[strings.ToUpper(params.WalletName)].Blocked,
				TotalBalance:   respStr.Wallets[strings.ToUpper(params.WalletName)].Balance,
				Currency:       params.WalletName,
			}}
		} else {
			return &brokerages.WalletResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: errors.New("get user profile error"),
				},
			}
		}
	} else {
		return &brokerages.WalletResponse{
			BasicResponse: brokerages.BasicResponse{
				Error: errors.New(resp.Status),
			},
		}
	}
}

func (config Config) WalletBalance(input brokerages.MustImplementAsFunctionParameter) interface{} {
	params := input.(WalletBalanceParams)
	req := networkManager.Request{
		Type:     networkManager.POST,
		Endpoint: "https://api.nobitex.ir/users/wallets/balance",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", config.Token)}},
		Params:   map[string]interface{}{"currency": params.Currency},
	}

	resp, err := req.Execute()
	if err != nil {
		return &brokerages.BalanceResponse{
			BasicResponse: brokerages.BasicResponse{Error: err},
		}
	}
	if resp.Code == 200 {
		respStr := struct {
			Status  string `json:"Status"`
			Balance string `json:"balance"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &brokerages.BalanceResponse{
				BasicResponse: brokerages.BasicResponse{Error: err},
			}
		}
		if respStr.Status == "ok" {
			return &brokerages.BalanceResponse{
				Symbol:  params.Currency,
				Balance: respStr.Balance,
			}
		} else {
			return &brokerages.BalanceResponse{
				BasicResponse: brokerages.BasicResponse{Error: errors.New("internal server error (wallet balance)")},
			}
		}
	} else {
		return &brokerages.BalanceResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New(resp.Status)},
		}
	}
}

func (config Config) TransactionList(input brokerages.MustImplementAsFunctionParameter) interface{} {
	params := input.(TransactionListParams)
	req := networkManager.Request{
		Type:     networkManager.POST,
		Endpoint: "https://api.nobitex.ir/users/wallets/transactions/list",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", config.Token)}},
		Params:   map[string]interface{}{"wallet": params.WalletID},
	}

	resp, err := req.Execute()
	if err != nil {
		return &brokerages.TransactionListResponse{
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
			return &brokerages.TransactionListResponse{
				BasicResponse: brokerages.BasicResponse{Error: err},
			}
		}
		if respStr.Status == "ok" {
			transactions := make([]todo.Transaction, len(respStr.Transactions))
			for i, transaction := range respStr.Transactions {
				transactions[i] = todo.Transaction{
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
			return &brokerages.TransactionListResponse{Transactions: transactions}
		} else {
			return &brokerages.TransactionListResponse{
				BasicResponse: brokerages.BasicResponse{Error: errors.New("internal server error (wallet balance)")},
			}
		}
	} else {
		return &brokerages.TransactionListResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New(resp.Status)},
		}
	}
}

func (config Config) NewOrder(input brokerages.MustImplementAsFunctionParameter) interface{} {
	params := input.(NewOrderParams)
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
		return &brokerages.OrderResponse{
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
			return &brokerages.OrderResponse{
				BasicResponse: brokerages.BasicResponse{Error: err},
			}
		}
		if respStr.Status == "ok" {
			price, err := strconv.ParseFloat(respStr.Order.Price, 64)
			if err != nil {
				return &brokerages.OrderResponse{
					BasicResponse: brokerages.BasicResponse{Error: err},
				}
			}
			return &brokerages.OrderResponse{
				Order: todo.Order{
					Model: gorm.Model{
						ID:        respStr.Order.Id,
						CreatedAt: respStr.Order.CreatedAt,
					},
					Fee:                 respStr.Order.Fee,
					User:                respStr.Order.User,
					Price:               price,
					Status:              todo.OrderStatus(respStr.Order.Status),
					Volume:              respStr.Order.Amount,
					OrderType:           todo.OrderType(respStr.Order.Type),
					TotalPrice:          respStr.Order.TotalPrice,
					MatchedVolume:       respStr.Order.Matched,
					SourceCurrency:      respStr.Order.Src,
					UnMatchedVolume:     respStr.Order.Unmatched,
					DestinationCurrency: respStr.Order.Dest,
				},
			}
		} else {
			return &brokerages.OrderResponse{
				BasicResponse: brokerages.BasicResponse{Error: errors.New(respStr.Message)},
			}
		}
	} else {
		return &brokerages.OrderResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New(resp.Status)},
		}
	}
}

func (config Config) OrderStatus(input brokerages.MustImplementAsFunctionParameter) interface{} {
	params := input.(OrderStatusParams)
	req := networkManager.Request{
		Type:     networkManager.POST,
		Endpoint: "https://api.nobitex.ir/market/orders/Status",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", config.Token)}},
		Params:   map[string]interface{}{"id": params.OrderId},
	}

	resp, err := req.Execute()
	if err != nil {
		return &brokerages.OrderResponse{
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
			return &brokerages.OrderResponse{
				BasicResponse: brokerages.BasicResponse{Error: err},
			}
		}
		if respStr.Status == "ok" {
			price, err := strconv.ParseFloat(respStr.Order.Price, 64)
			if err != nil {
				return &brokerages.OrderResponse{
					BasicResponse: brokerages.BasicResponse{Error: err},
				}
			}
			return &brokerages.OrderResponse{
				Order: todo.Order{
					Model: gorm.Model{
						ID:        respStr.Order.Id,
						CreatedAt: respStr.Order.CreatedAt,
					},
					Fee:                 respStr.Order.Fee,
					User:                respStr.Order.User,
					Price:               price,
					Status:              todo.OrderStatus(respStr.Order.Status),
					Volume:              respStr.Order.Amount,
					OrderType:           todo.OrderType(respStr.Order.Type),
					TotalPrice:          respStr.Order.TotalPrice,
					MatchedVolume:       respStr.Order.Matched,
					SourceCurrency:      respStr.Order.Src,
					UnMatchedVolume:     respStr.Order.Unmatched,
					DestinationCurrency: respStr.Order.Dest,
				},
			}
		} else {
			return &brokerages.OrderResponse{
				BasicResponse: brokerages.BasicResponse{Error: errors.New(respStr.Message)},
			}
		}
	} else {
		return &brokerages.OrderResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New(resp.Status)},
		}
	}
}

func (config Config) OrderList(input brokerages.MustImplementAsFunctionParameter) interface{} {
	params := input.(OrderListParams)
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
		return &brokerages.OrderListResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New("please specify Source Currency")},
		}
	}
	if params.Destination != "" {
		body["dstCurrency"] = params.Destination
	} else {
		return &brokerages.OrderListResponse{
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
		return &brokerages.OrderListResponse{
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
			return &brokerages.OrderListResponse{
				BasicResponse: brokerages.BasicResponse{Error: err},
			}
		}
		if respStr.Status == "ok" {
			orders := make([]todo.Order, len(respStr.Orders))
			for i, order := range respStr.Orders {
				price, err := strconv.ParseFloat(order.Price, 64)
				if err != nil {
					continue
				}
				orders[i] = todo.Order{
					Model: gorm.Model{
						ID:        order.Id,
						CreatedAt: order.CreatedAt,
					},
					Fee:                 order.Fee,
					User:                order.User,
					Price:               price,
					Status:              todo.OrderStatus(order.Status),
					Volume:              order.Amount,
					OrderType:           todo.OrderType(order.Type),
					TotalPrice:          order.TotalPrice,
					MatchedVolume:       order.Matched,
					SourceCurrency:      order.Src,
					UnMatchedVolume:     order.Unmatched,
					DestinationCurrency: order.Dest,
				}
			}
			return &brokerages.OrderListResponse{Orders: orders}
		} else {
			return &brokerages.OrderListResponse{
				BasicResponse: brokerages.BasicResponse{Error: errors.New(respStr.Message)},
			}
		}
	} else {
		return &brokerages.OrderListResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New(resp.Status)},
		}
	}
}

func (config Config) UpdateOrderStatus(input brokerages.MustImplementAsFunctionParameter) interface{} {
	params := input.(UpdateOrderStatusParams)
	body := make(map[string]interface{})
	if params.OrderId < 0 {
		body["Order"] = params.OrderId
	} else {
		return &brokerages.UpdateOrderStatusResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New("please specify Order id Currency")},
		}
	}
	if params.NewStatus != "" {
		body["Status"] = params.NewStatus
	} else {
		return &brokerages.UpdateOrderStatusResponse{
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
		return &brokerages.UpdateOrderStatusResponse{
			BasicResponse: brokerages.BasicResponse{Error: err},
		}
	}
	if resp.Code == 200 {
		respStr := struct {
			Status        string           `json:"Status"`
			Message       string           `json:"message"`
			UpdatedStatus todo.OrderStatus `json:"updatedStatus"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &brokerages.UpdateOrderStatusResponse{
				BasicResponse: brokerages.BasicResponse{Error: err},
			}
		}
		if respStr.Status == "ok" {
			return &brokerages.UpdateOrderStatusResponse{NewStatus: respStr.UpdatedStatus}
		} else {
			return &brokerages.UpdateOrderStatusResponse{
				BasicResponse: brokerages.BasicResponse{Error: errors.New(respStr.Message)},
			}
		}
	} else {
		return &brokerages.UpdateOrderStatusResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New(resp.Status)},
		}
	}
}

func (config Config) SubscribePeriodicOHLC(periodic PeriodicOHLC, period time.Duration, endSignal chan bool) {
	ticker := time.NewTicker(period)
	go func() {
		for {
			select {
			case <-endSignal:
				return
			case _ = <-ticker.C:
				now := time.Now().Unix()
				params := OHLCParams{
					Resolution: periodic.Resolution,
					Symbol:     periodic.Symbol,
					From:       now,
					To:         now,
				}
				periodic.Response <- config.OHLC(params)

			}
		}
	}()
}
