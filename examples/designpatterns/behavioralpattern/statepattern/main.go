package main

import "fmt"

type State interface {
	Handle(context *Context)
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle(context *Context) {
	fmt.Println("Handling state A")
	context.SetState(&ConcreteStateB{})
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle(context *Context) {
	fmt.Println("Handling state B")
	context.SetState(&ConcreteStateA{})
}

type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle(c)
}

func main() {
	context := &Context{state: &ConcreteStateA{}}

	context.Request()
	context.Request()

	context.Request()
	context.SetState(&ConcreteStateB{})
	context.Request()

}
