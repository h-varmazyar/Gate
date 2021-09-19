package models

import "time"

type Market struct {
	Id              uint16
	Value           string    `gorm:"size:50"`
	Brokerage       Brokerage `gorm:"foreignKey:BrokerageRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StartTime       time.Time
	StartTimeString string `yaml:"startTime" gorm:"-"`
	BrokerageRefer  uint8
}

func (market *Market) CreateOrLoad() error {
	err := db.Model(&Market{}).
		Where("brokerage_refer = ?", market.BrokerageRefer).
		Where("value LIKE ?", market.Value).
		First(&market).Error
	if err != nil && err.Error() == "record not found" {
		return db.Model(&Market{}).Create(&market).Error
	}
	return err
}
