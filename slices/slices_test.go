package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestOne struct {
	Name string
	Id   int
}

func TestUniq(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	arr := []string{"a", "a", "b", "c", "d"}
	result := Uniq(arr)
	assert.Equal(result, []string{"a", "b", "c", "d"})

}

func TestFilter(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	arr := []string{"a", "a", "b", "c", "d"}
	n1 := Filter(arr, func(index int, value string) bool {
		return value != "a"
	})

	assert.Equal(n1, []string{"b", "c", "d"})
}

func TestIn(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	arr := []TestOne{{"a", 1}, {"b", 1}}
	assert.True(In(TestOne{"a", 1}, arr))

	t1 := &TestOne{"a", 1}
	t2 := &TestOne{"a", 1}

	t.Log(t1 == t2)
}
