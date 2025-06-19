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
	url := fmt.Sprintf("%v/spot/ticker", c.cfg.APIBaseURL)
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
