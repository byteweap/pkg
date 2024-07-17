package logs

import (
	"testing"
)

func TestLogger(t *testing.T) {
	Init("debug", "", 6)
	Error().Any("table", []int{1, 2, 3, 4, 5}).Msg("------")
}

func BenchmarkNew(b *testing.B) {
	//Init("debug", "")
	Init("debug", "logs", 6)

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Error().Any("table", []int{1, 2, 3, 4, 5}).Msg("------")
	}
}
