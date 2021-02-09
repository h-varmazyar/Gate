package models

type Indicators struct {
	Stochastic    *Stochastic
	BollingerBand *BollingerBand
	RSI           *RSI
	ADX           *ADX
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

type ADX struct {
	DIPositive float64
	DINegative float64
	ADX        float64
	DmPositive float64
	DmNegative float64
	TR         float64
	DX         float64
}
