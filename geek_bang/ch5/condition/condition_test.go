package condition

import "testing"

// 多条件case
func TestSwitchMuiltCase(t *testing.T) {
	for i:=0;i<5;i++ {
		switch i {
		case 0,2:
			t.Log("偶数")
		case 1,3:
			t.Log("奇数")
		default:
			t.Log("超过3了")
		}
	}
}

// 带条件case
func TestSwitchCaseCondition(t *testing.T) {
	for i:=0;i<5;i++ {
		switch {
		case i%2 == 0:
			t.Log("偶数")
		case i%2 == 1:
			t.Log("奇数")
		default:
			t.Log("未知")
		}
	}
}


