package ch6

import "testing"

func TestSkiceInit(t *testing.T)  {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s2  := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2]) // 此时只能访问前三个元素，因为长度定义了3，只初始化了前三个
}

func TestSliceGrowing(t *testing.T) {
	s := []int {}
	for i:=0; i< 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s)) // 当cap容量不足时，会以倍数扩容（这里会有一定的性能问题）
	}
}

func TestSliceShareMemory(t *testing.T)  {
	months := [...]string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}

	q2 := months[3:6]
	summer := months[5:8]

	q2[2] = "jiawei" // 因为切片共享了内存，所以会同时修改summer的值

	t.Log(q2, len(q2), cap(q2))
	t.Log(summer, len(summer), cap(summer))
}


//func TestSliceCompare(t * testing.T) {
	//a := []int{1,2,3}
	//b := []int{1,2,3}
	//if a == b { // 会报错，slice不能进行比较
	//	t.Log('jiawei')
	//}
//}