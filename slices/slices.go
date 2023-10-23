package slices

// 去重
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

// 过滤器
func Filter[T any](collection []T, predicate func(index int, value T) bool) []T {

	result := make([]T, 0, len(collection))

	for index, value := range collection {
		if predicate(index, value) {
			result = append(result, value)
		}
	}
	return result
}
