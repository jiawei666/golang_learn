package ch6

import "testing"

func TestArrayInit(t *testing.T)  {
	var arr [3]int
	arr1 := [4]int{1,2,3,4}
	arr3 := [...]int{1,3,4,5}
	arr1[1] = 5

	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

func TestArrayTestTrval(t *testing.T)  {
	arr3 := [...]int{1,3,4,5}
	for idx, v := range arr3 {
		t.Log(idx, v)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1,3,4,5}
	arr3_sec := arr3[0:3]
	arr3_sec1 := arr3[3:4]
	t.Log(arr3_sec)
	t.Log(arr3_sec1)
}

