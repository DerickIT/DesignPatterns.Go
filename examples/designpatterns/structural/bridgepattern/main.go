package main

import "fmt"

type Abstraction interface {
	Operation() string
}

type RefinedAbstraction interface {
	Abstraction
	SetImplementor(implementor Implementor)
}

type Implementor interface {
	OperationImpl() string
}

type ConcreteImplementorA struct{}
type ConcreteImplementorB struct{}

func (implA *ConcreteImplementorA) OperationImpl() string {
	return "ConcreteImplementorA: OperationImpl"
}

func (implB *ConcreteImplementorB) OperationImpl() string {
	return "ConcreteImplementorB: OperationImpl"
}

type RefinedAbstractionX struct {
	implementor Implementor
}

type RefinedAbstractionY struct {
	implementor Implementor
}

func (raX *RefinedAbstractionX) Operation() string {
	return "RefinedAbstractionX:  " + raX.implementor.OperationImpl()
}

func (raY *RefinedAbstractionY) Operation() string {
	return "RefinedAbstractionY:  " + raY.implementor.OperationImpl()
}

func (raX *RefinedAbstractionX) SetImplementor(implementor Implementor) {
	raX.implementor = implementor
}

func (raY *RefinedAbstractionY) SetImplementor(implementor Implementor) {
	raY.implementor = implementor
}

func main() {

	implA := &ConcreteImplementorA{}
	implB := &ConcreteImplementorB{}
	raX := &RefinedAbstractionX{}
	raY := &RefinedAbstractionY{}

	raX.SetImplementor(implA)
	raY.SetImplementor(implB)

	fmt.Println(raX.Operation())
	fmt.Println(raY.Operation())

	fmt.Println("Changing implementor OF raX to implB")
	raX.SetImplementor(implB)
	fmt.Println(raX.Operation())

}
