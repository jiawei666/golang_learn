package main

import (
	"fmt"
	"math"
)

type Vex struct {
	x float64
	y float64
}


// 第一种写法
//func (v Vex) abs() float64  {
//	return math.Sqrt(v.x * v.x + v.y * v.y)
//}
//
//func main()  {
//	v := Vex{3, 4}
//	fmt.Println(v.abs())
//}

// 第二种写法（正常函数写法）
func abs(v Vex) float64  {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func main()  {
	v := Vex{3, 4}
	fmt.Println(abs(v))
}