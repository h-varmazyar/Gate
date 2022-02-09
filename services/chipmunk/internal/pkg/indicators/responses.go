package indicators

type BollingerBandsResponse struct {
	UpperBand float64
	LowerBand float64
	MA        float64
}

type MovingAverageResponse struct {
	Simple      float64
	Exponential float64
}

type StochasticResponse struct {
	IndexK float64
	IndexD float64
	FastK  float64
}

type RSIResponse struct {
	Gain float64
	Loss float64
	RSI  float64
}
