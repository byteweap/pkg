package mapx

import "sync"

// Mapx 线程安全的map
// desc: 通过互斥锁实现有性能损耗，必要时使用
type Mapx[Key comparable, Value comparable] struct {
	mux sync.RWMutex
	m   map[Key]Value
}

func New[Key comparable, Value comparable]() *Mapx[Key, Value] {
	return &Mapx[Key, Value]{m: make(map[Key]Value)}
}

func (m *Mapx[Key, Value]) Get(key Key) Value {
	m.mux.RLock()
	defer m.mux.RUnlock()
	return m.m[key]
}

func (m *Mapx[Key, Value]) Set(key Key, value Value) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.m[key] = value
}

func (m *Mapx[Key, Value]) Remove(key Key) {
	m.mux.Lock()
	defer m.mux.Unlock()
	delete(m.m, key)
}

func (m *Mapx[Key, Value]) Len() int {
	m.mux.RLock()
	defer m.mux.RUnlock()
	return len(m.m)
}

// 遍历
// 取数据操作保证安全, 解锁后执行fn
func (m *Mapx[Key, Value]) Range(fn func(Key, Value)) {
	m.mux.RLock()
	data := m.m
	m.mux.RUnlock()
	for k, v := range data {
		fn(k, v)
	}
}
