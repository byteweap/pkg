package mapx

import (
	"github.com/bytedance/sonic"
	"github.com/violin8/pkg/convert"
)

// KVMap 非线程安全的范型map
//
// 对标准库的扩展, key: 范型, value: 范型
// 当map的value是固定某一类型时
// 基本操作同标准库map
type KVMap[Key comparable, Value comparable] map[Key]Value

func (m KVMap[Key, Value]) ToJson() []byte {
	bytes, _ := sonic.Marshal(m)
	return bytes
}

func (m KVMap[Key, Value]) ToJsonString() string {
	return convert.Bytes2String(m.ToJson())
}
