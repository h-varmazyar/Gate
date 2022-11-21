package entity

import (
	"github.com/google/uuid"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/gormext"
)

type Brokerage struct {
	gormext.UniversalModel
	Title        string       `gorm:"not null"`
	Description  string       `gorm:"type:varchar(1000)"`
	Platform     api.Platform `gorm:"type:varchar(25);not null"`
	AuthType     api.AuthType `gorm:"type:varchar(25);not null"`
	Username     string       `gorm:"type:varchar(100)"`
	Password     string       `gorm:"type:varchar(100)"`
	AccessID     string       `gorm:"type:varchar(100)"`
	SecretKey    string       `gorm:"type:varchar(100)"`
	Status       api.Status   `gorm:"type:varchar(25);not null"`
	ResolutionID *uuid.UUID   `gorm:"not null"`
	//ResolutionID *uuid.UUID             `gorm:"type:uuid REFERENCES resolutions(id);not null;index"`
	//Resolution   *Resolution           `gorm:"->;foreignkey:ResolutionID;references:ID"`
	//Markets      []*Market  `gorm:"many2many:brokerage_markets"`
	StrategyID *uuid.UUID `gorm:"not null"`
}
