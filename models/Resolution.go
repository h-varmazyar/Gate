package models

type Resolution struct {
	Value interface{} `gorm:"-"`
	Label string      `gorm:"resolution_label"`
}
