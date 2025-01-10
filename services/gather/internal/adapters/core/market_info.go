package core

import (
	api "github.com/h-varmazyar/Gate/api/proto"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	"golang.org/x/net/context"
)

func (c Core) MarketInfo(ctx context.Context, marketKey string) (*coreApi.MarketInfo, error) {
	marketInfoReq := &coreApi.MarketInfoReq{
		Market: &chipmunkApi.Market{
			Platform: api.Platform_Coinex,
			Source: &chipmunkApi.Asset{
				Name:   marketKey,
				Symbol: marketKey,
			},
		},
	}
	marketInfo, err := c.functionService.GetMarketInfo(ctx, marketInfoReq)
	if err != nil {
		return nil, err
	}
	return marketInfo, nil
}
