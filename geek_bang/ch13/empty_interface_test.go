package ch13

import (
	"fmt"
	"testing"
)

func DoSome(p interface{}) {
	//if v,ok := p.(int); ok {
	//	fmt.Printf("int %d \n", v)
	//}else if v,ok := p.(string); ok {
	//	fmt.Printf("string %s \n", v)
	//} else {
	//	fmt.Println("unknow")
	//}

	switch v := p.(type) {
	case int:
		fmt.Printf("int %d \n", v)
	case string:
		fmt.Printf("string %s \n", v)
	default:
		fmt.Println("unknow")
	}
}

func TestEmptyInterface(t *testing.T)  {
	DoSome(10)
	DoSome("10")
}