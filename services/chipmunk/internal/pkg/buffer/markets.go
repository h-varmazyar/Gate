package buffer

import (
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
	"sync"
)

type markets struct {
	lock         *sync.Mutex
	data         map[uint32][]*repository.Candle
	BufferLength int
}

var Markets *markets

func init() {
	Markets = &markets{
		lock:         new(sync.Mutex),
		data:         make(map[uint32][]*repository.Candle),
		BufferLength: configs.Variables.CandleBufferLength,
	}
}

func (m *markets) Push(marketID uint32, candle *repository.Candle) {
	candles, ok := m.data[marketID]
	if !ok {
		candles = make([]*repository.Candle, m.BufferLength)
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	if candles[m.BufferLength-1] != nil && candles[m.BufferLength-1].Time.Equal(candle.Time) {
		candles[m.BufferLength-1] = candle
	} else {
		candles = append(candles[1:], candle)
	}
	m.data[marketID] = candles
}

func (m *markets) GetLastNCandles(marketID uint32, n int) []*repository.Candle {
	if candles, ok := m.data[marketID]; !ok || candles == nil {
		return nil
	} else {
		cloned := make([]*repository.Candle, n)
		for i := n; i > 0; i-- {
			c := *candles[m.BufferLength-i]
			cloned[i-1] = &c
		}
		return cloned
	}
}
