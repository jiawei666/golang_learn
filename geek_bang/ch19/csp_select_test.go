package ch16

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "done"
}

func otherTask() {
	fmt.Println("working...")
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("task is done")
}

func TestService(t *testing.T) {
	t.Log(service())
	otherTask()
}

func asyncService() chan string {
	ch := make(chan string)
	go func() {
		msg := service()
		ch <- msg
	}()

	return ch
}

func TestAsyncService(t *testing.T) {
	ch := <-asyncService()
	otherTask()
	t.Log(ch)
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-asyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Log("timeout")

	}
}
