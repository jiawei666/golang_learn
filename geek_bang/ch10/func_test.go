package ch10

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues()(int, int)  {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T)  {
	a, b := returnMultiValues()
	t.Log(a, b)
}


// 包裹函数的用法，函数式编程
func timeSpent(inner func(opt int)int) func(n int) int{
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

func Sum(ops ...int) int {
	ret := 0
	for _, v := range ops {
		ret += v
	}

	return ret
}

func TestSumFn(t *testing.T)  {
	t.Log(Sum(1,2,3,4))
}

func Clear() {
	fmt.Println("执行defer")
}

func TestDefer(t *testing.T) {
	defer Clear()
	t.Log("开始")
	panic("报个错")
}




