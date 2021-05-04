package ch9

import "testing"

func TestString(t *testing.T)  {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	// s[1] = 3 // 报错，因为string本质是[]byte的切片（只读，无法分赋值）
	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s)) // len()会计算byte的长度

}

func TestStringUnicode(t *testing.T)  {
	var s string
	s = "嘉炜"
	t.Log(len(s)) // len()会计算byte的长度

	c := []rune(s) // rune就是int32，byte就是int8，这里会把字符按每4个byte存放的方式放到rune
	t.Logf("长度%d;值=%d", len(c), c)

	t.Logf("嘉炜 unicode %x,%x", c[0], c[1])
	t.Logf("嘉炜 utf8 %x", s)

}

func TestStringTrave(t *testing.T)  {
	s := "嘉炜666"
	for _,v := range s {
		// 遍历后 v会变成rune类型，就是int32，
		// 我认为golang这样做的目的是兼容utf8编码
		t.Logf("%T %c %x",v, v, v)
	}
}


