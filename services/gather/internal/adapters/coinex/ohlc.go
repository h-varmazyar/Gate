package coinex

import (
	"encoding/json"
	"fmt"
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Candle struct {
	Open     float64 `json:"1"`
	High     float64 `json:"3"`
	Low      float64 `json:"4"`
	Close    float64 `json:"2"`
	Volume   float64 `json:"5"`
	Value    float64 `json:"6"`
	CreateAt int64   `json:"0"`
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
	fmt.Println("total", candleCount)
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

func (c Coinex) newOHLC(_ context.Context, market string, resolutionLabel string, limit int) ([]Candle, error) {
	url := fmt.Sprintf("%v/spot/kline?market=%v&period=%v&limit=%v", c.cfg.APIBaseURL, market, resolutionLabel, limit)
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
	fmt.Println("getting ", url, " -> ", len(respEntity.Data))

	return respEntity.Data, nil
}

func (c Coinex) reverseOHLC(_ context.Context, market string, resolution models.Resolution) ([]Candle, error) {
	candles := make([]Candle, 0)
	to := time.Now()
	from := to.Add(resolution.Duration * -1000)
	for {
		url := fmt.Sprintf("https://www.coinex.com/res/market/kline?market=%v&interval=%v&start_time=%v&end_time=%v",
			market, resolution.Value, from.Unix(), to.Unix())
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			continue
		}

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
		to = from.Add(resolution.Duration * -1000)
		if len(respEntity.Data) == 0 {
			break
		}
		time.Sleep(time.Second / 3)
		candles = append(candles, respEntity.Data...)
	}

	return candles, nil
}

func (c Coinex) historicalOHLC(_ context.Context, market string, resolution models.Resolution, from time.Time, reqCount int) ([]Candle, error) {
	candles := make([]Candle, 0)
	to := from.Add(resolution.Duration * 1000)
	fmt.Println("need istorical from - to - count:", from, to, reqCount)
	for i := 0; i < reqCount; {
		url := fmt.Sprintf("https://www.coinex.com/res/market/kline?market=%v&interval=%v&start_time=%v&end_time=%v",
			market, resolution.Value, from.Unix(), to.Unix())
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			fmt.Println("failed to create:", err)
			continue
		}

		resp, err := c.client.Do(req)
		if err != nil {
			fmt.Println("failed to go:", err)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			fmt.Println("invalid status code:", resp.StatusCode)
			continue
		}
		fmt.Println("done")

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("failed to read body:", err)
			resp.Body.Close()
			continue
		}
		respEntity := baseResponse[[][]any]{}
		err = json.Unmarshal(body, &respEntity)
		if err != nil {
			fmt.Println("failed to unmarshal:", err)
			resp.Body.Close()
			continue
		}
		resp.Body.Close()

		from = to
		to = from.Add(resolution.Duration * 1000)
		i++
		time.Sleep(time.Second / 3)
		for _, d := range respEntity.Data {
			candle := Candle{}
			candle.Open, _ = strconv.ParseFloat(d[1].(string), 64)
			candle.High, _ = strconv.ParseFloat(d[3].(string), 64)
			candle.Low, _ = strconv.ParseFloat(d[4].(string), 64)
			candle.Close, _ = strconv.ParseFloat(d[2].(string), 64)
			candle.Volume, _ = strconv.ParseFloat(d[5].(string), 64)
			candle.Value, _ = strconv.ParseFloat(d[6].(string), 64)
			candle.CreateAt = int64(d[0].(float64))
			candles = append(candles, candle)
		}
		//candles = append(candles, respEntity.Data...)
		fmt.Println("getting ", url, " -> ", len(respEntity.Data))
	}

	return candles, nil
}

//func (c Coinex) historicalOHLC(_ context.Context, market string, resolution models.Resolution, from time.Time, reqCount int) ([]Candle, error) {
//	candles := make([]Candle, 0)
//	to := from.Add(resolution.Duration * 1000)
//	for i := 0; i < reqCount; {
//		req := fasthttp.AcquireRequest()
//		url := fmt.Sprintf("https://www.coinex.com/res/market/kline?market=%v&interval=%v&start_time=%v&end_time=%v", market, resolution.Value, from.Unix(), to.Unix())
//		req.SetRequestURI(url)
//		req.Header.SetMethod(fasthttp.MethodGet)
//		resp := fasthttp.AcquireResponse()
//		err := c.client.Do(req, resp)
//		fasthttp.ReleaseRequest(req)
//		if err != nil {
//			continue
//		}
//		fasthttp.ReleaseResponse(resp)
//
//		statusCode := resp.StatusCode()
//		respBody := resp.Body()
//
//		if statusCode != fasthttp.StatusOK {
//			continue
//		}
//
//		respEntity := baseResponse[[]Candle]{}
//		err = json.Unmarshal(respBody, &respEntity)
//		if err != nil {
//			continue
//		}
//
//		from = to
//		to = from.Add(resolution.Duration * 1000)
//		i++
//		time.Sleep(time.Second / 3)
//		candles = append(candles, respEntity.Data...)
//	}
//
//	return candles, nil
//}
//
//func (c Coinex) newOHLC(_ context.Context, market string, resolutionLabel string, limit int) ([]Candle, error) {
//	req := fasthttp.AcquireRequest()
//	req.SetRequestURI(fmt.Sprintf("%v/spot/kline?market=%v&period=%v&limit=%v", c.cfg.BaseURL, market, resolutionLabel, limit))
//	req.Header.SetMethod(fasthttp.MethodGet)
//	resp := fasthttp.AcquireResponse()
//	err := c.client.Do(req, resp)
//	fasthttp.ReleaseRequest(req)
//	if err != nil {
//		return nil, err
//	}
//	fasthttp.ReleaseResponse(resp)
//
//	statusCode := resp.StatusCode()
//	respBody := resp.Body()
//
//	if statusCode != http.StatusOK {
//		return nil, fmt.Errorf("invalid HTTP response code: %d", statusCode)
//	}
//
//	respEntity := baseResponse[[]Candle]{}
//	err = json.Unmarshal(respBody, &respEntity)
//	if err != nil {
//		return nil, err
//	}
//
//	return respEntity.Data, nil
//}
