package brokerages

import (
	"context"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
)

type Responses interface {
	AsyncOHLC(ctx context.Context, response *networkAPI.Response, metadata *Metadata)
	AllMarkerStatistics(ctx context.Context, response *networkAPI.Response) (*coreApi.AllMarketStatisticsResp, error)
	GetMarketInfo(ctx context.Context, response *networkAPI.Response) (*coreApi.MarketInfo, error)
}
