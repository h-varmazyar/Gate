package core

import (
	api "github.com/h-varmazyar/Gate/api/proto"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	"github.com/h-varmazyar/Gate/services/gather/internal/workers/lastCandle"
	"golang.org/x/net/context"
	"time"
)

func (c Core) OHLC(ctx context.Context, param lastCandle.OHLCParam) (*chipmunkApi.Candles, error) {
	ohlcReq := &coreApi.OHLCReq{
		Platform: api.Platform_Coinex,
		Item: &coreApi.OHLCItem{
			Resolution: &chipmunkApi.Resolution{
				Label:    param.Resolution.Label,
				Value:    param.Resolution.Value,
				Duration: int64(param.Resolution.Duration),
			},
			Market:    &chipmunkApi.Market{Name: param.MarketKey},
			From:      param.From.Unix(),
			To:        param.To.Unix(),
			Timeout:   int64(param.Timeout),
			IssueTime: time.Now().Unix(),
		},
	}
	candles, err := c.functionService.OHLC(ctx, ohlcReq)
	if err != nil {
		return nil, err
	}

	return candles, nil
}
