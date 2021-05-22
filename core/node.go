package core

import (
	"errors"
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/models"
)

type Node struct {
	Symbol       brokerages.Symbol
	Currency     brokerages.Currency
	CacheCandles []models.Candle
	Brokerage    brokerages.Brokerage
}

func (n *Node) Start() error {
	if n.Symbol == "" {
		return errors.New("please set symbol first")
	}
	if n.Currency == "" {
		return errors.New("please set currency first")
	}
	if n.Brokerage == nil {
		return errors.New("please specify brokerage config first")
	}
}
