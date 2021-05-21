package ch20

import (
	"fmt"
	"sync"
	"testing"
)

func producer(c chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
		wg.Done()
	}()
}

func receiver(c chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-c; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	wg := sync.WaitGroup{}
	ch := make(chan int, 10)
	wg.Add(1)
	producer(ch, &wg)
	wg.Add(1)
	receiver(ch, &wg)
	wg.Wait()

}
