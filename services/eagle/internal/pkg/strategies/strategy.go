package strategies

import (
	"context"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
)

type Strategy interface {
	CheckForSignals(ctx context.Context, market *chipmunkApi.Market)
}
