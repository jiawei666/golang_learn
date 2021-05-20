package ch16

import (
	"sync"
	"testing"
	"time"
)

// mutex-线程安全
func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}

	time.Sleep(time.Millisecond * 100)

	t.Log(counter)
}

// 等待协程执行完成
func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup

	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				mut.Unlock()
			}()
			//time.Sleep(time.Second * 1)
			mut.Lock()
			counter++
		}()
	}

	wg.Wait()

	t.Log(counter)
}
