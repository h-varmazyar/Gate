package coinex

import (
	"context"
	"encoding/json"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/errors"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"net/http"
	"strconv"
)

type Response struct {
}

func NewResponse() (*Response, error) {
	r := &Response{}
	return r, nil
}

func (r *Response) AsyncOHLC(_ context.Context, responses *networkAPI.AsyncResponses) *coreApi.OHLCResponse {
	message := &coreApi.OHLCResponse{
		Items:    make([]*coreApi.OHLCResponseItem, 0),
		Platform: api.Platform_Coinex,
	}

	for _, response := range responses.Responses {
		metadata := new(brokerages.Metadata)
		item := new(coreApi.OHLCResponseItem)
		if err := json.Unmarshal([]byte(response.Metadata), metadata); err != nil {
			log.WithError(err).Error("failed to unmarshal coinex callback delivery")
			item.Error = err.Error()
			message.Items = append(message.Items, item)
			continue
		}
		if metadata.Platform != api.Platform_Coinex {
			item.Error = "invalid platform"
			message.Items = append(message.Items, item)
			continue
		}
		if response.Code != http.StatusOK {
			log.Errorf("ohlc request failed with code: %v - %v", response.Code, response.Body)
			item.Error = response.Body
			message.Items = append(message.Items, item)
			continue
		}
		data := make([][]interface{}, 0)
		if err := parseResponse(response.Body, &data); err != nil {
			log.WithError(err).Errorf("ohlc request parse failed: %v", response.Body)
			item.Error = err.Error()
			message.Items = append(message.Items, item)
			continue
		}
		item.Candles = make([]*coreApi.OHLCResponseItem_Candle, 0)
		for _, candle := range data {
			c := new(coreApi.OHLCResponseItem_Candle)
			c.Time = int64(candle[0].(float64))
			c.Open, _ = strconv.ParseFloat(candle[1].(string), 64)
			c.Close, _ = strconv.ParseFloat(candle[2].(string), 64)
			c.High, _ = strconv.ParseFloat(candle[3].(string), 64)
			c.Low, _ = strconv.ParseFloat(candle[4].(string), 64)
			c.Volume, _ = strconv.ParseFloat(candle[5].(string), 64)
			c.Amount, _ = strconv.ParseFloat(candle[6].(string), 64)
			item.Candles = append(item.Candles, c)
		}
		item.ResolutionID = metadata.ResolutionID
		item.MarketID = metadata.MarketID

		message.Items = append(message.Items, item)
	}
	return message
}

func (r *Response) AllMarkerStatistics(ctx context.Context, response *networkAPI.Response) (*coreApi.AllMarketStatisticsResp, error) {
	if response.Code != http.StatusOK {
		return nil, errors.New(ctx, codes.Canceled).AddDetails(response.Body)
	}
	data := struct {
		Date   float64 `json:"date"`
		Ticker map[string]struct {
			Buy        string `json:"buy"`
			BuyAmount  string `json:"buy_amount"`
			Open       string `json:"open"`
			High       string `json:"high"`
			Low        string `json:"low"`
			Last       string `json:"last"`
			Sell       string `json:"sell"`
			SellAmount string `json:"sell_amount"`
			Volume     string `json:"vol"`
		} `json:"ticker"`
	}{}
	if err := parseResponse(response.Body, &data); err != nil {
		return nil, err
	}
	resp := new(coreApi.AllMarketStatisticsResp)
	resp.Platform = api.Platform_Coinex
	resp.Date = int64(data.Date / 1000)
	resp.AllStatistics = make(map[string]*coreApi.MarketStatistics)

	var err error
	for key, value := range data.Ticker {
		marketStatistics := new(coreApi.MarketStatistics)
		marketStatistics.Date = int64(data.Date / 1000)
		if marketStatistics.Volume, err = strconv.ParseFloat(value.Volume, 64); err != nil {
			log.WithError(err).Error("failed to parse volume")
			return nil, err
		}
		if marketStatistics.Close, err = strconv.ParseFloat(value.Last, 64); err != nil {
			log.WithError(err).Error("failed to parse close")
			return nil, err
		}
		if marketStatistics.Open, err = strconv.ParseFloat(value.Open, 64); err != nil {
			log.WithError(err).Error("failed to parse open")
			return nil, err
		}
		if marketStatistics.High, err = strconv.ParseFloat(value.High, 64); err != nil {
			log.WithError(err).Error("failed to parse high")
			return nil, err
		}
		if marketStatistics.Low, err = strconv.ParseFloat(value.Low, 64); err != nil {
			log.WithError(err).Error("failed to parse low")
			return nil, err
		}
		resp.AllStatistics[key] = marketStatistics
	}
	return resp, nil
}

func (r *Response) GetMarketInfo(ctx context.Context, response *networkAPI.Response) (*coreApi.MarketInfo, error) {
	if response.Code != http.StatusOK {
		return nil, errors.New(ctx, codes.Canceled).AddDetails(response.Body)
	}
	data := struct {
		ShortName    string `json:"shortname"`
		FullName     string `json:"full_name"`
		IssueTime    int64  `json:"issue_time"`
		Logo         string `json:"logo"`
		WebsiteURL   string `json:"website_url"`
		Introduction string `json:"introduction"`
		Status       string `json:"status"`
	}{}

	if err := parseResponse(response.Body, &data); err != nil {
		return nil, err
	}
	marketInfo := &coreApi.MarketInfo{
		IssueDate:    data.IssueTime,
		ShortName:    data.ShortName,
		FullName:     data.FullName,
		Logo:         data.Logo,
		WebsiteURL:   data.WebsiteURL,
		Introduction: data.Introduction,
		Status:       data.Status,
	}
	return marketInfo, nil
}

func (r *Response) WalletsBalance(ctx context.Context, response *networkAPI.Response) (*chipmunkApi.Wallets, error) {
	if response.Code != http.StatusOK {
		return nil, errors.NewWithSlug(ctx, codes.Unknown, response.Body)
	}
	data := make(map[string]map[string]interface{})
	if err := parseResponse(response.Body, &data); err != nil {
		return nil, err
	}
	wallets := new(chipmunkApi.Wallets)
	wallets.Elements = make([]*chipmunkApi.Wallet, 0)
	var err error
	for key, value := range data {
		w := new(chipmunkApi.Wallet)
		w.AssetName = key
		w.ActiveBalance, err = strconv.ParseFloat(value["available"].(string), 64)
		if err != nil {
			continue
		}
		w.BlockedBalance, err = strconv.ParseFloat(value["frozen"].(string), 64)
		if err != nil {
			continue
		}
		w.TotalBalance = w.ActiveBalance + w.BlockedBalance
		wallets.Elements = append(wallets.Elements, w)
	}
	return wallets, nil
}
