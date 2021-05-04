package type_test

import (
	"math"
	"testing"
)


type jiaweiInt int64

// 不支持隐式转换
func TestType(t *testing.T)  {
	var a int = 1
	var b int64
	b = int64(a)

	var c jiaweiInt
	c = jiaweiInt(b)

	t.Log(a, b, c)
	t.Log(math.MaxInt8, math.MinInt8)
}

func TestPoint(t *testing.T){

	var a int = 1
	var aptr *int = &a
	//aptr = aptr + 1 // 报错 invalid operation: aptr + 1 (mismatched types *int and int)
	t.Logf("%T,%T", a, aptr)
}

func TestString(t *testing.T) {
	var s string
	t.Log(len(s), s == "")
}
