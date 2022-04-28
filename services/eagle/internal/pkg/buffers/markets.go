package buffers

import (
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	"sync"
)

type markets struct {
	lock         *sync.Mutex
	data         map[uuid.UUID][]*api.Candle
	BufferLength int
}

var Markets *markets

func init() {
	Markets = &markets{
		lock:         new(sync.Mutex),
		data:         make(map[uuid.UUID][]*api.Candle),
		BufferLength: configs.Variables.CandleBufferLength,
	}
}

func (m *markets) AddList(marketID uuid.UUID) {
	candles, ok := m.data[marketID]
	if !ok || candles == nil || len(candles) == 0 {
		m.lock.Lock()
		m.data[marketID] = make([]*api.Candle, m.BufferLength)
	}
	m.lock.Unlock()
}

func (m *markets) RemoveList(marketID uuid.UUID) {
	delete(m.data, marketID)
}

func (m *markets) Push(marketID uuid.UUID, candle *api.Candle) {
	candles, ok := m.data[marketID]
	if !ok || candles == nil || len(candles) == 0 {
		candles = make([]*api.Candle, m.BufferLength)
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	if candles[m.BufferLength-1] != nil && candles[m.BufferLength-1].Time == candle.Time {
		candles[m.BufferLength-1] = candle
	} else if candles[m.BufferLength-1].Time < candle.Time {
		candles = append(candles[1:], candle)
	} else {
		return
	}
	m.data[marketID] = candles
}

func (m *markets) GetLastNCandles(marketID uuid.UUID, n int) []*api.Candle {
	if candles, ok := m.data[marketID]; !ok || candles == nil {
		return nil
	} else {
		return m.data[marketID][len(m.data[marketID])-n:]
	}
}
