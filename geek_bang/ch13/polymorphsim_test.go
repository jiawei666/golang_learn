package ch13


import "testing"

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {}

func (g *GoProgrammer) WriteHelloWorld() Code  {
	return "i'm golang"
}

type JavaProgrammer struct {}

func (j *JavaProgrammer) WriteHelloWorld() Code  {
	return "i'm java"
}

func writeFirstProgram(p Programmer) Code  {
	return p.WriteHelloWorld()
}

func TestClient(t *testing.T)  {
	// go的接口实现不依赖定义，是Duck Type的形式
	// 接口是非入侵性
	//
	var p Programmer
	p = new(GoProgrammer)
	
	var j Programmer
	j = new(JavaProgrammer)
	// 实现多态
	t.Log(writeFirstProgram(p))
	t.Log(writeFirstProgram(j))
}