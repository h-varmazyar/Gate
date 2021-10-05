package coinex

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/networkManager"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var ErrMustBeImplemented = errors.New("must be implemented")

type Config struct {
	AccessId  string
	SecretKey string
}

func (config Config) Validate() error {
	return nil
}

func (config Config) Login(brokerages.LoginParams) *brokerages.BasicResponse {
	return nil
}

func (config Config) OrderBook(brokerages.OrderBookParams) *brokerages.OrderBookResponse {
	return &brokerages.OrderBookResponse{BasicResponse: brokerages.BasicResponse{Error: ErrMustBeImplemented}}
}

//market endpoints
func (config Config) OHLC(params brokerages.OHLCParams) *brokerages.OHLCResponse {
	req := networkManager.Request{
		Method:   networkManager.GET,
		Endpoint: "https://www.coinex.com/res/market/kline",
		Params: map[string]interface{}{
			"market":     params.Market.Name,
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

func (config Config) MarketList() *brokerages.MarketListResponse {
	req := networkManager.Request{
		Method:   networkManager.GET,
		Endpoint: "https://api.coinex.com/v1/market/list/",
	}

	resp, err := req.Execute()
	if err != nil {
		return &brokerages.MarketListResponse{
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
		if respStr.Code == ResponseSuccess {
			marketList.Markets = make([]models.Market, len(respStr.Markets))
			for i, market := range respStr.Markets {
				marketList.Markets[i].Name = market
			}
		} else {
			marketList.Error = errors.New("coinex response error: " + respStr.Message)
		}
	} else {
		marketList.Error = errors.New(resp.Status)
	}
	return &marketList
}

func (config Config) MarketInfo(params brokerages.MarketInfoParams) *brokerages.MarketInfoResponse {
	req := networkManager.Request{
		Method:   networkManager.GET,
		Endpoint: "https://api.coinex.com/v1/market/detail",
		Params:   map[string]interface{}{"market": params.MarketName},
	}

	resp, err := req.Execute()
	if err != nil {
		return &brokerages.MarketInfoResponse{
			BasicResponse: brokerages.BasicResponse{Error: err},
		}
	}
	marketInfo := brokerages.MarketInfoResponse{}
	if resp.Code == 200 {
		respStr := struct {
			Code           int    `json:"code"`
			Message        string `json:"message"`
			TackerFeeRate  string `json:"tacker_fee_rate"`
			MakerFeeRate   string `json:"maker_fee_rate"`
			MinAmount      string `json:"min_amount"`
			TradingName    string `json:"trading_name"`
			TradingDecimal int    `json:"trading_decimal"`
			PricingName    string `json:"pricing_name"`
			PricingDecimal int    `json:"pricing_decimal"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			marketInfo.Error = err
		}
		if respStr.Code == ResponseSuccess {
			marketInfo.Market.TakerFeeRate, err = strconv.ParseFloat(respStr.TackerFeeRate, 64)
			if err != nil {
				return &brokerages.MarketInfoResponse{
					BasicResponse: brokerages.BasicResponse{
						Error: err,
					},
				}
			}
			marketInfo.Market.MakerFeeRate, err = strconv.ParseFloat(respStr.MakerFeeRate, 64)
			if err != nil {
				return &brokerages.MarketInfoResponse{
					BasicResponse: brokerages.BasicResponse{
						Error: err,
					},
				}
			}
			marketInfo.Market.MinAmount, err = strconv.ParseFloat(respStr.MinAmount, 64)
			if err != nil {
				return &brokerages.MarketInfoResponse{
					BasicResponse: brokerages.BasicResponse{
						Error: err,
					},
				}
			}
			marketInfo.Market.TradingName = respStr.TradingName
			marketInfo.Market.TradingDecimal = respStr.TradingDecimal
			marketInfo.Market.PricingName = respStr.PricingName
			marketInfo.Market.PricingDecimal = respStr.PricingDecimal
		} else {
			marketInfo.Error = errors.New("coinex response error: " + respStr.Message)
		}
	} else {
		marketInfo.Error = errors.New(resp.Status)
	}
	return &marketInfo
}

//account endpoints
func (config Config) WalletList() *brokerages.WalletListResponse {
	req := networkManager.Request{
		Method:   networkManager.GET,
		Endpoint: "https://api.coinex.com/v1/balance/info",
		Headers:  map[string][]string{"authorization": {config.AccessId}},
		Params: map[string]interface{}{
			"access_id": config.AccessId,
			"tonce":     time.Now().UnixNano() / 1000,
		},
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
			Code int `json:"code"`
			Data map[string]struct {
				Available string `json:"available"`
				Frozen    string `json:"frozen"`
			} `json:"data"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &brokerages.WalletListResponse{
				BasicResponse: brokerages.BasicResponse{
					Error: err,
				},
			}
		}
		if respStr.Code == ResponseSuccess {
			var wallets []models.Wallet
			for key, value := range respStr.Data {
				wallet := models.Wallet{}
				wallet.Currency = key
				wallet.BlockedBalance, err = strconv.ParseFloat(value.Frozen, 64)
				if err != nil {
					continue
				}
				wallet.TotalBalance, err = strconv.ParseFloat(value.Available, 64)
				if err != nil {
					continue
				}
				wallet.ActiveBalance = wallet.TotalBalance - wallet.BlockedBalance
				wallets = append(wallets, wallet)
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

//trading endpoints
func (config Config) NewOrder(params brokerages.NewOrderParams) *brokerages.OrderResponse {
	endpoint := ""
	var queryParams map[string]interface{}
	switch params.OrderKind {
	case models.LimitOrderKind:
		endpoint = "https://api.coinex.com/v1/order/limit"
		queryParams = map[string]interface{}{
			"access_id": config.AccessId,
			"source_id": fmt.Sprintf("%d", rand.Int()),
			"market":    params.Market.Name,
			"amount":    fmt.Sprintf("%f", params.Amount),
			"option":    params.Option,
			"price":     fmt.Sprintf("%f", params.Price),
			"tonce":     time.Now().UnixNano() / 1000,
			"client_id": strings.ReplaceAll(params.ClientUUID.String(), "-", ""),
			"type":      params.BuyOrSell,
			"hide":      params.HideOrder,
		}
	case models.MarketOrderKind:
		endpoint = "https://api.coinex.com/v1/order/market"
		queryParams = map[string]interface{}{
			"access_id": config.AccessId,
			"market":    params.Market.Name,
			"type":      params.BuyOrSell,
			"amount":    fmt.Sprintf("%f", params.Amount),
			"tonce":     time.Now().UnixNano() / 1000,
			"client_id": strings.ReplaceAll(params.ClientUUID.String(), "-", ""),
			"source_id": fmt.Sprintf("%d", rand.Int()),
		}
	case models.StopLimitOrderKind:
		endpoint = "https://api.coinex.com/v1/order/stop/limit"
		queryParams = map[string]interface{}{
			"access_id":  config.AccessId,
			"market":     params.Market.Name,
			"type":       params.BuyOrSell,
			"amount":     fmt.Sprintf("%f", params.Amount),
			"price":      fmt.Sprintf("%f", params.Price),
			"stop_price": fmt.Sprintf("%f", params.StopPrice),
			"source_id":  fmt.Sprintf("%d", rand.Int()),
			"option":     params.Option,
			"tonce":      time.Now().UnixNano() / 1000,
			"client_id":  strings.ReplaceAll(params.ClientUUID.String(), "-", ""),
			"hide":       params.HideOrder,
		}
	case models.IOCOrderKind:
		endpoint = "https://api.coinex.com/v1/order/stop/limit"
		queryParams = map[string]interface{}{
			"access_id": config.AccessId,
			"market":    params.Market.Name,
			"type":      params.BuyOrSell,
			"amount":    fmt.Sprintf("%f", params.Amount),
			"price":     fmt.Sprintf("%f", params.Price),
			"source_id": fmt.Sprintf("%d", rand.Int()),
			"tonce":     time.Now().UnixNano() / 1000,
			"client_id": strings.ReplaceAll(params.ClientUUID.String(), "-", ""),
		}
	case models.MultipleLimitOrderKind:
		//todo: must be implemented
		return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{Error: errors.New("multiple limit order not implemented yet")}}
	default:
		return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{Error: errors.New("invalid order kind")}}
	}
	req := networkManager.Request{
		Method:   networkManager.POST,
		Endpoint: endpoint,
		Headers:  map[string][]string{"authorization": {config.AccessId}},
		Params:   queryParams,
	}

	resp, err := req.Execute()
	if err != nil {
		return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{Error: err}}
	}
	if resp.Code == 200 {
		respStr := struct {
			Code int `json:"code"`
			Data struct {
				Amount       string `json:"amount"`
				AssetFee     string `json:"asset_fee"`
				AvgPrice     string `json:"avg_price"`
				CreateTime   int64  `json:"create_time"`
				DealAmount   string `json:"deal_amount"`
				DealFee      string `json:"deal_fee"`
				DealMoney    string `json:"deal_money"`
				FeeAsset     string `json:"fee_asset"`
				FeeDiscount  string `json:"fee_discount"`
				FinishedTime int64  `json:"finished_time"`
				Id           int64  `json:"id"`
				Left         string `json:"left"`
				MakerFeeRate string `json:"maker_fee_rate"`
				MoneyFee     string `json:"money_fee"`
				Market       string `json:"market"`
				OrderType    string `json:"order_type"`
				Price        string `json:"price"`
				Status       string `json:"status"`
				StockFee     string `json:"stock_fee"`
				TakerFeeRate string `json:"taker_fee_rate"`
				Type         string `json:"type"`
				ClientId     string `json:"client_id"`
			} `json:"data"`
			Message string `json:"message"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{Error: err}}
		}
		if respStr.Code == ResponseSuccess {
			order := models.Order{
				ClientUUID:    respStr.Data.ClientId,
				ServerOrderId: respStr.Data.Id,
				CreatedAt:     time.Unix(respStr.Data.CreateTime, 0),
				FinishedAt:    time.Unix(respStr.Data.FinishedTime, 0),
				Status:        models.OrderStatus(respStr.Data.Status),
				Market:        params.Market,
				SellOrBuy:     models.OrderType(respStr.Data.Type),
				OrderKind:     models.OrderKind(respStr.Data.OrderType),
				FeeAsset:      models.Asset(respStr.Data.FeeAsset),
			}
			if respStr.Data.Amount != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.Amount, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("order amount parse failed")},
					}
				}
				order.Amount = tmp
			}

			if respStr.Data.DealAmount != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.DealAmount, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("executed amount parse failed")},
					}
				}
				order.ExecutedAmount = tmp
			}

			if respStr.Data.Left != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.Left, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("unexecuted amount parse failed")},
					}
				}
				order.UnExecutedAmount = tmp
			}

			if respStr.Data.DealMoney != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.DealMoney, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("executed price parse failed")},
					}
				}
				order.ExecutedPrice = tmp
			}

			if respStr.Data.MakerFeeRate != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.MakerFeeRate, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("maker fee rate parse failed")},
					}
				}
				order.MakerFeeRate = tmp
			}

			if respStr.Data.TakerFeeRate != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.TakerFeeRate, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("taker fee rate parse failed")},
					}
				}
				order.TakerFeeRate = tmp
			}

			if respStr.Data.AvgPrice != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.AvgPrice, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("average price parse failed")},
					}
				}
				order.AveragePrice = tmp
			}

			if respStr.Data.AssetFee != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.AssetFee, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("transaction fee parse failed")},
					}
				}
				order.TransactionFee = tmp
			}

			if respStr.Data.FeeDiscount != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.FeeDiscount, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("fee discount parse failed")},
					}
				}
				order.FeeDiscount = tmp
			}
			if respStr.Data.AssetFee != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.AssetFee, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("asset fee parse failed")},
					}
				}
				order.AssetFee = tmp
			}

			if respStr.Data.MoneyFee != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.MoneyFee, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("money fee parse failed")},
					}
				}
				order.MoneyFee = tmp
			}

			if respStr.Data.StockFee != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.StockFee, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("stock fee parse failed")},
					}
				}
				order.StockFee = tmp
			}
			return &brokerages.OrderResponse{Order: order}
		} else {
			return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
				Error: errors.New("get wallet list error")},
			}
		}
	} else {
		return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{Error: errors.New(resp.Status)}}
	}
}

func (config Config) CancelOrder(params brokerages.CancelOrderParams) *brokerages.OrderResponse {
	var queryParams map[string]interface{}
	queryParams = map[string]interface{}{
		"access_id":  config.AccessId,
		"account_id": 0,
		"market":     params.Market.Name,
		"tonce":      time.Now().UnixNano() / 1000,
		"client_id":  strings.ReplaceAll(params.ClientUUID.String(), "-", ""),
	}
	if !params.AllOrders {
		queryParams["id"] = params.ServerOrderId
		queryParams["type"] = params.IsBuy
	}

	req := networkManager.Request{
		Method:   networkManager.DELETE,
		Endpoint: "https://api.coinex.com/v1/order/pending",
		Headers:  map[string][]string{"authorization": {config.AccessId}},
		Params:   queryParams,
	}

	resp, err := req.Execute()
	if err != nil {
		return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{Error: err}}
	}
	if resp.Code == 200 {
		respStr := struct {
			Code int `json:"code"`
			Data struct {
				Amount       string `json:"amount"`
				AssetFee     string `json:"asset_fee"`
				AvgPrice     string `json:"avg_price"`
				CreateTime   int64  `json:"create_time"`
				DealAmount   string `json:"deal_amount"`
				DealFee      string `json:"deal_fee"`
				DealMoney    string `json:"deal_money"`
				FeeAsset     string `json:"fee_asset"`
				FeeDiscount  string `json:"fee_discount"`
				FinishedTime int64  `json:"finished_time"`
				Id           int64  `json:"id"`
				Left         string `json:"left"`
				MakerFeeRate string `json:"maker_fee_rate"`
				MoneyFee     string `json:"money_fee"`
				Market       string `json:"market"`
				OrderType    string `json:"order_type"`
				Price        string `json:"price"`
				Status       string `json:"status"`
				StockFee     string `json:"stock_fee"`
				TakerFeeRate string `json:"taker_fee_rate"`
				Type         string `json:"type"`
				ClientId     string `json:"client_id"`
			} `json:"data"`
			Message string `json:"message"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{Error: err}}
		}
		if respStr.Code == ResponseSuccess {
			order := models.Order{
				ClientUUID:    respStr.Data.ClientId,
				ServerOrderId: respStr.Data.Id,
				CreatedAt:     time.Unix(respStr.Data.CreateTime, 0),
				FinishedAt:    time.Unix(respStr.Data.FinishedTime, 0),
				Status:        models.OrderStatus(respStr.Data.Status),
				Market:        params.Market,
				SellOrBuy:     models.OrderType(respStr.Data.Type),
				OrderKind:     models.OrderKind(respStr.Data.OrderType),
				FeeAsset:      models.Asset(respStr.Data.FeeAsset),
			}
			if respStr.Data.Amount != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.Amount, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("order amount parse failed")},
					}
				}
				order.Amount = tmp
			}

			if respStr.Data.DealAmount != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.DealAmount, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("executed amount parse failed")},
					}
				}
				order.ExecutedAmount = tmp
			}

			if respStr.Data.Left != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.Left, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("unexecuted amount parse failed")},
					}
				}
				order.UnExecutedAmount = tmp
			}

			if respStr.Data.DealMoney != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.DealMoney, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("executed price parse failed")},
					}
				}
				order.ExecutedPrice = tmp
			}

			if respStr.Data.MakerFeeRate != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.MakerFeeRate, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("maker fee rate parse failed")},
					}
				}
				order.MakerFeeRate = tmp
			}

			if respStr.Data.TakerFeeRate != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.TakerFeeRate, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("taker fee rate parse failed")},
					}
				}
				order.TakerFeeRate = tmp
			}

			if respStr.Data.AvgPrice != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.AvgPrice, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("average price parse failed")},
					}
				}
				order.AveragePrice = tmp
			}

			if respStr.Data.AssetFee != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.AssetFee, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("transaction fee parse failed")},
					}
				}
				order.TransactionFee = tmp
			}

			if respStr.Data.FeeDiscount != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.FeeDiscount, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("fee discount parse failed")},
					}
				}
				order.FeeDiscount = tmp
			}
			if respStr.Data.AssetFee != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.AssetFee, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("asset fee parse failed")},
					}
				}
				order.AssetFee = tmp
			}

			if respStr.Data.MoneyFee != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.MoneyFee, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("money fee parse failed")},
					}
				}
				order.MoneyFee = tmp
			}

			if respStr.Data.StockFee != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.StockFee, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("stock fee parse failed")},
					}
				}
				order.StockFee = tmp
			}
			return &brokerages.OrderResponse{Order: order}
		} else {
			return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
				Error: errors.New("get wallet list error")},
			}
		}
	} else {
		return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{Error: errors.New(resp.Status)}}
	}
}

