package buffer

import (
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/services/chipmunk/configs"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/repository"
	"sync"
)

type markets struct {
	lock         *sync.RWMutex
	data         map[uuid.UUID][]*repository.Candle
	BufferLength int
}

var Markets *markets

func NewMarketInstance() {
	Markets = &markets{
		lock:         new(sync.RWMutex),
		data:         make(map[uuid.UUID][]*repository.Candle),
		BufferLength: configs.Variables.CandleBufferLength,
	}
}

func (m *markets) AddList(marketID uuid.UUID) {
	m.lock.Lock()
	defer m.lock.Unlock()
	candles, ok := m.data[marketID]
	if !ok || candles == nil || len(candles) == 0 {
		emptyCandles := make([]*repository.Candle, 0)
		for i := 0; i < m.BufferLength; i++ {
			emptyCandles = append(emptyCandles, new(repository.Candle))
		}
		m.data[marketID] = emptyCandles
	}
}

func (m *markets) RemoveList(marketID uuid.UUID) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.data, marketID)
}

func (m *markets) Push(marketID uuid.UUID, candle *repository.Candle) {
	m.lock.Lock()
	defer m.lock.Unlock()
	candles, ok := m.data[marketID]
	if !ok || candles == nil || len(candles) == 0 {
		candles = make([]*repository.Candle, m.BufferLength)
	}

	if candles[m.BufferLength-1] != nil && candles[m.BufferLength-1].Time.Equal(candle.Time) {
		candles[m.BufferLength-1] = candle
	} else {
		candles = append(candles[1:], candle)
	}
	m.data[marketID] = candles
}

func (m *markets) GetLastNCandles(marketID uuid.UUID, n int) []*repository.Candle {
	m.lock.Lock()
	defer m.lock.Unlock()
	if candles, ok := m.data[marketID]; !ok || candles == nil {
		return nil
	} else {
		cloned := make([]*repository.Candle, n)
		j := m.BufferLength - n
		for i := 0; i < n; i++ {
			c := *candles[j]
			cloned[i] = &c
			j++
		}
		return cloned
	}
}
