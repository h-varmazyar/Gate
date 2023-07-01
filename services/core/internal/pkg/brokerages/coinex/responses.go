package coinex

import (
	"context"
	"github.com/golang/protobuf/proto"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
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
	configs   *Configs
	ohlcQueue *amqpext.Queue
}

func NewResponse(configs *Configs, isAsync bool) (*Response, error) {
	r := &Response{
		configs: configs,
	}
	if isAsync {
		ohlcQueue, err := amqpext.Client.QueueDeclare(configs.ChipmunkOHLCQueue)
		if err != nil {
			return nil, err
		}
		r.ohlcQueue = ohlcQueue
	}
	return r, nil
}

func (r *Response) AsyncOHLC(_ context.Context, response *networkAPI.Response, metadata *brokerages.Metadata) {
	if response.Code != http.StatusOK {
		log.Errorf("ohlc request failed with code: %v - %v", response.Code, response.Body)
		return
	}
	data := make([][]interface{}, 0)
	if err := parseResponse(response.Body, &data); err != nil {
		log.WithError(err).Errorf("ohlc request parse failed: %v", response.Body)
		return
	}

	candles := make([]*chipmunkApi.Candle, 0)
	for _, item := range data {
		c := new(chipmunkApi.Candle)
		c.Time = int64(item[0].(float64))
		c.Open, _ = strconv.ParseFloat(item[1].(string), 64)
		c.Close, _ = strconv.ParseFloat(item[2].(string), 64)
		c.High, _ = strconv.ParseFloat(item[3].(string), 64)
		c.Low, _ = strconv.ParseFloat(item[4].(string), 64)
		c.Volume, _ = strconv.ParseFloat(item[5].(string), 64)
		c.Amount, _ = strconv.ParseFloat(item[6].(string), 64)
		c.ResolutionID = metadata.ResolutionID
		c.MarketID = metadata.MarketID
		candles = append(candles, c)
	}
	message := &chipmunkApi.Candles{
		Elements: candles,
		Count:    int64(len(candles)),
	}
	if message.Count > 0 {
		if bytes, err := proto.Marshal(message); err != nil {
			log.WithError(err).Errorf("faled to marshal coinex async ohls message")
			return
		} else {
			if publishErr := r.ohlcQueue.Publish(bytes, grpcext.ProtobufContentType); publishErr != nil {
				log.WithError(publishErr).Errorf("faled to publish coinex async ohlc")
			}
		}
	}
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
