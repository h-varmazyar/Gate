package todo

import "github.com/mrNobody95/Gate/models"

type Trade struct {
	Time   float64
	Price  float64
	Volume float64
	Type   models.OrderType
}
