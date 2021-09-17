package core

import (
	"github.com/fatih/color"
	"github.com/mrNobody95/Gate/brokerages/coinex"
	"github.com/mrNobody95/Gate/brokerages/nobitex"
	"github.com/mrNobody95/Gate/models"
)

func StartNewNode(brokerageName, configPath string) {
	brokerage, err := models.LoadBrokerage(brokerageName)
	if err != nil {
		color.Red("brokerage loading error: ", err.Error())
		return
	}
	node := Node{
		Brokerage: brokerage,
	}
	switch models.BrokerageName(brokerageName) {
	case models.Nobitex:
		node.Requests = nobitex.Config{Token: brokerage.Token}
	case models.Binance:
	case models.Coinex:
		node.Requests = coinex.Config{
			ApiId:   "",
			ApiHash: "",
		}
	}
	err = node.LoadConfig(configPath)
	if err != nil {
		color.Red("config loading failed: %s", err.Error())
		return
	}

	err = node.Validate()
	if err != nil {
		color.Red("brokerage not valid: %s", err.Error())
		return
	}
	err = node.Start()
	if err != nil {
		color.Red("brokerage starting failed: %s", err.Error())
		return
	}
}
