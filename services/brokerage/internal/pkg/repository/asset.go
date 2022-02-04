package repository

import (
	"gorm.io/gorm"
	"time"
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
* Date: 12.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type Asset struct {
	gorm.Model
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol" gorm:"unique"`
	IssueDate time.Time `json:"issue_date"`
}

type AssetRepository struct {
	db *gorm.DB
}

func (repository *AssetRepository) ReturnBySymbol(symbol string) (*Asset, error) {
	asset := new(Asset)
	return asset, repository.db.Model(&Asset{}).Where("symbol LIKE ?", symbol).First(asset).Error
}

func (repository *AssetRepository) Create(asset *Asset) (*Asset, error) {
	return asset, repository.db.Model(&Asset{}).Create(asset).Error
}

func (repository *AssetRepository) Set(asset *Asset) (*Asset, error) {
	return asset, repository.db.Model(&Asset{}).Save(asset).Error
}

func (repository *AssetRepository) List(page int) ([]*Asset, error) {
	assets := make([]*Asset, 0)
	return assets, repository.db.Model(&Asset{}).Offset((page - 1) * 25).Limit(25).Find(&assets).Error
}
