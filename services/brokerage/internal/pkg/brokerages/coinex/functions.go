package coinex

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/errors"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/brokerages"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/repository"
	networkAPI "github.com/mrNobody95/Gate/services/network/api"
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

func (c *Service) WalletList(ctx context.Context, runner brokerages.Handler) ([]*repository.Wallet, error) {
	request := new(networkAPI.Request)
	request.Type = networkAPI.Type_GET
	request.Endpoint = "https://api.coinex.com/v1/balance/info"
	request.Params = []*networkAPI.KV{
		{Key: "access_id", Value: c.Auth.AccessID},
		{Key: "tonce", Value: fmt.Sprintf("%d", time.Now().UnixNano()/1e6)},
	}
	request.Headers = []*networkAPI.KV{
		{Key: "authorization", Value: c.generateAuthorization(request.Params)},
		{Key: "tonce", Value: fmt.Sprintf("%d", time.Now().UnixNano()/1e6)},
	}
	resp, err := runner(ctx, request)
	if err != nil {
		return nil, err
	}
	fmt.Println("resp:", resp)
	return nil, nil
}

func (c *Service) OHLC(ctx context.Context, inputs brokerages.OHLCParams, runner brokerages.Handler) ([]*brokerageApi.Candle, error) {
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
		fmt.Println("after runner")
		return nil, err
	}
	if resp.Code != http.StatusOK {
		return nil, errors.NewWithSlug(ctx, codes.NotFound, resp.Response)
	}
	data := make([][]interface{}, 0)
	if err := parseResponse(resp.Response, &data); err != nil {
		return nil, err
	}
	candles := make([]*brokerageApi.Candle, 0)
	for _, item := range data {
		c := new(brokerageApi.Candle)
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

func (c *Service) generateAuthorization(params []*networkAPI.KV) string {
	urlParameters := url.Values{}
	for _, param := range params {
		urlParameters.Add(param.Key, param.Value)
	}
	queryParamsString := urlParameters.Encode()
	toEncodeParamsString := queryParamsString + "&secrect=" + c.Auth.SecretKey
	w := md5.New()
	io.WriteString(w, toEncodeParamsString)
	md5Str := fmt.Sprintf("%x", w.Sum(nil))
	md5Str = strings.ToUpper(md5Str)
	return md5Str
}
