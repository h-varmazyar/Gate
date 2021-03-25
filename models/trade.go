package models

type Trade struct {
	Time   float64
	Price  float64
	Volume float64
	Type   OrderType
}
