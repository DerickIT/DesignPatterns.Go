package main

import (
	"sync"
	"sync/atomic"
)

type Counter struct {
	count int64
}

func NewCounter() *Counter {
	return &Counter{count: 0}
}

func (c *Counter) Increment() {
	atomic.AddInt64(&c.count, 1)
}

func (c *Counter) GetCount() int64 {
	return atomic.LoadInt64(&c.count)
}

func main() {

	counter := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	println(counter.GetCount())

}
