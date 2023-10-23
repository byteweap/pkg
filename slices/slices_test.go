package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
