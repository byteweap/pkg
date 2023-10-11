package mapx

import "sync"

// Mapx 线程安全的map
type Mapx[Key comparable, Value comparable] struct {
	mux sync.RWMutex
	m   map[Key]Value
}

func New[Key comparable, Value comparable]() *Mapx[Key, Value] {
	return &Mapx[Key, Value]{m: make(map[Key]Value)}
}

func (m *Mapx[Key, Value]) UnsafeGet(key Key) Value {
	return m.m[key]
}

func (m *Mapx[Key, Value]) Get(key Key) Value {
	m.mux.RLock()
	defer m.mux.RUnlock()
	return m.UnsafeGet(key)
}

func (m *Mapx[Key, Value]) UnsafeSet(key Key, value Value) {
	m.m[key] = value
}

func (m *Mapx[Key, Value]) Set(key Key, value Value) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.UnsafeSet(key, value)
}

func (m *Mapx[Key, Value]) UnsafeDel(key Key) {
	delete(m.m, key)
}

func (m *Mapx[Key, Value]) Del(key Key) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.UnsafeDel(key)
}

func (m *Mapx[Key, Value]) UnsafeLen() int {
	if m.m == nil {
		return 0
	} else {
		return len(m.m)
	}
}

func (m *Mapx[Key, Value]) Len() int {
	m.mux.RLock()
	defer m.mux.RUnlock()
	return m.UnsafeLen()
}

func (m *Mapx[Key, Value]) UnsafeRange(f func(Key, Value)) {
	if m.m == nil {
		return
	}
	for k, v := range m.m {
		f(k, v)
	}
}

func (m *Mapx[Key, Value]) RLockRange(f func(Key, Value)) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	m.UnsafeRange(f)
}

func (m *Mapx[Key, Value]) LockRange(f func(Key, Value)) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.UnsafeRange(f)
}
