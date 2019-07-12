package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const max int32 = 1000

var pool = initPool()

func hitmesimple() {
	for i := 0; i < 100000; i++ {
		go simpleBufferOperation()
	}
	<-time.NewTicker(time.Duration(time.Second * 25)).C

}

func hitmepooled() {
	for i := 0; i < 100000; i++ {
		go pooledBufferOperation()
	}
	<-time.NewTicker(time.Duration(time.Second * 25)).C
}

func simpleBufferOperation() {
	buf := new(bytes.Buffer)
	buf.Write([]byte("roshan"))
	if buf.String() != "roshan" {
		fmt.Println("invalid output")
	}
}

func pooledBufferOperation() {
	atomic.AddInt32(&pool.counter, 1)
	buf := pool.p.Get().(*bytes.Buffer)
	buf.Write([]byte("roshan"))
	if buf.String() != "roshan" {
		fmt.Println("invalid output")
	}
	if atomic.AddInt32(&pool.counter, ^int32(0)) <= max {
		buf.Reset()
		pool.p.Put(buf)
	}
}

func initPool() *pBuf {
	p := new(pBuf)
	pool := new(sync.Pool)
	pool.New = func() interface{} {
		return new(bytes.Buffer)
	}
	p.p = pool
	return p
}

type pBuf struct {
	p       *sync.Pool
	counter int32
}
