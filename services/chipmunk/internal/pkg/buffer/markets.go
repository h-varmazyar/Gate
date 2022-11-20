package buffer

//type markets struct {
//	lock         *sync.RWMutex
//	data         map[uuid.UUID][]*entity.Candle
//	BufferLength int
//}
//
//var Markets *markets
//
//func NewMarketInstance() {
//	Markets = &markets{
//		lock:         new(sync.RWMutex),
//		data:         make(map[uuid.UUID][]*entity.Candle),
//		BufferLength: configs.Variables.CandleBufferLength,
//	}
//}
//
//func (m *markets) AddList(marketID uuid.UUID) {
//	m.lock.Lock()
//	defer m.lock.Unlock()
//	candles, ok := m.data[marketID]
//	if !ok || candles == nil || len(candles) == 0 {
//		emptyCandles := make([]*entity.Candle, 0)
//		for i := 0; i < m.BufferLength; i++ {
//			emptyCandles = append(emptyCandles, new(entity.Candle))
//		}
//		m.data[marketID] = emptyCandles
//	}
//}
//
//func (m *markets) RemoveList(marketID uuid.UUID) {
//	m.lock.Lock()
//	defer m.lock.Unlock()
//	delete(m.data, marketID)
//}
//
//func (m *markets) Push(marketID uuid.UUID, candle *entity.Candle) {
//	m.lock.Lock()
//	defer m.lock.Unlock()
//	candles, ok := m.data[marketID]
//	if !ok || candles == nil || len(candles) == 0 {
//		candles = make([]*entity.Candle, m.BufferLength)
//	}
//
//	if candles[m.BufferLength-1] != nil && candles[m.BufferLength-1].Time.Equal(candle.Time) {
//		candles[m.BufferLength-1] = candle
//	} else {
//		candles = append(candles[1:], candle)
//	}
//	m.data[marketID] = candles
//}
//
//func (m *markets) GetLastNCandles(marketID uuid.UUID, n int) []*entity.Candle {
//	m.lock.Lock()
//	defer m.lock.Unlock()
//	if candles, ok := m.data[marketID]; !ok || candles == nil {
//		return nil
//	} else {
//		cloned := make([]*entity.Candle, n)
//		j := m.BufferLength - n
//		for i := 0; i < n; i++ {
//			c := *candles[j]
//			cloned[i] = &c
//			j++
//		}
//		return cloned
//	}
//}
