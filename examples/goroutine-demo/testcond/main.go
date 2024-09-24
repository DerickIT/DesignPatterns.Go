package main

import (
	"fmt"
	"sync"
	"time"
)

func race() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num, "waiting")
			cond.L.Lock()
			cond.Wait()
			fmt.Println(num, "is run ....")
			cond.L.Unlock()
		}(i)
	}

	time.Sleep(3 * time.Second)
	go func() {
		defer wg.Done()
		fmt.Println("broadcast 裁判就wei")
		fmt.Println("broadcast ready ....")
		cond.Broadcast()
	}()

	wg.Wait()

}

func main() {
	race()
}
