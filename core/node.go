package core

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/indicators"
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/storage"
	"github.com/mrNobody95/Gate/strategies"
	"github.com/mrNobody95/Gate/wallets"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
	"time"
)

type Node struct {
	Markets           []models.Market
	Strategy          *strategies.Strategy
	Requests          brokerages.BrokerageRequests
	Brokerage         *models.Brokerage
	Resolutions       []models.Resolution
	PivotResolution   *models.Resolution
	FakeTrading       bool
	EnableTrading     bool
	IndicatorConfig   *indicators.Configuration
	dataChannel       chan models.Candle
	WalletsController *wallets.Controller
}

func (node *Node) Validate() error {
	color.HiGreen("Validating node")
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
	color.HiGreen("Starting node for %s(markets: %d)", node.Brokerage.Name, len(node.Markets))
	node.dataChannel = make(chan models.Candle, 100)
	for _, market := range node.Markets {
		pool, err := storage.NewPool(node.Strategy.BufferedCandleCount, market.Id, node.PivotResolution.Id)
		if err != nil {
			//color.Red("create candle pool failed for market %s in timeframe %s: %v",
			//	market.Name, node.PivotResolution.Label, err.Error())
			return err
		}
		thread := MarketThread{
			Node:            node,
			Market:          &market,
			StartFrom:       market.StartTime,
			Resolution:      node.PivotResolution,
			IndicatorConfig: node.IndicatorConfig,
			CandlePool:      pool,
		}
		if err = thread.CollectPrimaryData(); err != nil {
			//log.WithError(err).Errorf("failed to collect primary data for %s in timeframe %s",
			//	thread.Market.Name, thread.PivotResolution.Label)
			return err
		}
		//node.WalletsController = wallets.NewController(node.Requests)
		//time.Sleep(time.Minute)
		//go func(market models.Market) {
		//	thread.PeriodicOHLC()
		//}(market)
	}
	return nil
}

func (node *Node) LoadConfig(path string) error {
	color.HiGreen("Loading YAML config")
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
		resp := node.Requests.MarketList()
		if resp.Error != nil {
			return resp.Error
		}
		config.Markets = resp.Markets
	} else {
		fmt.Println("offline markets")
		config.Markets, err = models.GetBrokerageMarkets(node.Brokerage.Id)
		if err != nil {
			return err
		}
	}
	if len(config.Markets) == 0 {
		return errors.New("no market available")
	}
	for _, market := range config.Markets {
		if market.StartTimeString != "" {
			t, err := strconv.ParseInt(market.StartTimeString, 10, 64)
			if err != nil {
				return err
			}
			market.StartTime = time.Unix(t, 0)
		} else {
			market.StartTime = time.Unix(1594512000, 0)
		}
		market.BrokerageRefer = node.Brokerage.Id
		market.Brokerage = *node.Brokerage
		err = market.CreateOrLoad()
		if err != nil {
			fmt.Println("market load failed")
			return err
		}
		node.Markets = append(node.Markets, market)
	}
	fmt.Println(config.Resolutions)
	for _, resolution := range config.Resolutions {
		resolution.BrokerageRefer = node.Brokerage.Id
		err = resolution.CreateOrLoad()
		if err != nil {
			return err
		}
		node.Resolutions = append(node.Resolutions, resolution)
	}

	node.PivotResolution = &node.Resolutions[0]
	node.Strategy = &config.Strategy
	node.FakeTrading = config.FakeTrading
	node.EnableTrading = config.EnableTrading
	node.IndicatorConfig = &config.IndicatorConfigs
	return nil
}
