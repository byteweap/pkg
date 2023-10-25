package tool

import "testing"

func TestIf(t *testing.T) {

	a := If(true, 1, 2)
	t.Logf("a: %v", a)
}
