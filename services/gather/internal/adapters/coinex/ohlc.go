package coinex

import (
	"encoding/json"
	"fmt"
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Candle struct {
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	Volume   float64 `json:"volume"`
	Value    float64 `json:"value"`
	CreateAt int64   `json:"create_at"`
}

func (c Coinex) OHLC(ctx context.Context, market models.Market, resolution models.Resolution, from, to time.Time) ([]models.Candle, error) {
	candleCount := int(to.Sub(from) / resolution.Duration)

	reqCount := candleCount / 1000
	if candleCount%1000 != 0 {
		reqCount++
	}
	var (
		candles []Candle
		err     error
	)
	if candleCount <= 1000 {
		if candleCount == 0 {
			candleCount++
		}
		candles, err = c.newOHLC(ctx, market.Name, resolution.Label, candleCount)
	} else {
		candles, err = c.historicalOHLC(ctx, market.Name, resolution, from, reqCount)
	}
	if err != nil {
		return nil, err
	}
	resp := make([]models.Candle, len(candles))
	for i, candle := range candles {
		resp[i] = models.Candle{
			Open:         candle.Open,
			High:         candle.High,
			Low:          candle.Low,
			Close:        candle.Close,
			Volume:       candle.Volume,
			Amount:       candle.Value,
			MarketID:     market.ID,
			ResolutionID: resolution.ID,
			Time:         time.Unix(candle.CreateAt, 0),
		}
	}

	return resp, nil
}

func (c Coinex) historicalOHLC(_ context.Context, market string, resolution models.Resolution, from time.Time, reqCount int) ([]Candle, error) {
	candles := make([]Candle, 0)
	to := from.Add(resolution.Duration * 1000)
	for i := 0; i < reqCount; {
		//req := fasthttp.AcquireRequest()
		//req.SetRequestURI(fmt.Sprintf("https://www.coinex.com/res/market/kline?market=%v&interval=%v&start_time=%v&end_time=%v",
		//	market, resolution.Value, from.Unix(), to.Unix()))
		//req.Header.SetMethod(fasthttp.MethodGet)
		//resp := fasthttp.AcquireResponse()
		//err := c.client.Do(req, resp)
		//fasthttp.ReleaseRequest(req)
		//if err != nil {
		//	continue
		//}
		//fasthttp.ReleaseResponse(resp)

		url := fmt.Sprintf("https://www.coinex.com/res/market/kline?market=%v&interval=%v&start_time=%v&end_time=%v",
			market, resolution.Value, from.Unix(), to.Unix())
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			continue
		}
		//
		//statusCode := resp.StatusCode()
		//respBody := resp.Body()

		resp, err := c.client.Do(req)
		if err != nil {
			continue
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			continue
		}
		respEntity := baseResponse[[]Candle]{}
		err = json.Unmarshal(body, &respEntity)
		if err != nil {
			resp.Body.Close()
			continue
		}
		resp.Body.Close()

		from = to
		to = from.Add(resolution.Duration * 1000)
		i++
		time.Sleep(time.Second / 3)
		candles = append(candles, respEntity.Data...)
	}

	return candles, nil
}

func (c Coinex) newOHLC(_ context.Context, market string, resolutionLabel string, limit int) ([]Candle, error) {
	//req := fasthttp.AcquireRequest()
	//req.SetRequestURI(fmt.Sprintf("%v/spot/kline?market=%v&period=%v&limit=%v", c.cfg.BaseURL, market, resolutionLabel, limit))
	//req.Header.SetMethod(fasthttp.MethodGet)
	//resp := fasthttp.AcquireResponse()
	//err := c.client.Do(req, resp)
	//fasthttp.ReleaseRequest(req)
	//if err != nil {
	//	return nil, err
	//}
	//fasthttp.ReleaseResponse(resp)
	//
	//statusCode := resp.StatusCode()
	//respBody := resp.Body()

	url := fmt.Sprintf("%v/spot/kline?market=%v&period=%v&limit=%v", c.cfg.BaseURL, market, resolutionLabel, limit)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid HTTP response code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respEntity := baseResponse[[]Candle]{}
	err = json.Unmarshal(body, &respEntity)
	if err != nil {
		return nil, err
	}

	return respEntity.Data, nil
}
