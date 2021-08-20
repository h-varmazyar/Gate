package core

import (
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/brokerages/nobitex"
	"github.com/mrNobody95/Gate/indicators"
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/strategies"
	log "github.com/sirupsen/logrus"
	"time"
)

type Node struct {
	Strategy        strategies.Strategy
	Algorithm       interface{}
	NetworkManager  interface{}
	BufferedCandles map[brokerages.Symbol]map[string][]models.Candle
}

func (node *Node) Start() error {
	node.Strategy.Brokerage = nobitex.Config{}
	if err := node.Strategy.Validate(); err != nil {
		return err
	}
	if err := node.CollectPrimaryData(); err != nil {
		return err
	}

	return nil
}

func (node *Node) CollectPrimaryData() error {
	for _, symbol := range node.Strategy.Symbols {
		go func(symbol brokerages.Symbol) {
			if err := node.makeOHLCRequest(symbol, node.Strategy.PivotResolution[symbol]); err != nil {
				log.Errorf("ohlc request failed for symbol %s: %s", symbol, err.Error())
			}
		}(symbol)
		go func(symbol brokerages.Symbol) {
			if err := node.makeOHLCRequest(symbol, node.Strategy.HelperResolution[symbol]); err != nil {
				log.Errorf("ohlc request failed for symbol %s: %s", symbol, err.Error())
			}
		}(symbol)
	}
	return nil
}

func (node *Node) makeOHLCRequest(symbol brokerages.Symbol, resolution models.Resolution) error {
	conf := indicators.IndicatorConfig{
		Length: node.Strategy.IndicatorCalcLength,
	}
	firstTime := false
	candle := models.Candle{
		Symbol:     symbol,
		Resolution: resolution,
		Brokerage:  node.Strategy.Brokerage.GetName(),
	}
	err := candle.LoadLast()
	if err != nil {
		if err.Error() == "record not found" {
			candle.Time = time.Now().Add(-time.Hour * 24 * 365).Unix()
			firstTime = true
		} else {
			return err
		}
	} else {
		conf.Candles = append(conf.Candles, candle)
	}
	response := node.Strategy.Brokerage.OHLC(nobitex.OHLCParams{
		Resolution: candle.Resolution,
		Symbol:     candle.Symbol,
		From:       candle.Time,
		To:         time.Now().Unix(),
	})
	if response.Error != nil {
		return response.Error
	}
	conf.CalculateCandleIndicators(response.Candles, firstTime)
	node.UpdateBufferedData(conf.Candles, symbol, resolution.Label)
	return nil
}

func (node *Node) UpdateBufferedData(candles []models.Candle, symbol brokerages.Symbol, resolution string) {
	tmp := node.BufferedCandles[symbol][resolution]
	tmp = append(tmp, candles...)
	if diff := len(tmp) - node.Strategy.IndicatorCalcLength; diff > 0 {
		node.BufferedCandles[symbol][resolution] = tmp[diff:]
	}
}
