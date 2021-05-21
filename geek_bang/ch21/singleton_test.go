package ch21

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type singleObj interface {
}

var s *singleObj
var once = sync.Once{}

func getSingletonObj() *singleObj {
	once.Do(func() {
		fmt.Printf("hahah")
		s = new(singleObj)
	})

	return s
}

func TestSingleton(t *testing.T) {

	wg := sync.WaitGroup{}

	for i := 1; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			//fmt.Println(i)
			obj := getSingletonObj()
			fmt.Printf("%x \n", unsafe.Pointer(obj))
			wg.Done()
		}(i)
	}

	wg.Wait()

}
