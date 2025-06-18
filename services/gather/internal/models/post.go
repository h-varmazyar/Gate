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
	Content        string         `gorm:"type:varchar(4094)"`
	ParentId       uint           `gorm:"not null"`
	SenderUsername string         `gorm:"type:varchar(64)"`
	Provider       PostProvider   `gorm:"type:varchar(32)"`
	Tags           pq.StringArray `gorm:"type:varchar(32)[]"`
	Sentiment      string         `gorm:"type:varchar(32)"`
	LikeCount      *int           `gorm:"default:null"`
	RetwitCount    *int           `gorm:"default:null"`
	CommentCount   *int           `gorm:"default:null"`
	QuoteCount     *int           `gorm:"default:null"`
	Type           string         `gorm:"type:varchar(32)"`
	ParentID       *uint          `gorm:"default:null"`
}
