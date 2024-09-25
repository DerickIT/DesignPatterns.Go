package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- fmt.Sprint("buy配件", i)
		}
	}()
	time.Sleep(1 * time.Second)
	return out
}

func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for v := range in {
			out <- fmt.Sprint(v, "-build")
		}
	}()
	return out
}

func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for v := range in {
			out <- fmt.Sprint(v, "-pack")
		}
	}()
	return out
}

func sell(in <-chan string) {
	for v := range in {
		fmt.Println(v, "-sell")
	}
}

func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	p := func(in <-chan string) {
		defer wg.Done()
		for v := range in {
			out <- v
		}
	}
	wg.Add(len(ins))
	for _, v := range ins {
		go p(v)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func tailRecur(n, res int) int {
	if n == 0 {
		return res
	}
	return tailRecur(n-1, res+n)
}

func fibnacci(n int) int {

	if n == 1 || n == 2 {
		return n - 1
	}
	res := fibnacci(n-1) + fibnacci(n-2)
	return res
}

func forLoopRecur(n int) int {
	stack := list.New()
	res := 0
	for i := n; i > 0; i-- {
		stack.PushBack(i)
	}
	for stack.Len() != 0 {
		res += stack.Back().Value.(int)
		stack.Remove(stack.Back())
	}

	// res=1+2+3+4+5+6+7+8+9+10
	return res
}

func main() {
	v0 := forLoopRecur(10)
	fmt.Println(v0)
	aa := fibnacci(5)
	fmt.Println(aa)
	bb := tailRecur(100000, 0)
	fmt.Println(bb)
	coms := buy(100)
	phones1 := build(coms)
	phones2 := build(coms)
	phones3 := build(coms)
	phones := merge(phones1, phones2, phones3)
	packs := pack(phones)
	sell(packs)
}

// func main() {
// 	result := make(chan int)
// 	go func() {

// 		time.Sleep(3 * time.Second)
// 		result <- 100
// 	}()

// 	select {
// 	case v := <-result:
// 		fmt.Println(v)
// 	case <-time.After(2 * time.Second):
// 		fmt.Println("timeout")
// 	}

// }
