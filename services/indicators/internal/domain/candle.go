package domain

import "time"

type Candle struct {
	ID     uint      `json:"id"`
	Time   time.Time `json:"time"`
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
	Volume float64   `json:"volume"`
}

type CandlesConsumerPayload struct {
	MarketID     uint
	ResolutionID uint
	Candles      []Candle
}
