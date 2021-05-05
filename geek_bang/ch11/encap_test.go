package ch11

import (
	"fmt"
	"strconv"
	"testing"
	"unsafe"
)

type Employee struct {
	Id int
	Name string
	Age int
}

// 结构体行为（这种做法相当于函数传值）
func (e Employee) String() string {
	fmt.Printf("fn address is %x \n", unsafe.Pointer(&e.Id))
	return "Id: " + strconv.Itoa(e.Id) + ", Name: " + e.Name + ", Age" + strconv.Itoa(e.Age)
}

// 结构体行为（这种做法相当于函数传引用）
//func (e *Employee) String() string {
//	fmt.Printf("fn address is %x \n", unsafe.Pointer(&e.Id))
//	return "Id: " + strconv.Itoa(e.Id) + ", Name: " + e.Name + ", Age" + strconv.Itoa(e.Age)
//}

func TestCreateEmployee(t *testing.T) {
	e1 := Employee{1, "jiawei1", 18}
	e2 := Employee{Id: 2, Name: "jiawei2"}
	e3 := new(Employee)
	e3.Id = 3
	e3.Age = 18
	t.Log(e1)
	t.Log(e2)
	t.Log(e2.Id)
	t.Log(e3)
	t.Logf("e1 is %T", e1)
	t.Logf("e3 is %T", e3)
}


func TestStructOperations(t *testing.T) {
	e := Employee{1, "jiawei", 18}
	fmt.Printf("address is %x \n", unsafe.Pointer(&e.Id))
	t.Log(e.String())

}

