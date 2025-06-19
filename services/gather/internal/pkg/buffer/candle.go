package buffer

import (
	"errors"
	"fmt"
	"github.com/h-varmazyar/Gate/services/gather/configs"
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	"sync"
	"time"
)

type candleBuffer struct {
	lock         *sync.RWMutex
	data         map[uint]map[uint][]*models.Candle // first key is market id and second key is resolution id
	BufferLength int
}

var CandleBuffer *candleBuffer

func InitializeCandleBuffer(cfg configs.CandleBuffer) {
	if CandleBuffer == nil {
		CandleBuffer = &candleBuffer{
			lock:         new(sync.RWMutex),
			data:         make(map[uint]map[uint][]*models.Candle),
			BufferLength: cfg.CandleBufferLength,
		}

	}
}

func (buffer *candleBuffer) RemoveCandlePool(marketID, resolutionID uint) {
	buffer.lock.Lock()
	defer buffer.lock.Unlock()
	delete(buffer.data[marketID], resolutionID)
	if len(buffer.data[marketID]) == 0 {
		delete(buffer.data, marketID)
	}
}

func (buffer *candleBuffer) Push(candle *models.Candle) {
	buffer.lock.Lock()
	defer buffer.lock.Unlock()
	resolutions, ok := buffer.data[candle.MarketID]
	if !ok || resolutions == nil || len(resolutions) == 0 {
		buffer.data[candle.MarketID] = make(map[uint][]*models.Candle)
	}
	candles, ok := buffer.data[candle.MarketID][candle.ResolutionID]
	if !ok || candles == nil || len(candles) == 0 {
		candles = make([]*models.Candle, 0)
		for i := 0; i < buffer.BufferLength; i++ {
			candles = append(candles, new(models.Candle))
		}
	}

	if candles[buffer.BufferLength-1] != nil && candles[buffer.BufferLength-1].Time.Equal(candle.Time) {
		candles[buffer.BufferLength-1] = candle
	} else {
		candles = append(candles[1:], candle)
	}
	buffer.data[candle.MarketID][candle.ResolutionID] = candles
}

func (buffer *candleBuffer) Last(marketID, resolutionID uint) *models.Candle {
	candles := buffer.ReturnCandles(marketID, resolutionID, 1)
	if candles != nil {
		return candles[0]
	}
	return nil
}

func (buffer *candleBuffer) UpdateLast(marketID uint, lastPrice float64) error {
	buffer.lock.Lock()
	defer buffer.lock.Unlock()

	candleResolutions, ok := buffer.data[marketID]
	if !ok {
		return errors.New(fmt.Sprintf("candle resolutions not found for market %v", marketID))
	}

	for _, candles := range candleResolutions {
		length := len(candles)
		if length > 0 {
			last := candles[length-1]

			if last.Time.Add(last.Resolution.Duration).After(time.Now()) {
				if lastPrice < last.Low {
					last.Low = lastPrice
				}

				if lastPrice > last.High {
					last.High = lastPrice
				}

				last.Close = lastPrice
				//todo: publish candle
			}
		}
	}

	return nil
}

func (buffer *candleBuffer) Before(marketID, resolutionID uint, t time.Time, count int) []*models.Candle {
	buffer.lock.Lock()
	defer buffer.lock.Unlock()
	if candles, ok := buffer.data[marketID][resolutionID]; !ok || candles == nil {
		return nil
	} else if candles[0].Time.After(t) || candles[len(candles)-1].Time.Before(t) {
		return nil
	} else {
		cloned := make([]*models.Candle, count)
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

func (buffer *candleBuffer) ReturnCandles(marketID, resolutionID uint, n int) []*models.Candle {
	buffer.lock.Lock()
	defer buffer.lock.Unlock()
	if candles, ok := buffer.data[marketID][resolutionID]; !ok || candles == nil {
		return nil
	} else {
		cloned := make([]*models.Candle, n)
		j := buffer.BufferLength - n
		for i := 0; i < n; i++ {
			c := *candles[j]
			cloned[i] = &c
			j++
		}
		return cloned
	}
}
