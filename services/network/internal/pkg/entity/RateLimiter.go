package entity

import (
	"github.com/h-varmazyar/Gate/pkg/gormext"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	"time"
)

type RateLimiter struct {
	gormext.UniversalModel
	RequestCountLimit int64
	TimeLimit         time.Duration
	Type              networkAPI.RateLimiterType
}
