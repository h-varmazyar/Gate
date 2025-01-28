package candles

type Candle struct {
	Timestamp int64   `json:"timestamp"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Volume    float64 `json:"volume"`
}

type CandlePayload struct {
	MarketID     uint     `json:"market_id"`
	ResolutionID uint     `json:"resolution_id"`
	Candles      []Candle `json:"candles"`
}

type TickerPayload struct {
	MarketID  uint    `json:"market_id"`
	LastPrice float64 `json:"last_price"`
}
