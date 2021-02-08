package models

type Indicators struct {
	Stochastic    Stochastic
	BollingerBand BollingerBand
	RSI           RSI
}

type Stochastic struct {
	IndexK float64
	IndexD float64
}

type BollingerBand struct {
	UpperBond float64
	LowerBond float64
	MA        float64
}

type RSI struct {
	Gain float64
	Loss float64
	RSI  float64
}