func (config Config) OrderStatus(params brokerages.OrderStatusParams) *brokerages.OrderResponse {
	req := networkManager.Request{
		Method:   networkManager.GET,
		Endpoint: "https://api.coinex.com/v1/order/status",
		Headers:  map[string][]string{"authorization": {config.AccessId}},
		Params: map[string]interface{}{
			"access_id": config.AccessId,
			"id":        params.ServerOrderId,
			"market":    params.Market.Name,
			"tonce":     time.Now().UnixNano() / 1000,
			"client_id": strings.ReplaceAll(params.ClientUUID.String(), "-", ""),
		},
	}

	resp, err := req.Execute()
	if err != nil {
		return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{Error: err}}
	}
	if resp.Code == 200 {
		respStr := struct {
			Code int `json:"code"`
			Data struct {
				Amount       string `json:"amount"`
				AssetFee     string `json:"asset_fee"`
				AvgPrice     string `json:"avg_price"`
				CreateTime   int64  `json:"create_time"`
				DealAmount   string `json:"deal_amount"`
				DealFee      string `json:"deal_fee"`
				DealMoney    string `json:"deal_money"`
				FeeAsset     string `json:"fee_asset"`
				FeeDiscount  string `json:"fee_discount"`
				FinishedTime int64  `json:"finished_time"`
				Id           int64  `json:"id"`
				Left         string `json:"left"`
				MakerFeeRate string `json:"maker_fee_rate"`
				MoneyFee     string `json:"money_fee"`
				Market       string `json:"market"`
				OrderType    string `json:"order_type"`
				Price        string `json:"price"`
				Status       string `json:"status"`
				StockFee     string `json:"stock_fee"`
				TakerFeeRate string `json:"taker_fee_rate"`
				Type         string `json:"type"`
				ClientId     string `json:"client_id"`
			} `json:"data"`
			Message string `json:"message"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{Error: err}}
		}
		if respStr.Code == ResponseSuccess {
			order := models.Order{
				ClientUUID:    respStr.Data.ClientId,
				ServerOrderId: respStr.Data.Id,
				CreatedAt:     time.Unix(respStr.Data.CreateTime, 0),
				FinishedAt:    time.Unix(respStr.Data.FinishedTime, 0),
				Status:        models.OrderStatus(respStr.Data.Status),
				Market:        params.Market,
				SellOrBuy:     models.OrderType(respStr.Data.Type),
				OrderKind:     models.OrderKind(respStr.Data.OrderType),
				FeeAsset:      models.Asset(respStr.Data.FeeAsset),
			}
			if respStr.Data.Amount != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.Amount, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("order amount parse failed")},
					}
				}
				order.Amount = tmp
			}

			if respStr.Data.DealAmount != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.DealAmount, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("executed amount parse failed")},
					}
				}
				order.ExecutedAmount = tmp
			}

			if respStr.Data.Left != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.Left, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("unexecuted amount parse failed")},
					}
				}
				order.UnExecutedAmount = tmp
			}

			if respStr.Data.DealMoney != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.DealMoney, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("executed price parse failed")},
					}
				}
				order.ExecutedPrice = tmp
			}

			if respStr.Data.MakerFeeRate != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.MakerFeeRate, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("maker fee rate parse failed")},
					}
				}
				order.MakerFeeRate = tmp
			}

			if respStr.Data.TakerFeeRate != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.TakerFeeRate, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("taker fee rate parse failed")},
					}
				}
				order.TakerFeeRate = tmp
			}

			if respStr.Data.AvgPrice != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.AvgPrice, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("average price parse failed")},
					}
				}
				order.AveragePrice = tmp
			}

			if respStr.Data.AssetFee != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.AssetFee, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("transaction fee parse failed")},
					}
				}
				order.TransactionFee = tmp
			}

			if respStr.Data.FeeDiscount != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.FeeDiscount, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("fee discount parse failed")},
					}
				}
				order.FeeDiscount = tmp
			}
			if respStr.Data.AssetFee != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.AssetFee, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("asset fee parse failed")},
					}
				}
				order.AssetFee = tmp
			}

			if respStr.Data.MoneyFee != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.MoneyFee, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("money fee parse failed")},
					}
				}
				order.MoneyFee = tmp
			}

			if respStr.Data.StockFee != "" {
				tmp, parseErr := strconv.ParseFloat(respStr.Data.StockFee, 64)
				if parseErr != nil {
					return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
						Error: errors.New("stock fee parse failed")},
					}
				}
				order.StockFee = tmp
			}
			return &brokerages.OrderResponse{Order: order}
		} else {
			return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{
				Error: errors.New("get wallet list error")},
			}
		}
	} else {
		return &brokerages.OrderResponse{BasicResponse: brokerages.BasicResponse{Error: errors.New(resp.Status)}}
	}
}

func (config Config) OrderList(params brokerages.OrderListParams) *brokerages.OrderListResponse {
	endpoint := ""
	if params.IsExecuted {
		endpoint = "https://api.coinex.com/v1/order/finished"
	} else {
		endpoint = "https://api.coinex.com/v1/order/pending"
	}
	req := networkManager.Request{
		Method:   networkManager.GET,
		Endpoint: endpoint,
		Headers:  map[string][]string{"authorization": {config.AccessId}},
		Params: map[string]interface{}{
			"access_id":  config.AccessId,
			"market":     params.Market.Name,
			"type":       params.IsBuy,
			"page":       params.Page,
			"limit":      params.Limit,
			"tonce":      time.Now().UnixNano() / 1000,
			"account_id": 0,
			"client_id":  strings.ReplaceAll(params.ClientUUID.String(), "-", ""),
		},
	}

	resp, err := req.Execute()
	if err != nil {
		return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{Error: err}}
	}
	if resp.Code == 200 {
		respStr := struct {
			Code int `json:"code"`
			Data []struct {
				Amount       string `json:"amount"`
				AssetFee     string `json:"asset_fee"`
				AvgPrice     string `json:"avg_price"`
				CreateTime   int64  `json:"create_time"`
				DealAmount   string `json:"deal_amount"`
				DealFee      string `json:"deal_fee"`
				DealMoney    string `json:"deal_money"`
				FeeAsset     string `json:"fee_asset"`
				FeeDiscount  string `json:"fee_discount"`
				FinishedTime int64  `json:"finished_time"`
				Id           int64  `json:"id"`
				Left         string `json:"left"`
				MakerFeeRate string `json:"maker_fee_rate"`
				MoneyFee     string `json:"money_fee"`
				Market       string `json:"market"`
				OrderType    string `json:"order_type"`
				Price        string `json:"price"`
				Status       string `json:"status"`
				StockFee     string `json:"stock_fee"`
				TakerFeeRate string `json:"taker_fee_rate"`
				Type         string `json:"type"`
				ClientId     string `json:"client_id"`
			} `json:"data"`
			Message string `json:"message"`
		}{}
		if err := json.Unmarshal(resp.Body, &respStr); err != nil {
			return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{Error: err}}
		}
		if respStr.Code == ResponseSuccess {
			var orders []models.Order
			for _, data := range respStr.Data {
				order := models.Order{
					ClientUUID:    data.ClientId,
					ServerOrderId: data.Id,
					CreatedAt:     time.Unix(data.CreateTime, 0),
					FinishedAt:    time.Unix(data.FinishedTime, 0),
					Status:        models.OrderStatus(data.Status),
					Market:        params.Market,
					SellOrBuy:     models.OrderType(data.Type),
					OrderKind:     models.OrderKind(data.OrderType),
					FeeAsset:      models.Asset(data.FeeAsset),
				}
				if data.Amount != "" {
					tmp, parseErr := strconv.ParseFloat(data.Amount, 64)
					if parseErr != nil {
						return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{
							Error: errors.New("order amount parse failed")},
						}
					}
					order.Amount = tmp
				}

				if data.DealAmount != "" {
					tmp, parseErr := strconv.ParseFloat(data.DealAmount, 64)
					if parseErr != nil {
						return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{
							Error: errors.New("executed amount parse failed")},
						}
					}
					order.ExecutedAmount = tmp
				}

				if data.Left != "" {
					tmp, parseErr := strconv.ParseFloat(data.Left, 64)
					if parseErr != nil {
						return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{
							Error: errors.New("unexecuted amount parse failed")},
						}
					}
					order.UnExecutedAmount = tmp
				}

				if data.DealMoney != "" {
					tmp, parseErr := strconv.ParseFloat(data.DealMoney, 64)
					if parseErr != nil {
						return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{
							Error: errors.New("executed price parse failed")},
						}
					}
					order.ExecutedPrice = tmp
				}

				if data.MakerFeeRate != "" {
					tmp, parseErr := strconv.ParseFloat(data.MakerFeeRate, 64)
					if parseErr != nil {
						return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{
							Error: errors.New("maker fee rate parse failed")},
						}
					}
					order.MakerFeeRate = tmp
				}

				if data.TakerFeeRate != "" {
					tmp, parseErr := strconv.ParseFloat(data.TakerFeeRate, 64)
					if parseErr != nil {
						return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{
							Error: errors.New("taker fee rate parse failed")},
						}
					}
					order.TakerFeeRate = tmp
				}

				if data.AvgPrice != "" {
					tmp, parseErr := strconv.ParseFloat(data.AvgPrice, 64)
					if parseErr != nil {
						return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{
							Error: errors.New("average price parse failed")},
						}
					}
					order.AveragePrice = tmp
				}

				if data.AssetFee != "" {
					tmp, parseErr := strconv.ParseFloat(data.AssetFee, 64)
					if parseErr != nil {
						return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{
							Error: errors.New("transaction fee parse failed")},
						}
					}
					order.TransactionFee = tmp
				}

				if data.FeeDiscount != "" {
					tmp, parseErr := strconv.ParseFloat(data.FeeDiscount, 64)
					if parseErr != nil {
						return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{
							Error: errors.New("fee discount parse failed")},
						}
					}
					order.FeeDiscount = tmp
				}
				if data.AssetFee != "" {
					tmp, parseErr := strconv.ParseFloat(data.AssetFee, 64)
					if parseErr != nil {
						return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{
							Error: errors.New("asset fee parse failed")},
						}
					}
					order.AssetFee = tmp
				}

				if data.MoneyFee != "" {
					tmp, parseErr := strconv.ParseFloat(data.MoneyFee, 64)
					if parseErr != nil {
						return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{
							Error: errors.New("money fee parse failed")},
						}
					}
					order.MoneyFee = tmp
				}

				if data.StockFee != "" {
					tmp, parseErr := strconv.ParseFloat(data.StockFee, 64)
					if parseErr != nil {
						return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{
							Error: errors.New("stock fee parse failed")},
						}
					}
					order.StockFee = tmp
				}
				orders = append(orders, order)
			}
			return &brokerages.OrderListResponse{Orders: orders}
		} else {
			return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{
				Error: errors.New("get wallet list error")},
			}
		}
	} else {
		return &brokerages.OrderListResponse{BasicResponse: brokerages.BasicResponse{Error: errors.New(resp.Status)}}
	}
}

//todo: must be implement next methods
func (config Config) RecentTrades(brokerages.OrderBookParams) *brokerages.RecentTradesResponse {
	return &brokerages.RecentTradesResponse{BasicResponse: brokerages.BasicResponse{Error: ErrMustBeImplemented}}
}

func (config Config) UserInfo() *brokerages.UserInfoResponse {
	return &brokerages.UserInfoResponse{BasicResponse: brokerages.BasicResponse{Error: ErrMustBeImplemented}}
}

func (config Config) WalletInfo(brokerages.WalletInfoParams) *brokerages.WalletResponse {
	return &brokerages.WalletResponse{BasicResponse: brokerages.BasicResponse{Error: ErrMustBeImplemented}}
}

func (config Config) WalletBalance(brokerages.WalletBalanceParams) *brokerages.BalanceResponse {
	return &brokerages.BalanceResponse{BasicResponse: brokerages.BasicResponse{Error: ErrMustBeImplemented}}
}

func (config Config) TransactionList(brokerages.TransactionListParams) *brokerages.TransactionListResponse {
	return &brokerages.TransactionListResponse{BasicResponse: brokerages.BasicResponse{Error: ErrMustBeImplemented}}
}

func (config Config) UpdateOrderStatus(brokerages.UpdateOrderStatusParams) *brokerages.UpdateOrderStatusResponse {
	return &brokerages.UpdateOrderStatusResponse{BasicResponse: brokerages.BasicResponse{Error: ErrMustBeImplemented}}
}
