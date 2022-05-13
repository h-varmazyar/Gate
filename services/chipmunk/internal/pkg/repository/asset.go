package repository

import (
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"gorm.io/gorm"
	"time"
)

type Asset struct {
	gormext.UniversalModel
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol" gorm:"unique"`
	IssueDate time.Time `json:"issue_date"`
}

type AssetRepository struct {
	db *gorm.DB
}

func (repository *AssetRepository) ReturnBySymbol(symbol string) (*Asset, error) {
	asset := new(Asset)
	return asset, repository.db.Model(&Asset{}).Where("symbol = ?", symbol).First(asset).Error
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
