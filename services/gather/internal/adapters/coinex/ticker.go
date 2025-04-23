package coinex

import (
	"encoding/json"
	"fmt"
	"github.com/h-varmazyar/Gate/services/gather/internal/workers/candleTicker"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"strconv"
)

type Ticker struct {
	Last   string `json:"last"`
	Market string `json:"market"`
}

func (c Coinex) MarketsTicker(_ context.Context) ([]candleTicker.Ticker, error) {
	url := fmt.Sprintf("%v/spot/ticker", c.cfg.BaseURL)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("new ticker")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid HTTP response code: %d", resp.StatusCode)
	}

	tickers := baseResponse[[]Ticker]{}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &tickers)
	if err != nil {
		return nil, err
	}

	if tickers.Code == 0 {
		finalTickers := make([]candleTicker.Ticker, len(tickers.Data))
		for _, t := range tickers.Data {
			price, _ := strconv.ParseFloat(t.Last, 64)
			finalTickers = append(finalTickers, candleTicker.Ticker{
				MarketName: t.Market,
				LastPrice:  price,
			})
		}
		fmt.Println("ticker response:", len(finalTickers))
		return finalTickers, nil
	}
	return nil, fmt.Errorf("invalid response code: %d", tickers.Code)
}

//func (c Coinex) MarketsTicker(_ context.Context) ([]candleTicker.Ticker, error) {
//	url := fmt.Sprintf("%v/spot/ticker", c.cfg.BaseURL)
//	req := fasthttp.AcquireRequest()
//	req.SetRequestURI(url)
//	req.Header.SetMethod(fasthttp.MethodGet)
//	req.Header.SetContentType("application/json")
//	req.Header.Set("Connection", "Close")
//	//req.Header.Set("Accept-Encoding", "gzip, deflate, br")
//	//req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")
//	//req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
//
//	resp := fasthttp.AcquireResponse()
//	err := c.client.Do(req, resp)
//	defer fasthttp.ReleaseRequest(req)
//	defer fasthttp.ReleaseResponse(resp)
//	if err != nil {
//		return nil, err
//	}
//	statusCode := resp.StatusCode()
//	respBody := resp.Body()
//
//	if statusCode != http.StatusOK {
//		return nil, fmt.Errorf("invalid HTTP response code: %d", statusCode)
//	}
//
//	tickers := baseResponse[[]Ticker]{}
//
//	err = json.Unmarshal(respBody, &tickers)
//	if err != nil {
//		fmt.Println("unmarshal failed in ticker")
//		return nil, err
//	}
//
//	if tickers.Code == 0 {
//		finalTickers := make([]candleTicker.Ticker, len(tickers.Data))
//		for _, t := range tickers.Data {
//			price, _ := strconv.ParseFloat(t.Last, 64)
//			finalTickers = append(finalTickers, candleTicker.Ticker{
//				MarketName: t.Market,
//				LastPrice:  price,
//			})
//		}
//		return finalTickers, nil
//	}
//	return nil, fmt.Errorf("invalid response code: %d", tickers.Code)
//}
