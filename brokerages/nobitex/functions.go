package nobitex

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/brokerages/coinex"
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/models/todo"
	"github.com/mrNobody95/Gate/networkManager"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Token string
}

func (config Config) Validate() error {
	return nil
}

func (config Config) Login(params brokerages.LoginParams) *brokerages.BasicResponse {
	req := networkManager.Request{
		Method:   networkManager.POST,
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

func (config Config) OrderBook(params brokerages.OrderBookParams) *brokerages.OrderBookResponse {
	req := networkManager.Request{
		Method:   networkManager.GET,
		Endpoint: "https://api.nobitex.ir/v2/orderbook/" + params.Symbol.Name,
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
				Symbol: params.Symbol.Name,
				Bids:   make([]models.Order, len(respStr.Bids)),
				Asks:   make([]models.Order, len(respStr.Asks)),
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
				orderBook.Bids[i].AveragePrice = price
				orderBook.Bids[i].Amount, _ = strconv.ParseFloat(bid[1], 64)
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
				orderBook.Asks[i].AveragePrice = price
				orderBook.Asks[i].Amount, _ = strconv.ParseFloat(ask[1], 64)
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

func (config Config) MarketList() *brokerages.MarketListResponse {
	return &brokerages.MarketListResponse{
		BasicResponse: brokerages.BasicResponse{Error: errors.New("must be implemented")},
	}
}

func (config Config) MarketInfo(brokerages.MarketInfoParams) *brokerages.MarketInfoResponse {
	return &brokerages.MarketInfoResponse{
		BasicResponse: brokerages.BasicResponse{Error: errors.New("must be implemented")},
	}
}

func (config Config) RecentTrades(params brokerages.OrderBookParams) *brokerages.RecentTradesResponse {
	req := networkManager.Request{
		Method:   networkManager.GET,
		Endpoint: "https://api.nobitex.ir/v2/trades/" + params.Symbol.Name,
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
				Symbol: params.Symbol.Name,
				Trades: make([]todo.Trade, len(respStr.Trades)),
			}
			for i, trade := range respStr.Trades {
				recentTrade.Trades[i].Time = trade.Time
				recentTrade.Trades[i].Price, _ = strconv.ParseFloat(trade.Price, 64)
				recentTrade.Trades[i].Volume, _ = strconv.ParseFloat(trade.Volume, 64)
				recentTrade.Trades[i].Type = models.OrderType(trade.Type)
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

func (config Config) OHLC(params brokerages.OHLCParams) *brokerages.OHLCResponse {
	req := networkManager.Request{
		Method:   networkManager.GET,
		Endpoint: "https://api.nobitex.ir/market/udf/history",
		Params: map[string]interface{}{"symbol": params.Market.Name,
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
				Market:     params.Market,
				Resolution: params.Resolution,
				Status:     respStr.Status,
			}
			ohlc.Candles = make([]models.Candle, len(respStr.Time))
			for i := 0; i < len(respStr.Time); i++ {
				ohlc.Candles[i].Time = time.Unix(respStr.Time[i], 0)
				ohlc.Candles[i].Open = respStr.Open[i]
				ohlc.Candles[i].High = respStr.High[i]
				ohlc.Candles[i].Low = respStr.Low[i]
				ohlc.Candles[i].Close = respStr.Close[i]
				ohlc.Candles[i].Vol = respStr.Volume[i]
				ohlc.Candles[i].Market = params.Market
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

func (config Config) UserInfo() *brokerages.UserInfoResponse {
	req := networkManager.Request{
		Method:   networkManager.GET,
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

func (config Config) WalletList() *brokerages.WalletListResponse {
	req := networkManager.Request{
		Method:   networkManager.POST,
		Endpoint: "https://api.nobitex.ir/users/wallets/list",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", config.Token)}},
	}

	resp, err := req.Execute()
	if err != nil {
		return &brokerages.WalletListResponse{
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
			return &brokerages.WalletListResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: err,
				},
			}
		}
		if respStr.Status == "ok" {
			var wallets []models.Wallet
			for _, wallet := range respStr.Wallets {
				tmp := models.Wallet{}
				tmp.ReferenceCurrencyBalance = wallet.RialBalance
				tmp.BlockedBalance, err = strconv.ParseFloat(wallet.BlockedBalance, 64)
				if err != nil {
					continue
				}
				tmp.ActiveBalance, err = strconv.ParseFloat(wallet.ActiveBalance, 64)
				if err != nil {
					continue
				}
				tmp.TotalBalance, err = strconv.ParseFloat(wallet.Balance, 64)
				if err != nil {
					continue
				}
				tmp.Currency = wallet.Currency
				tmp.ID = wallet.Id
				wallets = append(wallets, tmp)
			}

			return &brokerages.WalletListResponse{Wallets: wallets}
		} else {
			return &brokerages.WalletListResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: errors.New("get wallet list error"),
				},
			}
		}
	} else {
		return &brokerages.WalletListResponse{
			BasicResponse: brokerages.BasicResponse{
				Error: errors.New(resp.Status),
			},
		}
	}
}

func (config Config) WalletInfo(params brokerages.WalletInfoParams) *brokerages.WalletResponse {
	req := networkManager.Request{
		Method:   networkManager.POST,
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
			wallet := models.Wallet{}
			wallet.ID = respStr.Wallets[strings.ToUpper(params.WalletName)].Id
			wallet.BlockedBalance, err = strconv.ParseFloat(respStr.Wallets[strings.ToUpper(params.WalletName)].Blocked, 64)
			if err != nil {
				return &brokerages.WalletResponse{BasicResponse: brokerages.BasicResponse{Error: err}}
			}
			wallet.TotalBalance, err = strconv.ParseFloat(respStr.Wallets[strings.ToUpper(params.WalletName)].Balance, 64)
			if err != nil {
				return &brokerages.WalletResponse{BasicResponse: brokerages.BasicResponse{Error: err}}
			}
			wallet.Currency = params.WalletName
			return &brokerages.WalletResponse{Wallet: wallet}
		} else {
			return &brokerages.WalletResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: errors.New("get wallet info error"),
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

func (config Config) WalletBalance(params brokerages.WalletBalanceParams) *brokerages.BalanceResponse {
	req := networkManager.Request{
		Method:   networkManager.POST,
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

func (config Config) TransactionList(params brokerages.TransactionListParams) *brokerages.TransactionListResponse {
	req := networkManager.Request{
		Method:   networkManager.POST,
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

func (config Config) NewOrder(params brokerages.NewOrderParams) *brokerages.OrderResponse {
	body := make(map[string]interface{})
	body["price"] = params.Price
	body["amount"] = params.Amount
	body["type"] = params.BuyOrSell
	body["srcCurrency"] = params.Market.Source
	body["destCurrency"] = params.Market.Destination
	req := networkManager.Request{
		Method:   networkManager.POST,
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
			order := models.Order{
				ClientUUID:       strings.ReplaceAll(params.ClientUUID.String(), "-", ""),
				ServerOrderId:    int64(respStr.Order.Id),
				CreatedAt:        respStr.Order.CreatedAt,
				AssetFee:         respStr.Order.Fee,
				User:             respStr.Order.User,
				AveragePrice:     price,
				Status:           models.OrderStatus(respStr.Order.Status),
				SellOrBuy:        models.OrderType(respStr.Order.Type),
				SourceAsset:      models.Asset(respStr.Order.Src),
				DestinationAsset: models.Asset(respStr.Order.Dest),
			}
			if respStr.Order.Amount != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Order.Amount, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("amount parse failed")},
					}
				}
				order.Amount = tmp
			}
			if respStr.Order.TotalPrice != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Order.TotalPrice, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("amount parse failed")},
					}
				}
				order.ExecutedPrice = tmp
			}
			if respStr.Order.Matched != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Order.Matched, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("amount parse failed")},
					}
				}
				order.ExecutedAmount = tmp
			}
			if respStr.Order.Unmatched != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Order.Unmatched, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("amount parse failed")},
					}
				}
				order.UnExecutedAmount = tmp
			}
			return &brokerages.OrderResponse{Order: order}
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

