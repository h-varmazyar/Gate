package models

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 04.12.21
* Github: https://github.com/h-varmazyar
* Email: hossein.varmazyar@yahoo.com
**/

type Indicators struct {
	BollingerBands
	MovingAverage
	Stochastic
	RSI
}

type BollingerBands struct {
	UpperBand float64
	LowerBand float64
	MA        float64
}

type MovingAverage struct {
	Simple      float64
	Exponential float64
}

type Stochastic struct {
	IndexK float64
	IndexD float64
	FastK  float64
}

type RSI struct {
	Gain float64
	Loss float64
	RSI  float64
}
