package entity

import (
	"github.com/google/uuid"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/gormext"
	"time"
)

type Market struct {
	gormext.UniversalModel
	Platform        api.Platform
	PricingDecimal  float64
	TradingDecimal  float64
	TakerFeeRate    float64
	MakerFeeRate    float64
	DestinationID   uuid.UUID
	Destination     *Asset `gorm:"->;foreignkey:DestinationID;references:ID"`
	IssueDate       time.Time
	MinAmount       float64
	SourceID        uuid.UUID
	Source          *Asset `gorm:"->;foreignkey:SourceID;references:ID"`
	IsAMM           bool
	Name            string
	Status          api.Status
	SourceName      string `gorm:"-"`
	DestinationName string `gorm:"-"`
}
