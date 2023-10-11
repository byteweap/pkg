package values

import "testing"

func TestGetIntArray(t *testing.T) {

	vm := ValueMap{
		"sucUsers": []int{1, 2, 3},
	}

	t.Logf("v: %v", vm)
	t.Logf("aaa: %v", vm.GetIntArray("sucUsers"))
}
