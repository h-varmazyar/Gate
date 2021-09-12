package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/mrNobody95/Gate/cli"
	"github.com/mrNobody95/Gate/core"
	log "github.com/sirupsen/logrus"
)

func main() {
	cli.Execute()
	return
	//nobitex.Config{}.OHLC(nobitex.OHLCParams{
	//	Resolution:                       models.Resolution{
	//		Label:          "hour",
	//		Value:          "60",
	//		Duration:       time.Hour,
	//	},
	//	Symbol:                           models.Symbol{
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
		//PivotResolution: map[models.Symbol]models.Resolution{
		//	nobitex.BTCIRT: {
		//		Value: "60",
		//	},
		//},
		//HelperResolution: map[models.Symbol]models.Resolution{
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
		//	//Symbols: []models.Symbol{
		//	//	nobitex.BTCIRT,
		//	//},
		//	StopLoss:            3,
		//	ReservePercentage:   0,
		//	CandleBufferLength:  400,
		//	IndicatorCalcLength: 14,
		//},
	}).Start()
	if err != nil {
		log.Panicf("core node start failed: %v", err)
	}

	select {}
}
