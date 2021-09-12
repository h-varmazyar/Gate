package core

import (
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/strategies"
)

type Node struct {
	Strategy  *strategies.Strategy
	Requests  brokerages.BrokerageRequests
	Brokerage *models.Brokerage
	//NetworkManager   interface{}
	//IndicatorConfig  indicators.Configuration
	//PivotResolution  map[models.Symbol]models.Resolution
	//HelperResolution map[models.Symbol]models.Resolution
	//BufferedCandles  map[models.Symbol]map[string][]models.Candle
}

func (node *Node) Validate() error {
	//if node.Brokerage == nil {
	//	return errors.New("you must declared working brokerage")
	//}
	//if err := node.Brokerage.Validate(); err != nil {
	//	return err
	//}
	//if len(node.PivotResolution) == 0 {
	//	return errors.New("pivot time frame must be declared")
	//}
	return nil
}

func (node *Node) Start() error {
	//if err := node.Validate(); err != nil {
	//	return err
	//}
	//if err := node.Strategy.Validate(); err != nil {
	//	return err
	//}
	//if err := node.CollectPrimaryData(); err != nil {
	//	return err
	//}

	return nil
}

//func (node *Node) CollectPrimaryData() error {
//	for _, symbol := range node.Strategy.Symbols {
//		go func(symbol models.Symbol) {
//			if err := node.makeOHLCRequest(symbol, node.PivotResolution[symbol]); err != nil {
//				log.Errorf("ohlc request failed for symbol %s: %s", symbol, err.Error())
//			}
//		}(symbol)
//		go func(symbol models.Symbol) {
//			if err := node.makeOHLCRequest(symbol, node.HelperResolution[symbol]); err != nil {
//				log.Errorf("ohlc request failed for symbol %s: %s", symbol, err.Error())
//			}
//		}(symbol)
//	}
//	return nil
//}

//func (node *Node) makeOHLCRequest(symbol models.Symbol, resolution models.Resolution) error {
//	log.Infof("make ohlc request:%s - %s - %s", symbol, resolution.Value, node.Brokerage.GetName())
//	conf := indicators.Configuration{
//		Length: node.Strategy.IndicatorCalcLength,
//	}
//	firstTime := false
//	candle := models.Candle{
//		Symbol:     symbol,
//		Resolution: resolution,
//	}
//	err := candle.LoadLast()
//	if err != nil {
//		if err.Error() == "record not found" {
//			candle.Time = time.Now().Add(-time.Hour * 24 * 365).Unix()
//			firstTime = true
//		} else {
//			return err
//		}
//	} else {
//		conf.Candles = append(conf.Candles, candle)
//	}
//	response := node.Brokerage.OHLC(nobitex.OHLCParams{
//		Resolution: candle.Resolution,
//		Symbol:     candle.Symbol,
//		From:       candle.Time,
//		To:         time.Now().Unix(),
//	})
//	if response.Error != nil {
//		return response.Error
//	}
//	fmt.Println(firstTime)
//	conf.CalculateIndicators(response.Candles, firstTime)
//	node.UpdateBufferedData(conf.Candles, symbol, resolution.Value)
//	return nil
//}
//
//func (node *Node) UpdateBufferedData(candles []models.Candle, symbol models.Symbol, resolution string) {
//	tmp := node.BufferedCandles[symbol][resolution]
//	tmp = append(tmp, candles...)
//	if diff := len(tmp) - node.Strategy.CandleBufferLength; diff > 0 {
//		node.BufferedCandles[symbol][resolution] = tmp[diff:]
//	}
//}
