package entity

import (
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"gorm.io/gorm"
	"time"
)

type Candle struct {
	gormext.UniversalModel
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Time            time.Time
	Open            float64
	High            float64
	Low             float64
	Close           float64
	Volume          float64
	Amount          float64
	MarketID        uuid.UUID
	ResolutionID    uuid.UUID
	IndicatorValues `gorm:"-"`
}
