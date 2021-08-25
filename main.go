package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mrNobody95/Gate/brokerages/nobitex"
	"github.com/mrNobody95/Gate/core"
	"github.com/mrNobody95/Gate/indicators"
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/strategies"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	fmt.Println(time.Hour.Milliseconds())
	fmt.Println((time.Hour * 24).Milliseconds())
	err := (&core.Node{
		Brokerage: nobitex.Config{
			Brokerage: models.Brokerage{
				Name:     models.Nobitex,
				Username: "hossein.varmazyar@yahoo.com",
				Password: "Hv0017773474",
			},
			Token: "19aa9b3edc7fb2e467819e32d7585bf3ea31c49a",
		},
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
		IndicatorConfig: indicators.Configuration{
			Length:            14,
			StochasticSmoothD: 3,
			StochasticSmoothK: 1,
			MacdFastLength:    12,
			MacdSlowLength:    26,
			MacdSignalLength:  9,
		},
		Strategy: strategies.Strategy{
			IsHFT:           false,
			PrimaryCurrency: nobitex.RLS,
			//Symbols: []models.Symbol{
			//	nobitex.BTCIRT,
			//},
			StopLoss:            3,
			ReservePercentage:   0,
			CandleBufferLength:  400,
			IndicatorCalcLength: 14,
		},
	}).Start()
	if err != nil {
		log.Panicf("core node start failed: %v", err)
	}

	for {

	}
}
