package constant_test

import (
	"testing"
)



func TestIota1(t *testing.T) {
	const (
		monday = 1 + iota
		tusday
		wensday
	)
	t.Log(monday,tusday,wensday)
}

func TestIota2(t *testing.T){
	const(
		readable = iota << 1
		writable
		execable
	)
	a := 1
	t.Log(readable, writable, execable)
	t.Log(a&readable == readable, a & writable == writable, a & execable == execable)
}
