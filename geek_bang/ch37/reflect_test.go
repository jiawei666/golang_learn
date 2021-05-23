package ch37

import (
	"fmt"
	"reflect"
	"testing"
)

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("浮点")
	case reflect.Int32:
		fmt.Println("32整型")
	case reflect.Int64:
		fmt.Println("64整型")
	case reflect.Int:
		fmt.Println("整型")
	default:
		fmt.Println("未知")
	}
	//fmt.Printf("%T\n", t.Kind())
	//fmt.Printf("%T\n", t)
}

func TestCheckType(t *testing.T) {
	CheckType(1)
}

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

type Employee struct {
	Id   int
	Name string `format:"hahha"`
	Age  int
}

func (e *Employee) UpdateAge(age int) {
	e.Age = age
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{1, "jiawei", 999}
	// 按名字获取成员
	t.Logf("name:%[1]v, type:%[1]T \n", reflect.ValueOf(*e).FieldByName("Name"))

	// reflect.Value还反射方法并执行
	reflect.ValueOf(e).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(888)})
	t.Logf("name:%v, age:%v", reflect.ValueOf(*e).FieldByName("Name"), reflect.ValueOf(*e).FieldByName("Age"))

	// reflect.Type可以反射标记tag
	if namefield, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("failed to get Name")
	} else {
		t.Log(namefield.Tag.Get("format"))
	}

}
