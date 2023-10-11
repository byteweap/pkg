package mathx

import (
	"math/rand"
	"time"
)

// RandInt [min, max) 随机数生成
func RandInt(min, max int) int {
	if min >= max || max == 0 {
		return max
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}

// RandInt64 [min, max) 随机数生成
func RandInt64(min, max int64) int64 {
	if min >= max || max == 0 {
		return max
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(max-min) + min
}

// CartesianProduct 组合算法
//（The cartesian product - 笛卡尔积算法）
func CartesianProduct(sets [][]int) [][]int {
	// 根据下标获取长度方法
	lenFn := func(index int) int {
		return len(sets[index])
	}
	// 下一坐标
	nextIndex := func(ix []int, lens func(i int) int) {
		for j := len(ix) - 1; j >= 0; j-- {
			ix[j]++
			if j == 0 || ix[j] < lens(j) {
				return
			}
			ix[j] = 0
		}
	}

	product := make([][]int, 0)
	// core
	for ix := make([]int, len(sets)); ix[0] < lenFn(0); nextIndex(ix, lenFn) {
		var r []int
		for j, k := range ix {
			r = append(r, sets[j][k])
		}
		product = append(product, r)
	}
	return product
}
