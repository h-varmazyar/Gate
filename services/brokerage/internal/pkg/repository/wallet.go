package repository

import "gorm.io/gorm"

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 20.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type Wallet struct {
	ReferenceCurrencyBalance float64
	BlockedBalance           float64
	ActiveBalance            float64
	TotalBalance             float64
	Asset                    string
	ID                       string
}

type WalletRepository struct {
	db *gorm.DB
}
