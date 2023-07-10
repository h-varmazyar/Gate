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
