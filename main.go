package main

import (
	"fmt"
	"github.com/jinzhu/copier"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mrNobody95/Gate/core"
	"github.com/mrNobody95/Gate/indicators"
	"github.com/mrNobody95/Gate/models"
	log "github.com/sirupsen/logrus"
)

func main() {
	//cli.Execute()
	conf := indicators.Configuration{Candles: []models.Candle{
		{
			ID: 1,
		},
	}}

	var other indicators.Configuration
	copier.Copy(&other, &conf)
	other.Candles = []models.Candle{{Open: 10}}
	other.Candles[0].ID = 2
	fmt.Println(conf.Candles[0].ID)
	fmt.Println(other.Candles[0].ID)
	return
	//nobitex.Config{}.OHLC(nobitex.OHLCParams{
	//	Resolution:                       models.Resolution{
	//		Label:          "hour",
	//		Value:          "60",
	//		Duration:       time.Hour,
	//	},
	//	Market:                           models.Market{
	//		Id:             0,
	//		Value:          "BTCIRT",
	//	},
	//	From:                             time.Now().Add(time.Hour*(-4)).Unix(),
	//	To:                               time.Now().Unix(),
	//})
	//return
	err := (&core.Node{
		//Brokerage: nobitex.Config{
		//	Brokerage: models.Brokerage{
		//		Name:     models.Nobitex,
		//		Username: "hossein.varmazyar@yahoo.com",
		//		Password: "Hv0017773474",
		//	},
		//	Token: "19aa9b3edc7fb2e467819e32d7585bf3ea31c49a",
		//},
		//PivotResolution: map[models.Market]models.Resolution{
		//	nobitex.BTCIRT: {
		//		Value: "60",
		//	},
		//},
		//HelperResolution: map[models.Market]models.Resolution{
		//	nobitex.BTCIRT: {
		//		Value: "D",
		//	},
		//},
		//IndicatorConfig: indicators.Configuration{
		//	Length:            14,
		//	StochasticLengthD: 3,
		//	StochasticSmoothK: 1,
		//	MacdFastLength:    12,
		//	MacdSlowLength:    26,
		//	MacdSignalLength:  9,
		//},
		//Strategy: strategies.Strategy{
		//	IsHFT:           false,
		//	PrimaryCurrency: nobitex.RLS,
		//	//Symbols: []models.Market{
		//	//	nobitex.BTCIRT,
		//	//},
		//	StopLoss:            3,
		//	ReservePercentage:   0,
		//	BufferedCandleCount:  400,
		//	IndicatorCalcLength: 14,
		//},
	}).Start()
	if err != nil {
		log.Panicf("core node start failed: %v", err)
	}

	select {}
}
