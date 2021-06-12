package models

import "time"

type Resolution struct {
	Duration time.Duration `gorm:"duration"`
	Value    interface{}   `gorm:"value"`
	Label    string        `gorm:"resolution_label"`
}
