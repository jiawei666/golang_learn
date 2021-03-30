package main

import "fmt"



func main() {

	var i interface{} = "hello"
	//
	s, ok := i.(string)
	fmt.Println(s, ok)
	fmt.Printf("(%v, %T)\n", i, i)
	//
	//f, ok := i.(float64)
	//fmt.Println(f, ok)
	//
	//f = i.(float64) // 报错(panic)
	//fmt.Println(f)
}
