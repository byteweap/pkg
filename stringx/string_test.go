package stringx

import (
	"testing"
)

func TestFilter(t *testing.T) {
	filter := Filter("Im a bcccccccoy", func(r string) bool {
		if r == "c" {
			return true
		}
		return false
	})
	t.Logf("result: %s", filter)
}

func TestRemove(t *testing.T) {
	arr := []string{`a`, `b`, `c`, `c`}
	t.Log(Remove(arr, "c", "c"))
}

func TestReverse(t *testing.T) {
	t.Log(Reverse("gfedcba"))
}

func TestUnion(t *testing.T) {

	t.Log(Union([]string{`a`, `b`, `c`, `c`}, []string{`c`, `e`, `d`}))
}

func TestRmDump(t *testing.T) {

	t.Log(RmDump([]string{`a`, `b`, `c`, `c`}))
}

func TestIntersect(t *testing.T) {

	t.Log(Intersect([]string{`a`, `b`, `c`}, []string{"c", "c"}))
}

func TestDifference(t *testing.T) {

	t.Log(Difference([]string{`a`, `b`, `c`, `c`}, []string{"c"}))
}
