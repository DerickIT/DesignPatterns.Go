package main

import "fmt"

type Memento struct {
	state string
}

func NewMemento(state string) *Memento {
	return &Memento{state: state}
}

func (m *Memento) GetState() string {
	return m.state
}

type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) SaveToMemento() *Memento {
	return NewMemento(o.state)
}

func (o *Originator) RestoreFromMemento(m *Memento) {
	o.state = m.GetState()
}

type Caretaker struct {
	mementos []*Memento
}

func (c *Caretaker) AddMemento(m *Memento) {
	c.mementos = append(c.mementos, m)
}

func (c *Caretaker) GetMemento(index int) *Memento {
	return c.mementos[index]
}

func main() {

	originator := &Originator{}
	caretaker := &Caretaker{}

	originator.SetState("State 1")
	fmt.Printf("Current state: %s\n", originator.GetState())
	caretaker.AddMemento(originator.SaveToMemento())

	originator.SetState("State 2")
	fmt.Printf("Current state: %s\n", originator.GetState())
	caretaker.AddMemento(originator.SaveToMemento())

	originator.SetState("State 3")
	fmt.Printf("Current state: %s\n", originator.GetState())
	caretaker.AddMemento(originator.SaveToMemento())

	originator.RestoreFromMemento(caretaker.GetMemento(1))

	fmt.Printf("Current state: %s\n", originator.GetState())

	originator.RestoreFromMemento(caretaker.GetMemento(0))
	fmt.Printf("Current state: %s\n", originator.GetState())
}
