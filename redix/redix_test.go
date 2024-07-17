package redix

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestNew(t *testing.T) {

	r, err := New("127.0.0.1:6379", "")
	assert.Equal(t, err, nil)

	abc, err := r.Get("abc")
	assert.Equal(t, err, nil)
	t.Logf("v: %v | err: %v", abc, err)
}

func TestRedix_ListenAppMsg(t *testing.T) {
	r, _ := New("127.0.0.1:6379", "")

	type msg struct {
		Name    string
		ReqType string
	}

	r.ListenAppMsg("lf", func(data []byte) {
		var m msg
		err := json.Unmarshal(data, &m)
		if err != nil {
			return
		}
		t.Logf("Get Msg: %v", m)
	})

	go func() {
		for {
			time.Sleep(time.Second)
			d := msg{Name: "LuckyFruit", ReqType: "req_1"}
			data, _ := json.Marshal(d)
			if err := r.LPush("lf", data); err != nil {
				t.Errorf("LPush err: %v", err)
			}
		}
	}()

	time.Sleep(time.Second * 5)

	r.Close()

}
