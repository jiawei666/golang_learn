package ch23

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("task %d is running", id)
}

func FirstResponse() string {
	taskNum := 10
	ch := make(chan string, 10) // 增加容量防止阻塞
	for i := 1; i <= taskNum; i++ {
		go func(i int) {
			ch <- runTask(i)
		}(i)
	}

	return <-ch
}

func AllResponse() string {
	taskNum := 10
	ch := make(chan string, 10) // 增加容量防止阻塞
	for i := 1; i <= taskNum; i++ {
		go func(i int) {
			ch <- runTask(i)
		}(i)
	}

	ret := ""
	for i := 1; i <= taskNum; i++ {
		ret += <-ch + "\n"
	}

	return ret
}

func TestFirstResponse(t *testing.T) {
	t.Log("before", runtime.NumGoroutine())
	t.Log(AllResponse())

	time.Sleep(time.Millisecond * 100)
	t.Log("after", runtime.NumGoroutine())
}
