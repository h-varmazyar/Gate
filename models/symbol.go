package models

type Symbol struct {
	Id             uint16
	Value          string    `gorm:"size:50"`
	Brokerage      Brokerage `gorm:"foreignKey:BrokerageRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BrokerageRefer uint8
}
