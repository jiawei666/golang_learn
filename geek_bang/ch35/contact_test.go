package ch34

import "testing"

func TestSquare(t *testing.T) {
	inputs := []int{1, 2, 3}
	expects := []int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expects[i] {
			t.Errorf("输入%d的期望值是%d，却得到了%d", inputs[i], expects[i], ret)
		}
	}
}
