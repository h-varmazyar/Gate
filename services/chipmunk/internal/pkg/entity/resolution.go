package entity

import (
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"time"
)

type Resolution struct {
	gormext.UniversalModel
	Platform api.Platform
	Duration time.Duration
	Label    string
	Value    string
}
