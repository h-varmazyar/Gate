package models

type Wallet struct {
	ReferenceCurrencyBalance float64
	BlockedBalance           string
	ActiveBalance            string
	TotalBalance             string
	Currency                 string
}
