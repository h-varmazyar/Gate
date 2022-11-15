package entity

import (
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"time"
)

type Asset struct {
	gormext.UniversalModel
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol" gorm:"unique"`
	IssueDate time.Time `json:"issue_date"`
}
