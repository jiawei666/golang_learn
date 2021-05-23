package ch39

import (
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	// 通过指针实现强制类型转换
	var a int64 = 10
	f := *(*float64)(unsafe.Pointer(&a))
	t.Log(f)
}
