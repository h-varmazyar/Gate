package buffer

import (
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	"sync"
)

type CandleBuffer struct {
	lock         *sync.RWMutex
	data         map[string]map[string][]*entity.Candle // first key is market id and second key is resolution id
	BufferLength int
}

var bufferInstance *CandleBuffer

func NewMarketInstance(configs *Configs) *CandleBuffer {
	if bufferInstance == nil {
		bufferInstance = &CandleBuffer{
			lock:         new(sync.RWMutex),
			data:         make(map[string]map[string][]*entity.Candle),
			BufferLength: configs.CandleBufferLength,
		}
	}
	return bufferInstance
}

//func (buffer *CandleBuffer) CreateCandleQueue(marketID, resolutionID uuid.UUID) {
//	buffer.lock.Lock()
//	defer buffer.lock.Unlock()
//	resolutions, ok := buffer.data[marketID]
//	if !ok || resolutions == nil || len(resolutions) == 0 {
//		buffer.data[marketID] = make(map[uuid.UUID][]*entity.Candle)
//	}
//	candles, ok := buffer.data[marketID][resolutionID]
//	if !ok || candles == nil || len(candles) == 0 {
//		emptyCandles := make([]*entity.Candle, 0)
//		for i := 0; i < buffer.BufferLength; i++ {
//			emptyCandles = append(emptyCandles, new(entity.Candle))
//		}
//		buffer.data[marketID][resolutionID] = emptyCandles
//	}
//}

func (buffer *CandleBuffer) RemoveCandlePool(marketID, resolutionID string) {
	buffer.lock.Lock()
	defer buffer.lock.Unlock()
	delete(buffer.data[marketID], resolutionID)
	if len(buffer.data[marketID]) == 0 {
		delete(buffer.data, marketID)
	}
}

func (buffer *CandleBuffer) Push(candle *entity.Candle) {
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

func (buffer *CandleBuffer) ReturnCandles(marketID, resolutionID string, n int) []*entity.Candle {
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
