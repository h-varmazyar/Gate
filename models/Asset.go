package models

import (
	"gorm.io/gorm"
	"strings"
	"time"
)

type Asset struct {
	gorm.Model
	Name      string
	Symbol    string `gorm:"unique"`
	Rank      int
	IssueDate time.Time
}

func GetAssetBySymbol(symbol string) (*Asset, error) {
	asset := new(Asset)
	return asset, db.Model(&Asset{}).Where("symbol LIKE ?", strings.ToUpper(symbol)).First(asset).Error
}

func CreateAsset(symbol string) (*Asset, error) {
	asset := &Asset{
		Name:   symbol,
		Symbol: symbol,
	}
	return asset, db.Model(&Asset{}).Save(asset).Error
}
