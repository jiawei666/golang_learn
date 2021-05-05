package ch14

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVsOsExit(t *testing.T) {
	defer func() {
		fmt.Println("now defer")
	}()
	fmt.Println("start")

	//os.Exit(-1)
	panic(errors.New("now panic"))
}

func TestPanicRecover(t *testing.T) {
	defer func() {
		// 相当于错误catch
		if err := recover(); err != nil {
			fmt.Println("错误信息: ")
			fmt.Println(err)
		}
	}()
	fmt.Println("start")

	panic(errors.New("now panic"))
}
