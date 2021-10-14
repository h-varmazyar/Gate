package models

import (
	"gorm.io/gorm"
	"strings"
	"time"
)

type Asset struct {
	gorm.Model
	Name      string
	Symbol    string   `gorm:"unique"`
	Markets   []Market `gorm:"-"`
	Rank      int
	IssueDate time.Time
}

func GetAssetBySymbol(symbol string) (*Asset, error) {
	asset := new(Asset)
	return asset, db.Where("symbol LIKE ?", strings.ToUpper(symbol)).First(&symbol).Error
}
