package uuid

import "unsafe"

func unsafeStr(b *[]byte) string {
	return *(*string)(unsafe.Pointer(b))
}
