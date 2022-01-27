package repository

import (
	"github.com/google/uuid"
	"github.com/mrNobody95/Gate/pkg/gormext"
	"github.com/mrNobody95/Gate/services/brokerage/api"
	"gorm.io/gorm"
)

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
	gormext.UniversalModel
	ReferenceCurrencyBalance float64
	BlockedBalance           float64
	ActiveBalance            float64
	TotalBalance             float64
	AssetRefer               uuid.UUID
	Asset                    *Asset `gorm:"foreignKey:AssetRefer"`
	BrokerageRefer           uuid.UUID
	Brokerage                *Brokerage `gorm:"foreignKey:BrokerageRefer"`
}

type WalletRepository struct {
	db *gorm.DB
}

func (r *WalletRepository) List(brokerageName string) ([]*Wallet, error) {
	wallets := make([]*Wallet, 0)
	tx := r.db.Model(new(Wallet)).Joins("Asset").Joins("Brokerage")
	if brokerageName != api.Names_All.String() {
		tx.Where("brokerages.name LIKE ?", brokerageName)
	}
	return wallets, tx.Find(&wallets).Error
}
