package entity

import (
	chipmunkAPI "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	Id             string               `gorm:"primarykey"`
	PostedAt       time.Time            `gorm:"not null"`
	Content        string               `gorm:"type:varchar(4094)"`
	LikeCount      uint32               `gorm:"default:0"`
	ParentId       string               `gorm:"type:varchar(32)"`
	SenderUsername string               `gorm:"type:varchar(64)"`
	Provider       chipmunkAPI.Provider `gorm:"type:varchar(32)"`
	Tags           pq.StringArray       `gorm:"type:varchar(32)[]"`
	Sentiment      chipmunkAPI.Polarity `gorm:"type:varchar(32)"`
}
