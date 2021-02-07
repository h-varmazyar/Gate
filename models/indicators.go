package models

type Indicators struct {
	Stochastic Stochastic
}

type Stochastic struct {
	IndexK float64
	IndexD float64
}
