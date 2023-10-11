package floatx

import "testing"

func TestFormatFloat(t *testing.T) {

	v1, _ := FormatFloat(1.0105, 3)
	v2, _ := FormatFloatCeil(1.0105, 3)
	v3, _ := FormatFloatFloor(1.0105, 3)
	t.Logf("1.0101 保留3位小数 : %v", v1)
	t.Logf("1.0101 保留3位小数Ceil: %v", v2)
	t.Logf("1.0101 保留3位小数Floor: %v", v3)

}
