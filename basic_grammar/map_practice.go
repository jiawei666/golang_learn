package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	ss := strings.Fields(s)

	//fmt.Println(ss)

	var myMap map[string]int
	myMap = make(map[string]int)
	for _,v := range ss {

		//_,k := myMap[v]
		myMap[v] = myMap[v] + 1

	}

	return myMap
}

func main() {
	fmt.Println(WordCount("Hello 世\n界!\tHe\vl\flo!"))
}
