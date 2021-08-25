package models

type Symbol struct {
	Id             uint16
	Value          string `gorm:"size:50"`
	Brokerage      Brokerage
	BrokerageRefer uint8 `gorm:"foreignKey:BrokerageRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
