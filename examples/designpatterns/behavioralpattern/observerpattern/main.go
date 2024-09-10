package main

import "fmt"

type Subject interface {
	Register(observer Observer)
	UnRegister(observer Observer)
	NotifyAll()
}

type Observer interface {
	Update(data interface{})
}

type ConcreteSubject struct {
	observers []Observer
	data      interface{}
}

func (s *ConcreteSubject) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) UnRegister(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) NotifyAll() {
	for _, obs := range s.observers {
		obs.Update(s.data)
	}
}

type ConcreteObserver struct {
	id int
}

func (o *ConcreteObserver) Update(data interface{}) {
	fmt.Printf("Observer %d received update with data: %v\n", o.id, data)
}

func main() {
	subject := &ConcreteSubject{}
	obs1 := &ConcreteObserver{id: 1}
	obs2 := &ConcreteObserver{id: 2}

	subject.Register(obs1)
	subject.Register(obs2)

	subject.data = "Hello"
	subject.NotifyAll()

	subject.UnRegister(obs1)

	subject.data = "World"
	subject.NotifyAll()
}
