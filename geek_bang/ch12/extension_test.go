package ch12

import (
	"fmt"
	"testing"
)

type Pet struct {

}

func (pet *Pet) Speak()  {
	fmt.Print("...")
}

func (pet *Pet) SpeakTo(host string)  {
	pet.Speak()
	fmt.Print(" ", host)
}

type Dog struct {
	Pet
}

func (dog *Dog) Speak()  {
	fmt.Print("汪汪")
}


func TestDog(t *testing.T) {
	// go没有lsp 没有面向对象的继承
	// 因此 go是也不是一个面向对象语言
	var dog Dog = Dog{}
	dog.Speak()
	t.Log("")
	dog.SpeakTo("sb")
}



