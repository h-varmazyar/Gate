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
//	rateLimiters, ok := m.data[marketID]
//	if !ok || rateLimiters == nil || len(rateLimiters) == 0 {
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
//	rateLimiters, ok := m.data[marketID]
//	if !ok || rateLimiters == nil || len(rateLimiters) == 0 {
//		rateLimiters = make([]*entity.Candle, m.BufferLength)
//	}
//
//	if rateLimiters[m.BufferLength-1] != nil && rateLimiters[m.BufferLength-1].Time.Equal(candle.Time) {
//		rateLimiters[m.BufferLength-1] = candle
//	} else {
//		rateLimiters = append(rateLimiters[1:], candle)
//	}
//	m.data[marketID] = rateLimiters
//}
//
//func (m *markets) GetLastNCandles(marketID uuid.UUID, n int) []*entity.Candle {
//	m.lock.Lock()
//	defer m.lock.Unlock()
//	if rateLimiters, ok := m.data[marketID]; !ok || rateLimiters == nil {
//		return nil
//	} else {
//		cloned := make([]*entity.Candle, n)
//		j := m.BufferLength - n
//		for i := 0; i < n; i++ {
//			c := *rateLimiters[j]
//			cloned[i] = &c
//			j++
//		}
//		return cloned
//	}
//}
