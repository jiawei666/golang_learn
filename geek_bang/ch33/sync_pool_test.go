package ch33

import (
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	sp := sync.Pool{
		New: func() interface{} {
			return "jiawei"
		},
	}

	sp.Put("jiaweiyuan")
	sp.Put("jiaweiyuan")
	sp.Put("jiaweiyuan")
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			t.Log(sp.Get())
			wg.Done()
		}()
	}
	wg.Wait()
}
