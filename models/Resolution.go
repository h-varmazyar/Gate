package models

import (
	"time"
)

type Resolution struct {
	Id             uint   `gorm:"primarykey"`
	Label          string `gorm:"size:50"`
	Value          string `gorm:"size:50"`
	Duration       time.Duration
	Brokerage      Brokerage `gorm:"foreignKey:BrokerageRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BrokerageRefer uint8
}

func (resolution *Resolution) CreateOrLoad() error {
	err := db.Model(&Resolution{}).
		Where("brokerage_refer = ?", resolution.BrokerageRefer).
		Where("value LIKE ?", resolution.Value).
		First(&resolution).Error
	if err != nil && err.Error() == "record not found" {
		return db.Model(&Resolution{}).Create(&resolution).Error
	}
	return err
}
