package buffer

import (
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	"sync"
	"time"
)

type candleBuffer struct {
	lock         *sync.RWMutex
	data         map[string]map[string][]*entity.Candle // first key is market id and second key is resolution id
	BufferLength int
}

var CandleBuffer *candleBuffer

func InitializeCandleBuffer(configs *Configs) {
	if CandleBuffer == nil {
		CandleBuffer = &candleBuffer{
			lock:         new(sync.RWMutex),
			data:         make(map[string]map[string][]*entity.Candle),
			BufferLength: configs.CandleBufferLength,
		}

	}
}

func (buffer *candleBuffer) RemoveCandlePool(marketID, resolutionID string) {
	buffer.lock.Lock()
	defer buffer.lock.Unlock()
	delete(buffer.data[marketID], resolutionID)
	if len(buffer.data[marketID]) == 0 {
		delete(buffer.data, marketID)
	}
}

func (buffer *candleBuffer) Push(candle *entity.Candle) {
	buffer.lock.Lock()
	defer buffer.lock.Unlock()
	resolutions, ok := buffer.data[candle.MarketID.String()]
	if !ok || resolutions == nil || len(resolutions) == 0 {
		buffer.data[candle.MarketID.String()] = make(map[string][]*entity.Candle)
	}
	candles, ok := buffer.data[candle.MarketID.String()][candle.ResolutionID.String()]
	if !ok || candles == nil || len(candles) == 0 {
		candles = make([]*entity.Candle, 0)
		for i := 0; i < buffer.BufferLength; i++ {
			candles = append(candles, new(entity.Candle))
		}
	}

	if candles[buffer.BufferLength-1] != nil && candles[buffer.BufferLength-1].Time.Equal(candle.Time) {
		candles[buffer.BufferLength-1] = candle
	} else {
		candles = append(candles[1:], candle)
	}
	buffer.data[candle.MarketID.String()][candle.ResolutionID.String()] = candles
}

func (buffer *candleBuffer) Last(marketID, resolutionID string) *entity.Candle {
	candles := buffer.ReturnCandles(marketID, resolutionID, 1)
	if candles != nil {
		return candles[0]
	}
	return nil
}

func (buffer *candleBuffer) Before(marketID, resolutionID string, t time.Time, count int) []*entity.Candle {
	buffer.lock.Lock()
	defer buffer.lock.Unlock()
	if candles, ok := buffer.data[marketID][resolutionID]; !ok || candles == nil {
		return nil
	} else if candles[0].Time.After(t) || candles[len(candles)-1].Time.Before(t) {
		return nil
	} else {
		cloned := make([]*entity.Candle, count)
		founded := -1
		for i := len(candles) - 1; i >= 0; i-- {
			if candles[i].Time.After(t) {
				continue
			}
			founded = i
		}

		for j := count - 1; j >= 0; j-- {
			c := *candles[founded]
			cloned[j] = &c
			founded--
			if founded == -1 {
				break
			}
		}
		return cloned
	}
}

func (buffer *candleBuffer) ReturnCandles(marketID, resolutionID string, n int) []*entity.Candle {
	buffer.lock.Lock()
	defer buffer.lock.Unlock()
	if candles, ok := buffer.data[marketID][resolutionID]; !ok || candles == nil {
		return nil
	} else {
		cloned := make([]*entity.Candle, n)
		j := buffer.BufferLength - n
		for i := 0; i < n; i++ {
			c := *candles[j]
			cloned[i] = &c
			j++
		}
		return cloned
	}
}
