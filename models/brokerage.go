package models

type BrokerageName string

const (
	Binance BrokerageName = "binance"
	Coinex  BrokerageName = "coinex"
	Nobitex BrokerageName = "nobitex"
)

type Brokerage struct {
	Id       uint8         `gorm:"primarykey"`
	Name     BrokerageName `gorm:"size:50"`
	Token    string        `gorm:"size:150"`
	Username string        `gorm:"size:50"`
	Password string        `gorm:"size:100"`
}

func LoadBrokerage(brokerageName string) (*Brokerage, error) {
	br := &Brokerage{}
	return br, db.Model(&Brokerage{}).Where("name LIKE ?", brokerageName).First(&br).Error
}
