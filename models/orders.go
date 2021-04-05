package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Fee                 float64
	Qty                 int
	User                string
	Price               float64
	Status              OrderStatus
	Volume              string
	OrderType           OrderType
	TotalPrice          string
	AveragePrice        string
	MatchedVolume       string
	SourceCurrency      string
	UnMatchedVolume     string
	DestinationCurrency string
}

type OrderType string
type OrderStatus string

const (
	Ask  OrderType = "ask"
	Bid  OrderType = "bid"
	Buy  OrderType = "buy"
	Sell OrderType = "sell"
)

const (
	All      OrderStatus = "all"
	New      OrderStatus = "new"
	Open     OrderStatus = "open"
	Done     OrderStatus = "done"
	Close    OrderStatus = "Close"
	Active   OrderStatus = "active"
	Canceled OrderStatus = "canceled"
)
