package main

import "fmt"

type Mediator interface {
	Notify(sender, event string)
}

type ConcreteMediator struct {
	components map[string]Component
}

func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{components: make(map[string]Component)}
}

func (m *ConcreteMediator) AddComponent(name string, component Component) {
	m.components[name] = component
}

func (m *ConcreteMediator) Notify(sender, event string) {
	for name, component := range m.components {
		if name != sender {
			component.Receive(event)
		}
	}
}

type Component interface {
	Send(event string)
	Receive(event string)
}

type ConcreteComponent struct {
	name     string
	mediator Mediator
}

func NewConcreteComponent(name string, mediator Mediator) *ConcreteComponent {
	return &ConcreteComponent{name: name, mediator: mediator}
}

func (c *ConcreteComponent) Send(event string) {
	fmt.Printf("%s send event: %s\n", c.name, event)
	c.mediator.Notify(c.name, event)
}

func (c *ConcreteComponent) Receive(event string) {
	fmt.Printf("%s receive event: %s\n", c.name, event)
	// c.mediator.
}

func main() {
	mediator := NewConcreteMediator()
	component1 := NewConcreteComponent("Component1", mediator)
	component2 := NewConcreteComponent("Component2", mediator)
	component3 := NewConcreteComponent("Component3", mediator)

	mediator.AddComponent("Component1", component1)
	mediator.AddComponent("Component2", component2)
	mediator.AddComponent("Component3", component3)

	component1.Send("Hello")
	component2.Send("Hello2")
	component3.Send("Hello3")
}
