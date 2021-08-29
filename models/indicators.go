package models

type Trend int

type Indicators struct {
	MovingAverage
	BollingerBand
	Stochastic
	RSI
	ADX
	ParabolicSAR
	MACD
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

type ParabolicSAR struct {
	SAR          float64
	Trend        Trend
	TrendFlipped bool
}

type MACD struct {
	MACD   float64
	Signal float64
}

type MovingAverage struct {
	Simple      float64
	Exponential float64
}
