package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mrNobody95/Gate/core"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//cli.Execute()

	exit := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		core.Stop()
		done <- true
	}()
	core.StartNewNode("coinex", "")

	<-done
	fmt.Println("exiting")

	return
	//
	//fmt.Println(time.Unix(1606780800, 0))

	//color.HiRed("salam")
	//fmt.Println("second:",color.HiRedString("salam2"))
	//return
	//nobitex.Config{}.OHLC(nobitex.OHLCParams{
	//	Resolution:                       models.Resolution{
	//		Label:          "hour",
	//		Name:          "60",
	//		Duration:       time.Hour,
	//	},
	//	Market:                           models.Market{
	//		Id:             0,
	//		Name:          "BTCIRT",
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
		//		Name: "60",
		//	},
		//},
		//HelperResolution: map[models.Market]models.Resolution{
		//	nobitex.BTCIRT: {
		//		Name: "D",
		//	},
		//},
		//IndicatorConfig: indicators.Configuration{
		//	Length:            14,
		//	StochasticD: 3,
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
