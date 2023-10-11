package stringx

// Contains 判断字符串数组是否包含某一元素
func Contains(arr []string, s string) bool {
	for _, item := range arr {
		if item == s {
			return true
		}
	}
	return false
}

// Filter 字符串过滤器
// filter 返回 true 时，会过滤掉 r，反之
func Filter(s string, filter func(r string) bool) string {
	var n int
	chars := []rune(s)

	for i, x := range chars {
		if n < i {
			chars[n] = x
		}
		if !filter(string(x)) {
			n++
		}
	}

	return string(chars[:n])
}

// HasEmpty 判断是否有空值
func HasEmpty(args ...string) bool {
	for _, arg := range args {
		if len(arg) == 0 {
			return true
		}
	}

	return false
}

// NotEmpty 判断是否有空值
func NotEmpty(args ...string) bool {
	return !HasEmpty(args...)
}

// Remove 删除指定元素
func Remove(strings []string, outs ...string) []string {
	outArr := append([]string{}, outs...)

	temp := make([]string, 0)
	for _, str := range strings {
		if !Contains(outArr, str) {
			temp = append(temp, str)
		}
	}

	return temp
}

// Reverse 反转字符串
func Reverse(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

// Union 并集
func Union(slice1, slice2 []string) []string {
	result := append(slice1, slice2...)
	return RmDump(result)
}

// Difference 差集 a -> b
func Difference(a, b []string) []string {
	result := make([]string, 0)
	for _, s := range a {
		if !Contains(b, s) {
			result = append(result, s)
		}
	}
	return RmDump(result)
}

// Intersect 交集
func Intersect(slice1, slice2 []string) []string {
	result := make([]string, 0)
	for _, s := range slice1 {
		if Contains(slice2, s) {
			result = append(result, s)
		}
	}
	return RmDump(result)
}

// RmDump 去重
func RmDump(arr []string) []string {
	result := make([]string, 0)

	for _, s := range arr {
		if !Contains(result, s) {
			result = append(result, s)
		}
	}
	return result
}
