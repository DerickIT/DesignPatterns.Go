package main

import (
	"fmt"
	"sync"
	"time"
)

func race() {

	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	ready := false
	wg.Add(11)
	for i := 0; i < 10; i++ {
		// wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println("wait", num)
			cond.L.Lock()
			for !ready {
				cond.Wait()
			}
			fmt.Println("go", num)
			cond.L.Unlock()
		}(i)
	}
	time.Sleep(1 * time.Second)
	// wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("裁判员准备")

		cond.L.Lock()
		ready = true

		cond.Broadcast()
		cond.L.Unlock()
		fmt.Println("开始比赛")
	}()

	wg.Wait()

}

func main() {
	race()
}
