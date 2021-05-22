package ch24

import (
	"errors"
	"testing"
	"time"
)

type connObj struct {
}

type objPool struct {
	ch chan *connObj
}

func newPool(num int) *objPool {
	thech := make(chan *connObj, num)
	for i := 0; i < num; i++ {
		thech <- new(connObj)
	}
	pool := objPool{ch: thech}

	return &pool
}

func (pool *objPool) getConn(timeout time.Duration) (*connObj, error) {
	select {
	case ret := <-pool.ch:
		return ret, nil
	default:
		<-time.After(timeout)
		return nil, errors.New("time out")
	}
}

func (pool *objPool) releaseConn(conn *connObj) error {
	select {
	case pool.ch <- conn:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestPool(t *testing.T) {
	pool := newPool(10)
	//err := pool.releaseConn(new(connObj))
	//t.Log(err)
	for i := 0; i < 11; i++ {
		if conn, err := pool.getConn(10); err != nil {
			t.Log(err.Error())
		} else {
			t.Logf("类型%T \n", conn)
			if err := pool.releaseConn(conn); err != nil {
				t.Logf(err.Error())
			}
		}
	}
}
