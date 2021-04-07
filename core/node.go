package core

import (
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/models"
)

type Node struct {
	Symbol       brokerages.Symbol
	Currency     brokerages.Currency
	CacheCandles []models.Candle
	Brokerage    brokerages.Brokerage
}

//func (n *Node) Start() {
//
//}
