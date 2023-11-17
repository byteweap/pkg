package mapx

// Mapx
// generic type map which is thread safe or not
type Mapx[Key comparable, Value any] interface {
	Get(Key) (Value, bool)     // Get value by key.
	Set(Key, Value)            // Set value to key.
	Delete(Key)                // Delete by key.
	Len() int                  // Get length.
	Range(fn func(Key, Value)) // Range.
	Keys() []Key               // Get keys.
}

// New[Key comparable, Value any](threadSafe bool) returns a new Mapx[Key, Value] based on the value of threadSafe.
//
// threadSafe: a boolean indicating whether the mapx should be thread safe or not.
// Returns: an instance of Mapx[Key, Value], either a thread safe or non-thread safe depending on the value of threadSafe.
func New[Key comparable, Value any](threadSafe bool) Mapx[Key, Value] {

	if threadSafe {
		return newSafeMapx[Key, Value]() // 线程安全mapx
	} else {
		return newUnSafeMapx[Key, Value]() // 非线程安全mapx
	}
}
