package fib

import(
	"testing"
)

func TestFibList(t *testing.T) {
	var a int
	var b int
	a = 1
	b = 1
	t.Log(a)
	t.Log(b)
	for i := 0; i < 5; i++ {
		a, b = b, a + b
		t.Log(b)
	}
}


func TestChange(t *testing.T) {

	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}
