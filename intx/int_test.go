package intx

import "testing"

func TestEqual(t *testing.T) {

	a1 := []int{1, 2, 3, 4, 5}
	a2 := []int{3, 4, 1, 2, 5}
	a3 := []int{3, 4, 1, 2, 5, 6}

	t.Logf("1-2: %v", Equal(a1, a2))
	t.Logf("1-3: %v", Equal(a1, a3))
}
