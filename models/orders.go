package models

type Order struct {
	Price  float64
	Volume float64
	Qty    int
}

type OrderType string

const (
	Bid = "bid"
	Ask = "ask"
)
