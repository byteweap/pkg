package convert

import "testing"

func TestConvert(t *testing.T) {

	str := "9999999"
	bs := String2Bytes(str)

	data := Bytes2String(bs)

	t.Logf("data: %v", data)

}

func BenchmarkConvert(b *testing.B) {
	str := "9999999"

	// 性能测试对比
	for i := 0; i < b.N; i++ {
		a := String2Bytes(str)
		// a := []byte(str)
		_ = a
	}
}
