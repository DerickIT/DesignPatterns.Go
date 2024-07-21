package main

import "fmt"

type Vehicle interface {
	Drive() string
}

type Car struct{}

func (c *Car) Drive() string {
	return "Driving a car"
}

type Motorcycle struct{}

func (m *Motorcycle) Drive() string {
	return "Driving a motorcycle"
}

type VehicleFacoory interface {
	CreateVehicle() Vehicle
}

type CarFactory struct{}

func (cf *CarFactory) CreateVehicle() Vehicle {
	return &Car{}
}

type MotorcycleFactory struct{}

func (mf *MotorcycleFactory) CreateVehicle() Vehicle {
	return &Motorcycle{}
}

func main() {
	carFactory := &CarFactory{}
	car := carFactory.CreateVehicle()
	fmt.Println(car.Drive())

	motorFactory := &MotorcycleFactory{}
	motorcycle := motorFactory.CreateVehicle()
	fmt.Println(motorcycle.Drive())
}
