package strategies

import (
	"context"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
)

type Strategy interface {
	CheckForSignals(ctx context.Context, market *brokerageApi.Market)
}
