package models

import (
	api "github.com/h-varmazyar/Gate/api/proto"
	"gorm.io/gorm"
	"time"
)

const ResolutionTableName = "resolutions"

type Resolution struct {
	gorm.Model
	Platform api.Platform
	Duration time.Duration
	Label    string
	Value    string
}

func (o Resolution) Table() string {
	return ResolutionTableName
}
