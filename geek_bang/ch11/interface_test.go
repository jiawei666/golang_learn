package ch11

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {}

func (g *GoProgrammer) WriteHelloWorld() string  {
	return "i'm golang"
}

func TestClient(t *testing.T)  {
	// go的接口实现不依赖定义，是Duck Type的形式
	// 接口是非入侵性
	//
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}