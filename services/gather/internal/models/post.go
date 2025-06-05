package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type PostProvider string

const (
	PostProviderTwitter       PostProvider = "TWITTER"
	PostProviderSahamyab      PostProvider = "SAHAMYAB"
	PostProviderRahavard      PostProvider = "RAHAVARD"
	PostProviderCoinMarketCap PostProvider = "COIN_MARKET_CAP"
)

type Post struct {
	gorm.Model
	PostedAt       time.Time      `gorm:"not null"`
	Id             string         `gorm:"primarykey"`
	Content        string         `gorm:"type:varchar(4094)"`
	ParentId       string         `gorm:"type:varchar(32)"`
	SenderUsername string         `gorm:"type:varchar(64)"`
	Provider       PostProvider   `gorm:"type:varchar(32)"`
	Tags           pq.StringArray `gorm:"type:varchar(32)[]"`
	Sentiment      float32        `gorm:"type:varchar(32)"`
	LikeCount      uint32         `gorm:"default:0"`
}
