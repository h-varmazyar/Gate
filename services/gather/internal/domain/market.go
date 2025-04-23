package domain

type CoinexMarket struct {
	Market                      string  `json:"market"`
	BaseCurrency                string  `json:"base_ccy"`
	QuoteCurrency               string  `json:"quote_ccy"`
	MakerFeeRate                float64 `json:"maker_fee_rate"`
	TakerFeeRate                float64 `json:"taker_fee_rate"`
	MinAmount                   float64 `json:"min_amount"`
	BaseCurrencyPrecision       uint8   `json:"base_ccy_precision"`
	QuoteCurrencyPrecision      uint8   `json:"quote_ccy_precision"`
	IsAmmAvailable              bool    `json:"is_amm_available"`
	IsApiTradingAvailable       bool    `json:"is_api_trading_available"`
	IsMarginAvailable           bool    `json:"is_margin_available"`
	IsPremarketTradingAvailable bool    `json:"is_premarket_trading_available"`
}

type CoinexMarkets struct {
	List []CoinexMarket `json:"list"`
}
