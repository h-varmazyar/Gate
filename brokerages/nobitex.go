package brokerages

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/models"
	"strconv"
	"strings"
)

type Nobitex struct {
	Username  string
	Password  string
	Token     string
	LongToken bool
}

func (n *Nobitex) Login(totp int) error {
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

func (n *Nobitex) OrderBook(symbol string) (*api.OrderBookResponse, error) {
	req := api.Request{
		Type:     api.GET,
		Endpoint: "https://api.nobitex.ir/v2/orderbook/" + symbol,
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
				Symbol: symbol,
				Bids:   make([]models.Order, len(respStr.Bids)),
				Asks:   make([]models.Order, len(respStr.Asks)),
			}
			for i, bid := range respStr.Bids {
				price, err := strconv.ParseFloat(bid[0], 64)
				if err != nil {
					return nil, err
				}
				volume, err := strconv.ParseFloat(bid[1], 64)
				if err != nil {
					return nil, err
				}
				orderBook.Bids[i].Price = price
				orderBook.Bids[i].Volume = volume
			}
			for i, ask := range respStr.Asks {
				price, err := strconv.ParseFloat(ask[0], 64)
				if err != nil {
					return nil, err
				}
				volume, err := strconv.ParseFloat(ask[1], 64)
				if err != nil {
					return nil, err
				}
				orderBook.Asks[i].Price = price
				orderBook.Asks[i].Volume = volume
			}
			return &orderBook, nil
		} else {
			return nil, errors.New("nobitex tesponse error")
		}
	} else {
		return nil, errors.New(resp.Status)
	}
}

func (n *Nobitex) RecentTrades(symbol string) (*api.RecentTradesResponse, error) {
	req := api.Request{
		Type:     api.GET,
		Endpoint: "https://api.nobitex.ir/v2/trades/" + symbol,
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
				Symbol: symbol,
				Trades: make([]models.Trade, len(respStr.Trades)),
			}
			for i, trade := range respStr.Trades {
				recentTrade.Trades[i].Time = trade.Time
				recentTrade.Trades[i].Price, _ = strconv.ParseFloat(trade.Price, 64)
				recentTrade.Trades[i].Volume, _ = strconv.ParseFloat(trade.Volume, 64)
				recentTrade.Trades[i].Type = models.TradeType(trade.Type)
			}
			return &recentTrade, nil
		} else {
			return nil, errors.New("nobitex tesponse error")
		}
	} else {
		return nil, errors.New(resp.Status)
	}
}

func (n *Nobitex) OHLC(symbol string, resolution *models.Resolution, from, to float64) (*api.OHLCResponse, error) {
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
				Symbol:     symbol,
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

func (n *Nobitex) UserInfo() (*api.UserInfoResponse, error) {
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

func (n *Nobitex) WalletList() (*api.WalletsResponse, error) {
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
			}

			return &api.WalletsResponse{Wallets: wallets}, nil
		} else {
			return nil, errors.New("get user profile error")
		}
	} else {
		return nil, errors.New(resp.Status)
	}
}

func (n *Nobitex) WalletInfo(walletName string) (*api.WalletResponse, error) {
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
