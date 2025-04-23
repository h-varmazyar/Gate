package coinex

import (
	"encoding/json"
	"fmt"
	"github.com/h-varmazyar/Gate/services/gather/internal/domain"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"strconv"
)

type marketInfo struct {
	Market                      string `json:"market"`
	BaseCurrency                string `json:"base_ccy"`
	QuoteCurrency               string `json:"quote_ccy"`
	MakerFeeRate                string `json:"maker_fee_rate"`
	TakerFeeRate                string `json:"taker_fee_rate"`
	MinAmount                   string `json:"min_amount"`
	BaseCurrencyPrecision       uint8  `json:"base_ccy_precision"`
	QuoteCurrencyPrecision      uint8  `json:"quote_ccy_precision"`
	IsAmmAvailable              bool   `json:"is_amm_available"`
	IsApiTradingAvailable       bool   `json:"is_api_trading_available"`
	IsMarginAvailable           bool   `json:"is_margin_available"`
	IsPremarketTradingAvailable bool   `json:"is_premarket_trading_available"`
}

func (c Coinex) MarketList(ctx context.Context) (domain.CoinexMarkets, error) {
	url := fmt.Sprintf("%v/spot/market", c.cfg.BaseURL)
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

	respEntity := baseResponse[[]marketInfo]{}
	err = json.Unmarshal(body, &respEntity)
	if err != nil {
		return domain.CoinexMarkets{}, err
	}

	if respEntity.Code != 0 {
		return domain.CoinexMarkets{}, fmt.Errorf("invalid response code: %d", respEntity.Code)
	}

	markets := make([]domain.CoinexMarket, len(respEntity.Data))
	for i, market := range respEntity.Data {
		markets[i] = domain.CoinexMarket{
			Market:                      market.Market,
			BaseCurrency:                market.BaseCurrency,
			QuoteCurrency:               market.QuoteCurrency,
			BaseCurrencyPrecision:       market.BaseCurrencyPrecision,
			QuoteCurrencyPrecision:      market.QuoteCurrencyPrecision,
			IsAmmAvailable:              market.IsAmmAvailable,
			IsApiTradingAvailable:       market.IsApiTradingAvailable,
			IsMarginAvailable:           market.IsMarginAvailable,
			IsPremarketTradingAvailable: market.IsPremarketTradingAvailable,
		}
		markets[i].MakerFeeRate, err = strconv.ParseFloat(market.MakerFeeRate, 64)
		if err != nil {
			continue
		}
		markets[i].TakerFeeRate, err = strconv.ParseFloat(market.TakerFeeRate, 64)
		if err != nil {
			continue
		}
		markets[i].MinAmount, err = strconv.ParseFloat(market.MinAmount, 64)
		if err != nil {
			continue
		}
	}

	coinexMarkets := domain.CoinexMarkets{
		List: markets,
	}
	return coinexMarkets, nil
}
