package values

import (
	"github.com/bytedance/sonic"
	"github.com/spf13/cast"
)

// 字典类型
type ValueMap map[string]any

type Value any

// NewValueMapArray creates a new array of ValueMap objects from a byte array.
//
// It takes a byte array as input and returns an array of ValueMap objects and an error.
func NewValueMapArray(data []byte) ([]ValueMap, error) {
	if len(data) == 0 {
		return []ValueMap{}, nil
	}
	var values []ValueMap
	err := sonic.Unmarshal(data, &values)
	return values, err
}

// NewValuesFromJSON creates a new ValueMap from JSON data.
//
// It takes a byte slice of JSON data as a parameter and returns a ValueMap and an error.
func NewValuesFromJSON(data []byte) (ValueMap, error) {
	if len(data) == 0 {
		return ValueMap{}, nil
	}
	var vm ValueMap
	err := sonic.Unmarshal(data, &vm)
	return vm, err
}

// GetValueMap returns a ValueMap for the given name.
func (vm ValueMap) GetValueMap(name string) ValueMap {
	value := vm[name]
	return cast.ToStringMap(value)
}

// GetString returns the string value associated with the given name in the ValueMap.
func (vm ValueMap) GetString(name string) string {
	value := vm[name]
	return cast.ToString(value)
}

// GetInt returns the integer value associated with the given name in the ValueMap.
func (vm ValueMap) GetInt(name string) int {
	value := vm[name]
	return cast.ToInt(value)
}

// GetInt64 returns the int64 value associated with the given name in the ValueMap.
func (vm ValueMap) GetInt64(name string) int64 {
	value := vm[name]
	return cast.ToInt64(value)
}

// GetIntArray returns an array of integers from the given key in the ValueMap.
func (vm ValueMap) GetIntArray(key string) []int {
	value := vm[key]
	return cast.ToIntSlice(value)
}

// GetFloat64 returns the float64 value associated with the given name in the ValueMap.
func (vm ValueMap) GetFloat64(name string) float64 {
	value := vm[name]
	return cast.ToFloat64(value)
}

// GetBool retrieves a boolean value from the ValueMap.
func (vm ValueMap) GetBool(name string) bool {
	value := vm[name]
	return cast.ToBool(value)
}

// ToJSON converts the ValueMap to a JSON byte slice.
//
// It returns a byte slice containing the JSON representation of the ValueMap.
func (vm ValueMap) ToJSON() []byte {
	ret, _ := sonic.Marshal(vm)
	return ret
}
