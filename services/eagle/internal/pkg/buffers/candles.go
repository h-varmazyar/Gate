package buffers

import (
	"fmt"
	"github.com/mrNobody95/Gate/services/eagle/configs"
	"github.com/mrNobody95/Gate/services/eagle/internal/pkg/models"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 09.12.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type candleBuffer struct {
	candles map[string][]*models.Candle
}

var (
	Candles *candleBuffer
)

func init() {
	Candles = new(candleBuffer)
	Candles.candles = make(map[string][]*models.Candle)
}

func (buffer *candleBuffer) AddList(marketID, resolutionID string) {
	buffer.candles[key(marketID, resolutionID)] = make([]*models.Candle, configs.Variables.CandleBufferLength)
}

func (buffer *candleBuffer) RemoveList(marketID, resolutionID string) {
	delete(buffer.candles, key(marketID, resolutionID))
}

func (buffer *candleBuffer) List(marketID, resolutionID string) []*models.Candle {
	return buffer.candles[key(marketID, resolutionID)]
}

func (buffer *candleBuffer) Enqueue(candle *models.Candle) {
	list, ok := buffer.candles[key(candle.MarketID, candle.ResolutionID)]
	if !ok {
		list = make([]*models.Candle, configs.Variables.CandleBufferLength)
	}
	last := len(list) - 1
	if candle.Time.Equal(list[last].Time) {
		list[last] = candle
	} else {
		list = append(list[1:], candle)
	}
	buffer.candles[key(candle.MarketID, candle.ResolutionID)] = list
}

func key(marketID, resolutionID string) string {
	return fmt.Sprintf("%s > %s", marketID, resolutionID)
}
