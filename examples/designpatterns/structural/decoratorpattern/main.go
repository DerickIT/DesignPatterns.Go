package main

import "fmt"

type Component interface {
	Operation() string
}

type ConcreteComponent struct{}

func (cc *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}

type Decorator interface {
	Component
	SetComponent(component Component)
}

type ConcreteDecoratorA struct {
	component Component
}

func (decoratorA *ConcreteDecoratorA) Operation() string {
	return "ConcreteDecoratorA: " + decoratorA.component.Operation()
}

func (decoratorA *ConcreteDecoratorA) SetComponent(component Component) {
	decoratorA.component = component
}

type ConcreteDecoratorB struct {
	component Component
}

func (decoratorB *ConcreteDecoratorB) Operation() string {
	return "ConcreteDecoratorB: " + decoratorB.component.Operation()
}

func (decoratorB *ConcreteDecoratorB) SetComponent(component Component) {
	decoratorB.component = component
}

func main() {

	concreteComponent := &ConcreteComponent{}
	decoratorA := &ConcreteDecoratorA{}

	decoratorA.SetComponent(concreteComponent)

	decoratorB := &ConcreteDecoratorB{}
	decoratorB.SetComponent(decoratorA)

	result := decoratorB.Operation()
	fmt.Println(result)

	// 循环依赖
	// decoratorA.SetComponent(decoratorB)
	// result2 := decoratorA.Operation()
	// fmt.Println(result2)

}
