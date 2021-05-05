package ch14

import (
	"errors"
	"testing"
)

func FibList(n int) (int, error) {

	// 及早失败，避免嵌套（）
	if n < 0 || n > 100 {
		return 0, errors.New("入参错误")
	}

	var a int
	var b int
	a = 1
	b = 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return b, nil
}

func TestGetFib(t *testing.T) {
	if v, err := FibList(5); err != nil {
		// 及早发现原则，把错误放到前面，出错则return
		t.Log(err.Error())
	} else {
		t.Log(v)
	}
}
