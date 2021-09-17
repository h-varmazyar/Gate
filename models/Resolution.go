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

func (resolution *Resolution) CreateOrUpdate() error {
	count := int64(0)
	err := db.Model(&Resolution{}).
		Where("brokerage_refer = ?", resolution.BrokerageRefer).
		Where("value LIKE ?", resolution.Value).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		return db.Model(&Resolution{}).Create(&resolution).Error
	}
	return nil
}
