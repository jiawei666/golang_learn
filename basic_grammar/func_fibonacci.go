package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	lastRes := 0
	res := 0
	count := 0
	return func() int {
		count++
		if count == 1 {
			return res
		}

		if count == 2 {
			res = 1
			return res
		}

		returnRes := res + lastRes
		lastRes = res
		res = returnRes
		return returnRes
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
