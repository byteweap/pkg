package randx

import (
	cryptorand "crypto/rand"
	"math/big"
	mathrand "math/rand"
	"time"
)

// Int [min, max) 随机数生成
// crypto/rand 生成安全的随机数,相比math/rand性能更好，推荐使用
func Int(min, max int) int {

	if min >= max || max == 0 {
		return max
	}
	result, _ := cryptorand.Int(cryptorand.Reader, big.NewInt(int64(max-min)))
	return int(result.Int64()) + min
}

// Int64 [min, max) 随机数生成
// crypto/rand 生成安全的随机数,相比math/rand性能更好，推荐使用
func Int64(min, max int64) int64 {

	if min >= max || max == 0 {
		return max
	}
	result, _ := cryptorand.Int(cryptorand.Reader, big.NewInt(int64(max-min)))
	return result.Int64() + min
}

// IsBingo 是否中奖
// 		percent: 概率 (%)
func IsBingo(percent int) bool {
	return Int(0, 10000) < percent*100
}

// IsBingo64 是否中奖
// 		percent: 概率 (%)
func IsBingo64(percent int64) bool {
	return Int64(0, 10000) < percent*100
}

// Intx [min, max) 随机数生成.
// math/rand 生成的随机数相比crypto/rand性能更差.
// Deprecated: use Int() instead.
func Intx(min, max int) int {
	if min >= max || max == 0 {
		return max
	}
	r := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}

// Int64x [min, max) 随机数生成.
// math/rand 生成的随机数相比crypto/rand性能更差.
// Deprecated: use Int64() instead.
func Int64x(min, max int64) int64 {
	if min >= max || max == 0 {
		return max
	}
	r := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	return r.Int63n(max-min) + min
}
