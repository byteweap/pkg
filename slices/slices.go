package slices

import "math/rand"

// Uniq 去重
func Uniq[T comparable](collection []T) []T {

	result := make([]T, 0, len(collection))
	seen := make(map[T]struct{}, len(collection))

	for _, item := range collection {
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		result = append(result, item)
	}
	return result
}

// UniqBy 去重
// 通过uniqKey函数指定唯一条件
func UniqBy[T any, U comparable](collection []T, uniqKey func(iterm T) U) []T {
	result := make([]T, 0, len(collection))
	seen := make(map[U]struct{}, len(collection))

	for _, item := range collection {
		key := uniqKey(item)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		result = append(result, item)
	}
	return result
}

// Filter 过滤器
func Filter[T any](collection []T, okFunc func(index int, value T) bool) []T {

	result := make([]T, 0, len(collection))

	for index, value := range collection {
		if okFunc(index, value) {
			result = append(result, value)
		}
	}
	return result
}

// In 判断v是否在collection中
// T 需要是可比较的类型,指针类型不适用
func In[T comparable](v T, collection []T) bool {

	for i := 0; i < len(collection); i++ {
		if collection[i] == v {
			return true
		}
	}
	return false
}

// InBy 判断v是否在collection中
// equal函数为判定条件, src是v, target是collection[i]
func InBy[T any](v T, collection []T, equal func(src, target T) bool) bool {

	for i := 0; i < len(collection); i++ {
		if equal(v, collection[i]) {
			return true
		}
	}
	return false
}

// ForEach 遍历
func ForEach[T any](collection []T, anything func(index int, value T)) {
	for index, value := range collection {
		anything(index, value)
	}
}

// Chunk 均分
// collection被均分为长度为size的组，如果不能均分，则最后一组为剩余元素
func Chunk[T any](collection []T, size int) [][]T {
	if size <= 0 {
		panic("size parameter must be greater than 0")
	}

	chunksNum := len(collection) / size
	if len(collection)%size != 0 {
		chunksNum += 1
	}

	result := make([][]T, 0, chunksNum)

	for i := 0; i < chunksNum; i++ {
		last := (i + 1) * size
		if last > len(collection) {
			last = len(collection)
		}
		result = append(result, collection[i*size:last])
	}

	return result
}

// Shuffle 洗牌
func Shuffle[T any](collection []T) []T {

	rand.Shuffle(len(collection), func(i, j int) {
		collection[i], collection[j] = collection[j], collection[i]
	})
	return collection
}

// Count 元素个数
func Count[T comparable](collection []T, value T) (count int) {

	for _, item := range collection {
		if item == value {
			count++
		}
	}
	return
}

// CountBy 元素个数
// 接收is函数为判定条件, true计数,false不计数
func CountBy[T any](collection []T, okFunc func(item T) bool) (count int) {

	for _, item := range collection {
		if okFunc(item) {
			count++
		}
	}
	return
}
