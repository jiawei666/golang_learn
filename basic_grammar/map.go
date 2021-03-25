package main
import "fmt"
type student struct {
	id int
	name string
}

var myMap map[string]student = map[string]student {"asdasd" : {1, "jiawei"}}

func main()  {

	fmt.Println(myMap)
}
