package buffer

import (
	"fmt"
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
	"sync"
)

type candleBuffer struct {
	candles map[string][]*repository.Candle
}

var (
	candlesLock *sync.Mutex
	Candles     *candleBuffer
)

func init() {
	candlesLock = new(sync.Mutex)
	Candles = new(candleBuffer)
	Candles.candles = make(map[string][]*repository.Candle)
}

func (buffer *candleBuffer) AddList(marketID, resolutionID uint32) {
	buffer.candles[key(marketID, resolutionID)] = make([]*repository.Candle, configs.Variables.CandleBufferLength)
}

func (buffer *candleBuffer) RemoveList(marketID, resolutionID uint32) {
	delete(buffer.candles, key(marketID, resolutionID))
}

func (buffer *candleBuffer) Last(marketID, resolutionID uint32) *repository.Candle {
	list := buffer.candles[key(marketID, resolutionID)]
	return list[len(list)-1]
}

func (buffer *candleBuffer) Enqueue(candle *repository.Candle) {
	list, ok := buffer.candles[key(candle.MarketID, candle.ResolutionID)]
	if !ok || list == nil || len(list) == 0 {
		list = make([]*repository.Candle, configs.Variables.CandleBufferLength)
	}
	last := len(list) - 1
	if list[last] != nil && candle.Time.Equal(list[last].Time) {
		list[last] = candle
	} else {
		list = append(list[1:], candle)
	}
	candlesLock.Lock()
	buffer.candles[key(candle.MarketID, candle.ResolutionID)] = list
	candlesLock.Unlock()
}

func key(marketID, resolutionID uint32) string {
	return fmt.Sprintf("%d > %d", marketID, resolutionID)
}
