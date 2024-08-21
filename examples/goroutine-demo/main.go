package main

import (
	"fmt"
	"sync"
)

const (
	numWorkers = 3
	numCycles  = 100
)

func worker(id int, prevCh, nextCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numCycles; i++ {
		<-prevCh
		fmt.Printf("Worker %d: %d\n", id+1, i+1)
		nextCh <- 1
	}
}

func main() {
	channels := make([]chan int, numWorkers)
	for i := range channels {
		channels[i] = make(chan int, 1)
	}

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(i, channels[i], channels[(i+1)%numWorkers], &wg)
	}

	// 启动第一个 worker
	channels[0] <- 1

	wg.Wait()
	fmt.Println("All workers finished")
}
