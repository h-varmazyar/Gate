package models

import "time"

type Resolution struct {
	Id             uint
	Label          string `gorm:"size:50"`
	Value          string `gorm:"size:50"`
	Duration       time.Duration
	Brokerage      Brokerage `gorm:"foreignKey:BrokerageRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BrokerageRefer uint8
}
