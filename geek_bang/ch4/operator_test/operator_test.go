package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

func TestBit(t *testing.T) {
	const(
		readable = 1 << iota
		writable
		execable
	)

	a := 7 // 0111

	a = a &^ readable

	t.Log(a & readable == readable,
		a & writable == writable,
		a & execable == execable)
}
