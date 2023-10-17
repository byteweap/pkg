package mapx

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/violin8/pkg/cast"
	"github.com/violin8/pkg/convert"
)

// KMap 非线程安全的范型(key)map
//
//		I. 对标准库的扩展, key: 范型, value: 固定为any类型
//	 II. 当map的value不是固定某一类型时
//		III. 基本操作同标准库map, 以下对特殊操作进行了扩展
type KMap[Key comparable] map[Key]any

func (m KMap[Key]) GetInt(key Key) int {
	val := m[key]
	return cast.ToInt(val)
}

func (m KMap[Key]) GetInt64(key Key) int64 {
	val := m[key]
	return cast.ToInt64(val)
}

func (m KMap[Key]) GetString(key Key) string {
	val := m[key]
	return cast.ToString(val)
}

func (m KMap[Key]) GetFloat64(key Key) float64 {
	val := m[key]
	return cast.ToFloat64(val)
}

func (m KMap[Key]) GetTime(key Key) time.Time {
	val := m[key]
	return cast.ToTime(val)
}

func (m KMap[Key]) GetBool(key Key) bool {
	val := m[key]
	return cast.ToBool(val)
}

func (m KMap[Key]) GetIntSlice(key Key) []int {
	val := m[key]
	return cast.ToIntSlice(val)
}

// ToJson return json bytes
func (m KMap[Key]) ToJson() []byte {
	bytes, _ := sonic.Marshal(m)
	return bytes
}

// ToJsonString return json string
func (m KMap[Key]) ToJsonString() string {
	return convert.Bytes2String(m.ToJson())
}
