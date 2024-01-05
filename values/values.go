package values

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/spf13/cast"
)

// 字典类型
type ValueMap map[string]any

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
func (vm ValueMap) GetValueMap(key string) ValueMap {
	value := vm[key]
	return cast.ToStringMap(value)
}

// GetString returns the string value associated with the given name in the ValueMap.
func (vm ValueMap) GetString(key string) string {
	value := vm[key]
	return cast.ToString(value)
}

// GetInt returns the integer value associated with the given name in the ValueMap.
func (vm ValueMap) GetInt(key string) int {
	value := vm[key]
	return cast.ToInt(value)
}

// GetInt64 returns the int64 value associated with the given name in the ValueMap.
func (vm ValueMap) GetInt64(key string) int64 {
	value := vm[key]
	return cast.ToInt64(value)
}

// GetIntArray returns an array of integers from the given key in the ValueMap.
func (vm ValueMap) GetIntArray(key string) []int {
	value := vm[key]
	return cast.ToIntSlice(value)
}

// GetStringArray returns an array of strings from the given key in the ValueMap.
func (vm ValueMap) GetStringArray(key string) []string {
	value := vm[key]
	return cast.ToStringSlice(value)
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

// GetTime retrieves a time.Time value from the ValueMap.
func (vm ValueMap) GetTime(name string) time.Time {
	value := vm[name]
	return cast.ToTime(value)
}

// ToJSON converts the ValueMap to a JSON byte slice.
//
// It returns a byte slice containing the JSON representation of the ValueMap.
func (vm ValueMap) ToJSON() []byte {
	ret, _ := sonic.Marshal(vm)
	return ret
}
