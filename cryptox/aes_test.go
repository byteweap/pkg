package cryptox

import "testing"

func TestEncryptAES(t *testing.T) {

	// key := "CtTSPzn9lgPBgKCuTzfJMw=="
	// data := "hello"
	// r, err := EncryptAES(key, data)

	// t.Logf(" data: %v, 加密后: %v, err: %v", string(data), r, err)

	d, err := DecryptAes("tl+fVCFFeRxOoVpRV177CA==", "q15CpbOL91W0OLMIcQy4BYjuuWN76r0FJkvy+eQZYGA=")
	t.Logf("解密: %v, err: %v", d, err)
}

func BenchmarkEncryptAES(b *testing.B) {

	b.ResetTimer()
	b.ReportAllocs()

	key := "12345678abcdefgh12345678"
	data := "hello"

	for i := 0; i < b.N; i++ {
		EncryptAES(key, data)
	}
}