func (config Config) OrderStatus(params brokerages.OrderStatusParams) *brokerages.OrderResponse {
	req := networkManager.Request{
		Method:   networkManager.POST,
		Endpoint: "https://api.nobitex.ir/market/orders/Status",
		Headers:  map[string][]string{"Authorization": {fmt.Sprintf("Token %s", config.Token)}},
		Params:   map[string]interface{}{"id": params.ServerOrderId},
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
			order := models.Order{
				ClientUUID:       strings.ReplaceAll(params.ClientUUID.String(), "-", ""),
				ServerOrderId:    int64(respStr.Order.Id),
				CreatedAt:        respStr.Order.CreatedAt,
				AssetFee:         respStr.Order.Fee,
				User:             respStr.Order.User,
				AveragePrice:     price,
				Status:           models.OrderStatus(respStr.Order.Status),
				SellOrBuy:        models.OrderType(respStr.Order.Type),
				SourceAsset:      models.Asset(respStr.Order.Src),
				DestinationAsset: models.Asset(respStr.Order.Dest),
			}
			if respStr.Order.Amount != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Order.Amount, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("amount parse failed")},
					}
				}
				order.Amount = tmp
			}
			if respStr.Order.TotalPrice != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Order.TotalPrice, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("amount parse failed")},
					}
				}
				order.ExecutedPrice = tmp
			}
			if respStr.Order.Matched != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Order.Matched, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("amount parse failed")},
					}
				}
				order.ExecutedAmount = tmp
			}
			if respStr.Order.Unmatched != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Order.Unmatched, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("amount parse failed")},
					}
				}
				order.UnExecutedAmount = tmp
			}
			return &brokerages.OrderResponse{Order: order}
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

