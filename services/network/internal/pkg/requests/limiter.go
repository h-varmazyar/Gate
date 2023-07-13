package requests

import (
	"github.com/google/uuid"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	"time"
)

type Limiter struct {
	ID                uuid.UUID
	RequestCountLimit int64
	TimeLimit         time.Duration
	Type              networkAPI.RateLimiterType
	RequestChannel    chan *BucketRequest
}

func (l *Limiter) IntervalDuration() time.Duration {
	switch l.Type {
	case networkAPI.RateLimiter_Spread:
		return time.Duration(int64(l.TimeLimit) / l.RequestCountLimit)
	case networkAPI.RateLimiter_Immediate:
		return 0
	}
	return 0
}
