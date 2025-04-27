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

type newCandle struct {
	Open      string `json:"open"`
	Close     string `json:"close"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Volume    string `json:"volume"`
	Value     string `json:"value"`
	CreatedAt int64  `json:"created_at"`
}

func (c Coinex) OHLC(ctx context.Context, market models.Market, resolution models.Resolution, from, to time.Time) ([]models.Candle, error) {
	candleCount := int(to.Sub(from) / resolution.Duration)

	reqCount := candleCount / 1000
	if candleCount%1000 != 0 {
		reqCount++
	}
	var (
		candles []models.Candle
		err     error
	)
	fmt.Println("total", candleCount)
	if candleCount <= 1000 {
		if candleCount == 0 {
			candleCount++
		}
		candles, err = c.newOHLC(ctx, market, resolution, candleCount)
	} else {
		candles, err = c.historicalOHLC(ctx, market.Name, resolution, from, reqCount)
	}
	if err != nil {
		return nil, err
	}

	return candles, nil
}

func (c Coinex) newOHLC(_ context.Context, market models.Market, resolution models.Resolution, limit int) ([]models.Candle, error) {
	url := fmt.Sprintf("%v/spot/kline?market=%v&period=%v&limit=%v", c.cfg.APIBaseURL, market, resolution.Label, limit)
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

	respEntity := baseResponse[[]newCandle]{}
	err = json.Unmarshal(body, &respEntity)
	if err != nil {
		return nil, err
	}

	candles := make([]models.Candle, 0)
	for _, candle := range respEntity.Data {
		c := models.Candle{}
		c.Open, _ = strconv.ParseFloat(candle.Open, 64)
		c.High, _ = strconv.ParseFloat(candle.High, 64)
		c.Low, _ = strconv.ParseFloat(candle.Low, 64)
		c.Close, _ = strconv.ParseFloat(candle.Close, 64)
		c.Volume, _ = strconv.ParseFloat(candle.Volume, 64)
		c.Amount, _ = strconv.ParseFloat(candle.Value, 64)
		c.Time = time.Unix(candle.CreatedAt, 0)
		candles = append(candles, c)
	}

	return candles, nil
}

//func (c Coinex) reverseOHLC(_ context.Context, market string, resolution models.Resolution) ([]Candle, error) {
//	candles := make([]Candle, 0)
//	to := time.Now()
//	from := to.Add(resolution.Duration * -1000)
//	for {
//		url := fmt.Sprintf("https://www.coinex.com/res/market/kline?market=%v&interval=%v&start_time=%v&end_time=%v",
//			market, resolution.Value, from.Unix(), to.Unix())
//		req, err := http.NewRequest(http.MethodGet, url, nil)
//		if err != nil {
//			continue
//		}
//
//		resp, err := c.client.Do(req)
//		if err != nil {
//			continue
//		}
//		if resp.StatusCode != http.StatusOK {
//			resp.Body.Close()
//			continue
//		}
//
//		body, err := ioutil.ReadAll(resp.Body)
//		if err != nil {
//			resp.Body.Close()
//			continue
//		}
//		respEntity := baseResponse[[]Candle]{}
//		err = json.Unmarshal(body, &respEntity)
//		if err != nil {
//			resp.Body.Close()
//			continue
//		}
//		resp.Body.Close()
//
//		from = to
//		to = from.Add(resolution.Duration * -1000)
//		if len(respEntity.Data) == 0 {
//			break
//		}
//		time.Sleep(time.Second / 3)
//		candles = append(candles, respEntity.Data...)
//	}
//
//	return candles, nil
//}

func (c Coinex) historicalOHLC(_ context.Context, market string, resolution models.Resolution, from time.Time, reqCount int) ([]models.Candle, error) {
	to := from.Add(resolution.Duration * 1000)
	candles := make([]models.Candle, 0)
	for i := 0; i < reqCount; {
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
		respEntity := baseResponse[[][]any]{}
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
		for _, d := range respEntity.Data {
			candle := models.Candle{}
			candle.Open, _ = strconv.ParseFloat(d[1].(string), 64)
			candle.High, _ = strconv.ParseFloat(d[3].(string), 64)
			candle.Low, _ = strconv.ParseFloat(d[4].(string), 64)
			candle.Close, _ = strconv.ParseFloat(d[2].(string), 64)
			candle.Volume, _ = strconv.ParseFloat(d[5].(string), 64)
			candle.Amount, _ = strconv.ParseFloat(d[6].(string), 64)
			candle.Time = time.Unix(int64(d[0].(float64)), 0)
			candles = append(candles, candle)
		}
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
