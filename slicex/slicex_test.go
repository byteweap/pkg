package slicex

import "testing"

func TestSlicex(t *testing.T) {

	arr := New[int]()
	arr.Append(1)
	a := arr.Pop(0, 2)
	t.Logf("%v", a)
}

func BenchmarkSlicex(b *testing.B) {
	arr := New[int]()
	for i := 0; i < b.N; i++ {
		arr.Append(1)
	}

	b.Logf("%v", arr.Len())
}
