package repository

type BollingerBandsValue struct {
	UpperBand float64
	LowerBand float64
	MA        float64
}

type MovingAverageValue struct {
	Simple      float64
	Exponential float64
}

type StochasticValue struct {
	IndexK float64
	IndexD float64
	FastK  float64
}

type RSIValue struct {
	Gain float64
	Loss float64
	RSI  float64
}

type IndicatorValue struct {
	BB         *BollingerBandsValue
	MA         *MovingAverageValue
	Stochastic *StochasticValue
	RSI        *RSIValue
}
