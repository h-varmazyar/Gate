package markets

type MarketReq struct {
	PricingDecimal    float64 `json:"pricing_decimal"`
	TradingDecimal    float64 `json:"trading_decimal"`
	TakerFeeRate      float64 `json:"taker_fee_rate"`
	MakerFeeRate      float64 `json:"maker_fee_rate"`
	DestinationSymbol string  `json:"destination_symbol"`
	IssueDate         int64   `json:"issue_date"`
	MinAmount         float64 `json:"min_amount"`
	SourceSymbol      string  `json:"source_symbol"`
	IsAMM             bool    `json:"is_amm"`
	Name              string  `json:"name"`
	Status            string  `json:"status" enums:"NoneS,Enable,Disable"`
	Platform          string  `json:"platform" enums:"Coinex,UnknownBrokerage,Nobitex,Mazdax,Binance"`
}

type Platform struct {
	Platform string `json:"platform" enums:"Coinex,UnknownBrokerage,Nobitex,Mazdax,Binance"`
}
