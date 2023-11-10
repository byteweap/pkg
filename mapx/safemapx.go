package mapx

import "sync"

// generic type map which is thread safe.
type safeMapx[Key comparable, Value any] struct {
	mux sync.RWMutex
	m   map[Key]Value
}

// newSafeMapx creates a new instance of safeMapx.
func newSafeMapx[Key comparable, Value any]() *safeMapx[Key, Value] {
	return &safeMapx[Key, Value]{m: make(map[Key]Value)}
}

// Get get value by key.
func (m *safeMapx[Key, Value]) Get(key Key) Value {
	m.mux.RLock()
	defer m.mux.RUnlock()
	return m.m[key]
}

// Set assigns a value to the key.
func (m *safeMapx[Key, Value]) Set(key Key, value Value) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.m[key] = value
}

// Delete delete the value by key.
func (m *safeMapx[Key, Value]) Delete(key Key) {
	m.mux.Lock()
	defer m.mux.Unlock()
	delete(m.m, key)
}

// Len returns the length of the safeMapx.
func (m *safeMapx[Key, Value]) Len() int {
	m.mux.RLock()
	defer m.mux.RUnlock()
	return len(m.m)
}

// Range iterates over the map and calls the given function for each key-value pair.
func (m *safeMapx[Key, Value]) Range(fn func(Key, Value)) {
	m.mux.RLock()
	data := m.m
	m.mux.RUnlock()
	for k, v := range data {
		fn(k, v)
	}
}

// Keys returns a slice of keys in the safeMapx[Key, Value].
func (m *safeMapx[Key, Value]) Keys() []Key {
	m.mux.RLock()
	defer m.mux.RUnlock()

	keys := make([]Key, 0, len(m.m))
	for k := range m.m {
		keys = append(keys, k)
	}
	return keys
}

// IsExist checks if the key exists in the safeMapx.
func (m *safeMapx[Key, Value]) IsExist(key Key) bool {
	m.mux.RLock()
	defer m.mux.RUnlock()

	_, ok := m.m[key]
	return ok
}
