package markets

type MarketCreateReq struct {
	PricingDecimal    float64
	TradingDecimal    float64
	TakerFeeRate      float64
	MakerFeeRate      float64
	DestinationSymbol string
	IssueDate         int64
	MinAmount         float64
	SourceSymbol      string
	IsAMM             bool
	Name              string
	Status            string `enums:"Enable,Disable"`
	Platform          string `enums:"Coinex,UnknownBrokerage,Nobitex,Mazdax,Binance"`
}
