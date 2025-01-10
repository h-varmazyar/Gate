package models

import "gorm.io/gorm"

const AssetTableName = "assets"

type Asset struct {
	gorm.Model
	Name   string
	Symbol string `gorm:"unique"`
}

func (o Asset) Table() string {
	return AssetTableName
}