func (config Config) OrderList(params brokerages.OrderListParams) *brokerages.OrderListResponse {
	body := make(map[string]interface{})
	if params.Status != "" {
		body["status"] = params.Status
	}
	if params.Type != "" {
		body["type"] = params.Type
	}
	if params.Market.Source != "" {
		body["srcCurrency"] = params.Market.Source
	} else {
		return &brokerages.OrderListResponse{
			BasicResponse: brokerages.BasicResponse{Error: errors.New("please specify Source Currency")},
		}
	}
	if params.Market.Destination != "" {
		body["dstCurrency"] = params.Market.Destination
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
		Method:   networkManager.POST,
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
			orders := make([]models.Order, len(respStr.Orders))
			for i, order := range respStr.Orders {
				price, err := strconv.ParseFloat(order.Price, 64)
				if err != nil {
					continue
				}
				orders[i] = models.Order{
					ClientUUID:       strings.ReplaceAll(params.ClientUUID.String(), "-", ""),
					ServerOrderId:    int64(order.Id),
					CreatedAt:        order.CreatedAt,
					AssetFee:         order.Fee,
					User:             order.User,
					AveragePrice:     price,
					Status:           models.OrderStatus(order.Status),
					SellOrBuy:        models.OrderType(order.Type),
					SourceAsset:      models.Asset(order.Src),
					DestinationAsset: models.Asset(order.Dest),
				}
				if order.Amount != "" {
					tmp, parseErr := strconv.ParseFloat(order.Amount, 64)
					if parseErr != nil {
						return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{
							Error: errors.New("amount parse failed")},
						}
					}
					orders[i].Amount = tmp
				}
				if order.TotalPrice != "" {
					tmp, parseErr := strconv.ParseFloat(order.TotalPrice, 64)
					if parseErr != nil {
						return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{
							Error: errors.New("amount parse failed")},
						}
					}
					orders[i].ExecutedPrice = tmp
				}
				if order.Matched != "" {
					tmp, parseErr := strconv.ParseFloat(order.Matched, 64)
					if parseErr != nil {
						return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{
							Error: errors.New("amount parse failed")},
						}
					}
					orders[i].ExecutedAmount = tmp
				}
				if order.Unmatched != "" {
					tmp, parseErr := strconv.ParseFloat(order.Unmatched, 64)
					if parseErr != nil {
						return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{
							Error: errors.New("amount parse failed")},
						}
					}
					orders[i].UnExecutedAmount = tmp
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

func (config Config) UpdateOrderStatus(params brokerages.UpdateOrderStatusParams) *brokerages.UpdateOrderStatusResponse {
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
		Method:   networkManager.POST,
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
			Status        string             `json:"Status"`
			Message       string             `json:"message"`
			UpdatedStatus models.OrderStatus `json:"updatedStatus"`
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

func (config Config) CancelOrder(brokerages.CancelOrderParams) *brokerages.OrderResponse {
	return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{Error: coinex.ErrMustBeImplemented}}
}
