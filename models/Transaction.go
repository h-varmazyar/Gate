package models

import "time"

type Transaction struct {
	ID            uint64
	Volume        string
	Currency      string
	CreatedAt     time.Time
	Description   string
	CalculatedFee string
}
