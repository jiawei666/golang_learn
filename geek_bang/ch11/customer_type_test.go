package ch11

import (
	"fmt"
	"testing"
	"time"
)

type myFn func(n int) int

func timeSpent(inner myFn) myFn{
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println(
			"耗时：",
			time.Since(start).Seconds(),
		)
		return ret
	}
}

func slowFun(opt int)int {
	time.Sleep(time.Second * 4)
	return opt * 2
}

func TestFn1(t *testing.T) {
	newSlowFun := timeSpent(slowFun)
	ret := newSlowFun(10)
	t.Log(ret)
}
