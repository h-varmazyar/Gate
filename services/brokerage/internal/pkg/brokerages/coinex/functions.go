package coinex

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/errors"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/brokerages"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/repository"
	networkAPI "github.com/mrNobody95/Gate/services/network/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 19.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type Service struct {
	Auth *api.Auth
}

func (service *Service) WalletList(ctx context.Context, runner brokerages.Handler) ([]*repository.Wallet, error) {
	request := new(networkAPI.Request)
	request.Type = networkAPI.Type_GET
	request.Endpoint = "https://api.coinex.com/v1/balance/info"
	request.Params = []*networkAPI.KV{
		{Key: "access_id", Value: service.Auth.AccessID},
		{Key: "tonce", Value: fmt.Sprintf("%d", time.Now().UnixNano()/1e6)},
	}
	request.Headers = []*networkAPI.KV{
		{Key: "authorization", Value: service.generateAuthorization(request.Params)},
		{Key: "tonce", Value: fmt.Sprintf("%d", time.Now().UnixNano()/1e6)},
	}
	_, err := runner(ctx, request)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (service *Service) OHLC(ctx context.Context, inputs brokerages.OHLCParams, runner brokerages.Handler) ([]*api.Candle, error) {
	request := new(networkAPI.Request)
	request.Type = networkAPI.Type_GET
	count := (inputs.To.Sub(inputs.From)) / inputs.Resolution.Duration
	if int64((inputs.To.Sub(inputs.From))%inputs.Resolution.Duration) > 0 {
		count++
	}
	if int64(count) >= 1000 {
		request.Params = []*networkAPI.KV{
			{Key: "market", Value: inputs.Market.Name},
			{Key: "interval", Value: inputs.Resolution.Value},
			{Key: "start_time", Value: fmt.Sprintf("%v", inputs.From.Unix())},
			{Key: "end_time", Value: fmt.Sprintf("%v", inputs.To.Unix())}}
		request.Endpoint = "https://www.coinex.com/res/market/kline"
	} else {
		request.Params = []*networkAPI.KV{
			{Key: "market", Value: inputs.Market.Name},
			{Key: "type", Value: inputs.Resolution.Label},
			{Key: "limit", Value: fmt.Sprintf("%v", int64(count))},
		}
		request.Endpoint = "https://api.coinex.com/v1/market/kline"
	}
	resp, err := runner(ctx, request)
	if err != nil {
		return nil, err
	}
	if resp.Code != http.StatusOK {
		return nil, errors.NewWithSlug(ctx, codes.NotFound, resp.Response)
	}
	data := make([][]interface{}, 0)
	if err := parseResponse(resp.Response, &data); err != nil {
		return nil, err
	}
	candles := make([]*api.Candle, 0)
	for _, item := range data {
		c := new(api.Candle)
		var err error
		c.Time = int64(item[0].(float64))
		c.Open, err = strconv.ParseFloat(item[1].(string), 64)
		if err != nil {
			continue
		}
		c.Close, err = strconv.ParseFloat(item[2].(string), 64)
		if err != nil {
			continue
		}
		c.High, err = strconv.ParseFloat(item[3].(string), 64)
		if err != nil {
			continue
		}
		c.Low, err = strconv.ParseFloat(item[4].(string), 64)
		if err != nil {
			continue
		}
		c.Volume, err = strconv.ParseFloat(item[5].(string), 64)
		if err != nil {
			continue
		}
		c.Amount, err = strconv.ParseFloat(item[6].(string), 64)
		if err != nil {
			continue
		}
		candles = append(candles, c)
	}
	return candles, nil
}

func (service *Service) UpdateMarket(ctx context.Context, runner brokerages.Handler) ([]*repository.Market, error) {
	request := new(networkAPI.Request)
	request.Type = networkAPI.Type_GET
	request.Endpoint = "https://api.coinex.com/v1/market/info"

	resp, err := runner(ctx, request)
	if err != nil {
		return nil, err
	}
	if resp.Code != http.StatusOK {
		return nil, errors.NewWithSlug(ctx, codes.NotFound, resp.Response)
	}
	//data := make(map[string]struct {
	//	Name           string `json:"name"`
	//	MinAmount      string `json:"min_amount"`
	//	MakerFeeRate   string `json:"maker_fee_rate"`
	//	TakerFeeRate   string `json:"taker_fee_rate"`
	//	PricingName    string `json:"pricing_name"`
	//	PricingDecimal int    `json:"pricing_decimal"`
	//	TradingName    string `json:"trading_name"`
	//	TradingDecimal int    `json:"trading_decimal"`
	//})
	tmp := new(responseModel)
	if err := json.Unmarshal([]byte(resp.Response), tmp); err != nil {
		return nil, err
	}
	if tmp.Code != 0 {
		log.WithError(err)
		return nil, errors.New(ctx, codes.Canceled)
	}
	data := tmp.Data.(map[string]interface{})
	markets := make([]*repository.Market, 0)
	for _, value := range data {
		item := value.(map[string]interface{})
		m := new(repository.Market)
		m.ID = uuid.New()
		m.BrokerageName = brokerageApi.Names_Coinex.String()
		m.PricingDecimal = int(item["pricing_decimal"].(float64))
		m.TradingDecimal = int(item["trading_decimal"].(float64))
		num, err := strconv.ParseFloat(item["taker_fee_rate"].(string), 64)
		if err != nil {
			log.WithError(err).WithField("taker_fee_rate", item["taker_fee_rate"]).Error("failed to add market")
			continue
		}
		m.TakerFeeRate = num
		num, err = strconv.ParseFloat(item["maker_fee_rate"].(string), 64)
		if err != nil {
			log.WithError(err).WithField("maker_fee_rate", item["maker_fee_rate"]).Error("failed to add market")
			continue
		}
		m.MakerFeeRate = num
		num, err = strconv.ParseFloat(item["min_amount"].(string), 64)
		if err != nil {
			log.WithError(err).WithField("min_amount", item["min_amount"]).Error("failed to add market")
			continue
		}
		m.MinAmount = num
		m.SourceName = item["trading_name"].(string)
		m.DestinationName = item["pricing_name"].(string)
		m.StartTime = time.Unix(1641025800, 0)
		m.IsAMM = false
		m.Name = item["name"].(string)
		m.Status = api.Status_Enable
		markets = append(markets, m)
	}
	return markets, nil
}

func (service *Service) generateAuthorization(params []*networkAPI.KV) string {
	urlParameters := url.Values{}
	for _, param := range params {
		urlParameters.Add(param.Key, param.Value)
	}
	queryParamsString := urlParameters.Encode()
	toEncodeParamsString := queryParamsString + "&secrect=" + service.Auth.SecretKey
	w := md5.New()
	io.WriteString(w, toEncodeParamsString)
	md5Str := fmt.Sprintf("%x", w.Sum(nil))
	md5Str = strings.ToUpper(md5Str)
	return md5Str
}
