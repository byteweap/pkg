package mapx

import "sync"

// 安全的泛型map
// 使用互斥锁实现
type safeMapx[Key comparable, Value any] struct {
	mux sync.RWMutex
	m   map[Key]Value
}

// 初始化
func newSafeMapx[Key comparable, Value any]() *safeMapx[Key, Value] {
	return &safeMapx[Key, Value]{m: make(map[Key]Value)}
}

// 取值
func (m *safeMapx[Key, Value]) Get(key Key) Value {
	m.mux.RLock()
	defer m.mux.RUnlock()
	return m.m[key]
}

// 赋值
func (m *safeMapx[Key, Value]) Set(key Key, value Value) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.m[key] = value
}

// 删除
func (m *safeMapx[Key, Value]) Delete(key Key) {
	m.mux.Lock()
	defer m.mux.Unlock()
	delete(m.m, key)
}

// 长度
func (m *safeMapx[Key, Value]) Len() int {
	m.mux.RLock()
	defer m.mux.RUnlock()
	return len(m.m)
}

// 遍历
// 取数据操作保证安全, 解锁后执行fn
func (m *safeMapx[Key, Value]) Range(fn func(Key, Value)) {
	m.mux.RLock()
	data := m.m
	m.mux.RUnlock()
	for k, v := range data {
		fn(k, v)
	}
}

// 获取所有key
func (m *safeMapx[Key, Value]) Keys() []Key {
	m.mux.RLock()
	defer m.mux.RUnlock()

	keys := make([]Key, 0, len(m.m))
	for k := range m.m {
		keys = append(keys, k)
	}
	return keys
}

// 是否存在
func (m *safeMapx[Key, Value]) IsExist(key Key) bool {
	m.mux.RLock()
	defer m.mux.RUnlock()

	_, ok := m.m[key]
	return ok
}
