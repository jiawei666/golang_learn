package ch8

import "testing"

// map的value可以是匿名函数
func TestMapWithFunValue(t *testing.T)  {
	m := map[int]func(arg int)int{}
	m[1] = func(arg int) int {
		return arg
	}
	m[2] = func(arg int) int {
		return arg * 2
	}
	m[3] = func(arg int) int {
		return arg * arg
	}

	t.Log(m[1](3), m[2](3), m[3](3))
}

// 用map实现set集合的基本功能
func TestMapForTest(t *testing.T)  {
	// 判断元素是否存在
	m := map[int]bool{}
	m[1] = true
	if m[1] {
		t.Log("exist")
	} else {
		t.Log("not found")
	}

}
