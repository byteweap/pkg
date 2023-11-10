package mapx

// generic type map which is not thread safe.
type unsafeMapx[Key comparable, Value any] map[Key]Value

// newSafeMapx creates a new instance of unsafeMapx.
func newUnSafeMapx[Key comparable, Value any]() unsafeMapx[Key, Value] {
	return make(unsafeMapx[Key, Value])
}

// Get returns the value and a boolean indicating if the key exists in the map.
// Returns:
//    Value: The value associated with the key.
//    bool: true if the key exists in the map, false otherwise.
func (m unsafeMapx[Key, Value]) Get(key Key) (Value, bool) {
	val, ok := m[key]
	return val, ok
}

// Set assigns a value to the key.
func (m unsafeMapx[Key, Value]) Set(key Key, value Value) {
	m[key] = value
}

// Delete delete the value by key.
func (m unsafeMapx[Key, Value]) Delete(key Key) {
	delete(m, key)
}

// Len returns the length of the safeMapx.
func (m unsafeMapx[Key, Value]) Len() int {
	return len(m)
}

// Range iterates over the map and calls the given function for each key-value pair.
func (m unsafeMapx[Key, Value]) Range(fn func(Key, Value)) {
	for k, v := range m {
		fn(k, v)
	}
}

// Keys returns a slice of keys in the safeMapx[Key, Value].
func (m unsafeMapx[Key, Value]) Keys() []Key {
	keys := make([]Key, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
