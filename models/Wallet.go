package models

type Wallet struct {
	ID                       int
	ReferenceCurrencyBalance float64
	BlockedBalance           string
	ActiveBalance            string
	TotalBalance             string
	Currency                 string
}
