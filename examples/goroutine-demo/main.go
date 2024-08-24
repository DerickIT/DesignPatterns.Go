package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	totalPassengers = 200
	minProcessTime  = 100
	maxProcessTime  = 1000
	concurrentLimit = 5
)

func process(name string, in <-chan int, out chan<- int, wg *sync.WaitGroup, done <-chan bool) {
	defer wg.Done()
	for {
		select {
		case passenger, ok := <-in:
			if !ok {
				return
			}
			processTime := time.Duration(rand.Intn(maxProcessTime-minProcessTime+1)+minProcessTime) * time.Millisecond
			time.Sleep(processTime)
			fmt.Printf("%s processed passenger %d\n", name, passenger)
			if out != nil {
				out <- passenger
			}
		case <-done:
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()

	securityChan := make(chan int, concurrentLimit)
	ticketChan := make(chan int, concurrentLimit)
	idChan := make(chan int, concurrentLimit)
	doneChan := make(chan int, totalPassengers)
	done := make(chan bool)

	var wg sync.WaitGroup

	// 启动处理 goroutines
	for i := 0; i < concurrentLimit; i++ {
		wg.Add(3)
		go process("Security", securityChan, ticketChan, &wg, done)
		go process("Ticket", ticketChan, idChan, &wg, done)
		go process("ID", idChan, doneChan, &wg, done)
	}

	// 发送乘客
	go func() {
		for i := 1; i <= totalPassengers; i++ {
			securityChan <- i
		}
		close(securityChan)
	}()

	// 等待所有乘客处理完成
	go func() {
		for i := 0; i < totalPassengers; i++ {
			<-doneChan
		}
		close(done)
	}()

	wg.Wait()

	close(ticketChan)
	close(idChan)
	close(doneChan)

	elapsed := time.Since(start)
	fmt.Printf("Total time: %v\n", elapsed)
	fmt.Printf("Average time per passenger: %v\n", elapsed/time.Duration(totalPassengers))
}
