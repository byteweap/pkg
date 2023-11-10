package intx

import (
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {

	a1 := []int{1, 2, 3, 4, 5}
	a2 := []int{3, 4, 1, 2, 5}
	a3 := []int{3, 4, 1, 2, 5, 6}

	t.Logf("1-2: %v", Equal(a1, a2))
	t.Logf("1-3: %v", Equal(a1, a3))
}

func TestSonic(t *testing.T) {
	str := "{}"

	m := make(map[string]string)

	err := sonic.Unmarshal([]byte(str), m)
	if !assert.Error(t, err) {
		t.Log(m)
	}
}
