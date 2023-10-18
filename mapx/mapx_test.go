package mapx

import (
	"testing"
)

func TestMapx(t *testing.T) {

	km := New[string, int](false)
	v := km.Get("foo")

	t.Logf("v: %v", v)
}

func BenchmarkSafeMapx(b *testing.B) {
	km := New[string, int](true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		km.Get("foo")
	}
}

func BenchmarkUnsafeMapx(b *testing.B) {
	km := New[string, int](false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		km.Get("foo")
	}
}
