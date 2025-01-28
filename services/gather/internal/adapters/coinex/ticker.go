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
	//req := fasthttp.AcquireRequest()
	//req.SetRequestURI(fmt.Sprintf("%v/spot/ticker", c.cfg.BaseURL))
	//req.Header.SetMethod(fasthttp.MethodGet)
	//req.Header.SetContentType("application/json")
	//resp := fasthttp.AcquireResponse()

	url := fmt.Sprintf("%v/spot/ticker", c.cfg.BaseURL)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	//req.Header.Set("content-type", "application/json")
	//req.Header.Set("Connection", "Keep-Alive")
	//req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	//req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")
	//req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")

	resp, err := c.client.Do(req)
	//fasthttp.ReleaseRequest(req)
	if err != nil {
		fmt.Println("failedddddddd")
		return nil, err
	}
	defer resp.Body.Close()
	//fasthttp.ReleaseResponse(resp)

	//statusCode := resp.StatusCode()
	//respBody := resp.Body()

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
		return finalTickers, nil
	}
	return nil, fmt.Errorf("invalid response code: %d", tickers.Code)
}
