package main

import (
	"sync"
	"sync/atomic"
)

type Count struct {
	count int64
}

func NewCount() *Count {
	return &Count{count: 0}
}

func (c *Count) Increment() {
	atomic.AddInt64(&c.count, 1)
}

func (c *Count) GetCount() int64 {
	return atomic.LoadInt64(&c.count)
}

func main() {

	counter := NewCount()
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	if count := counter.GetCount(); count == 1000 {
		println("Counter is 1000, got:", count)
	}

}
