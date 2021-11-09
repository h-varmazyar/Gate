package core

import (
	"github.com/fatih/color"
	"github.com/mrNobody95/Gate/brokerages/coinex"
	"github.com/mrNobody95/Gate/models"
)

var nodeArr []*Node

func Stop() {
	for _, node := range nodeArr {
		close(node.dataChannel)
	}
}

func StartNewNode(brokerageName, configPath string) {
	brokerage, err := models.LoadBrokerage(brokerageName)
	if err != nil {
		color.Red("brokerage loading error: ", err.Error())
		return
	}
	node := new(Node)
	node.Brokerage = brokerage
	switch models.BrokerageName(brokerageName) {
	case models.Nobitex:
		//node.Requests = nobitex.Config{Token: brokerage.Token}
	case models.Coinex:
		node.Requests = coinex.Config{
			AccessId:  "",
			SecretKey: "",
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
	nodeArr = append(nodeArr, node)
	err = node.Start()
	if err != nil {
		color.Red("brokerage starting failed: %s", err.Error())
		return
	}
}
