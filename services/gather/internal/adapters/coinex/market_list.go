package coinex

import (
	"encoding/json"
	"fmt"
	"github.com/h-varmazyar/Gate/services/gather/internal/domain"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"strconv"
	"time"
)

type market struct {
	BuyAssetType  string `json:"buy_asset_type"`
	CreateTime    int64  `json:"create_time"`
	IsPre         bool   `json:"is_pre"`
	LeastAmount   string `json:"least_amount"`
	MakerFeeRate  string `json:"maker_fee_rate"`
	Market        string `json:"market"`
	SellAssetType string `json:"sell_asset_type"`
	TakerFeeRate  string `json:"taker_fee_rate"`
}

type marketList struct {
	DefaultTradingArea string            `json:"default_trading_area"`
	MarketInfo         map[string]market `json:"market_info"`
}

func (c Coinex) MarketList(ctx context.Context) (domain.CoinexMarkets, error) {
	url := fmt.Sprintf("%v/res/market/", c.cfg.BaseURL)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return domain.CoinexMarkets{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return domain.CoinexMarkets{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.CoinexMarkets{}, fmt.Errorf("invalid HTTP response code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.CoinexMarkets{}, err
	}

	respEntity := baseResponse[marketList]{}
	err = json.Unmarshal(body, &respEntity)
	if err != nil {
		return domain.CoinexMarkets{}, err
	}

	if respEntity.Code != 0 {
		return domain.CoinexMarkets{}, fmt.Errorf("invalid response code: %d", respEntity.Code)
	}

	markets := make([]domain.CoinexMarket, 0)
	for _, market := range respEntity.Data.MarketInfo {
		m := domain.CoinexMarket{
			Market:                      market.Market,
			BaseCurrency:                market.SellAssetType,
			QuoteCurrency:               market.BuyAssetType,
			IsPremarketTradingAvailable: market.IsPre,
			IssueDate:                   time.Unix(market.CreateTime, 0),
		}
		m.MakerFeeRate, err = strconv.ParseFloat(market.MakerFeeRate, 64)
		if err != nil {
			continue
		}
		m.TakerFeeRate, err = strconv.ParseFloat(market.TakerFeeRate, 64)
		if err != nil {
			continue
		}
		m.MinAmount, err = strconv.ParseFloat(market.LeastAmount, 64)
		if err != nil {
			continue
		}
		markets = append(markets, m)
	}

	coinexMarkets := domain.CoinexMarkets{
		List: markets,
	}
	return coinexMarkets, nil
}
