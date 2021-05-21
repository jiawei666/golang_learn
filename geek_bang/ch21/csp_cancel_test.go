package ch21

import (
	"fmt"
	"testing"
	"time"
)

func isCancel(ch chan int) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func cancel1(ch chan int) {
	ch <- 1
}

func cancel2(ch chan int) {
	close(ch)
}

func TestCspCancel(t *testing.T) {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go func(i int, ch chan int) {
			for {
				if isCancel(ch) {
					fmt.Printf("task %d done \n", i)
					break
				}
				//fmt.Printf("task %d still running \n", i)
			}
		}(i, ch)
	}

	time.Sleep(time.Millisecond * 1)
	cancel1(ch)
	//cancel2(ch)
}
