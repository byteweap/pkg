package values

import (
	"strconv"
	"strings"

	"github.com/bytedance/sonic"
)

// Value 值类型
// type Value interface{}

// Value 任何类型
type Value interface{}

// // NewValueFromJSON 从 JSON 解析值
// func NewValueFromJSON(data []byte) (Value, error) {
// 	var value Value
// 	err := json.Unmarshal(data, &value)
// 	return value, err
// }

// NewValueMapArray 返回 ValueＭap 数组
func NewValueMapArray(data []byte) ([]ValueMap, error) {
	var values []interface{}
	err := sonic.Unmarshal(data, &values)
	if err != nil {
		return nil, err
	}
	var vms []ValueMap
	for _, item := range values {
		vms = append(vms, item.(map[string]interface{}))
	}
	return vms, nil
}

// ValueMap 字典类型
type ValueMap map[string]interface{}

// NewValuesFromJSON 创建新数据
func NewValuesFromJSON(data []byte) (ValueMap, error) {
	var vm ValueMap
	err := sonic.Unmarshal(data, &vm)
	return vm, err
}

// GetValueMap 返回 ValueMap
func (vm ValueMap) GetValueMap(name string) ValueMap {
	value := vm[name]
	// fmt.Println(reflect.TypeOf(value))
	switch v := value.(type) {
	case map[string]interface{}:
		return value.(map[string]interface{})
	case ValueMap:
		return value.(ValueMap)
	case string:
		vm, _ := NewValuesFromJSON([]byte(v))
		return vm
	}
	return nil
}

// GetString 获取字符串数据
func (vm ValueMap) GetString(name string) string {
	value := vm[name]
	if value == nil {
		return ""
	}
	switch value.(type) {
	case int:
		return strconv.Itoa(value.(int))
	case string:
		return strings.TrimSpace(value.(string))
	case []interface{}:
		data, _ := sonic.Marshal(value)
		return string(data)
	case interface{}:
		data, _ := sonic.Marshal(value)
		return string(data)
	}
	return ""
}

// GetInt 获取 int 类型数据
func (vm ValueMap) GetInt(name string) int {
	value := vm[name]
	if value == nil {
		return 0
	}
	// fmt.Println(reflect.TypeOf(value))
	switch value.(type) {
	case int:
		return value.(int)
	case float64:
		return int(value.(float64))
	case string:
		ret, err := strconv.Atoi(value.(string))
		if err != nil {
			return 0
		}
		return ret
	}
	return 0
}

// GetInt64 获取 int 类型数据
func (vm ValueMap) GetInt64(name string) int64 {
	value := vm[name]
	if value == nil {
		return 0
	}
	// fmt.Println(reflect.TypeOf(value))
	switch value.(type) {
	case int:
		return int64(value.(int))
	case int64:
		return value.(int64)
	case float64:
		return int64(value.(float64))
	case string:
		ret, err := strconv.Atoi(value.(string))
		if err != nil {
			return 0
		}
		return int64(ret)
	}
	return 0
}

func (vm ValueMap) GetIntArray(key string) []int {
	value, ok := vm[key]
	if !ok {
		return []int{}
	}
	switch v := value.(type) {
	case []int:
		return v
	case []int32:
		ret := make([]int, len(v))
		for i := 0; i < len(v); i++ {
			ret[i] = int(v[i])
		}
		return ret
	case []int64:
		ret := make([]int, len(v))
		for i := 0; i < len(v); i++ {
			ret[i] = int(v[i])
		}
		return ret
	case []interface{}:
		bs, err := sonic.Marshal(v)
		if err != nil {
			return []int{}
		}
		arr := make([]int, 0)
		err = sonic.Unmarshal(bs, &arr)
		if err != nil {
			return []int{}
		}
		return arr
	case string:
		ret := make([]int, 0)
		_ = sonic.Unmarshal([]byte(v), &ret)
		return ret
	default:
		return []int{}
	}

}

// GetFloat64 获取 float64 数据
func (vm ValueMap) GetFloat64(name string) float64 {
	value := vm[name]
	if value == nil {
		return 0
	}
	switch value.(type) {
	case float64:
		return value.(float64)
	case string:
		ret, err := strconv.ParseFloat(value.(string), 64)
		if err != nil {
			return 0
		}
		return ret
	}
	return 0
}

// GetBool 获取 bool 值
func (vm ValueMap) GetBool(name string) bool {
	value := vm[name]
	if value == nil {
		return false
	}
	switch value.(type) {
	case bool:
		return value.(bool)
	case int:
		return value.(int) == 1
	case string:
		return value.(string) == "true"
	}
	return false
}

// // SetString 设置数据
//
//	func (f Values) SetString(name, value string) {
//		f[name] = value
//	}
func (vm ValueMap) ToJSON() []byte {
	ret, _ := sonic.Marshal(vm)
	return ret
}
