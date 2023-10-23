package limit

import (
	"testing"
)

func TestRun(t *testing.T) {

	l := New(3000)
	l.Run(func(datas []any) []any {
		return datas
	})

	go func() {
		//for i := 0; i < 10; i++ {
		//	go func(v int) {
		//		resp := <-l.Action(v)
		//		t.Logf("----- %v", resp)
		//	}(i)
		//}

		resp1 := l.Action(1)
		resp2 := l.Action(2)
		resp3 := l.Action(3)
		resp4 := l.Action(4)
		d1 := <-resp1
		d2 := <-resp2
		d3 := <-resp3
		d4 := <-resp4
		t.Logf("d1: %v, d2: %v, d3: %v, d4: %v", d1, d2, d3, d4)
	}()

	select {}
}
