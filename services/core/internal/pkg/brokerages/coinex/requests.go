package coinex

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
)

type Requests struct {
	Auth    *api.Auth
	configs *Configs
}

func NewRequest(configs *Configs, auth *api.Auth) brokerages.Requests {
	return &Requests{
		configs: configs,
		Auth:    auth,
	}
}

var (
	ErrNilResolution  = errors.New("resolution must be declared")
	ErrNilMarket      = errors.New("market must be declared")
	ErrWrongStartTime = errors.New("start time must be grater than 0")
	ErrWrongEndTime   = errors.New("end time must be grater than 0")
)

func (r *Requests) AsyncOHLC(_ context.Context, inputs *brokerages.OHLCParams) (*networkAPI.Request, error) {
	if inputs.Resolution == nil {
		return nil, ErrNilResolution
	}
	if inputs.Market == nil {
		return nil, ErrNilMarket
	}
	if inputs.From.IsZero() {
		return nil, ErrWrongStartTime
	}
	if inputs.To.IsZero() {
		return nil, ErrWrongEndTime
	}

	request := new(networkAPI.Request)
	request.Method = networkAPI.Request_GET
	request.CallbackQueue = r.configs.CoinexCallbackQueue
	resolutionSeconds := inputs.Resolution.Duration
	count := int64(inputs.To.Sub(inputs.From)) / resolutionSeconds
	if count < 0 {
		count *= -1
	}
	if (int64(inputs.To.Sub(inputs.From)) % resolutionSeconds) > 0 {
		count++
	}
	if int64(count) >= 1000 {
		request.Params = []*networkAPI.KV{
			networkAPI.NewKV("market", inputs.Market.Name),
			networkAPI.NewKV("interval", inputs.Resolution.Value),
			networkAPI.NewKV("start_time", fmt.Sprintf("%v", inputs.From.Unix())),
			networkAPI.NewKV("end_time", fmt.Sprintf("%v", inputs.To.Unix()))}
		request.Endpoint = "https://www.coinex.com/res/market/kline"
	} else {
		request.Params = []*networkAPI.KV{
			networkAPI.NewKV("markets", inputs.Market.Name),
			networkAPI.NewKV("type", inputs.Resolution.Label),
			networkAPI.NewKV("limit", fmt.Sprintf("%v", int64(count))),
		}
		request.Endpoint = "https://api.coinex.com/v1/market/kline"
	}

	metadataBytes, _ := json.Marshal(&brokerages.Metadata{
		Method:       MethodOHLC,
		Platform:     api.Platform_Coinex,
		MarketID:     inputs.Market.ID,
		ResolutionID: inputs.Resolution.ID,
	})
	request.Metadata = string(metadataBytes)
	return request, nil
}

func (r *Requests) AllMarketStatistics(_ context.Context, _ *brokerages.AllMarketStatisticsParams) (*networkAPI.Request, error) {
	request := new(networkAPI.Request)
	request.Method = networkAPI.Request_GET
	request.Endpoint = "https://api.coinex.com/v1/market/ticker/all"

	return request, nil
}

func (r *Requests) GetMarketInfo(_ context.Context, inputs *brokerages.MarketInfoParams) (*networkAPI.Request, error) {
	request := new(networkAPI.Request)
	request.Method = networkAPI.Request_GET
	request.Endpoint = fmt.Sprintf("https://www.coinex.com/res/vote2/project/%v", inputs.Market.Name)

	return request, nil
}
