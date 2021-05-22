package main

import (
	"github.com/mrNobody95/Gate/brokerages"
)

func main() {
	input(brokerages.NobitexConfig{})
	//s:=strategies.Strategy{
	//	IsHFT:              false,
	//	Symbols:            nil,
	//	StopLoss:           0,
	//	Brokerage:          brokerages.NobitexConfig{},
	//	TimeFrame:          "",
	//	MinBenefit:         0,
	//	MaxBenefit:         0,
	//	PrimaryAmount:      0,
	//	CurrentAmount:      0,
	//	PrimaryCandles:     nil,
	//	BufferedCandles:    nil,
	//	PrimaryCurrency:    "",
	//	CurrentCurrency:    "",
	//	MaxDailyBenefit:    0,
	//	ReservePercentage:  0,
	//	CandleBufferLength: 0,
	//}
}

func input(brok brokerages.Brokerage) {

}
