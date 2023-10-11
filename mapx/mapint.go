package mapx

import (
	"sync"
)

type MapInt struct {
	sync.RWMutex
	m map[int]int64
}

func NewMapInt() *MapInt {
	return &MapInt{
		m: make(map[int]int64),
	}
}

func (m *MapInt) init() {
	if m.m == nil {
		m.m = make(map[int]int64)
	}
}

func (m *MapInt) UnsafeGet(key int) int64 {
	if m.m == nil {
		return 0
	} else {
		return m.m[key]
	}
}

func (m *MapInt) Get(key int) int64 {
	m.RLock()
	defer m.RUnlock()
	return m.UnsafeGet(key)
}

func (m *MapInt) UnsafeSet(key int, value int64) {
	m.init()
	m.m[key] = value
}

func (m *MapInt) Set(key int, value int64) {
	m.Lock()
	defer m.Unlock()
	m.UnsafeSet(key, value)
}

func (m *MapInt) UnsafeDel(key int) {
	m.init()
	delete(m.m, key)
}

func (m *MapInt) Del(key int) {
	m.Lock()
	defer m.Unlock()
	m.UnsafeDel(key)
}

func (m *MapInt) UnsafeLen() int {
	if m.m == nil {
		return 0
	} else {
		return len(m.m)
	}
}

func (m *MapInt) Len() int {
	m.RLock()
	defer m.RUnlock()
	return m.UnsafeLen()
}

func (m *MapInt) UnsafeRange(f func(k int, value int64)) {
	if m.m == nil {
		return
	}
	for k, v := range m.m {
		f(k, v)
	}
}

func (m *MapInt) RLockRange(f func(k int, value int64)) {
	m.RLock()
	defer m.RUnlock()
	m.UnsafeRange(f)
}

func (m *MapInt) LockRange(f func(k int, value int64)) {
	m.Lock()
	defer m.Unlock()
	m.UnsafeRange(f)
}

func (m *MapInt) IsExist(key int) bool {
	m.RLock()
	defer m.RUnlock()
	_, ok := m.m[key]
	return ok
}

// 修改数据
// desc: 如果已有数据,则修改;没有数据,则Set
func (m *MapInt) Update(key int, incr int64) {
	m.Lock()
	defer m.Unlock()
	m.UnsafeSet(key, m.UnsafeGet(key)+incr)
}
