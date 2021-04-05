package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Volume        string
	Currency      string
	Description   string
	CalculatedFee string
}
