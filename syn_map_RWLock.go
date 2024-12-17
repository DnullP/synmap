package syncmap

import "sync"

type SyncMapRWLock[T comparable, U any] struct {
	lock sync.RWMutex
	data map[T]U
}

func NewSyncMapRWLock[T comparable, U any]() *SyncMapRWLock[T, U] {
	return &SyncMapRWLock[T, U]{
		data: make(map[T]U),
	}
}

func (m *SyncMapRWLock[T, U]) Store(key T, value U) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.data[key] = value
}

func (m *SyncMapRWLock[T, U]) Load(key T) (value U, ok bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	value, ok = m.data[key]
	return value, ok
}

func (m *SyncMapRWLock[T, U]) Reset() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.data = make(map[T]U)
}
