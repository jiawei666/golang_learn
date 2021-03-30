// 接口的隐式实现
package main

import (
	"fmt"
)

type I interface {
	M() string
}

type T struct {
	S string
}

func (v T) M() string {
	return v.S
}

func main() {
	var aa I = T{"yuanjiawei"}
	fmt.Println(aa.M())
}


