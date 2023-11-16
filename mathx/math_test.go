package mathx

import "testing"

// 笛卡尔积 组合算法
func BenchmarkCartesianProduct(b *testing.B) {

	b.ResetTimer()
	arr := make([][]int, 0)

	arr = append(arr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	arr = append(arr, []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	arr = append(arr, []int{21, 22, 23, 24, 25, 26, 27, 28, 29, 30})

	for i := 0; i < b.N; i++ {
		CartesianProduct(arr)
	}
}

func TestCartesianProduct(t *testing.T) {
	arr := make([][]int, 0)
	arr = append(arr, []int{1, 2, 3})
	arr = append(arr, []int{1, 2, 3})
	arr = append(arr, []int{7, 8, 9})
	for _, ints := range CartesianProduct(arr) {
		t.Log(ints)
	}
}
