package ch20

import (
	"context"
	"fmt"
	"testing"
)

func isCancel(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCspCancel(t *testing.T) {
	ctx, cancelfunc := context.WithCancel(context.Background())

	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancel(ctx) {
					fmt.Printf("task %d done \n", i)
					break
				}
				//fmt.Printf("task %d still running \n", i)
			}
		}(i, ctx)
	}
	cancelfunc()
	//time.Sleep(time.Millisecond * 1)

}
