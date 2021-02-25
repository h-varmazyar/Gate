package brokerages

import (
	"encoding/json"
	"errors"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/models"
	"strconv"
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
		req.Headers = map[string]interface{}{"X-TOTP": totp}
	}
	resp := req.Execute()
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
		return errors.New(resp.ErrorMessage)
	}
}

func (n *Nobitex) OrderBook(symbol string) (*api.OrderBookResponse, error) {
	req := api.Request{
		Type:     api.GET,
		Endpoint: "https://api.nobitex.ir/v2/orderbook/" + symbol,
	}

	resp := req.Execute()
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
		return nil, errors.New(resp.ErrorMessage)
	}
}

func (n *Nobitex) RecentTrades(symbol string) (*api.RecentTradesResponse, error) {
	req := api.Request{
		Type:     api.GET,
		Endpoint: "https://api.nobitex.ir/v2/trades/" + symbol,
	}

	resp := req.Execute()
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
		return nil, errors.New(resp.ErrorMessage)
	}
}

func (n *Nobitex) OHLC(symbol string, resolution *models.Resolution, from, to float64) (*api.OHLCResponse, error) {
	req := api.Request{
		Type:     api.GET,
		Endpoint: "https://api.nobitex.ir/market/udf/history",
		Params:   map[string]interface{}{"symbol": symbol, "resolution": resolution.Value, "from": from, "to": to},
	}

	resp := req.Execute()
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
		return nil, errors.New(resp.ErrorMessage)
	}
}
