package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	totalPassengers = 100
	processTime     = 200 * time.Millisecond
	concurrentLimit = 5 // 每个检查点的并发处理数
)

func process(name string, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for passenger := range in {
		time.Sleep(processTime)
		fmt.Printf("%s processed passenger %d\n", name, passenger)
		if out != nil {
			out <- passenger
		}
	}
}

func main() {
	start := time.Now()
	securityChan := make(chan int, concurrentLimit)
	ticketChan := make(chan int, concurrentLimit)
	idChan := make(chan int, concurrentLimit)
	doneChan := make(chan int, totalPassengers)

	var wg sync.WaitGroup

	// 启动安检处理
	for i := 0; i < concurrentLimit; i++ {
		wg.Add(1)
		go process("Security", securityChan, ticketChan, &wg)
	}

	// 启动票检处理
	for i := 0; i < concurrentLimit; i++ {
		wg.Add(1)
		go process("Ticket", ticketChan, idChan, &wg)
	}

	// 启动身份检查处理
	for i := 0; i < concurrentLimit; i++ {
		wg.Add(1)
		go process("ID", idChan, doneChan, &wg)
	}

	// 发送乘客
	go func() {
		for i := 1; i <= totalPassengers; i++ {
			securityChan <- i
		}
		close(securityChan)
	}()

	// 等待所有处理完成并关闭通道
	go func() {
		wg.Wait()
		close(ticketChan)
		close(idChan)
		close(doneChan)
	}()

	// 统计处理完的乘客
	passengersProcessed := 0
	for range doneChan {
		passengersProcessed++
	}

	elapsed := time.Since(start)
	fmt.Printf("Total passengers processed: %d 耗时：%v\n", passengersProcessed, elapsed)
}
