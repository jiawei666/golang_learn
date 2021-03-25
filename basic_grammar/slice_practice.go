package main

import "fmt"

//import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var res [][]uint8

	for i:=0;i<dy;i++ {

		var child []uint8
		for j:=0;j<dx;j++ {
			child = append(child, uint8(0))
		}
		res = append(res, child)

	}
	return res
}

func main() {
	a := Pic(3, 3)
	for _,v:=range a {
		printSlice(v)
	}

}

func printSlice(s []uint8) {
	fmt.Printf("len=%d cap=%d %u\n", len(s), cap(s), s)
}
