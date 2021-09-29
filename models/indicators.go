package models

type Trend int

type Indicators struct {
	MovingAverage
	BollingerBand
	ParabolicSAR
	Stochastic
	MACD
	ATR
	ADX
	RSI
}

type Stochastic struct {
	IndexK float64
	IndexD float64
	FastK  float64
}

func (stochastic *Stochastic) SignalStrength() float64 {
	return float64(0)
}

type BollingerBand struct {
	UpperBond float64
	LowerBond float64
	MA        float64
}

func (bollingerBand *BollingerBand) SignalStrength() float64 {
	return float64(0)
}

type RSI struct {
	Gain float64
	Loss float64
	RSI  float64
}

func (rsi *RSI) SignalStrength() float64 {
	return float64(0)
}

type ATR struct {
	TR  float64
	ATR float64
}

func (atr *ATR) SignalStrength() float64 {
	return float64(0)
}

type ADX struct {
	DIPositive float64
	DINegative float64
	ADX        float64
	DmPositive float64
	TR         float64
	DmNegative float64
	DX         float64
}

func (adx *ADX) SignalStrength() float64 {
	return float64(0)
}

type ParabolicSAR struct {
	SAR          float64
	Trend        Trend
	TrendFlipped bool
}

func (pSar *ParabolicSAR) SignalStrength() float64 {
	return float64(0)
}

type MACD struct {
	MACD    float64
	Signal  float64
	SlowEMA float64
	FastEMA float64
}

func (macd *MACD) SignalStrength() float64 {
	return float64(0)
}

type MovingAverage struct {
	Simple      float64
	Exponential float64
}

func (movingAverage *MovingAverage) SignalStrength() float64 {
	return float64(0)
}
