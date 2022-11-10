package rateLimiter

//type IPManager struct {
//	IPPool          map[uuid.UUID]*IP
//	cancelFunctions map[uuid.UUID]context.CancelFunc
//	lock            *sync.Mutex
//}
//
//func NewIPManager() *IPManager {
//	manager := &IPManager{
//		lock:            new(sync.Mutex),
//		IPPool:          make(map[uuid.UUID]*IP),
//		cancelFunctions: make(map[uuid.UUID]context.CancelFunc),
//	}
//	return manager
//}
//
//func (m *IPManager) DeleteIP(ctx context.Context, id uuid.UUID) error {
//	if id == uuid.Nil {
//		return errors.New(ctx, codes.FailedPrecondition).AddDetailF("invalid IP id")
//	}
//	cancelFunc, ok := m.cancelFunctions[id]
//	if ok {
//		m.lock.Lock()
//		cancelFunc()
//		m.lock.Unlock()
//	}
//
//	_, ok = m.IPPool[id]
//	if !ok {
//		return errors.New(ctx, codes.NotFound)
//	}
//
//	m.lock.Lock()
//	defer m.lock.Unlock()
//
//	delete(m.IPPool, id)
//	return nil
//}
