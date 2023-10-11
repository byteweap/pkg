package qo

import (
	"fmt"
	"testing"
	"time"
)

func TestQo(t *testing.T) {
	fmt.Println("qo testing ...")

	q := New()

	start := time.Now()
	for i := 0; i < 10; i++ {
		v := i
		q.Go(func() {
			time.Sleep(time.Millisecond * 5)
			this := time.Now().UnixNano() - start.UnixNano()
			fmt.Println(fmt.Sprintf("%d", v), this)
			if v == 5 {
				b := 0
				a := 83 / b
				_ = a
			}
		})
	}

	q = New()
	for i := 0; i < 10; i++ {
		v := i
		q.Go(func() {
			time.Sleep(time.Millisecond * 5)
			this := time.Now().UnixNano() - start.UnixNano()
			fmt.Println(fmt.Sprintf("-- %d", v), this)
		})
	}

	q = New()
	for i := 0; i < 10; i++ {
		v := i
		q.Go(func() {
			time.Sleep(time.Millisecond * 5)
			this := time.Now().UnixNano() - start.UnixNano()
			fmt.Println(fmt.Sprintf("---- %d", v), this)
		})
	}

	all := time.Now().UnixNano() - start.UnixNano()

	fmt.Println("--------------- ", all)
	time.Sleep(time.Second * 3)
}
