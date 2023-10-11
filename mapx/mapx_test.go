package mapx

import (
	"testing"
)

func TestMapx(t *testing.T) {

	// string - int
	// mint := New[string, int]()
	// mint.Set("name", 1)
	// t.Log(mint.Get("name"))
	//
	mstring := New[*Map, *Map]()
	m1 := &Map{}
	m1.Set(2, 2)
	m2 := &Map{}
	m2.Set(1, 2)
	mstring.Set(m1, m2)
	t.Log(mstring.Get(m1))
}
