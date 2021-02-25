package models

type Trade struct {
	Time   float64
	Price  float64
	Volume float64
	Type   TradeType
}

type TradeType string

const (
	Buy  = "buy"
	Sell = "sell"
)
