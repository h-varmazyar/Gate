package core

import (
	"errors"
	"github.com/fatih/color"
	"github.com/jinzhu/copier"
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/indicators"
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/strategies"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type Node struct {
	Markets         []models.Market
	Strategy        *strategies.Strategy
	Requests        brokerages.BrokerageRequests
	Brokerage       *models.Brokerage
	Resolutions     []models.Resolution
	FakeTrading     bool
	EnableTrading   bool
	IndicatorConfig indicators.Configuration
	//NetworkManager   interface{}
	//IndicatorConfig  indicators.Configuration
	//PivotResolution  map[models.Market]models.Resolution
	//HelperResolution map[models.Market]models.Resolution
	//BufferedCandles  map[models.Market]map[string][]models.Candle
}

func (node *Node) Validate() error {
	if node.Brokerage == nil {
		return errors.New("you must declared working brokerage")
	}
	if node.Requests == nil {
		return errors.New("brokerage request API not declared")
	}
	if node.Strategy == nil {
		return errors.New("trading strategy must be declared")
	}
	return nil
}

func (node *Node) Start() error {
	for _, market := range node.Markets {
		go func(market models.Market) {
			indicatorConfigs := make(map[uint]*indicators.Configuration)
			for _, resolution := range node.Resolutions {
				tmp := indicators.Configuration{}
				err := copier.Copy(&tmp, &node.IndicatorConfig)
				if err != nil {
					color.Red("copy indicator config failed for %s in resolution %s: %s",
						market.Value, resolution.Value, err.Error())
					return
				}
				tmp.Candles = make([]models.Candle, 0)
				indicatorConfigs[resolution.Id] = &tmp
			}
			thread := BrokerageSymbolThread{
				Node:                         node,
				Market:                       market,
				StartFrom:                    time.Time{},
				IndicatorConfigPerResolution: indicatorConfigs,
			}
			thread.CollectPrimaryData()
			thread.PeriodicOHLC()
			if node.EnableTrading || node.FakeTrading {
				thread.checkForSignals()
			}
		}(market)
	}
	return nil
}

func (node *Node) LoadConfig(path string) error {
	if path == "" {
		path = node.Brokerage.DefaultConfigPath
	}
	confBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var config NodeConfig
	if err = yaml.Unmarshal(confBytes, &config); err != nil {
		return err
	}
	if config.LoadMarketsOnline {
		resp := node.Requests.MarketList(nil).(brokerages.MarketListResponse)
		if resp.Error != nil {
			return resp.Error
		}
		for _, market := range resp.Markets {
			err = market.CreateOrLoad()
			if err != nil {
				return err
			}
			node.Markets = append(node.Markets, market)
		}
	}
	node.Resolutions = config.Resolutions
	node.Strategy = &config.Strategy
	node.FakeTrading = config.FakeTrading
	node.EnableTrading = config.EnableTrading
	return nil
}

//func (node *Node) CollectPrimaryData() error {
//	for _, symbol := range node.Strategy.Symbols {
//		go func(symbol models.Market) {
//			if err := node.makeOHLCRequest(symbol, node.PivotResolution[symbol]); err != nil {
//				log.Errorf("ohlc request failed for symbol %s: %s", symbol, err.Error())
//			}
//		}(symbol)
//		go func(symbol models.Market) {
//			if err := node.makeOHLCRequest(symbol, node.HelperResolution[symbol]); err != nil {
//				log.Errorf("ohlc request failed for symbol %s: %s", symbol, err.Error())
//			}
//		}(symbol)
//	}
//	return nil
//}

//func (node *Node) makeOHLCRequest(symbol models.Market, resolution models.Resolution) error {
//	log.Infof("make ohlc request:%s - %s - %s", symbol, resolution.Value, node.Brokerage.GetName())
//	conf := indicators.Configuration{
//		Length: node.Strategy.IndicatorCalcLength,
//	}
//	firstTime := false
//	candle := models.Candle{
//		Market:     symbol,
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
//		Market:     candle.Market,
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
//func (node *Node) UpdateBufferedData(candles []models.Candle, symbol models.Market, resolution string) {
//	tmp := node.BufferedCandles[symbol][resolution]
//	tmp = append(tmp, candles...)
//	if diff := len(tmp) - node.Strategy.BufferedCandleCount; diff > 0 {
//		node.BufferedCandles[symbol][resolution] = tmp[diff:]
//	}
//}
