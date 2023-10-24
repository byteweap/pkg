package intx

import "sort"

// Contains 判断数组是否包含某一元素
func Contains(arr []int, s int) bool {
	for _, item := range arr {
		if item == s {
			return true
		}
	}
	return false
}

// SumInt 求和
func SumInt(arr []int) int {

	sum := 0
	for _, item := range arr {
		sum += item
	}
	return sum
}

// SumPre 求某一元素及之前所有元素的和
func SumPre(index int, arr []int) int {
	ret := 0
	if index >= 0 && index <= len(arr)-1 {
		ret = SumInt(arr[:index+1])
	}
	return ret
}

// Equal 相同数组(不计顺序)
func Equal(arr1, arr2 []int) bool {
	sort.Ints(arr1)
	sort.Ints(arr2)

	if len(arr1) == len(arr2) {
		for i, a := range arr1 {
			if arr2[i] != a {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

// IndexOf 下标, 没有则返回-1
func IndexOf(arr []int, s int) int {
	for i, v := range arr {
		if v == s {
			return i
		}
	}
	return -1
}
