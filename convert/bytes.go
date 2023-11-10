package convert

import (
	"reflect"
	"unsafe"
)

// Bytes2String converts byte slice to a string without memory allocation.
func Bytes2String(b []byte) string {
	/* #nosec G103 */
	return *(*string)(unsafe.Pointer(&b))
}

// String2Bytes converts string to a byte slice without memory allocation.
func String2Bytes(s string) (b []byte) {
	/* #nosec G103 */
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	/* #nosec G103 */
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}
