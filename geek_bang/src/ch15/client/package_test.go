package client

import (
	"ch15/jiawei_test"
	"fmt"
	"testing"
)

// 使用包函数前会执行init init可以存在多个
func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func TestPackage(t *testing.T) {
	t.Log(jiawei_test.GetFibo(5))
	jiawei_test.GetName()
}

func TestHahah(t *testing.T) {
	t.Log("jijiji")
}
