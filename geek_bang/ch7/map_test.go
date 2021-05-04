package ch7

import "testing"

// 初始化
func TestInitMap(t *testing.T)  {
	m1 := map[int]int{1:1,2:4,3:9}
	t.Log(m1[2])
	t.Logf("m1 len=%d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("m2 len=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("m3 len=%d", len(m3))
}

// 判断是否存在
func TestMapKeyExist(t *testing.T)  {
	m1 := map[int]int{}
	m1[1] = 2
	_, ok1 := m1[1]
	_, ok2 := m1[2]
	t.Log(ok1)
	t.Log(ok2)
}

// 遍历
func TestMapTrave(t *testing.T)  {
	m1 := map[int]string{1:"1", 2:"2", 3:"3"}
	for idx,v := range m1 {
		t.Log(idx, v)
	}
}

