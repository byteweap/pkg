package logs

import (
	"testing"
)

func TestLogger(t *testing.T) {
	Init("debug", "", 6)
	Error("-------------- 游戏开始!!! GameCt: %v/%v", 1, 10)
	Errorx().Any("table", []int{1, 2, 3, 4, 5}).Msg("------")
}

func BenchmarkNew(b *testing.B) {
	//Init("debug", "")
	Init("debug", "logs", 6)

	for i := 0; i < b.N; i++ {
		Error("table: %d-------------- 游戏开始!!! GameCt: %v/%v", 21312, 1, 10)
	}
}
