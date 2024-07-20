package main

import (
	"fmt"
	"singletonfactory"
)

func main() {
	fmt.Println("Singleton Factory")
	singletonfactory.GetInstance()
	singletonFactory := singletonfactory.GetInstance()
	singletonFactory2 := singletonfactory.GetInstance()
	singletonFactory.SetPrice(10)
	fmt.Println(singletonFactory.GetPrice())
	singletonFactory2.SetPrice(20)
	fmt.Println(singletonFactory2.GetPrice())
}
