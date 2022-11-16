package entity

import (
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"time"
)

type Resolution struct {
	gormext.UniversalModel
	BrokerageName string
	Duration      time.Duration
	Label         string
	Value         string
}
