package entity

import (
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"time"
)

type Market struct {
	gormext.UniversalModel
	BrokerageID     uuid.UUID
	PricingDecimal  float64
	TradingDecimal  float64
	TakerFeeRate    float64
	MakerFeeRate    float64
	DestinationID   uuid.UUID
	Destination     *Asset `gorm:"->;foreignkey:DestinationID;references:ID"`
	StartTime       time.Time
	MinAmount       float64
	SourceID        uuid.UUID
	Source          *Asset `gorm:"->;foreignkey:SourceID;references:ID"`
	IsAMM           bool
	Name            string
	Status          api.Status
	SourceName      string `gorm:"-"`
	DestinationName string `gorm:"-"`
}
