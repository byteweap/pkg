package msg

import (
	"testing"
	"time"

	"github.com/byteweap/pkg/redix"
)

func TestRequestResponse(t *testing.T) {
	rdx, err := redix.New("127.0.0.1:6379", "")
	if err != nil {
		t.Error(err)
		return
	}

	a2b := NewMsgChan("a", "b", rdx)

	// a 请求 b
	go func() {
		for {
			a2b.Request([]byte("12345"))
			t.Logf("request: %v", "12345")
			time.Sleep(time.Second * 2)
		}
	}()

	// a 监听 b的响应
	a2b.ListenResponse(func(data []byte) {
		t.Logf("response: %v ", string(data))
	})

	b2a := NewMsgChan("b", "a", rdx)

	// b接收a的请求
	b2a.ListenRequest(func(data []byte) {
		// b响应给a
		b2a.Response([]byte("I am b resp"))
	})

	select {}

}

func TestPush(t *testing.T) {

	rdx, err := redix.New("127.0.0.1:6379", "")
	if err != nil {
		t.Error(err)
		return
	}
	a2b := NewMsgChan("a", "b", rdx)
	go func() {
		for {
			a2b.Push([]byte("123"))
			time.Sleep(time.Second * 2)
		}
	}()

	b2a := NewMsgChan("b", "a", rdx)

	b2a.ListenPush(func(data []byte) {

		t.Logf("get push data : %v ", string(data))
	})

	select {}
}
