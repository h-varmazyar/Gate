package core

import (
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	"golang.org/x/net/context"
)

func (c Core) MarketList(ctx context.Context) ([]*proto.Market, error) {
	listReq := &coreApi.MarketListReq{Platform: api.Platform_Coinex}
	markets, err := c.functionService.MarketList(ctx, listReq)
	if err != nil {
		return nil, err
	}

	return markets.GetElements(), nil
}
