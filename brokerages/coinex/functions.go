package coinex

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
	ApiId   string
	ApiHash string
	Token   string //remove it
}

type PeriodicOHLC struct {
	models.Market
	models.Resolution
	Response chan *brokerages.OHLCResponse
}

func (config Config) Validate() error {
	return nil
}

func (config Config) Login(input brokerages.MustImplementAsFunctionParameter) interface{} {
	params := input.(brokerages.LoginParams)
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

func (config Config) OrderBook(input brokerages.MustImplementAsFunctionParameter) interface{} {
	params := input.(brokerages.OrderBookParams)
	req := networkManager.Request{
		Method:   networkManager.GET,
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

func (config Config) MarketList(_ brokerages.MustImplementAsFunctionParameter) interface{} {
	req := networkManager.Request{
		Method:   networkManager.GET,
		Endpoint: "https://api.coinex.com/v1/market/list/",
	}

	resp, err := req.Execute()
	if err != nil {
		return &brokerages.RecentTradesResponse{
			BasicResponse: brokerages.BasicResponse{Error: err},
		}
	}
	marketList := brokerages.MarketListResponse{}
	if resp.Code == 200 {
		respStr := struct {
			Code    int      `json:"code"`
			Markets []string `json:"data"`
			Message string   `json:"message"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			marketList.Error = err
		}
		if respStr.Code == 0 {
			marketList.Markets = make([]models.Market, len(respStr.Markets))
			for i, market := range respStr.Markets {
				marketList.Markets[i].Value = market
			}
		} else {
			marketList.Error = errors.New("coinex response error: " + respStr.Message)
		}
	} else {
		marketList.Error = errors.New(resp.Status)
	}
	return marketList
}

func (config Config) RecentTrades(input brokerages.MustImplementAsFunctionParameter) interface{} {
	params := input.(brokerages.OrderBookParams)
	req := networkManager.Request{
		Method:   networkManager.GET,
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
	params := input.(brokerages.OHLCParams)
	req := networkManager.Request{
		Method:   networkManager.GET,
		Endpoint: "https://www.coinex.com/res/market/kline",
		Params: map[string]interface{}{
			"market":     params.Market.Value,
			"interval":   params.Resolution.Value,
			"start_time": params.From,
			"end_time":   params.To},
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
			Code    int             `json:"code"`
			Data    [][]interface{} `json:"data"`
			Message string          `json:"message"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &brokerages.OHLCResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: err,
				},
			}
		}
		if respStr.Code == 0 {
			ohlc := brokerages.OHLCResponse{
				Market:     params.Market,
				Resolution: params.Resolution,
				Status:     respStr.Message,
			}
			ohlc.Candles = make([]models.Candle, len(respStr.Data))
			for i := 0; i < len(respStr.Data); i++ {
				ohlc.Candles[i].Time = time.Unix(int64((respStr.Data[i][0]).(float64)), 0)
				num, err := strconv.ParseFloat(respStr.Data[i][1].(string), 64)
				if err != nil {
					return &brokerages.OHLCResponse{
						BasicResponse: brokerages.BasicResponse{
							Error: err,
						},
					}
				}
				ohlc.Candles[i].Open = num
				num, err = strconv.ParseFloat(respStr.Data[i][2].(string), 64)
				if err != nil {
					return &brokerages.OHLCResponse{
						BasicResponse: brokerages.BasicResponse{
							Error: err,
						},
					}
				}
				ohlc.Candles[i].Close = num
				num, err = strconv.ParseFloat(respStr.Data[i][3].(string), 64)
				if err != nil {
					return &brokerages.OHLCResponse{
						BasicResponse: brokerages.BasicResponse{
							Error: err,
						},
					}
				}
				ohlc.Candles[i].High = num
				num, err = strconv.ParseFloat(respStr.Data[i][4].(string), 64)
				if err != nil {
					return &brokerages.OHLCResponse{
						BasicResponse: brokerages.BasicResponse{
							Error: err,
						},
					}
				}
				ohlc.Candles[i].Low = num
				num, err = strconv.ParseFloat(respStr.Data[i][5].(string), 64)
				if err != nil {
					return &brokerages.OHLCResponse{
						BasicResponse: brokerages.BasicResponse{
							Error: err,
						},
					}
				}
				ohlc.Candles[i].Vol = num
				ohlc.Candles[i].Market = params.Market
				ohlc.Candles[i].Resolution = params.Resolution
			}
			return &ohlc
		} else {
			return &brokerages.OHLCResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: errors.New(fmt.Sprintf("error occured (%d): %s", respStr.Code, respStr.Message)),
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

func (config Config) WalletList(brokerages.MustImplementAsFunctionParameter) interface{} {
	req := networkManager.Request{
		Method:   networkManager.POST,
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
	params := input.(brokerages.WalletInfoParams)
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
	params := input.(brokerages.WalletBalanceParams)
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

func (config Config) TransactionList(input brokerages.MustImplementAsFunctionParameter) interface{} {
	params := input.(brokerages.TransactionListParams)
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

func (config Config) NewOrder(input brokerages.MustImplementAsFunctionParameter) interface{} {
	params := input.(brokerages.NewOrderParams)
	body := make(map[string]interface{})
	body["price"] = params.Order.Price
	body["amount"] = params.Order.Volume
	body["type"] = params.Order.OrderType
	body["srcCurrency"] = params.Order.SourceCurrency
	body["destCurrency"] = params.Order.DestinationCurrency
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
	params := input.(brokerages.OrderStatusParams)
	req := networkManager.Request{
		Method:   networkManager.POST,
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
	params := input.(brokerages.OrderListParams)
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
	params := input.(brokerages.UpdateOrderStatusParams)
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
				params := brokerages.OHLCParams{
					Resolution: periodic.Resolution,
					Market:     periodic.Market,
					From:       now,
					To:         now,
				}
				periodic.Response <- config.OHLC(params)

			}
		}
	}()
}
