package randx

import (
	"testing"
)

func TestInt(t *testing.T) {

	for i := 0; i < 10; i++ {
		t.Log(Int(0, 10))
	}
}

func TestInt64(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(Int64(0, 10))
	}
}

func BenchmarkInt(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Int(0, 10)
	}
}

func BenchmarkInt64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Int64(0, 10)
	}
}

func TestIntx(t *testing.T) {

	for i := 0; i < 10; i++ {
		t.Log(Intx(0, 10))
	}
}

func TestInt64x(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(Int64x(0, 10))
	}
}

func BenchmarkIntx(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Intx(0, 10)
	}
}

func BenchmarkInt64x(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Int64x(0, 10)
	}
}
