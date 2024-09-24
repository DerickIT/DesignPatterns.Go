package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func watchDog(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "停止工作")
			return
		default:
			fmt.Println(name, "汪汪汪-工作中")

		}
		time.Sleep(1 * time.Second)
	}
}

func getUser(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("获取用户", "getUser停止工作")
			return
		default:
			userId := ctx.Value("userId")
			fmt.Println("获取用户", "userId:", userId)
			time.Sleep(1 * time.Second)
		}
	}
}
func main() {

	ctx, stop := context.WithCancel(context.Background())

	valCtx := context.WithValue(ctx, "userId", 2)

	wg.Add(4)
	go func() {
		defer wg.Done()
		getUser(valCtx)
	}()
	go func() {
		defer wg.Done()
		watchDog(ctx, "dog1")
	}()
	go func() {
		defer wg.Done()
		watchDog(ctx, "dog2")
	}()
	go func() {
		defer wg.Done()
		watchDog(ctx, "dog3")
	}()
	time.Sleep(5 * time.Second)
	stop()
	wg.Wait()
}
