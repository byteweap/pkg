package mapx

// Mapx
// New()入参threadSafe指定是否为线程安全
type Mapx[Key comparable, Value any] interface {
	Get(key Key) Value                 // 取值
	Set(key Key, value Value)          // 赋值
	Delete(key Key)                    // 删除
	Len() int                          // 长度
	Range(fn func(key Key, val Value)) // 遍历
	Keys() []Key                       // 获取所有key
	IsExist(key Key) bool              // 是否存在
}

// New
// 泛型key和value
// threadSafe: 是否线程安全
func New[Key comparable, Value any](threadSafe bool) Mapx[Key, Value] {

	if threadSafe {
		return newSafeMapx[Key, Value]() // 线程安全mapx
	} else {
		return newUnSafeMapx[Key, Value]() // 非线程安全mapx
	}
}
